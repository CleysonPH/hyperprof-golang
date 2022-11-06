package validator

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/cleysonph/hyperprof/internal/database"
	"github.com/cleysonph/hyperprof/internal/model"
)

func ValidateAluno(aluno *model.Aluno) error {
	validationErr := &model.ValidationError{}

	validationErr.AddErrorIf(aluno.Nome == "", "nome", "is required")
	validationErr.AddErrorIf(len(aluno.Nome) < 3, "nome", "must be at least 3 characters")
	validationErr.AddErrorIf(len(aluno.Nome) > 100, "nome", "must be at most 100 characters")
	validationErr.AddErrorIf(aluno.Email == "", "email", "is required")
	validationErr.AddErrorIf(len(aluno.Email) < 3, "email", "must be at least 3 characters")
	validationErr.AddErrorIf(len(aluno.Email) > 255, "email", "must be at most 255 characters")
	validationErr.AddErrorIf(aluno.DataAula.IsZero(), "data_aula", "is required")
	validationErr.AddErrorIf(aluno.DataAula.Before(time.Now()), "data_aula", "must be in the future")

	if validationErr.HasErrors() {
		return validationErr
	}
	return nil
}

func ValidateProfessor(professor *model.Professor, passwordConfirmation string) error {
	validationErr := &model.ValidationError{}

	validationErr.AddErrorIf(professor.Nome == "", "nome", "is required")
	validationErr.AddErrorIf(len(professor.Nome) < 3, "nome", "must be at least 3 characters")
	validationErr.AddErrorIf(len(professor.Nome) > 100, "nome", "must be at most 100 characters")
	validationErr.AddErrorIf(professor.Email == "", "email", "is required")
	validationErr.AddErrorIf(len(professor.Email) < 3, "email", "must be at least 3 characters")
	validationErr.AddErrorIf(len(professor.Email) > 255, "email", "must be at most 255 characters")
	validationErr.AddErrorIf(professor.Idade < 18, "idade", "must be at least 18 years old")
	validationErr.AddErrorIf(professor.Idade > 100, "idade", "must be at most 100 years old")
	validationErr.AddErrorIf(professor.Descricao == "", "descricao", "is required")
	validationErr.AddErrorIf(len(professor.Descricao) < 10, "descricao", "must be at least 100 characters")
	validationErr.AddErrorIf(len(professor.Descricao) > 500, "descricao", "must be at most 500 characters")
	validationErr.AddErrorIf(professor.ValorHora < 10, "valor_hora", "must be at least 10")
	validationErr.AddErrorIf(professor.ValorHora > 500, "valor_hora", "must be at most 500")
	validationErr.AddErrorIf(professor.Password == "", "password", "is required")
	validationErr.AddErrorIf(len(professor.Password) < 6, "password", "must be at least 6 characters")
	validationErr.AddErrorIf(passwordConfirmation == "", "password_confirmation", "is required")
	validationErr.AddErrorIf(professor.Password != passwordConfirmation, "password_confirmation", "must match password")
	validationErr.AddErrorIf(professor.ID <= 0 && database.ExistsProfessorByEmail(professor.Email), "email", "already exists")
	validationErr.AddErrorIf(professor.ID > 0 && database.ExistsProfessorByEmailAndNotID(professor.Email, professor.ID), "email", "already exists")

	if validationErr.HasErrors() {
		return validationErr
	}
	return nil
}

func ValidateLogin(email, password string) error {
	validationErr := &model.ValidationError{}

	validationErr.AddErrorIf(email == "", "email", "is required")
	validationErr.AddErrorIf(len(email) < 3, "email", "must be at least 3 characters")
	validationErr.AddErrorIf(len(email) > 255, "email", "must be at most 255 characters")
	validationErr.AddErrorIf(password == "", "password", "is required")
	validationErr.AddErrorIf(len(password) < 6, "password", "must be at least 6 characters")

	if validationErr.HasErrors() {
		return validationErr
	}
	return nil
}

func ValidateRefresh(refreshToken string) error {
	validationErr := &model.ValidationError{}

	validationErr.AddErrorIf(refreshToken == "", "refresh_token", "is required")

	if validationErr.HasErrors() {
		return validationErr
	}
	return nil
}

func ValidateProfessorFoto(file *multipart.FileHeader) error {
	validationErr := &model.ValidationError{}

	validationErr.AddErrorIf(file == nil, "foto", "is required")
	validationErr.AddErrorIf(file.Size > 2*1024*1024, "foto", "must be at most 2MB")
	validationErr.AddErrorIf(!strings.HasPrefix(file.Header.Get("Content-Type"), "image/"), "foto", "must be an image")

	if validationErr.HasErrors() {
		return validationErr
	}
	return nil
}
