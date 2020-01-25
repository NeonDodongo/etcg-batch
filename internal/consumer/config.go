package consumer

import (
	"etcg-batch/pkg/mongo"
	"etcg-batch/pkg/s3"
)

// Config holds the application configuration
type Config struct {
	Environment string         `json:"environment"`
	Port        string         `json:"port"`
	Timeout     int            `json:"timeout"`
	LogLevel    string         `json:"logLevel"`
	Mongo       mongo.ConnInfo `json:"mongo"`
	S3          s3.ConnInfo    `json:"s3"`
}
