package handler

import (
	"echo-practice/internal/config"
	"echo-practice/internal/repository"
	"echo-practice/internal/service"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes 统一注册所有路由
func RegisterRoutes(e *echo.Echo, cfg *config.Config, db *gorm.DB) {
	// 初始化 Huma API
	api := humaecho.New(e, huma.DefaultConfig("Echo Practice API", "1.0.0"))

	// 依赖注入
	userRepo := repository.NewGormUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userSvc)

	if cfg.Server.Env == "development" {
		e.Logger.Info("Enabling development mode")
	}

	// 注册 Huma 路由
	huma.Register(api, huma.Operation{
		OperationID:   "get-user",
		Method:        http.MethodGet,
		Path:          "/api/v1/users/{id}",
		Summary:       "获取用户",
		Description:   "根据 ID 获取用户信息",
		Tags:          []string{"Users"},
	}, userHandler.GetUser)

	huma.Register(api, huma.Operation{
		OperationID:   "create-user",
		Method:        http.MethodPost,
		Path:          "/api/v1/users",
		Summary:       "创建用户",
		Description:   "创建一个新的用户",
		Tags:          []string{"Users"},
	}, userHandler.CreateUser)

	huma.Register(api, huma.Operation{
		OperationID:   "update-user",
		Method:        http.MethodPut,
		Path:          "/api/v1/users/{id}",
		Summary:       "更新用户",
		Description:   "更新指定 ID 的用户信息",
		Tags:          []string{"Users"},
	}, userHandler.UpdateUser)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-user",
		Method:        http.MethodDelete,
		Path:          "/api/v1/users/{id}",
		Summary:       "删除用户",
		Description:   "根据 ID 删除用户",
		Tags:          []string{"Users"},
	}, userHandler.DeleteUser)
}
