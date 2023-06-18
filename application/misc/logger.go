package misc

import "github.com/labstack/echo/v4"

var log echo.Logger

func InitLogger(e *echo.Echo) {
	log = e.Logger
}

// LogEf error with format
func LogEf(format string, i ...interface{}) {
	log.Errorf(format, i)
}
