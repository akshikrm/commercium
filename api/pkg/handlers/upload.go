package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
)

type UploadApi struct{}

func (a *UploadApi) Upload(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	file, _, _ := r.FormFile("file")
	_, err := io.ReadAll(file)
	log.Println("initiating upload")
	if err != nil {
		log.Printf("failed to decode the file due to %s", err)
		return serverError(w)
	}
	log.Println("upload completed")
	return writeJson(w, http.StatusCreated, "uploaded")
}

func newUpload() *UploadApi {
	return &UploadApi{}
}
