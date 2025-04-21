package mySQL

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {

	db, err := sql.Open("mysql", "fatemeh:Mirarab@(localhost:3306)/mydb")
	if err != nil {
		panic(fmt.Errorf("can not open mysql %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MySQLDB{db}
}
