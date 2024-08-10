package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/hotspot-checker-backend/internal/models"
	"codeberg.org/tomkoid/hotspot-checker-backend/internal/storage"
	"codeberg.org/tomkoid/hotspot-checker-backend/internal/tools"
)

var (
	writeSave = true
)

func countRoute(c echo.Context) error {
	SentConsecutiveMessages = 0

	err := tools.ValidateBody(c, secretPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	Count++

	if writeSave {
		err = storage.SaveCount(&Count)
		if err != nil {
			writeSave = false
		}
	}

	LastUpdateTime = time.Now()

	resp := models.Response{
		Count:                 Count,
		LastUpdateTimeSeconds: time.Since(LastUpdateTime).Seconds(),
	}

	return c.JSON(http.StatusOK, resp)
}
