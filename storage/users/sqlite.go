package users

import (
	"fmt"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteUserStorage struct{}

func (s *SQLiteUserStorage) NewUser(u *User) error {
	tx, err := storage.SQLiteConn.Begin()
	if err != nil {
		return err
	}

	domain, err := helpers.MailDomainPart(helpers.MailAddress(u.Mail))
	if err != nil {
		return err
	}

	var domainID int
	rw := tx.QueryRow("SELECT id FROM domains WHERE domain = ?;", domain)
	err = rw.Scan(&domainID)
	if err != nil {
		return fmt.Errorf("Domain %s not found", domain)
	}

	_, err = tx.Exec("INSERT INTO users (user_id, domain_id, mail, password_hash) VALUES (null, ?, ?, ?, ?);", domainID, u.Mail, u.PasswordHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteUserStorage) UpdatePassword(u *User) error {
	return nil
}

func (s *SQLiteUserStorage) DeleteUser(u *User) error {
	return nil
}

type SQLiteDomainStorage struct{}

func (s *SQLiteDomainStorage) NewDomain(d *Domain) error {
	tx, err := storage.SQLiteConn.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO domains (id, domain) VALUES (null, ?);", string(d.Name))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
