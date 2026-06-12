package service

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/repository"
	"errors"
)

type EmployeeService interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id uint) (*model.Employee, error)
	List(ctx context.Context) ([]model.Employee, error)
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id uint) error
}

type employeeService struct {
	repo     repository.EmployeeRepository
	deptRepo repository.DepartmentRepository
}

func NewEmployeeService(repo repository.EmployeeRepository, deptRepo repository.DepartmentRepository) EmployeeService {
	return &employeeService{
		repo:     repo,
		deptRepo: deptRepo,
	}
}

func (s *employeeService) Create(ctx context.Context, employee *model.Employee) error {
	// 校验部门是否存在
	_, err := s.deptRepo.GetByID(ctx, employee.DepartmentID)
	if err != nil {
		return errors.New("所选部门不存在")
	}
	return s.repo.Create(ctx, employee)
}

func (s *employeeService) GetByID(ctx context.Context, id uint) (*model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *employeeService) List(ctx context.Context) ([]model.Employee, error) {
	return s.repo.List(ctx)
}

func (s *employeeService) Update(ctx context.Context, employee *model.Employee) error {
	// 校验部门是否存在
	_, err := s.deptRepo.GetByID(ctx, employee.DepartmentID)
	if err != nil {
		return errors.New("所选部门不存在")
	}
	return s.repo.Update(ctx, employee)
}

func (s *employeeService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
