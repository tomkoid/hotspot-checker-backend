package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/storage"
)

var (
	Count                   = storage.GetCount()
	LastUpdateTime          = time.Now()
	SentConsecutiveMessages = 0 // should not be higher than 2
	Stopped                 = false
	secretPassword          = ""
)

func loadSecretPassword() {
	secretPassword = os.Getenv("PASSWORD")

	if secretPassword == "" {
		fmt.Println("password: PASSWORD is not set")
	}
}

func RegisterRoutes(e *echo.Echo) {
	loadSecretPassword()

	e.GET("/", indexRoute)
	e.POST("/count", countRoute)
	e.POST("/start", startRoute)
	e.POST("/stop", stopRoute)
}
