package types

import "time"

type Upload struct {
	ID        uint32     `json:"id"`
	Path      string     `json:"path"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UploadModeler interface {
	Create(string) (*Upload, error)
}
