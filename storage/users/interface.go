package users

import "github.com/GeilMail/geilmail/helpers"

var (
	UserProvider   UserStorage
	DomainProvider DomainStorage
)

type DomainName string

type User struct {
	ID           int
	Domain       DomainName
	Mail         string
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
