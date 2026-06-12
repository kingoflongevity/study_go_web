package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestStatus() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			// 3. 【后置处理】计算耗时
			duration := time.Since(start)
			path := c.Path()
			method := c.Request().Method
			status := c.Response().Status
			if err != nil {
				return echo.NewHTTPError(http.StatusBadGateway, err.Error())
			}
			// 基础日志打印
			fmt.Printf("[Stats] %s %s | Status: %d | Latency: %v\n", method, path, status, duration)
			return err
		}
	}
}
