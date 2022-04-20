package routers

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"io"
	"net/http"
	"os"
	"strings"
)

const avatarsPath string = "uploads/avatars/"
const avatarPermissions os.FileMode = 0666

func UploadAvatar(writer http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath = avatarsPath + LoggedUser.ID.String() + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, avatarPermissions)

	if err != nil {
		http.Error(writer, "Could not upload your avatar: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(writer, "Could not copy your avatar to the system: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	var status bool

	user.Avatar = strings.ReplaceAll(filePath, avatarsPath, "")
	status, err = bd.Update(user, LoggedUser.ID.String())

	if err != nil || !status {
		http.Error(writer, "Could not save your avatar into DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
