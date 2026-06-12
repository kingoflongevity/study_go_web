package repository

import (
	"context"
	"echo-practice/internal/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id uint) (*model.Employee, error)
	List(ctx context.Context) ([]model.Employee, error)
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id uint) error
}

type gormEmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &gormEmployeeRepository{db: db}
}

func (r *gormEmployeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

func (r *gormEmployeeRepository) GetByID(ctx context.Context, id uint) (*model.Employee, error) {
	var emp model.Employee
	err := r.db.WithContext(ctx).Preload("Department").First(&emp, id).Error
	return &emp, err
}

func (r *gormEmployeeRepository) List(ctx context.Context) ([]model.Employee, error) {
	var emps []model.Employee
	err := r.db.WithContext(ctx).Preload("Department").Find(&emps).Error
	return emps, err
}

func (r *gormEmployeeRepository) Update(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Save(employee).Error
}

func (r *gormEmployeeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Employee{}, id).Error
}
