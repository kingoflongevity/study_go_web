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

	deptRepo := repository.NewDepartmentRepository(db)
	deptSvc := service.NewDepartmentService(deptRepo)
	deptHandler := NewDepartmentHandler(deptSvc)

	empRepo := repository.NewEmployeeRepository(db)
	empSvc := service.NewEmployeeService(empRepo, deptRepo)
	empHandler := NewEmployeeHandler(empSvc)

	if cfg.Server.Env == "development" {
		e.Logger.Info("Enabling development mode")
	}

	// 注册 Huma 路由
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        "/api/v1/users/{id}",
		Summary:     "获取用户",
		Description: "根据 ID 获取用户信息",
		Tags:        []string{"Users"},
	}, userHandler.GetUser)

	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/v1/users",
		Summary:     "创建用户",
		Description: "创建一个新的用户",
		Tags:        []string{"Users"},
	}, userHandler.CreateUser)

	huma.Register(api, huma.Operation{
		OperationID: "update-user",
		Method:      http.MethodPut,
		Path:        "/api/v1/users/{id}",
		Summary:     "更新用户",
		Description: "更新指定 ID 的用户信息",
		Tags:        []string{"Users"},
	}, userHandler.UpdateUser)

	huma.Register(api, huma.Operation{
		OperationID: "delete-user",
		Method:      http.MethodDelete,
		Path:        "/api/v1/users/{id}",
		Summary:     "删除用户",
		Description: "根据 ID 删除用户",
		Tags:        []string{"Users"},
	}, userHandler.DeleteUser)

	// ========== 员工路由 ==========

	huma.Register(api, huma.Operation{
		OperationID: "list-employees",
		Method:      http.MethodGet,
		Path:        "/api/v1/employees",
		Summary:     "员工列表",
		Description: "获取所有员工列表（包含部门信息）",
		Tags:        []string{"Employees"},
	}, empHandler.ListEmployees)

	huma.Register(api, huma.Operation{
		OperationID: "create-employee",
		Method:      http.MethodPost,
		Path:        "/api/v1/employees",
		Summary:     "创建员工",
		Description: "创建一个新员工并关联部门",
		Tags:        []string{"Employees"},
	}, empHandler.CreateEmployee)

	huma.Register(api, huma.Operation{
		OperationID: "get-employee",
		Method:      http.MethodGet,
		Path:        "/api/v1/employees/{id}",
		Summary:     "获取员工详情",
		Description: "根据 ID 获取员工信息及所属部门",
		Tags:        []string{"Employees"},
	}, empHandler.GetEmployee)

	huma.Register(api, huma.Operation{
		OperationID: "update-employee",
		Method:      http.MethodPut,
		Path:        "/api/v1/employees/{id}",
		Summary:     "更新员工",
		Description: "更新指定 ID 的员工信息",
		Tags:        []string{"Employees"},
	}, empHandler.UpdateEmployee)

	huma.Register(api, huma.Operation{
		OperationID: "delete-employee",
		Method:      http.MethodDelete,
		Path:        "/api/v1/employees/{id}",
		Summary:     "删除员工",
		Description: "根据 ID 删除员工",
		Tags:        []string{"Employees"},
	}, empHandler.DeleteEmployee)

	// ========== 部门路由 ==========

	huma.Register(api, huma.Operation{
		OperationID: "list-departments",
		Method:      http.MethodGet,
		Path:        "/api/v1/departments",
		Summary:     "部门列表",
		Description: "获取部门树形列表",
		Tags:        []string{"Departments"},
	}, deptHandler.ListDepartments)

	huma.Register(api, huma.Operation{
		OperationID: "create-department",
		Method:      http.MethodPost,
		Path:        "/api/v1/departments",
		Summary:     "创建部门",
		Description: "创建一个新部门",
		Tags:        []string{"Departments"},
	}, deptHandler.CreateDepartment)

	huma.Register(api, huma.Operation{
		OperationID: "get-department",
		Method:      http.MethodGet,
		Path:        "/api/v1/departments/{id}",
		Summary:     "获取部门",
		Description: "根据 ID 获取部门详情（含子部门）",
		Tags:        []string{"Departments"},
	}, deptHandler.GetDepartment)

	huma.Register(api, huma.Operation{
		OperationID: "update-department",
		Method:      http.MethodPut,
		Path:        "/api/v1/departments/{id}",
		Summary:     "更新部门",
		Description: "更新指定 ID 的部门信息",
		Tags:        []string{"Departments"},
	}, deptHandler.UpdateDepartment)

	huma.Register(api, huma.Operation{
		OperationID: "delete-department",
		Method:      http.MethodDelete,
		Path:        "/api/v1/departments/{id}",
		Summary:     "删除部门",
		Description: "根据 ID 删除部门",
		Tags:        []string{"Departments"},
	}, deptHandler.DeleteDepartment)
}
