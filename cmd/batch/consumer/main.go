package main

import (
	"etcg-batch/internal/consumer"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := loadConfig("./cfg/config.json"); err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Configuration successful")
	log.Info().Msg("Application Start")

	if err := consumer.Start(config); err != nil {
		log.Fatal().Err(err)
	}
}
