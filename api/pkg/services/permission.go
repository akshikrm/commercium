package services

import (
	"akshidas/e-com/pkg/types"
)

type PermissionService struct {
	repository types.PermissionRepository
}

func (r *PermissionService) GetAll() ([]*types.Permission, error) {
	return r.repository.GetAll()
}

func (r *PermissionService) GetOne(id int) (*types.Permission, error) {
	return r.repository.GetOne(id)
}

func (r *PermissionService) Create(newPermission *types.CreateNewPermission) error {
	return r.repository.Create(newPermission)
}

func (r *PermissionService) Update(id int, newPermission *types.CreateNewPermission) (*types.Permission, error) {
	return r.repository.Update(id, newPermission)
}

func (r *PermissionService) Delete(id int) error {
	return r.repository.Delete(id)
}

func NewPermissionService(repository types.PermissionRepository) *PermissionService {
	return &PermissionService{
		repository: repository,
	}
}
