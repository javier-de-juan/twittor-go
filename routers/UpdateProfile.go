package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"net/http"
)

func UpdateProfile(writer http.ResponseWriter, request *http.Request) {
	defer responseOnError(writer)
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		panic(err.Error())
		return
	}

	updated, err := bd.Update(user, LoggedUser.ID.Hex())

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(writer, "User was not updated", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
