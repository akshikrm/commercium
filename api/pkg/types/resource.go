package types

import "time"

type Resource struct {
	ID          uint32     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CreateResourceRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResourceRepository interface {
	GetAll() ([]*Resource, error)
	GetOne(int) (*Resource, error)
	Create(*CreateResourceRequest) error
	Update(int, *CreateResourceRequest) (*Resource, error)
	Delete(int) error
}
