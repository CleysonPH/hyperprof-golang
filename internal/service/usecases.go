package service

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/cleysonph/hyperprof/internal/database"
	"github.com/cleysonph/hyperprof/internal/model"
	"github.com/cleysonph/hyperprof/internal/validator"
)

func FindAllProfessores(q string) ([]*model.Professor, error) {
	professores, err := database.FindAllProfessores(q)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}
	return professores, nil
}

func FindProfessorByID(professorID int64) (*model.Professor, error) {
	professor, err := database.FindProfessorByID(professorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.ProfessorNotFoundError{
				Message: fmt.Sprintf("Professor with ID %d not found", professorID),
			}
		}
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}
	return professor, nil
}

func GetProfessorByToken(token string) (*model.Professor, error) {
	email, err := getSubFromAccessToken(token)
	if err != nil {
		return nil, &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	professor, err := database.FindProfessorByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.JwtTokenError{
				Message: fmt.Sprintf("Professor with email %s not found", email),
			}
		}
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return professor, nil
}

func CreateProfessor(professor *model.Professor, passwordConfirmation string) (*model.Professor, error) {
	if err := validator.ValidateProfessor(professor, passwordConfirmation); err != nil {
		return nil, err
	}

	hash, err := hashPassword(professor.Password)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	professor.Password = hash
	professor, err = database.CreateProfessor(professor)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return professor, nil
}

func UpdateProfessorByToken(token string, professorData *model.Professor, passwordConfirmation string) (*model.Professor, error) {
	professorEmail, err := getSubFromAccessToken(token)
	if err != nil {
		return nil, &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	professor, err := database.FindProfessorByEmail(professorEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.JwtTokenError{
				Message: fmt.Sprintf("Professor with email %s not found", professorEmail),
			}
		}
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	professorData.ID = professor.ID
	if err := validator.ValidateProfessor(professorData, passwordConfirmation); err != nil {
		return nil, err
	}

	hash, err := hashPassword(professorData.Password)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}
	professorData.Password = hash

	professor, err = database.UpdateProfessor(professorData)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return professor, nil
}

func UpdateProfessorFotoByToken(token string, file multipart.File, fileHeader *multipart.FileHeader) error {
	if err := validator.ValidateProfessorFoto(fileHeader); err != nil {
		return err
	}

	professorEmail, err := getSubFromAccessToken(token)
	if err != nil {
		return &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	filepath, err := uploadFile(file, fileHeader)
	if err != nil {
		return &model.ApplicationError{
			Message: err.Error(),
		}
	}

	err = database.UpdateProfessorFotoPerfilByEmail(professorEmail, filepath)
	if err != nil {
		return &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return nil
}

func DeleteProfessorByToken(token string) error {
	professorEmail, err := getSubFromAccessToken(token)
	if err != nil {
		return &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	err = database.DeleteAlunosByProfessorEmail(professorEmail)
	if err != nil {
		return &model.ApplicationError{
			Message: err.Error(),
		}
	}

	err = database.DeleteProfessorByEmail(professorEmail)
	if err != nil {
		return &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return nil
}

func CreateAluno(aluno *model.Aluno) (*model.Aluno, error) {
	if !database.ExistsProfessorByID(aluno.ProfessorID) {
		return nil, &model.ProfessorNotFoundError{
			Message: fmt.Sprintf("Professor with ID %d not found", aluno.ProfessorID),
		}
	}

	if err := validator.ValidateAluno(aluno); err != nil {
		return nil, err
	}

	aluno, err := database.CreateAluno(aluno)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return aluno, nil
}

func GetAlunosByProfessorToken(token string) ([]*model.Aluno, error) {
	professorEmail, err := getSubFromAccessToken(token)
	if err != nil {
		return nil, &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	alunos, err := database.FindAlunosByProfessorEmail(professorEmail)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return alunos, nil
}

func Login(email, password string) ([]string, error) {
	if err := validator.ValidateLogin(email, password); err != nil {
		return nil, err
	}

	professor, err := database.FindProfessorByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.BadCredentialsError{
				Message: "Invalid credentials",
			}
		}
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	match := checkPasswordHash(password, professor.Password)
	if !match {
		return nil, &model.BadCredentialsError{
			Message: "Invalid credentials",
		}
	}

	accessToken, err := generateAccessToken(professor.Email)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}
	refreshtoken, err := generateRefreshToken(professor.Email)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return []string{accessToken, refreshtoken}, nil
}

func Refresh(refreshToken string) ([]string, error) {
	if err := validator.ValidateRefresh(refreshToken); err != nil {
		return nil, err
	}

	email, err := getSubFromRefreshToken(refreshToken)
	if err != nil {
		return nil, &model.JwtTokenError{
			Message: err.Error(),
		}
	}
	_, err = database.FindProfessorByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &model.JwtTokenError{
				Message: "Invalid refresh token",
			}
		}
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	invalidateTokens(refreshToken)

	accessToken, err := generateAccessToken(email)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}
	newRefreshToken, err := generateRefreshToken(email)
	if err != nil {
		return nil, &model.ApplicationError{
			Message: err.Error(),
		}
	}

	return []string{accessToken, newRefreshToken}, nil
}

func Logout(token, refreshToken string) error {
	if err := validator.ValidateRefresh(refreshToken); err != nil {
		return err
	}

	_, err := getSubFromRefreshToken(refreshToken)
	if err != nil {
		return &model.JwtTokenError{
			Message: err.Error(),
		}
	}

	invalidateTokens(token, refreshToken)
	return nil
}
