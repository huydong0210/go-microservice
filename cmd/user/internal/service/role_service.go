package service

import (
	"go-microservices/cmd/user/internal/repository"
	"go-microservices/cmd/user/pkg/model"
)

type RoleServiceInterface interface {
	FindAllRolesByUserId(userId uint) ([]model.Role, error)
	FindRoleByName(roleName string) (model.Role, error)
}
type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repository *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repository}
}
func (s RoleService) FindAllRolesByUserId(userId uint) ([]model.Role, error) {
	return s.repo.FindAllRolesByUserId(userId)
}
func (s RoleService) FindRoleByName(roleName string) (model.Role, error) {
	return s.repo.FindRoleByRoleName(roleName)
}
