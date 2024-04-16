package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/models"
	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func sendStartNotification() {
	notifReq := models.NotificationRequest{
		Title:    "Status",
		Message:  "ESP is up!",
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

func startRoute(c echo.Context) error {
	err := tools.ValidateBody(c, secretPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	LastUpdateTime = time.Now()
	Stopped = false

	go sendStartNotification()

	return c.String(http.StatusOK, "Started")
}
