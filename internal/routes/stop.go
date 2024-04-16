package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func stopRoute(c echo.Context) error {
	err := tools.ValidateBody(c, secretPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	LastUpdateTime = time.Now()
	Stopped = true
	return c.String(http.StatusOK, "Stopped")
}
