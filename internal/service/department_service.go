package service

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/repository"
	"fmt"
	"log/slog"
)

type DepartmentService struct {
	repo   repository.DepartmentRepository
	logger *slog.Logger
}

func NewDepartmentService(repo repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo:   repo,
		logger: slog.With(slog.String("service", "department")),
	}
}

func (svc *DepartmentService) CreateDepartment(ctx context.Context, department *model.Department) error {
	if department.ParentID != nil {
		_, err := svc.repo.GetByID(ctx, *department.ParentID)
		if err != nil {
			return fmt.Errorf("父部门 %d 不存在", *department.ParentID)
		}
	}
	return svc.repo.Create(ctx, department)
}

func (svc *DepartmentService) GetDepartmentByID(ctx context.Context, id uint) (*model.Department, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *DepartmentService) UpdateDepartment(ctx context.Context, department *model.Department) error {
	return svc.repo.Update(ctx, department)
}

func (svc *DepartmentService) DeleteDepartment(ctx context.Context, id uint) error {
	return svc.repo.Delete(ctx, id)
}

func (svc *DepartmentService) ListDepartments(ctx context.Context) ([]model.Department, error) {
	return svc.repo.List(ctx)
}
