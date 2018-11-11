package main

import (
	"log"
	"net/http"

	"./common"
	"./database"
	"./routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	database.InitDb()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
