package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Professor struct {
	ID         int64      `json:"id"`
	Nome       string     `json:"nome"`
	Email      string     `json:"email"`
	Idade      int32      `json:"idade"`
	Descricao  string     `json:"descricao"`
	ValorHora  float64    `json:"valor_hora"`
	FotoPerfil NullString `json:"foto_perfil"`
	Password   string     `json:"-"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Aluno struct {
	ID          int64     `json:"id"`
	ProfessorID int64     `json:"-"`
	Nome        string    `json:"nome"`
	Email       string    `json:"email"`
	DataAula    time.Time `json:"data_aula"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}
