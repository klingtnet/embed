package main

import (
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	embedMigrate "github.com/klingtnet/embed/migrate"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sourceDriver, err := embedMigrate.WithInstance(Embeds)
	if err != nil {
		log.Fatal("err")
	}
	m, err := migrate.NewWithInstance("embed", sourceDriver, "", dbDriver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO t VALUES("a test value")`)
	if err != nil {
		log.Fatal(err)
	}
	var msg string
	err = db.QueryRow(`SELECT a FROM t LIMIT 1`).Scan(&msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(msg)
}
