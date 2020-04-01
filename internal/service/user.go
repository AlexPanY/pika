package service

import (
	"fmt"
	"pika/internal/model"
	"pika/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}

//GetUserInfoByID - get
func (s *UserService) GetUserInfoByID(userID uint64) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	if user == nil {
		return nil, fmt.Errorf("not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
