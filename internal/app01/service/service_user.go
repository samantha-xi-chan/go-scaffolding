package service

import (
	"go-scaffolding/internal/app01/model"
	"go-scaffolding/internal/app01/repo"
	"go-scaffolding/util/idgen"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepo(db),
	}
}

func (s *UserService) CreateUser(item model.User) (*model.User, error) {
	item.Id = idgen.GetIdWithPref("user")
	item.CreateAt = time.Now().UnixMilli()
	log.Println("item: ", item)

	return s.repo.CreateItem(item)
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}
