package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

const (
	// TODO: fetch this from env
	DB_USER     = "admin"
	DB_PASSWORD = "123456"
	DB_NAME     = "golang-prototype-db"
	DB_HOST     = "localhost:5432"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     DB_USER,
		Password: DB_PASSWORD,
		Addr:     DB_HOST,
		Database: DB_NAME,
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect database")
		os.Exit(100)
	}

	log.Print("Connected to database")

	return db
}
