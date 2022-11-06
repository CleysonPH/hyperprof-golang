package rest

import (
	"net/http"

	"github.com/cleysonph/hyperprof/internal/service"
)

func getProfessores(w http.ResponseWriter, r *http.Request) {
	professores, err := service.FindAllProfessores(getStringQueryParam(w, r, "q"))
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, professores)
}

func getProfessorByID(w http.ResponseWriter, r *http.Request) {
	professorID, err := getInt64UrlParam(w, r, "professorID")
	if err != nil {
		writeError(w, err)
		return
	}

	professor, err := service.FindProfessorByID(professorID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, professor)
}

func postProfessor(w http.ResponseWriter, r *http.Request) {
	professorRequest := &professorRequest{}
	if err := readJSON(r, &professorRequest); err != nil {
		writeError(w, err)
		return
	}

	professor, err := service.CreateProfessor(professorRequest.ToModel(), professorRequest.PasswordConfirmation)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, professor)
}

func deleteProfessor(w http.ResponseWriter, r *http.Request) {
	err := service.DeleteProfessorByToken(getTokenFromHeader(r))
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func putProfessor(w http.ResponseWriter, r *http.Request) {
	professorRequest := &professorRequest{}
	if err := readJSON(r, &professorRequest); err != nil {
		writeError(w, err)
		return
	}

	professor, err := service.UpdateProfessorByToken(
		getTokenFromHeader(r),
		professorRequest.ToModel(),
		professorRequest.PasswordConfirmation,
	)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, professor)
}

func postProfessorFoto(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("foto")
	if err != nil {
		writeError(w, err)
		return
	}
	defer file.Close()

	err = service.UpdateProfessorFotoByToken(getTokenFromHeader(r), file, handler)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, messageResponse{Message: "Foto atualizada com sucesso"})
}

func postAluno(w http.ResponseWriter, r *http.Request) {
	professorID, err := getInt64UrlParam(w, r, "professorID")
	if err != nil {
		writeError(w, err)
		return
	}

	alunoRequest := &alunoRequest{}
	if err := readJSON(r, &alunoRequest); err != nil {
		writeError(w, err)
		return
	}

	alunoModel := alunoRequest.ToModel()
	alunoModel.ProfessorID = professorID
	aluno, err := service.CreateAluno(alunoModel)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, aluno)
}

func getProfessorAlunos(w http.ResponseWriter, r *http.Request) {
	alunos, err := service.GetAlunosByProfessorToken(getTokenFromHeader(r))
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, alunos)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	loginRequest := &loginRequest{}
	if err := readJSON(r, &loginRequest); err != nil {
		writeError(w, err)
		return
	}

	tokens, err := service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		writeError(w, err)
		return
	}

	response := &loginResponse{
		Token:        tokens[0],
		RefreshToken: tokens[1],
	}

	writeJSON(w, http.StatusOK, response)
}

func postRefresh(w http.ResponseWriter, r *http.Request) {
	refreshRequest := &refreshRequest{}
	if err := readJSON(r, &refreshRequest); err != nil {
		writeError(w, err)
		return
	}

	tokens, err := service.Refresh(refreshRequest.RefreshToken)
	if err != nil {
		writeError(w, err)
		return
	}

	response := &loginResponse{
		Token:        tokens[0],
		RefreshToken: tokens[1],
	}

	writeJSON(w, http.StatusOK, response)
}

func getMe(w http.ResponseWriter, r *http.Request) {
	professor, err := service.GetProfessorByToken(getTokenFromHeader(r))
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, professor)
}

func postLogout(w http.ResponseWriter, r *http.Request) {
	refreshToken := &refreshRequest{}
	if err := readJSON(r, &refreshToken); err != nil {
		writeError(w, err)
		return
	}

	if err := service.Logout(getTokenFromHeader(r), refreshToken.RefreshToken); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusResetContent)
}
