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

func Update(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	userCollection := db.Collection(collection)

	userUpdateString := bson.M{
		"$set": getUserInfoToUpdate(user),
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err := userCollection.UpdateOne(ctx, filter, userUpdateString)

	if err != nil {
		return false, err
	}

	return true, err
}

func getUserInfoToUpdate(user models.User) map[string]interface{} {
	userUpdated := make(map[string]interface{})

	if len(user.Name) > 0 {
		userUpdated["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		userUpdated["lastName"] = user.LastName
	}

	if len(user.Banner) > 0 {
		userUpdated["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		userUpdated["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		userUpdated["location"] = user.Location
	}

	if len(user.Website) > 0 {
		userUpdated["website"] = user.Website
	}

	userUpdated["birthday"] = user.Birthday

	return userUpdated
}

func GetUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	userCollection := db.Collection(collection)

	condition := bson.M{"email": email}

	var user models.User

	err := userCollection.FindOne(ctx, condition).Decode(&user)

	if err != nil {
		return user, false, user.ID.Hex()
	}

	return user, true, user.ID.Hex()
}

func GetUserById(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), getMaxExecutionTime())

	defer cancel()

	db := MongoCN.Database(dbName)
	userCollection := db.Collection(collection)
	id, _ := primitive.ObjectIDFromHex(userId)

	condition := bson.M{"_id": id}

	var user models.User

	err := userCollection.FindOne(ctx, condition).Decode(&user)

	if err != nil {
		return user, err
	}

	user.Password = ""

	return user, err
}

func getMaxExecutionTime() time.Duration {
	return 15 * time.Second
}
