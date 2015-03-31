package users

import (
	"testing"

	"github.com/GeilMail/geilmail/storage"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facebookgo/ensure"
)

func TestCreateDomainAndUser(t *testing.T) {
	var err error
	storage.SQLiteConn, err = sqlmock.New()
	if err != nil {
		t.Fail()
	}

	d := &Domain{
		ID:   0,
		Name: "example.com",
	}
	sqlmock.ExpectBegin()
	sqlmock.ExpectExec(`INSERT INTO domains \(id, domain\) VALUES \(null, (.*)\);`).WithArgs("example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlmock.ExpectCommit()

	domainStorage := &SQLiteDomainStorage{}
	ensure.Nil(t, domainStorage.NewDomain(d))

	//TODO: create user
}
