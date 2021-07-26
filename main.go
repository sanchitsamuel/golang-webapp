package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sanchitsamuel/webapp/database"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var router *mux.Router

func main() {
	godotenv.Load()

	db, err = database.Connect()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected to database")
	}

	database.Migrate(db)

	AppRouter()
	http.ListenAndServe(":8080", router)
}
