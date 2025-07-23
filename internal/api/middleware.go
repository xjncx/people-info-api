package api

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xjncx/people-info-api/pkg/logger"
	"go.uber.org/zap"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		req := c.Request()
		logger.Log.Info("Request started",
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.String("query", req.URL.RawQuery),
		)

		err := next(c)

		logger.Log.Info("Request completed",
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status", c.Response().Status),
		)

		return err
	}
}
