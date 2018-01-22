package firewood

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Pg *sql.DB

func InitPg(addr, user, password, database string) error {
	if Pg != nil {
		return nil
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, addr, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Pg = db

	return nil
}
