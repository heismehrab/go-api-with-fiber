package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToMysql() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

func DisconnectFromMysql() {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}
}
