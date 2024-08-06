package controllers

import (
	"net/http"
	"provaSuficiencia/config"
	"provaSuficiencia/models"

	"github.com/gin-gonic/gin"
)

//Controller de cada função;

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	config.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func CreateAluno(c *gin.Context) {
	var input models.Aluno
	// Faz o bind do json para o input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Cria um novo aluno no banco
	aluno := models.Aluno{Nome: input.Nome, Telefone: input.Telefone}
	config.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
