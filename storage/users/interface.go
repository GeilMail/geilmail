package users

import "github.com/GeilMail/geilmail/helpers"

var (
	UserProvider   UserStorage
	DomainProvider DomainStorage
)

type DomainName string
type UserID int

type User struct {
	ID           UserID
	Domain       DomainName
	Mail         string
	PasswordHash string
}

type Domain struct {
	ID   int
	Name DomainName
}

type UserStorage interface {
	NewUser(user *User) error
	GetUserByAddress(addr helpers.MailAddress) (*User, error)
	UpdatePassword(user *User, newPW string) error
	DeleteUser(user *User) error
}

type DomainStorage interface {
	NewDomain(domain *Domain) error
}
