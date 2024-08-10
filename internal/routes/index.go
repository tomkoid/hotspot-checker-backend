package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/hotspot-checker-backend/internal/models"
)

func indexRoute(c echo.Context) error {
	resp := models.Response{
		Count:                   Count,
		LastUpdateTimeSeconds:   time.Since(LastUpdateTime).Seconds(),
		Stopped:                 Stopped,
		SentConsecutiveMessages: SentConsecutiveMessages,
	}

	return c.JSON(http.StatusOK, resp)
}
