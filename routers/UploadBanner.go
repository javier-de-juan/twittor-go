package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"io"
	"net/http"
	"os"
	"strings"
)

const bannersPath string = "uploads/banners/"
const bannerPermissions os.FileMode = 0666

func UploadBanner(writer http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath = bannersPath + LoggedUser.ID.String() + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, bannerPermissions)

	if err != nil {
		http.Error(writer, "Could not upload your banner: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(writer, "Could not copy your banner to the system: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	var status bool

	user.Banner = strings.ReplaceAll(filePath, bannersPath, "")
	status, err = bd.Update(user, LoggedUser.ID.String())

	if err != nil || !status {
		http.Error(writer, "Could not save your banner into DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
