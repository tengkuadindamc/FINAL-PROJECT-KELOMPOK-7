package database

import (
	"fmt"
	"fp-2/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "santoso"
	dbPort   = 5433
	dbname   = "fp-2"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	if err := db.Debug().AutoMigrate(entities.User{}, entities.Photo{}, entities.Comment{}, entities.SocialMedia{}); err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	log.Println("Connected to DB")

}

func GetPostgresInstance() *gorm.DB {
	return db
}
