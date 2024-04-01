package main

import (
	"log"
	"net/http"
	"transactions/controllers"
	"transactions/database"
	"transactions/repository"
	"transactions/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Establecer la conexión a la base de datos
	err := database.SetupDB()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Crear un cliente MongoDB
	client := database.GetClient()

	// Crear un repositorio de transacciones
	transactionRepo := repository.NewTransactionRepository(client, "tach2", "transactions")

	// Crear un controlador de transacciones
	transactionController := controllers.NewTransactionController(transactionRepo)

	// Crear un enrutador
	r := mux.NewRouter()

	// Configurar las rutas
	routes.ConfigureRoutes(r, transactionController)

	// Iniciar el servidor HTTP
	log.Fatal(http.ListenAndServe(":8081", r))
}
