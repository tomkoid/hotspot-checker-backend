package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/models"
)

func indexRoute(c echo.Context) error {
	resp := models.Response{
		Count:                   Count,
		LastUpdateTimeSeconds:   time.Since(LastUpdateTime).Seconds(),
		Stopped:                 Stopped,
		SentConsecutiveMessages: SentConsecutiveMessages,
	}

	fmt.Printf("Count: %d\n", Count)

	return c.JSON(http.StatusOK, resp)
}
