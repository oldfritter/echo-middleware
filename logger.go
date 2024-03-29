package middleware

import (
	"github.com/oldfritter/echo/v4"
	mw "github.com/oldfritter/echo/v4/middleware"
	"io"
)

// LoggerWithOutput returns a Logger middleware with output.
// See: `Logger()`.
func LoggerWithOutput(w io.Writer) echo.MiddlewareFunc {
	config := mw.DefaultLoggerConfig
	config.Output = w

	return mw.LoggerWithConfig(config)
}
