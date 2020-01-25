package mongo

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection is an object used to perform MongoDB transactions
type Connection struct {
	Client         *mongo.Client
	Database       string
	CardCollection string
	SetCollection  string
}

// Upsert will insert/update a document in MongoDB collection determined by the filter
func (db Connection) Upsert(t interface{}, filter bson.M, c string) error {
	collection := db.Client.Database(db.Database).Collection(c)
	update := bson.M{"$set": t}

	r, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("Failed to insert one to %s collection [ %v ]", c, err)
	}

	if r.MatchedCount > 1 || r.ModifiedCount > 1 {
		log.Warn().Msgf("Duplicate files detected during upsert to collection %s using filter %v", c, filter)
	}

	return nil
}

// Get will return a slice of documents determined by the filter
// Target interface must be a pointer to a slice of desired documents
func (db Connection) Get(t interface{}, filter bson.M, c string) error {

	if v := reflect.ValueOf(t); v.Kind() != reflect.Ptr && v.Kind() != reflect.Slice {
		return errors.New("Target for MongoDB operation must be a pointer to a slice")
	}

	col := db.Client.Database(db.Database).Collection(c)
	ctx := context.Background()

	cur, err := col.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to find documents based on filter [ %w ]", err)
	}

	if err := cur.All(ctx, t); err != nil {
		return fmt.Errorf("failed to read results to target interface [ %w ]", err)
	}

	return nil
}

// Count will return a count of the documents in the provided collection
func (db Connection) Count(c string, filter bson.M) (int64, error) {
	col := db.Client.Database(db.Database).Collection(c)

	return col.CountDocuments(context.Background(), filter)
}
