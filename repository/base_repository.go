package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type DbConfig struct {
	DbIP     string
	DbPort   int
	DbName   string
	Username string
	Password string
}

func Connect(config DbConfig) {

	if config.DbIP == "" {
		log.Fatalln("Database IP is required!")
	}

	if config.DbPort == 0 {
		log.Fatalln("Database port is required!")
	}

	if config.DbName == "" {
		log.Fatalln("Database name is required!")
	}

	if config.Username == "" {
		log.Fatalln("User name is required!")
	}

	if config.Password == "" {
		log.Fatalln("Password is required!")
	}

	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.DbIP, config.DbPort, config.DbName)

	_db, err := sqlx.Connect("mysql", dataSource)

	if err != nil {
		log.Fatalln(err)
	}

	db = _db
}
