package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"net/http"
)

func Profile(writer http.ResponseWriter, request *http.Request) {
	defer responseOnError(writer)

	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		panic("ID is required")
	}

	if ID == LoggedUser.ID.String() {
		returnProfile(writer, LoggedUser)
		return
	}

	user, err := bd.GetUserById(ID)

	if err != nil {
		http.Error(writer, "User not found", http.StatusNotFound)
		return
	}

	returnProfile(writer, user)
}

func returnProfile(writer http.ResponseWriter, user models.User) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(user)

	if err != nil {
		return
	}
}
