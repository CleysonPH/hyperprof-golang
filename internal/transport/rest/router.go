package rest

import (
	"net/http"

	"github.com/cleysonph/hyperprof/config"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/professores", getProfessores).Methods(http.MethodGet)
	router.HandleFunc("/api/professores", postProfessor).Methods(http.MethodPost)
	router.HandleFunc("/api/professores", putProfessor).Methods(http.MethodPut)
	router.HandleFunc("/api/professores", deleteProfessor).Methods(http.MethodDelete)
	router.HandleFunc("/api/professores/foto", postProfessorFoto).Methods(http.MethodPost)
	router.HandleFunc("/api/professores/alunos", getProfessorAlunos).Methods(http.MethodGet)
	router.HandleFunc("/api/professores/{professorID}", getProfessorByID).Methods(http.MethodGet)
	router.HandleFunc("/api/professores/{professorID}/alunos", postAluno).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/login", postLogin).Methods(http.MethodPost)
	router.HandleFunc("/api/me", getMe).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/refresh", postRefresh).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/logout", postLogout).Methods(http.MethodPost)

	// Serve static files in development mode
	if config.IsDev() {
		router.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads"))))
	}

	return router
}
