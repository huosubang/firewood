package firewood

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitPg(addr, user, password, database string) error {
	if Db != nil {
		return nil
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, addr, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Db = db

	return nil
}

func InitMysql(addr, user, password, database string) error {
	if Db != nil {
		return nil
	}
	connStr := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, database)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	Db = db

	return nil
}
