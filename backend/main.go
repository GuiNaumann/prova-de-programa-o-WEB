package main

import (
	"log"
	"net/http"
	"provaSuficiencia/config"
	"provaSuficiencia/models"
	"provaSuficiencia/routes"

	"github.com/rs/cors"
)

func main() {
	// Conectar ao banco de dados
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// iniciar o servidor
	r := routes.SetupRouter()

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
