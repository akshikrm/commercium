package handlers

import (
	"akshidas/e-com/pkg/utils"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type ApiResponse struct {
	Data any `json:"data"`
}

func accessDenied(w http.ResponseWriter) error {
	return writeError(w, http.StatusUnauthorized, errors.New("access denied"))
}

func invalidId(w http.ResponseWriter) error {
	return writeError(w, http.StatusUnprocessableEntity, errors.New("invalid id"))
}

func conflict(w http.ResponseWriter) error {
	return writeError(w, http.StatusConflict, errors.New("conflict"))
}

func invalidRequest(w http.ResponseWriter) error {
	return writeError(w, http.StatusUnprocessableEntity, errors.New("invalid request"))
}

func notFound(w http.ResponseWriter) error {
	return writeError(w, http.StatusNotFound, errors.New("not found"))
}

func serverError(w http.ResponseWriter) error {
	return writeError(w, http.StatusInternalServerError, errors.New("something went wrong"))
}

func Cors(w http.ResponseWriter) error {
	return writeJson(w, http.StatusNoContent, errors.New("no content"))
}

func writeJson(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&ApiResponse{Data: value})
}

func writeError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&ApiError{Error: err.Error()})
}

func parseId(id string) (int, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to convert param %s to integer", id)
		return 0, utils.InvalidParam
	}
	return parsedId, nil
}

func DecodeBody(body io.ReadCloser, a any) error {
	if err := json.NewDecoder(body).Decode(a); err != nil {
		log.Printf("failed to decode request due to %s", err)
		if err == io.EOF {
			return utils.InvalidRequest
		}
		return err
	}
	return nil
}
