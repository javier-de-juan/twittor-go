package routers

import (
	"errors"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/javier-de-juan/twittor-go/bd"
	"github.com/javier-de-juan/twittor-go/jwt"
	"github.com/javier-de-juan/twittor-go/models"
	"github.com/javier-de-juan/twittor-go/models/requestModel"
	"strings"
)

const splitter string = "Bearer"
var LoggedUser models.User

func IsValidRequestToken(requestToken string) (*requestModel.Claim, bool, string, error) {
	token := []byte(jwt.Key)
	claims := &requestModel.Claim{}

	splittedToken := strings.Split(requestToken, splitter)

	if len(splittedToken) != 2 {
		return claims, false, "", errors.New("token format is invalid")
	}

	requestToken = strings.TrimSpace(splittedToken[1])

	parsedToken, err := jwt2.ParseWithClaims(requestToken, claims, func(jsonWebToken *jwt2.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return claims, false, "", err
	}

	if !parsedToken.Valid {
		return claims, false, "", errors.New("token can be parsed with claim")
	}

	user, found, _ := bd.GetUserByEmail(claims.Email)

	if found {
		LoggedUser = user
	}

	return claims, found, LoggedUser.ID.Hex(), nil
}
