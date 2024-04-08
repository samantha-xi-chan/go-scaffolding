package service

import (
	"go-scaffolding/internal/app01/model"
	"go-scaffolding/internal/app01/repo"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repo.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepository(db),
	}
}

func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	return s.repo.FindByID(userID)
}
