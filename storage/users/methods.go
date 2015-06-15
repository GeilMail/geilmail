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
	u := User{
		Domain:       d,
		Mail:         mailAddr,
		PasswordHash: pwHash,
	}
	err := db.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func LookUpDomain(domainName string) (Domain, error) {
	d, err := db.Get(&Domain{}, domainName)
	if err != nil {
		return Domain{}, ErrNotFound
	}
	return d.(Domain), nil
}

func CheckPassword(mailAddr helpers.MailAddress, pw []byte) bool {
	u := &User{}
	db.Select(u, "SELECT password_hash FROM users WHERE mail = ?", mailAddr)
	return checkPassword(pw, u.PasswordHash)
}
