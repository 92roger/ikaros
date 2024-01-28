package main

import (
	"ikaros/backend/api/routes"
	"ikaros/backend/pkg/database"
	"log"
	"net/http"
)

func main() {
	// Inicialize a conexão com o MongoDB
	database.ConnectMongoDB()
	// Crie os índices necessários
	database.CreateIndices()
	router := routes.SetupRouter() // Configura as rotas
	log.Fatal(http.ListenAndServe(":8080", router))
}
