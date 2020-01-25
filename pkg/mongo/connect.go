package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnInfo holds the connection info required to interface with MongoDB
type ConnInfo struct {
	URI            string `json:"uri"`
	Database       string `json:"database"`
	CardCollection string `json:"cardCollection"`
	SetCollection  string `json:"setCollection"`
}

// DecodeConnInfo not implemented yet
func DecodeConnInfo(ci ConnInfo) ConnInfo {
	// TODO: Implement decoding encoded secrets (base64 probably)

	return ConnInfo{}
}

// New creates an instance of a MongoDB Connection
func New(c ConnInfo) (Connection, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI))
	if err != nil {
		return Connection{}, fmt.Errorf("Failed to create MongoDB client [ %w ]", err)
	}

	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		return Connection{}, fmt.Errorf("Failed to connect to MongoDB [%v]", err)
	}

	db := Connection{
		Client:         client,
		Database:       c.Database,
		CardCollection: c.CardCollection,
		SetCollection:  c.SetCollection,
	}

	return db, nil
}
