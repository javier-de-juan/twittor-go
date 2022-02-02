package routers

import (
	"encoding/json"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/jwt"
	"github.com/javier-de-juan/twittor-go/models/requestModel"
	"github.com/javier-de-juan/twittor-go/models/responseModel"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-type", "application/json")
	defer responseOnError(writer)

	var userCredentials requestModel.Login

	err := json.NewDecoder(request.Body).Decode(&userCredentials)

	if err != nil {
		panic("invalid user / password" + err.Error())
		return
	}

	if len(userCredentials.Email) == 0 {
		panic("User's email is required")
	}

	user, found := bd.Login(userCredentials.Email, userCredentials.Password)

	if !found {
		http.Error(writer, "invalid user / password"+err.Error(), http.StatusUnauthorized)
		return
	}

	userJWTKey, err := jwt.GetJWT(user)

	if err != nil {
		http.Error(writer, "JWT could not be generated. The problem was:"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := responseModel.Login{
		Token: userJWTKey,
	}

	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, "Error encoding the response. The problem was:"+err.Error(), http.StatusInternalServerError)
		return
	}

	setJwtCookie(writer, userJWTKey)
}

func setJwtCookie(writer http.ResponseWriter, userJWTKey string) {
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   userJWTKey,
		Expires: time.Now().Add(24 * time.Hour),
	})
}
