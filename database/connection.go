package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")
	dbPass := os.Getenv("DBPASS")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
