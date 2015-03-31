package users

type DomainName string

type User struct {
	ID       int
	Domain   DomainName
	Mail     string
	Salt     string
	Password string
}

type Domain struct {
	ID   int
	Name DomainName
}

type UserStorage interface {
	NewUser(*User) error
	UpdatePassword(*User) error
	DeleteUser(*User) error
}

type DomainStorage interface {
	NewDomain(*Domain) error
}
