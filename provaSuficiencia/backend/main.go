package main

import (
	"provaSuficiencia/config"
	"provaSuficiencia/models"
	"provaSuficiencia/routes"
)

func main() {
	// primeiro de tudo faz a conexão com o banco
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Aluno{})

	// depois inicia o servido
	r := routes.SetupRouter()
	//na porta 8080
	r.Run(":8080")
}
