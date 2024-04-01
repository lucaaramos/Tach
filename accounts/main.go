package main

import (
	"Tach/controllers"
	"Tach/database"
	"Tach/repository"
	"Tach/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := database.SetupDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	client := database.GetClient()
	accountRepo := repository.NewAccountRepository(client, "tach", "accounts")

	accountController := controllers.NewAccountController(accountRepo)
	r := mux.NewRouter()

	routes.ConfigureRoutes(r, accountController)
	log.Fatal(http.ListenAndServe(":8080", r))
}
