package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"net/http"
	"strconv"
)

func GetTweets(writer http.ResponseWriter, request *http.Request) {
	defer responseOnError(writer)

	userId := request.URL.Query().Get("id")
	pageParam := request.URL.Query().Get("page")

	if len(userId) < 1 {
		panic("Id is required.")
	}

	if len(pageParam) == 0 {
		pageParam = "1"
	}

	page, err := strconv.Atoi(pageParam)

	if err != nil || page < 1 {
		panic("Page should be a valid number starting from 1")
	}

	pageNumber := int64(page)

	tweets, found := bd.GetTweetsFromUser(userId, pageNumber)

	if !found {
		http.Error(writer, "Tweets not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tweets)
}
