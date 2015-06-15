package users

import (
	"errors"

	"github.com/GeilMail/geilmail/helpers"
)

var (
	ErrNotFound = errors.New("no record found")
	ErrInternal = errors.New("internal error")
)

func New(d Domain, mailAddr string, pwHash []byte) error {
	_, err := GetDomainOrCreate(d.DomainName)
	if err != nil {
		return err
	}
	u := User{
		Domain:       d.DomainName,
		Mail:         mailAddr,
		PasswordHash: pwHash,
	}
	err = db.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func GetDomainOrCreate(domainName string) (*Domain, error) {
	domain, err := db.Get(Domain{}, domainName)
	if err != nil {
		return nil, err
	}
	if domain == nil {
		d := &Domain{
			DomainName: domainName,
		}
		err = db.Insert(d)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
	return domain.(*Domain), nil
}

func CheckPassword(mailAddr helpers.MailAddress, pw []byte) bool {
	u := &User{}
	db.Select(u, "SELECT password_hash FROM users WHERE mail = ?", mailAddr)
	return checkPassword(pw, u.PasswordHash)
}
