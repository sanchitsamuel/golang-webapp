package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sanchitsamuel/webapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var router *mux.Router

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	godotenv.Load()

	host := os.Getenv("HOST")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")
	dbPass := os.Getenv("DBPASS")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbUser, dbPass, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to database")
	}

	db.AutoMigrate(&models.User{})

	AppRouter()
	http.ListenAndServe(":8080", router)
}
