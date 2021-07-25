package main

import (
	"github.com/gorilla/mux"
	"github.com/sanchitsamuel/webapp/controllers"
)

func AppRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/ping", controllers.Pong)
}
