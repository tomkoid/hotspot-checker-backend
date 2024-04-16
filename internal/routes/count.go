package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/models"
	"codeberg.org/tomkoid/http-server-count/internal/storage"
	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func countRoute(c echo.Context) error {
	SentConsecutiveMessages = 0

	err := tools.ValidateBody(c, secretPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	Count++
	storage.SaveCount(&Count)

	LastUpdateTime = time.Now()

	resp := models.Response{
		Count:                 Count,
		LastUpdateTimeSeconds: time.Since(LastUpdateTime).Seconds(),
	}

	return c.JSON(http.StatusOK, resp)
}
