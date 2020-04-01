package repository

import "pika/internal/model"

type UserRepository interface {
	FindByID(userID uint64) (*model.User, error)
}

type UserRepo struct {
}

func NewUserRepository() *UserRepo {
	return &UserRepo{}
}
func (repo *UserRepo) FindByID(userID uint64) (*model.User, error) {
	return model.UserModel().FindByID(userID)
}
