package service

import "go-microservices/internal/repository"

type WhiteIpServiceInterface interface {
	IsWhiteIp(ip string) bool
}

type WhiteIpService struct {
	repo *repository.WhiteIPRepository
}

func NewWhiteService(repo *repository.WhiteIPRepository) *WhiteIpService {
	return &WhiteIpService{}
}
func (s *WhiteIpService) IsWhiteIp(ip string) bool {
	_, err := s.repo.FindWhiteIPByIP(ip)
	return err == nil
}
