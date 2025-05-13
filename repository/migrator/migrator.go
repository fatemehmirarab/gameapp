package migrator

import (
	"database/sql"
	"fmt"

	mySQL "github.com/fatemehmirarab/gameapp/repository/mySQL/db"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	Migrations *migrate.FileMigrationSource
	Config     mySQL.Config
}

func New(config mySQL.Config) Migrator {
	//read migrations from a folder
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mySQL/migrations",
	}
	return Migrator{Migrations: migrations, Config: config}
}

func (m Migrator) Up() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", m.Config.UserName,
		m.Config.PassWord, m.Config.Host, m.Config.Port, m.Config.DBName))
	if err != nil {
		panic(fmt.Errorf("can not open mysql %v", err))
	}

	n, err := migrate.Exec(db, "postgres", m.Migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can not execute migrations %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", m.Config.UserName,
		m.Config.PassWord, m.Config.Host, m.Config.Port, m.Config.DBName))
	if err != nil {
		panic(fmt.Errorf("can not open mysql %v", err))
	}

	n, err := migrate.Exec(db, "postgres", m.Migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can not execute migrations %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
func (m Migrator) Status() {

}
