package service

import (
	"echo-practice/internal/model"
	"echo-practice/internal/repository"

	"github.com/labstack/gommon/log"
)

type UserService struct {
	repo repository.UserRepository
	slog *log.Logger
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(u *model.User) error {
	return s.repo.Create(u)
}

func (s *UserService) DeleteUser(id string) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(u *model.User) error {
	s.slog.Info(u)
	return s.repo.Update(u)
}
