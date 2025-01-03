package services

import (
	"akshidas/e-com/pkg/types"
)

type role struct {
	repository types.RoleRepository
}

func (r *role) GetAll() ([]*types.Role, error) {
	return r.repository.GetAll()
}

func (r *role) GetOne(id int) (*types.Role, error) {
	return r.repository.GetOne(id)
}

func (r *role) Create(newRole *types.CreateRoleRequest) error {
	return r.repository.Create(newRole)
}

func (r *role) Update(id int, newRole *types.CreateRoleRequest) (*types.Role, error) {
	return r.repository.Update(id, newRole)
}

func (r *role) Delete(id int) error {
	return r.repository.Delete(id)
}

func newRoleService(repository types.RoleRepository) *role {
	return &role{
		repository: repository,
	}
}
