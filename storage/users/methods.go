package users

import (
	"errors"

	"github.com/GeilMail/geilmail/helpers"
)

var (
	ErrNotFound = errors.New("no record found")
	ErrInternal = errors.New("internal error")
)

func New(mailAddr helpers.MailAddress, password string) error {
	pwHash, err := HashPassword([]byte(password))
	if err != nil {
		return err
	}
	u := User{
		Mail:         string(mailAddr),
		PasswordHash: pwHash,
	}
	err = db.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func CheckPassword(mailAddr helpers.MailAddress, pw []byte) bool {
	u := &User{}
	err := db.SelectOne(u, "SELECT passwordHash FROM users WHERE mail = ?;", string(mailAddr))
	if err != nil {
		return false
	}
	return checkPassword(pw, u.PasswordHash)
}

// AllDomains retrieves all active domains that have mailboxes.
func AllDomains() (domains []string, err error) {
	var addrs []string
	_, err = db.Select(&addrs, "SELECT mail FROM users;")
	if err != nil {
		return
	}
	mSet := map[string]struct{}{}
	for _, ad := range addrs {
		dp, err := helpers.MailAddress(ad).DomainPart()
		if err != nil {
			return nil, err
		}
		mSet[dp] = struct{}{}
	}
	for ad := range mSet {
		domains = append(domains, ad)
	}
	return
}
