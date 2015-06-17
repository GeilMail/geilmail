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
	err := New("test@example.com", "123456")
	ensure.Nil(t, err)

	ensure.True(t, CheckPassword("test@example.com", []byte("123456")), "oink")
	ensure.False(t, CheckPassword("test@example.com", []byte("")))
	ensure.False(t, CheckPassword("test2@example.com", []byte("123456")))
	testTeardownDB(t)
}

func TestDomainListing(t *testing.T) {
	testBuildDB(t)
	var err error
	err = New("a@a.example.com", "123456")
	ensure.Nil(t, err)
	err = New("b@a.example.com", "123456")
	ensure.Nil(t, err)
	err = New("b@c.example.com", "123456")
	ensure.Nil(t, err)
	err = New("f@d.example.com", "123456")
	ensure.Nil(t, err)

	dm, err := AllDomains()
	ensure.Nil(t, err)
	ensure.DisorderedSubset(t, dm, []string{"a.example.com", "c.example.com", "d.example.com"})

	testTeardownDB(t)
}
