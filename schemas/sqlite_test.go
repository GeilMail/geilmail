package schemas

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/facebookgo/ensure"
	_ "github.com/mattn/go-sqlite3"
)

func TestSQLite(t *testing.T) {
	filepath := path.Join(os.TempDir(), "sqlite_test.db")
	db, err := sql.Open("sqlite3", filepath)
	ensure.Nil(t, err)

	cwd, err := os.Getwd()
	ensure.Nil(t, err)

	schemaPath := path.Join(cwd, "sqlite.sql")
	buf, err := ioutil.ReadFile(schemaPath)
	ensure.Nil(t, err)

	_, err = db.Exec(fmt.Sprintf("%s", buf))
	ensure.Nil(t, err)

	os.Remove(filepath)
}
