package bd

import (
	"context"
	"github.com/javier-de-juan/twittor-go/models"
	"github.com/javier-de-juan/twittor-go/models/responseModel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const tweetCollection = "tweets"
const rowsPerPage = 20

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

func GetTweetsFromUser(ID string, page int64) ([]*responseModel.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	dbCollection := db.Collection(tweetCollection)

	var tweets []*responseModel.Tweet

	filter := bson.M{
		"user_id": ID,
	}

	options := mongoOptions.Find()
	options.SetLimit(rowsPerPage)
	options.SetSort(bson.D{{Key: "created_at", Value: -1}})
	options.SetSkip((page - 1) * rowsPerPage)

	cursor, err := dbCollection.Find(ctx, filter, options)

	if err != nil {
		log.Fatal(err.Error())

		return tweets, false
	}

	// No me interesa un contexto. Con "TODO" lo que puedo hacer es esperar indefinidamente
	for cursor.Next(context.TODO()) {
		var tweet responseModel.Tweet
		err := cursor.Decode(&tweet)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		tweets = append(tweets, &tweet)
	}

	return tweets, true
}

func DeleteTweet(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	dbCollection := db.Collection(tweetCollection)

	tweetId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M {
		"_id": tweetId,
		"user_id": UserId,
	}

	_, err := dbCollection.DeleteOne(ctx, condition)

	return err
}
