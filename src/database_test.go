package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:33769)/go_mysql")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
