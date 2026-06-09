package main

import (
	"echo-practice/internal/config"
	"echo-practice/internal/handler"
	"echo-practice/internal/handler/middlewares"
	"echo-practice/internal/model"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. 强制开启颜色输出
	color.Enable()

	// 2. 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 3. 初始化数据库
	db, err := gorm.Open(mysql.Open(cfg.Database.Source), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// 4. 自动迁移
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}

	// 5. 初始化 Echo 实例
	e := echo.New()

	// 6. 注册校验器
	e.Validator = middlewares.NewCustomValidator()

	// 7. 基础中间件
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\x1b[36m${time_rfc3339}\x1b[0m | ${color}${status}${reset} | \x1b[33m${latency_human}\x1b[0m | ${methodColor}${method}${reset} | ${uri}\n",
	}))
	e.Use(middleware.Recover())

	// 8. 自定义全局错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Printf("\x1b[31m[ERROR]\x1b[0m %v\n", err)
		e.DefaultHTTPErrorHandler(err, c)
	}

	// 9. 路由注册 (依赖注入：e, config 指针, db 指针)
	handler.RegisterRoutes(e, &cfg, db)

	// 10. 启动服务器
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
