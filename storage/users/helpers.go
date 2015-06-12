package users

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 6

func checkPassword(input string, dbval []byte) bool {
	err := bcrypt.CompareHashAndPassword(dbval, []byte(input))
	if err == nil {
		return true
	}
	return false
}
