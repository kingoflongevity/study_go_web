package repository

import (
	"context"
	"echo-practice/internal/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(ctx context.Context, department *model.Department) error
	GetByID(ctx context.Context, id uint) (*model.Department, error)
	Update(ctx context.Context, department *model.Department) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (d *departmentRepository) Create(ctx context.Context, department *model.Department) error {
	return d.db.WithContext(ctx).Create(department).Error
}

func (d *departmentRepository) GetByID(ctx context.Context, id uint) (*model.Department, error) {
	var department model.Department
	if err := d.db.WithContext(ctx).Preload("Children").First(&department, id).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (d *departmentRepository) Update(ctx context.Context, department *model.Department) error {
	return d.db.WithContext(ctx).Model(department).Updates(department).Error
}

func (d *departmentRepository) Delete(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&model.Department{}, id).Error
}

func (d *departmentRepository) List(ctx context.Context) ([]model.Department, error) {
	var departments []model.Department
	err := d.db.WithContext(ctx).Where("parent_id IS NULL").Preload("Children").Find(&departments).Error
	return departments, err
}
