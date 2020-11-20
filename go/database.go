package openapi

import (
	"context"
	"fmt"
	"sort"
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
func DbConnect(dbName string, password string) (*Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	atlasUrl := fmt.Sprintf("mongodb://mongoadmin:%s@some-mongo/admin", password)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(atlasUrl))
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection("queue")

	return &Database{client: client, collection: collection}, nil
}

func (db *Database) InsertNewVideo(code string, video Video) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", code}}
	update := bson.D{
		{"$push", bson.D{{"queue", video}}},
		{"$inc", bson.D{{"counter", 1}}}}

	_, err := db.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (db *Database) DeleteVideo(code string, videoId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": code}
	update := bson.M{
		"$pull": bson.M{
			"queue": bson.M{"_id": videoId},
		},
		"$inc": bson.M{
			"counter": 1,
		},
	}

	_, err := db.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (db *Database) GetVideoQueue(code string) (VideoQueue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"_id", code}}
	var result VideoQueue
	err := db.collection.FindOne(ctx, filter).Decode(&result)
	sortVideoQueue(&result)
	return result, err
}

func (db *Database) IncQueueCounter(code string) (VideoQueue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result VideoQueue
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.D{{"_id", code}}
	update := bson.D{{"$inc", bson.D{{"counter", 1}}}}
	err := db.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	return result, err
}

func (db *Database) UpdateVideoIndex(code string, videoId string, index int) (VideoQueue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result VideoQueue
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.D{{"_id", code}, {"queue._id", videoId}}
	update := bson.M{
		"$set": bson.M{"queue.$.index": index},
	}

	err := db.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	sortVideoQueue(&result)
	return result, err
}

func sortVideoQueue(videoQueue *VideoQueue) {
	sort.Slice(videoQueue.Queue, func(i, j int) bool {
		if videoQueue.Queue[i].Index == videoQueue.Queue[j].Index {
			return videoQueue.Queue[i].AddedAt.Before(videoQueue.Queue[j].AddedAt)
		}
		return videoQueue.Queue[i].Index > videoQueue.Queue[j].Index
	})
}
