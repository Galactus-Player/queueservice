package openapi

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// RemoveVideo - Remove video from the queue
func DbConnect(password string, dbName string) (*Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://BSathvik:<pass>@cluster0.hrnmp.mongodb.net/firstTest?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}

	collection := client.Database("testing").Collection("queue")

	return &Database{client: client, collection: collection}, nil
}

func (db *Database) InsertNewVideo(code string, video Video) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// _, err := db.collection.InsertOne(ctx, video)

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", code}}
	update := bson.D{{"$push", bson.D{{"queue", video}}}}

	_, err := db.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (db *Database) DeleteVideo(code string, videoId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// _, err := db.collection.InsertOne(ctx, video)

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", code}}
	update := bson.D{{"$pull", bson.D{{"queue", bson.D{{"_id", videoId}}}}}}

	_, err := db.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
