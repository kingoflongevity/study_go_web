package repository

import (
	"echo-practice/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
	Create(user *model.User) error
	DeleteUser(id string) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) DeleteUser(id string) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormUserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}
