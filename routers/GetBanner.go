package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"io"
	"net/http"
	"os"
)

func GetBanner(writer http.ResponseWriter, request *http.Request) {
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

	profileBanner, err := os.Open(bannersPath + profile.Banner)

	if err != nil || len(profile.Banner) < 1 {
		http.Error(writer, "Banner not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(writer, profileBanner)

	if err != nil {
		http.Error(writer, "Could not upload read banner file: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
