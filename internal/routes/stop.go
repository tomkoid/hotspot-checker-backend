package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/models"
	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func sendStopNotification() {
	notifReq := models.NotificationRequest{
		Title:    "Status",
		Message:  "ESP has stopped!",
		Tags:     "",
		Priority: "min",
	}

	reqResp, err := tools.SendNotification(&notifReq)
	if err != nil {
		fmt.Println(err)
	}

	if reqResp.StatusCode == 200 {
		fmt.Println("start: sent ntfy notification")
	}
}

func stopRoute(c echo.Context) error {
	err := tools.ValidateBody(c, secretPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if Stopped {
		return c.String(http.StatusBadRequest, "Already stopped")
	}

	LastUpdateTime = time.Now()
	Stopped = true

	go sendStopNotification()

	return c.String(http.StatusOK, "Stopped")
}
