package rest

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cleysonph/hyperprof/internal/model"
	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		writeError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func writeError(w http.ResponseWriter, err error) {
	switch t := err.(type) {
	case *model.ApplicationError:
		e := createJsonError(w, http.StatusInternalServerError, t)
		writeJSON(w, http.StatusInternalServerError, e)
	case *model.ProfessorNotFoundError:
		e := createJsonError(w, http.StatusNotFound, t)
		writeJSON(w, http.StatusNotFound, e)
	case *model.ConversionError:
		e := createJsonError(w, http.StatusBadRequest, t)
		writeJSON(w, http.StatusBadRequest, e)
	case *model.JsonError:
		e := createJsonError(w, http.StatusBadRequest, t)
		writeJSON(w, http.StatusBadRequest, e)
	case *model.ValidationError:
		e := createJsonValidationError(w, http.StatusBadRequest, t)
		writeJSON(w, http.StatusBadRequest, e)
	case *model.BadCredentialsError:
		e := createJsonError(w, http.StatusUnauthorized, t)
		writeJSON(w, http.StatusUnauthorized, e)
	case *model.JwtTokenError:
		e := createJsonError(w, http.StatusUnauthorized, t)
		writeJSON(w, http.StatusUnauthorized, e)
	default:
		e := createJsonError(w, http.StatusInternalServerError, t)
		writeJSON(w, http.StatusInternalServerError, e)
	}
}

func createJsonError(w http.ResponseWriter, status int, err error) errorResponse {
	return errorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().UTC(),
		Status:    status,
		Error:     http.StatusText(status),
		Cause:     reflect.TypeOf(err).String(),
	}
}

func createJsonValidationError(w http.ResponseWriter, status int, err *model.ValidationError) validationErrorResponse {
	return validationErrorResponse{
		errorResponse: createJsonError(w, status, err),
		Errors:        err.Errors,
	}
}

func getStringQueryParam(w http.ResponseWriter, r *http.Request, key string) string {
	return strings.Trim(r.URL.Query().Get(key), " ")
}

func getInt64UrlParam(w http.ResponseWriter, r *http.Request, key string) (int64, error) {
	params := mux.Vars(r)
	param, ok := params[key]
	if !ok {
		return 0, &model.ConversionError{Message: "Param not found"}
	}
	i, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, &model.ConversionError{
			Message: key + " must be an integer",
		}
	}
	return i, nil
}

func readJSON(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return &model.JsonError{
			Message: err.Error(),
		}
	}
	return nil
}

func getTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")
	return strings.TrimPrefix(token, "Bearer ")
}
