package service

import (
	"go-microservices/cmd/user/internal/repository"
	"go-microservices/cmd/user/pkg/model"
	request2 "go-microservices/internal/api/request"
	helper2 "go-microservices/internal/helper"
)

type UserServiceInterface interface {
	FindUserByUserName(username string) (*model.User, error)
	CreateUser(request *request2.UserCreationRequest) error
	FindAllUsers() ([]model.User, error)
}

type UserService struct {
	repo        *repository.UserRepository
	roleService RoleServiceInterface
}

func NewUserService(repo *repository.UserRepository, roleService RoleServiceInterface) *UserService {
	return &UserService{repo: repo, roleService: roleService}
}

func (s *UserService) FindUserByUserName(username string) (*model.User, error) {
	return s.repo.FindUserByUsername(username)
}
func (s *UserService) CreateUser(request *request2.UserCreationRequest) error {

	hashPassword, err := helper2.HashPassword(request.Password)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: request.Username,
		Password: hashPassword,
		Email:    request.Email,
	}

	role, err := s.roleService.FindRoleByName("USER")
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	err = s.repo.InsertUserRole(user.ID, role.ID)
	return err
}
func (s *UserService) FindAllUsers() ([]model.User, error) {
	return s.repo.FindAllUser()
}
