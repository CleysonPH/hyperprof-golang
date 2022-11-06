package rest

import (
	"time"

	"github.com/cleysonph/hyperprof/internal/model"
)

type messageResponse struct {
	Message string `json:"message"`
}

type alunoRequest struct {
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	DataAula time.Time `json:"data_aula"`
}

func (a *alunoRequest) ToModel() *model.Aluno {
	return &model.Aluno{
		Nome:     a.Nome,
		Email:    a.Email,
		DataAula: a.DataAula,
	}
}

type professorRequest struct {
	Nome                 string  `json:"nome"`
	Email                string  `json:"email"`
	Idade                int32   `json:"idade"`
	Descricao            string  `json:"descricao"`
	ValorHora            float64 `json:"valor_hora"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
}

func (p *professorRequest) ToModel() *model.Professor {
	return &model.Professor{
		Nome:      p.Nome,
		Email:     p.Email,
		Idade:     p.Idade,
		Descricao: p.Descricao,
		ValorHora: p.ValorHora,
		Password:  p.Password,
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type errorResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Cause     string    `json:"cause"`
}

type validationErrorResponse struct {
	errorResponse
	Errors map[string][]string `json:"errors"`
}
