package repository

import (
	"go-microservices/internal/models"
	"gorm.io/gorm"
)

type WhiteIPRepository struct {
	db *gorm.DB
}

func NewWhiteIpRepository(db *gorm.DB) *WhiteIPRepository {
	return &WhiteIPRepository{db: db}
}

func (r *WhiteIPRepository) FindWhiteIPByIP(ip string) (*models.WhiteIP, error) {
	var whiteIp models.WhiteIP
	result := r.db.Where(" ip = ?", ip).First(&whiteIp)
	return &whiteIp, result.Error
}
