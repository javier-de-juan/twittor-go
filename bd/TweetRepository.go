package bd

import (
	"context"
	"github.com/javier-de-juan/twittor-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const tweetCollection = "tweets"

func SaveTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	dbCollection := db.Collection(tweetCollection)
	
	result, err := dbCollection.InsertOne(ctx, tweet)

	if err != nil {
		return "", false, err
	}

	tweetID, _ := result.InsertedID.(primitive.ObjectID)

	return tweetID.String(), true, nil
}