package handlers

import (
	"akshidas/e-com/pkg/types"
	"io"
	"log"
	"net/http"
	"os"
)

type UploadApi struct {
	model types.UploadModeler
}

func (a *UploadApi) Upload(w http.ResponseWriter, r *http.Request) error {
	tempFile, err := os.CreateTemp("storage", "temp-*.jpeg")
	if err != nil {
		log.Printf("failed to create temp file due to %s", err)
		return serverError(w)
	}

	file, _, _ := r.FormFile("file")
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("failed to decode the file due to %s", err)
		return serverError(w)
	}

	tempFile.Write(fileBytes)
	fileName := tempFile.Name()
	uploaded, err := a.model.Create(fileName)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, uploaded)
}

func NewUploadApi(store types.UploadModeler) *UploadApi {
	return &UploadApi{
		model: store,
	}
}
