package users

import "github.com/GeilMail/geilmail/helpers"

type DomainName string

type User struct {
	ID           int
	Domain       DomainName
	Mail         string
	Salt         string
	PasswordHash string
}

type Domain struct {
	ID   int
	Name DomainName
}

type UserStorage interface {
	NewUser(*User) error
	UpdatePassword(*User) error
	DeleteUser(*User) error
	GetUserByAddress(addr helpers.MailAddress)
}

type DomainStorage interface {
	NewDomain(*Domain) error
}
