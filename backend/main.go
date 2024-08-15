package main

import (
	"fmt"
	"log"
	"net/http"
  "github.com/pivot/database"
  //"github.com/pivot/middleware"
  "github.com/pivot/routes"
	"github.com/pivot/utils"

  "github.com/joho/godotenv"

)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//r := mux.NewRouter()
  //r.Use(middleware.AuthMiddleware)

	routes.SetupRoutes()

	if err := database.Connect(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// DB is a global variable in database folder
	if err := database.RunMigrations(database.DB); err != nil {
		log.Printf("Error running database migrations: %v", err)
		log.Fatalf("Error running database migrations: %v", err)
	}

	utils.InitLogging()

	fmt.Println("Server running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

