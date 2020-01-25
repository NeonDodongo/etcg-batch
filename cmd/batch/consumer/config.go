package main

import (
	"etcg-batch/internal/consumer"
	"etcg-batch/pkg/env"
)

var config consumer.Config

func loadConfig(filePath string) error {
	if err := env.LoadConfig(&config, filePath); err != nil {
		return err
	}

	return nil
}
