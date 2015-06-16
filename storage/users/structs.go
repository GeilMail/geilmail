package users

import "gopkg.in/gorp.v1"

var db *gorp.DbMap

type User struct {
	ID           uint
	Domain       string
	Mail         string
	PasswordHash []byte
}

func Prepare(dbm *gorp.DbMap) {
	db = dbm
	db.AddTableWithName(User{}, "users").SetKeys(true, "ID")
}
