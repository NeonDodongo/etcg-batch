package consumer

import (
	"bufio"
	"etcg-batch/internal/consumer/dto"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

var _reader *bufio.Reader

// Start will run the batch program with loaded Config
func Start(config Config) error {
	if err := initController(config); err != nil {
		return err
	}

	_reader = bufio.NewReader(os.Stdin)

	fmt.Println("/**************************************/")
	fmt.Println("Welcome to ETC-Batch | PTCG API Consumer")

	for {
		topMenu()

		c, _, err := _reader.ReadLine()
		if err != nil {
			return fmt.Errorf("failed to process user input [ %w ]", err)
		}

		in := strings.TrimSpace(string(c))

		if !validInput(in) {
			log.Error().Msgf(`Invalid user input: '%s'`, in)
			continue
		}

		if err := doTask(in); err != nil {
			log.Error().Err(err)
			continue
		}

		log.Info().Msg("Task completed successfully")
	}
}

func topMenu() {
	fmt.Println("/**************************************/")
	fmt.Println("Please choose an option below...")
	fmt.Println("1. Update Sets")
	fmt.Println("2. Update Cards by Existing Set")
	fmt.Println("3. Audit Card Count per Set")
	fmt.Println("0. Exit")
}

func doTask(in string) error {
	switch in {
	case "1":
		if err := updateSetsTask(); err != nil {
			return err
		}
	case "2":
		if err := updateCardsTask(); err != nil {
			return err
		}
	case "3":
		if err := auditCardsTask(); err != nil {
			return err
		}
	case "0":
		fmt.Println("/* CLOSING APPLICATION... */")
		time.Sleep(time.Second)
		os.Exit(0)
	}

	return nil
}

func updateSetsTask() error {
	log.Info().Msg("CHECKING DATABASE...")

	count, err := _controller.Mongo.Count(_controller.Mongo.SetCollection, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to count documents in collection: %s [ %w ]", _controller.Mongo.SetCollection, err)
	}

	log.Info().Msgf("Document Count (%s): %d", _controller.Mongo.SetCollection, count)
	log.Info().Msgf("Downloading SET data from PTCG API...")

	sets, err := _controller.downloadAllSets()
	if err != nil {
		return fmt.Errorf("error occurred while downloading data from PTCG API [ %w ]", err)
	}

	log.Info().Msgf("Available Count (%s): %d", _controller.Mongo.SetCollection, len(sets.Sets))

	if len(sets.Sets) > int(count) {
		for _, s := range sets.Sets {
			if err := _controller.Mongo.Upsert(s, bson.M{"code": s.Code}, _controller.Mongo.SetCollection); err != nil {
				return fmt.Errorf("Failed to upsert set: %s [ %w ]", s.Code, err)
			}
		}
	}
	return nil
}

func updateCardsTask() error {
	log.Info().Msg("CHECKING DATABASE...")

	// Get sets from MongoDB
	var sets []dto.Set
	if err := _controller.Mongo.Get(&sets, bson.M{}, _controller.Mongo.SetCollection); err != nil {
		return err
	}

	log.Debug().Msgf("Retrieved %d sets from MongoDB", len(sets))

	// Iterate through sets and download all cards per set from PTCG API
	log.Info().Msgf("Downloading CARD data from PTCG API...")
	for _, s := range sets {

		cards, err := _controller.downloadCardsBySetCode(s.Code)
		if err != nil {
			return err
		}

		if len(cards.Cards) != s.TotalCards {
			log.Warn().Msgf(`Discrepancy detected in '%s' - expected: %d, actual: %d`, s.Code, s.TotalCards, len(cards.Cards))
		}

		count := 0

		for _, c := range cards.Cards {
			if err := _controller.Mongo.Upsert(c, bson.M{"id": c.ID}, _controller.Mongo.CardCollection); err != nil {
				return err
			}
			count++
		}

		log.Info().Msgf("%d Cards were saved from Set %s", count, s.Code)
	}
	// Save each set of cards in MongoDB before moving onto the next set

	return nil
}

func auditCardsTask() error {
	log.Info().Msg("CHECKING DATABASE...")

	// Get sets from MongoDB
	var sets []dto.Set
	if err := _controller.Mongo.Get(&sets, bson.M{}, _controller.Mongo.SetCollection); err != nil {
		return err
	}

	for _, s := range sets {

		count, err := _controller.Mongo.Count(_controller.Mongo.CardCollection, bson.M{"setCode": s.Code})
		if err != nil {
			return err
		}

		if int(count) != s.TotalCards {
			log.Warn().Msgf(`Discrepancy detected in '%s' - expected: %d, actual: %d`, s.Code, s.TotalCards, count)
		}
	}

	return nil
}

func validInput(in string) bool {

	// if len(in) > 2 {
	// 	return false
	// }

	valid := map[string]string{
		"0": "",
		"1": "",
		"2": "",
		"3": "",
	}

	_, ok := valid[in]

	return ok
}
