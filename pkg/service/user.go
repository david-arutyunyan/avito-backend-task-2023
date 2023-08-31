package service

import (
	avito "avito-backend-task-2023"
	"avito-backend-task-2023/pkg/repository"
	_ "errors"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user avito.User) (string, error) {
	return s.repo.CreateUser(user)

}

func (s *UserService) DeleteUser(userId string) error {
	return s.repo.DeleteUser(userId)

}
