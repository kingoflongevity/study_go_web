package service

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/repository"
	"log/slog"
)

type UserService struct {
	repo   repository.UserRepository
	logger *slog.Logger
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo:   repo,
		logger: slog.With("user_service"),
	}
}

func (s *UserService) GetUserProfile(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, u *model.User) error {
	return s.repo.Create(ctx, u)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, u *model.User) error {
	s.logger.Info("user updated", "user", u)
	return s.repo.Update(ctx, u)
}
