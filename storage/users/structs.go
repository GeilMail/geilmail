package users

import "gopkg.in/gorp.v1"

var db *gorp.DbMap

type User struct {
	ID           uint
	Domain       string
	Mail         string
	PasswordHash []byte
}

type Domain struct {
	DomainName string `sql:"primary_key"`
}

func Prepare(dbm *gorp.DbMap) {
	db = dbm
	db.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	db.AddTableWithName(Domain{}, "domains").SetKeys(false, "DomainName")
}
