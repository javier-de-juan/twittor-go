package bd

import (
	"github.com/javier-de-juan/twittor-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := GetUserByEmail(email)

	if !found {
		return user, false
	}

	loginPassword := []byte(password)
	userPassword := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(userPassword, loginPassword)

	if err != nil {
		return user, false
	}

	return user, true
}
