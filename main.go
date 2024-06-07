package main

import (
	"flag"
	"fmt"
	"log"
	"migration/config"
	"migration/handlers"
	"migration/migrations"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrate {
		fmt.Println("Running migrations...")
		migrations.Migrate()
		fmt.Println("Migrations completed!")
		return
	}

	config.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
