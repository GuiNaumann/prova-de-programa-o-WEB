package models

import "gorm.io/gorm"

// entidade de aluno
type Aluno struct {
	gorm.Model
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}
