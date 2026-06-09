package handler

import (
	"echo-practice/internal/config"
	"echo-practice/internal/repository"
	"echo-practice/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoutes 统一注册所有路由
func RegisterRoutes(e *echo.Echo, cfg *config.Config, db *gorm.DB) {
	// 依赖注入 (Dependency Injection)
	// 1. 初始化 Repository (数据层)
	userRepo := repository.NewGormUserRepository(db)

	// 2. 初始化 Service (业务层)，注入 Repository
	userSvc := service.NewUserService(userRepo)

	// 3. 初始化 Handler (接口层)，注入 Service
	userHandler := NewUserHandler(userSvc)

	// 你可以在这里使用 cfg 做一些判断
	if cfg.Server.Env == "development" {
		e.Logger.Info("Enabling debug mode from config")
	}

	// 分组路由
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users/:id", userHandler.GetUser)
		v1.POST("/users", userHandler.CreateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
