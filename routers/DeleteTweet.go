package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"net/http"
)

func DeleteTweet(writer http.ResponseWriter, request *http.Request) {
	defer responseOnError(writer)

	tweetId := request.URL.Query().Get("id")

	if len(tweetId) < 1 {
		panic("Id is required.")
	}

	err := bd.DeleteTweet(tweetId, LoggedUser.ID.Hex())

	if err!=nil {
		http.Error(writer, "Error trying to remove the tweet: "+ err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
