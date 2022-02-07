package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/javier-de-juan/twittor-go/models"
	"time"
)

const Key string = "SomosLosPutosAmos_rishmawiTeam"

func GetJWT(user models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"lastname":   user.LastName,
		"birthday":   user.Birthday,
		"biography":  user.Biography,
		"location":   user.Location,
		"_id":        user.ID.Hex(),
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(Key))

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, err
}
