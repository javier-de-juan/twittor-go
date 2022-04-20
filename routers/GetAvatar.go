package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"io"
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
		http.Error(writer, "User not found: ", http.StatusNotFound)
		return
	}

	profileAvatar, err := os.Open(avatarsPath + profile.Avatar)

	if err != nil || len(profile.Avatar) < 1 {
		http.Error(writer, "Avatar not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(writer, profileAvatar)

	if err != nil {
		http.Error(writer, "Could not read avatar file: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
