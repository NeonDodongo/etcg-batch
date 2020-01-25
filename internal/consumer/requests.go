package consumer

import (
	"encoding/json"
	"errors"
	"etcg-batch/internal/consumer/dto"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// EndpointURI is the endpoint for the API this application consumes from
const EndpointURI = "https://api.pokemontcg.io/v1"

func (c Controller) downloadAllSets() (dto.Sets, error) {
	r, err := http.Get(EndpointURI + "/sets")
	if err != nil {
		return dto.Sets{}, fmt.Errorf("error getting sets from pokemon tcg api [ %w ]", err)
	}

	defer r.Body.Close()

	//Read body into byte slice
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return dto.Sets{}, fmt.Errorf("error reading response body [ %w ]", err)
	}

	var sets dto.Sets

	//Unmarshal byte slice into sets struct
	if err := json.Unmarshal(body, &sets); err != nil {
		return dto.Sets{}, fmt.Errorf("error unmarshalling into sets [ %v ]", err)
	}

	return sets, nil
}

func (c Controller) downloadCardsBySetCode(setCode string) (dto.Cards, error) {
	setCode = strings.TrimSpace(setCode)

	if setCode == "" {
		return dto.Cards{}, errors.New("a set code must be provided")
	}

	cards, err := c.paginateCards(setCode)
	if err != nil {
		return dto.Cards{}, err
	}

	return cards, nil
}

func (c Controller) paginateCards(setCode string) (dto.Cards, error) {
	actualCards := make([]dto.Card, 0)
	count := 1

	for {
		//build query with setCode and replace <count>
		query := EndpointURI + "/cards?page=<count>&setCode=" + setCode
		uri := strings.Replace(query, "<count>", strconv.Itoa(count), 1)

		//call to pokemon tcg api
		resp, err := http.Get(uri)
		if err != nil {
			return dto.Cards{}, fmt.Errorf("error in call to uri: %s [ %w ]", uri, err)
		}

		var cards dto.Cards

		//decode response from call to dynamic tcg uri
		err = c.decode(resp.Body, &cards)
		if err != nil {
			return dto.Cards{}, fmt.Errorf("error decoding interface [ %w ]", err)
		}

		//append cards to temporary card array
		actualCards = append(actualCards, cards.Cards...)

		links := resp.Header.Get("link")
		if !strings.Contains(links, `rel="next"`) {
			break
		}

		count++
	}

	//set cards.Cards equal to the temp card array
	cds := dto.Cards{
		Cards: actualCards,
	}

	return cds, nil
}
