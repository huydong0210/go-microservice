package database

import (
	model2 "go-microservices/cmd/to_do/pkg/model"
	"go-microservices/cmd/user/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize(databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model2.TodoItem{},
	)
}
