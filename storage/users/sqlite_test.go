package users

import (
	"fmt"
	"testing"

	"github.com/GeilMail/geilmail/storage"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facebookgo/ensure"
)

func TestCreateDomainAndUser(t *testing.T) {
	var err error
	storage.SQLiteConn, err = sqlmock.New()
	ensure.Nil(t, err)

	d := &Domain{
		ID:   0,
		Name: "example.com",
	}
	u := &User{
		ID:           0,
		Domain:       "example.com",
		Mail:         "test@example.com",
		Salt:         "ThisIsSomeRandomSalt",
		PasswordHash: "ThisIsAPasswordHash",
	}

	sqlmock.ExpectBegin()
	sqlmock.ExpectExec(`INSERT INTO domains \(id, domain\) VALUES \(null, (.*)\);`).WithArgs("example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlmock.ExpectCommit()

	domainStorage := &SQLiteDomainStorage{}
	ensure.Nil(t, domainStorage.NewDomain(d))

	sqlmock.ExpectBegin()
	sqlmock.ExpectQuery(`SELECT id FROM domains WHERE domain = \?;`).WithArgs("example.com").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	sqlmock.ExpectExec(`INSERT INTO users \(user_id, domain_id, mail, salt, password_hash\) VALUES \(null, \?, \?, \?, \?\);`).WithArgs(1, u.Mail, u.Salt, u.PasswordHash).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlmock.ExpectCommit()

	//TODO: create user
	userStorage := SQLiteUserStorage{}
	ensure.Nil(t, userStorage.NewUser(u))
}

func TestCreatingUserForInexistentDomain(t *testing.T) {
	var err error
	storage.SQLiteConn, err = sqlmock.New()
	ensure.Nil(t, err)

	u := &User{
		ID:           0,
		Domain:       "inexistent.example.com",
		Mail:         "inexistent@inexistent.example.com",
		Salt:         "random",
		PasswordHash: "oink",
	}

	sqlmock.ExpectBegin()
	sqlmock.ExpectQuery(`SELECT id FROM domains WHERE domain = \?;`).WithArgs("inexistent.example.com").WillReturnError(fmt.Errorf("No row found"))

	UserStorage := SQLiteUserStorage{}
	ensure.NotNil(t, UserStorage.NewUser(u))
}
