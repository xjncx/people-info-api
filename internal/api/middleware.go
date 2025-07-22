package api

import (
	"time"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		req := c.Request()
		if req.Method == "GET" && req.URL.RawQuery != "" {
			logger.Log.Info("[%s] %s?%s", req.Method, req.URL.Path, req.URL.RawQuery)
		} else {
			logger.Log.Info("[%s] %s", req.Method, req.URL.Path)
		}

		err := next(c)

		logger.Log.Info("Request completed in %v", time.Since(start))

		return err
	}
}
