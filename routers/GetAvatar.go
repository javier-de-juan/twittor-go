package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"net/http"
	"os"
)

func GetAvatar(writer http.ResponseWriter, request *http.Request) {
	defer responseOnError(writer)

	userId := request.URL.Query().Get("id")

	if len(userId) < 1 {
		panic("Id is required.")
	}

	profile, err := bd.GetUserById(userId)

	if err != nil {
		panic("User not found")
	}

	OpenFile, err := os.Open(avatarsPath + profile.Avatar)
}
