package users

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 6

func checkPassword(input []byte, dbval []byte) bool {
	err := bcrypt.CompareHashAndPassword(dbval, input)
	if err == nil {
		return true
	}
	return false
}

func HashPassword(input []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(input, bcryptCost)
}
