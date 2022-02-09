package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"github.com/javier-de-juan/twittor-go/models/requestModel"
	"net/http"
	"time"
)

func Tweet(writer http.ResponseWriter, request *http.Request) {
	var tweetRequest requestModel.Tweet

	err := json.NewDecoder(request.Body).Decode(&tweetRequest)

	if err != nil {
		http.Error(writer, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer responseOnError(writer)

	if len(tweetRequest.Text) <= 0 || len(tweetRequest.Text) > 240 {
		panic("Tweet has no valid length (1 to 240 characters).")
	}

	tweet := models.Tweet{
		UserID: LoggedUser.ID.Hex(),
		Text: tweetRequest.Text,
		CreatedAt: time.Now(),
	}

	_, status, err := bd.SaveTweet(tweet)

	if err != nil {
		http.Error(writer, "Internal error saving tweet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(writer, "Tweet was not saved", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}