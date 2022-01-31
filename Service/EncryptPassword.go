package Service

import "golang.org/x/crypto/bcrypt"

const encryptionCost int = 8 // It's going to use 2^8 times to encrypt. (For normal users, 6 is recommended. For superAdmins, 8 is recommended)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), encryptionCost)

	return string(bytes), err
}
