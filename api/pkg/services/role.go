package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
)

type role struct {
	repository types.RoleRepository
}

func (r *role) GetAll() ([]*types.Role, error) {
	roles, ok := r.repository.GetAll()
	if !ok {
		return nil, utils.ServerError
	}
	return roles, nil
}

func (r *role) GetOne(id int) (*types.Role, error) {
	role, ok := r.repository.GetOne(id)
	if !ok {
		return nil, utils.ServerError
	}
	if role == nil {
		return nil, utils.NotFound
	}
	return role, nil
}

func (r *role) Create(newRole *types.CreateRoleRequest) error {
	ok := r.repository.Create(newRole)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func (r *role) Update(id int, newRole *types.CreateRoleRequest) (*types.Role, error) {
	updatedRole, ok := r.repository.Update(id, newRole)
	if !ok {
		return nil, utils.ServerError
	}
	if updatedRole == nil {
		return nil, utils.NotFound
	}
	return updatedRole, nil
}

func (r *role) Delete(id int) error {
	ok := r.repository.Delete(id)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newRoleService(repository types.RoleRepository) *role {
	return &role{
		repository: repository,
	}
}
