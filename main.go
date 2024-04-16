package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"codeberg.org/tomkoid/http-server-count/internal/routes"
	"codeberg.org/tomkoid/http-server-count/internal/tools"
)

func main() {
	// load .env variable
	tools.EnvLoad()

	e := echo.New()
	e.HideBanner = true

	routes.RegisterRoutes(e)

	go func() {
		for {
			timeSince := time.Since(routes.LastUpdateTime)
			fmt.Printf("Last update: %.2fs\nStopped: %t\n", timeSince.Seconds(), routes.Stopped)

			if timeSince.Seconds() > 35.0 && routes.SentConsecutiveMessages < 2 && !routes.Stopped {
				routes.SentConsecutiveMessages++
				fmt.Println("too soon")

				// send ntfy notification
				req, err := http.NewRequest(
					"POST",
					fmt.Sprintf("https://ntfy.sh/%s", os.Getenv("NTFY_ROOM")),
					bytes.NewBuffer([]byte(os.Getenv("NTFY_MSG"))),
				)

				req.Header.Set("Priority", "urgent")
				req.Header.Set("Title", os.Getenv("NTFY_TITLE"))
				req.Header.Set("Tags", "warning,skull")

				reqResp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println(err)
				}

				if reqResp.StatusCode == 200 {
					fmt.Println("sent ntfy notification")
					routes.LastUpdateTime = time.Now()
				}
			}

			time.Sleep(time.Second)
		}
	}()

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
