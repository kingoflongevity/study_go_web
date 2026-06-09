package handler

import (
	"context"
	"echo-practice/internal/model"
	"echo-practice/internal/service"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		svc: s,
	}
}

// 定义请求和响应结构体

type CreateUserInput struct {
	Body struct {
		Name  string `json:"name" doc:"User name" minLength:"1"`
		Email string `json:"email" doc:"User email" format:"email"`
	}
}

type UserIDInput struct {
	ID string `path:"id" doc:"User ID"`
}

type UserResponse struct {
	Body *model.User
}

type SuccessResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(ctx context.Context, input *UserIDInput) (*UserResponse, error) {
	user, err := h.svc.GetUserProfile(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("User %s not found", input.ID))
	}
	return &UserResponse{Body: user}, nil
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(ctx context.Context, input *CreateUserInput) (*UserResponse, error) {
	u := &model.User{
		Name:  input.Body.Name,
		Email: input.Body.Email,
	}

	if err := h.svc.CreateUser(ctx, u); err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return &UserResponse{Body: u}, nil
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(ctx context.Context, input *UserIDInput) (*SuccessResponse, error) {
	if err := h.svc.DeleteUser(ctx, input.ID); err != nil {
		return nil, huma.Error404NotFound(err.Error())
	}

	resp := &SuccessResponse{}
	resp.Body.Message = fmt.Sprintf("User %s deleted", input.ID)
	return resp, nil
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(ctx context.Context, input *struct {
	UserIDInput
	Body struct {
		Name  string `json:"name" doc:"New user name"`
		Email string `json:"email" doc:"New user email" format:"email"`
	}
}) (*UserResponse, error) {
	// 1. 先查找是否存在
	user, err := h.svc.GetUserProfile(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("User not found")
	}

	// 2. 更新数据
	user.Name = input.Body.Name
	user.Email = input.Body.Email

	// 3. 执行更新
	if err := h.svc.UpdateUser(ctx, user); err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return &UserResponse{Body: user}, nil
}
