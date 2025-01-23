package types

import "time"

type CreateRoleRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	ID          uint32     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type RoleRepository interface {
	GetAll() ([]*Role, bool)
	GetOne(int) (*Role, bool)
	Create(*CreateRoleRequest) bool
	Update(int, *CreateRoleRequest) (*Role, bool)
	Delete(int) bool
}

type RoleService interface {
	GetAll() ([]*Role, error)
	GetOne(int) (*Role, error)
	Create(*CreateRoleRequest) error
	Update(int, *CreateRoleRequest) (*Role, error)
	Delete(int) error
}
