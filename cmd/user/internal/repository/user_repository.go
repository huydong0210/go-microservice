package repository

import (
	"go-microservices/cmd/user/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (repo *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := repo.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
func (repo *UserRepository) FindUserById(id int64) (*model.User, error) {
	var user model.User
	result := repo.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}
func (repo *UserRepository) CreateUser(user *model.User) error {
	return repo.db.Create(user).Error
}
func (repo *UserRepository) FindAllUser() (users []model.User, err error) {
	err = repo.db.Find(&users).Error
	return
}
func (repo *UserRepository) InsertUserRole(userId uint, roleId uint) error {
	return repo.db.Exec("insert into user_role(user_id, role_id) value ( ? , ?)", userId, roleId).Error
}
