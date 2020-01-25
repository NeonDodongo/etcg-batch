package consumer

import (
	"encoding/json"
	"etcg-batch/pkg/mongo"
	"etcg-batch/pkg/s3"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Controller holds the required connections for consumption and storage of this application
type Controller struct {
	Mongo  mongo.Connection
	S3     s3.Connection
	Client http.Client
}

var _controller Controller

func initController(config Config) error {
	m, err := mongo.New(config.Mongo)
	if err != nil {
		return err
	}

	s, err := s3.New(config.S3)
	if err != nil {
		return err
	}

	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       60 * time.Second,
	}

	_controller.Mongo = m
	_controller.S3 = s
	_controller.Client = c

	return nil
}

func (c Controller) decode(r io.Reader, t interface{}) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read while decoding content [ %w ]", err)
	}

	err = json.Unmarshal(content, &t)
	if err != nil {
		return fmt.Errorf("failed to unmarshal content to target interface [ %w ]", err)
	}

	return nil
}
