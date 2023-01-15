package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	controllers "github.com/rizperdana/prototype-golang-rest/controllers"
	models "github.com/rizperdana/prototype-golang-rest/models"
)

const (
	// TODO: fetch this from env
	DB_USER     = "admin"
	DB_PASSWORD = "123456"
	DB_NAME     = "golang_db"
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
	log.Printf("Connected to database")

	log.Printf("Initiate table creation if any")
	models.CreateTodoTable(db)

	log.Printf("Initiate controller to connect db")
	controllers.InitiateDB(db)
	return db
}
