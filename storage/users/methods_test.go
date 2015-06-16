package users

import (
	"database/sql"
	"os"
	"testing"

	"github.com/facebookgo/ensure"
	"gopkg.in/gorp.v1"

	// driver import
	_ "github.com/mattn/go-sqlite3"
)

var testDBPath = "test.db"

func testBuildDB(t *testing.T) {
	sqldb, err := sql.Open("sqlite3", testDBPath)
	ensure.Nil(t, err)

	gdb := &gorp.DbMap{Db: sqldb, Dialect: gorp.SqliteDialect{}}
	Prepare(gdb)
	err = gdb.CreateTablesIfNotExists()
	ensure.Nil(t, err)
}

func testTeardownDB(t *testing.T) {
	os.Remove(testDBPath)
}

func TestUserCreateAndLogin(t *testing.T) {
	testBuildDB(t)
	pwHash, err := HashPassword([]byte("123456"))
	ensure.Nil(t, err, pwHash)
	err = New("test@example.com", pwHash)
	ensure.Nil(t, err)

	ensure.True(t, CheckPassword("test@example.com", []byte("123456")), "oink")
	ensure.False(t, CheckPassword("test@example.com", []byte("")))
	ensure.False(t, CheckPassword("test2@example.com", []byte("123456")))
	testTeardownDB(t)
}

func TestDomainListing(t *testing.T) {
	testBuildDB(t)
	pwHash, err := HashPassword([]byte("123456"))
	ensure.Nil(t, err)
	err = New("a@a.example.com", pwHash)
	ensure.Nil(t, err)
	err = New("b@a.example.com", pwHash)
	ensure.Nil(t, err)
	err = New("b@c.example.com", pwHash)
	ensure.Nil(t, err)
	err = New("f@d.example.com", pwHash)
	ensure.Nil(t, err)

	dm, err := AllDomains()
	ensure.Nil(t, err)
	ensure.DisorderedSubset(t, dm, []string{"a.example.com", "c.example.com", "d.example.com"})

	testTeardownDB(t)
}
