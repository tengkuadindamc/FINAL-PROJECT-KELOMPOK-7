package config

import (
	"finalproject4/database/seeder"
	"finalproject4/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func LoadConfig() model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverPort := os.Getenv("SERVER_PORT")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	config := model.Config{
		ServerPort: serverPort,
		Database: model.Database{
			Host:     databaseHost,
			Port:     databasePort,
			Username: databaseUser,
			Name:     databaseName,
			Password: databasePassword,
		},
	}

	return config
}

func ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{}, &model.TransactionHistory{})
	seeder.DBSeed(db)
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
