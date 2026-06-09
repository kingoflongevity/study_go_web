package handler

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/service"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

type DepartmentHandler struct {
	svc *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{svc: s}
}

// ---------- 请求 / 响应结构体 ----------

type CreateDepartmentInput struct {
	Body struct {
		Name     string `json:"name" minLength:"1" doc:"部门名称"`
		ParentID *uint  `json:"parent_id,omitempty" doc:"父部门ID，为空表示顶级部门"`
	}
}

type DepartmentIDInput struct {
	ID uint `path:"id" doc:"部门ID"`
}

type DepartmentResponse struct {
	Body *model.Department
}

type DepartmentListResponse struct {
	Body []model.Department
}

// ---------- Handler 方法 ----------

// CreateDepartment 创建部门
func (h *DepartmentHandler) CreateDepartment(ctx context.Context, input *CreateDepartmentInput) (*DepartmentResponse, error) {
	dept := &model.Department{
		Name:     input.Body.Name,
		ParentID: input.Body.ParentID,
	}

	if err := h.svc.CreateDepartment(ctx, dept); err != nil {
		return nil, huma.Error500InternalServerError("创建部门失败: " + err.Error())
	}

	return &DepartmentResponse{Body: dept}, nil
}

// GetDepartment 获取部门详情
func (h *DepartmentHandler) GetDepartment(ctx context.Context, input *DepartmentIDInput) (*DepartmentResponse, error) {
	dept, err := h.svc.GetDepartmentByID(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("部门 %d 不存在", input.ID))
	}

	return &DepartmentResponse{Body: dept}, nil
}

// UpdateDepartment 更新部门
func (h *DepartmentHandler) UpdateDepartment(ctx context.Context, input *struct {
	DepartmentIDInput
	Body struct {
		Name     string `json:"name" minLength:"1" doc:"部门名称"`
		ParentID *uint  `json:"parent_id,omitempty" doc:"父部门ID"`
	}
}) (*DepartmentResponse, error) {
	// 1. 先查是否存在
	dept, err := h.svc.GetDepartmentByID(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("部门 %d 不存在", input.ID))
	}

	// 2. 更新字段
	dept.Name = input.Body.Name
	dept.ParentID = input.Body.ParentID

	// 3. 执行更新
	if err := h.svc.UpdateDepartment(ctx, dept); err != nil {
		return nil, huma.Error500InternalServerError("更新部门失败: " + err.Error())
	}

	return &DepartmentResponse{Body: dept}, nil
}

// DeleteDepartment 删除部门
func (h *DepartmentHandler) DeleteDepartment(ctx context.Context, input *DepartmentIDInput) (*SuccessResponse, error) {
	if err := h.svc.DeleteDepartment(ctx, input.ID); err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("删除部门 %d 失败: %v", input.ID, err))
	}

	resp := &SuccessResponse{}
	resp.Body.Message = fmt.Sprintf("部门 %d 已删除", input.ID)
	return resp, nil
}

// ListDepartments 获取部门列表（树形结构）
func (h *DepartmentHandler) ListDepartments(ctx context.Context, input *struct{}) (*DepartmentListResponse, error) {
	depts, err := h.svc.ListDepartments(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("获取部门列表失败: " + err.Error())
	}

	return &DepartmentListResponse{Body: depts}, nil
}
