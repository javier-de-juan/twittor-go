package bd

import (
	"context"
	"github.com/javier-de-juan/twittor-go/Service"
	"github.com/javier-de-juan/twittor-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const collection = "users"

func Save(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	userCollection := db.Collection(collection)

	user.Password, _ = Service.EncryptPassword(user.Password)

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	userID, _ := result.InsertedID.(primitive.ObjectID)

	return userID.String(), true, nil
}

func GetUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	userCollection := db.Collection(collection)

	condition := bson.M{"email":email}

	var user models.User

	err := userCollection.FindOne(ctx, condition).Decode(&user)

	if err != nil {
		return user, false, user.ID.Hex()
	}

	return user, true, user.ID.Hex()
}

func getMaxExecutionTime() time.Duration {
	return 15 * time.Second
}
