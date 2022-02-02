package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/models"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user) // Una vez que uso el body, no se puede volver a usar porque es de tipo STREAM

	if err != nil {
		http.Error(writer, "Bad requestModel: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer responseOnError(writer)

	validate(user)

	_, status, err := bd.Save(user)

	if err != nil {
		http.Error(writer, "Internal error saving user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(writer, "User was not saved", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func validate(user models.User) {
	if len(user.Email) == 0 {
		panic("User's email is required")
	}

	if len(user.Password) < 6 {
		panic("User's password has a min length of 6 chars.")
	}

	_, userFound, _ := bd.GetUserByEmail(user.Email)

	if userFound {
		panic("User already exists.")
	}
}
