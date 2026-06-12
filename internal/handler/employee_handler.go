package handler

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/service"

	"github.com/danielgtaylor/huma/v2"
)

type EmployeeHandler struct {
	svc service.EmployeeService
}

func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}

// --- DTOs ---

type EmployeeDeptBrief struct {
	ID   uint   `json:"id" doc:"部门ID"`
	Name string `json:"name" doc:"部门名称"`
}

type EmployeeOutput struct {
	ID           uint              `json:"id" doc:"员工ID"`
	Name         string            `json:"name" doc:"员工姓名"`
	DepartmentID uint              `json:"department_id" doc:"所属部门ID"`
	Department   EmployeeDeptBrief `json:"department" doc:"所属部门信息"`
}

type CreateEmployeeInput struct {
	Body struct {
		Name         string `json:"name" doc:"员工姓名" required:"true"`
		DepartmentID uint   `json:"department_id" doc:"所属部门ID" required:"true"`
	}
}

type UpdateEmployeeInput struct {
	ID   uint `path:"id" doc:"员工ID"`
	Body struct {
		Name         string `json:"name" doc:"员工姓名" required:"true"`
		DepartmentID uint   `json:"department_id" doc:"所属部门ID" required:"true"`
	}
}

// --- Handlers ---

func (h *EmployeeHandler) ListEmployees(ctx context.Context, input *struct{}) (*struct{ Body []EmployeeOutput }, error) {
	emps, err := h.svc.List(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("获取员工列表失败: " + err.Error())
	}

	output := make([]EmployeeOutput, 0, len(emps))
	for _, emp := range emps {
		output = append(output, EmployeeOutput{
			ID:           emp.ID,
			Name:         emp.Name,
			DepartmentID: emp.DepartmentID,
			Department: EmployeeDeptBrief{
				ID:   emp.Department.ID,
				Name: emp.Department.Name,
			},
		})
	}

	return &struct{ Body []EmployeeOutput }{Body: output}, nil
}

func (h *EmployeeHandler) CreateEmployee(ctx context.Context, input *CreateEmployeeInput) (*struct{ Body EmployeeOutput }, error) {
	emp := &model.Employee{
		Name:         input.Body.Name,
		DepartmentID: input.Body.DepartmentID,
	}

	if err := h.svc.Create(ctx, emp); err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	// 重新获取以包含部门信息
	newEmp, _ := h.svc.GetByID(ctx, emp.ID)

	return &struct{ Body EmployeeOutput }{
		Body: EmployeeOutput{
			ID:           newEmp.ID,
			Name:         newEmp.Name,
			DepartmentID: newEmp.DepartmentID,
			Department: EmployeeDeptBrief{
				ID:   newEmp.Department.ID,
				Name: newEmp.Department.Name,
			},
		},
	}, nil
}

func (h *EmployeeHandler) GetEmployee(ctx context.Context, input *struct {
	ID uint `path:"id"`
}) (*struct{ Body EmployeeOutput }, error) {
	emp, err := h.svc.GetByID(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("员工不存在")
	}

	return &struct{ Body EmployeeOutput }{
		Body: EmployeeOutput{
			ID:           emp.ID,
			Name:         emp.Name,
			DepartmentID: emp.DepartmentID,
			Department: EmployeeDeptBrief{
				ID:   emp.Department.ID,
				Name: emp.Department.Name,
			},
		},
	}, nil
}

func (h *EmployeeHandler) UpdateEmployee(ctx context.Context, input *UpdateEmployeeInput) (*struct{ Body EmployeeOutput }, error) {
	emp := &model.Employee{
		Name:         input.Body.Name,
		DepartmentID: input.Body.DepartmentID,
	}
	emp.ID = input.ID

	if err := h.svc.Update(ctx, emp); err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	newEmp, _ := h.svc.GetByID(ctx, emp.ID)

	return &struct{ Body EmployeeOutput }{
		Body: EmployeeOutput{
			ID:           newEmp.ID,
			Name:         newEmp.Name,
			DepartmentID: newEmp.DepartmentID,
			Department: EmployeeDeptBrief{
				ID:   newEmp.Department.ID,
				Name: newEmp.Department.Name,
			},
		},
	}, nil
}

func (h *EmployeeHandler) DeleteEmployee(ctx context.Context, input *struct {
	ID uint `path:"id"`
}) (*struct{}, error) {
	if err := h.svc.Delete(ctx, input.ID); err != nil {
		return nil, huma.Error500InternalServerError("删除失败")
	}
	return nil, nil
}
