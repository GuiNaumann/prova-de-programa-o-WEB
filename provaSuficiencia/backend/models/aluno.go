package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}
