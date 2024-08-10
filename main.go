package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/models"
	"codeberg.org/tomkoid/http-server-count/internal/routes"
	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func main() {
	// load .env variable
	tools.EnvLoad()

	e := echo.New()
	e.HideBanner = true

	routes.RegisterRoutes(e)

	waitTimeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			timeSince := time.Since(routes.LastUpdateTime)

			if os.Getenv("DEBUG") == "true" {
				fmt.Printf("Last update: %.2fs, Stopped: %t\n", timeSince.Seconds(), routes.Stopped)
			}

			if timeSince.Seconds() > float64(waitTimeout) && routes.SentConsecutiveMessages < 2 && !routes.Stopped {
				routes.SentConsecutiveMessages++
				log.Println("Too late!")

				// send ntfy notification
				notifReq := models.NotificationRequest{
					Title:    os.Getenv("NTFY_TITLE"),
					Message:  os.Getenv("NTFY_MSG"),
					Tags:     "warning,skull",
					Priority: "urgent",
				}

				reqResp, err := tools.SendNotification(&notifReq)
				if err != nil {
					log.Println(err)
				}

				if reqResp.StatusCode == 200 {
					log.Println("Sent ntfy notification!")
					routes.LastUpdateTime = time.Now()
				}
			}

			time.Sleep(time.Second)
		}
	}()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
}
