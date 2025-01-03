package services

import (
	"akshidas/e-com/pkg/types"
)

type ResourceService struct {
	repository types.ResourceRepository
}

func (r *ResourceService) GetAll() ([]*types.Resource, error) {
	return r.repository.GetAll()
}

func (r *ResourceService) GetOne(id int) (*types.Resource, error) {
	return r.repository.GetOne(id)
}

func (r *ResourceService) Create(newResource *types.CreateResourceRequest) error {
	return r.repository.Create(newResource)
}

func (r *ResourceService) Update(id int, newResource *types.CreateResourceRequest) (*types.Resource, error) {
	return r.repository.Update(id, newResource)
}

func (r *ResourceService) Delete(id int) error {
	return r.repository.Delete(id)
}

func NewResourceService(repository types.ResourceRepository) *ResourceService {
	return &ResourceService{
		repository: repository,
	}
}
