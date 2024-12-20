package storage

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"log"
)

type UploadModal struct {
	store *sql.DB
}

func (m *UploadModal) Create(path string) (*types.Upload, error) {
	query := "INSERT INTO uploads(path) VALUES($1) RETURNING *"
	row := m.store.QueryRow(query, path)
	var uploaded types.Upload

	err := row.Scan(&uploaded.ID, &uploaded.Path, &uploaded.CreatedAt, &uploaded.UpdatedAt, &uploaded.DeletedAt)
	if err != nil {
		log.Printf("failed to save image to database due to %s", err)
		return nil, utils.ServerError
	}
	return &uploaded, nil
}

func NewUploadStorage(store *sql.DB) *UploadModal {
	return &UploadModal{store: store}

}
