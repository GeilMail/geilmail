package users

import "github.com/GeilMail/geilmail/helpers"

func NewUser(d Domain, mailAddr string, pwHash []byte) error {
	u := User{
		Domain:       d,
		Mail:         mailAddr,
		PasswordHash: pwHash,
	}
	db.NewRecord(&u)
	//TODO: error handling
	return nil
}

func CheckPassword(mailAddr helpers.MailAddress, pw []byte) bool {
	u := &User{}
	db.Where("mail = ?", mailAddr).Find(u)
	return checkPassword(pw, u.PasswordHash)
}
