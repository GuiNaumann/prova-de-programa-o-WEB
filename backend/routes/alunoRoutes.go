package routes

import (
	"provaSuficiencia/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Aqui são as rotas get da lisgem de alunos e post para cadastro
	r.GET("/aluno", controllers.GetAlunos)
	r.POST("/aluno", controllers.CreateAluno)

	return r
}
