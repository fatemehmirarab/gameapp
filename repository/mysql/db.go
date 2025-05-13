package mySQL

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	UserName string
	PassWord string
	Port     int
	Host     string
	DBName   string
}

type MySQLDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySQLDB {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", config.UserName,
		config.PassWord, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(fmt.Errorf("can not open mysql %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MySQLDB{config, db}
}
