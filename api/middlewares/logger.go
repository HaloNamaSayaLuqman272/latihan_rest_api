package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoggerConfig struct {
	Config middleware.RequestLoggerConfig
}

func (c *LoggerConfig) Init() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(c.Config)
}
