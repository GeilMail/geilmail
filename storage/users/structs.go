package users

import "github.com/jinzhu/gorm"

var db gorm.DB

type User struct {
	ID           uint
	Domain       Domain
	Mail         string `sql:"unique"`
	PasswordHash []byte
}

type Domain struct {
	DomainName string `sql:"primary_key"`
}

func Prepare(gdb gorm.DB) error {
	db = gdb
	d := db.AutoMigrate(&User{}, &Domain{})
	return d.Error
}
