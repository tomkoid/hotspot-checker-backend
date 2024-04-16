package tools

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"codeberg.org/tomkoid/http-server-count/internal/models"
)

func SendNotification(notifReq *models.NotificationRequest) (*http.Response, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://ntfy.sh/%s", os.Getenv("NTFY_ROOM")),
		bytes.NewBuffer([]byte(notifReq.Message)),
	)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Priority", notifReq.Priority)
	req.Header.Set("Title", notifReq.Title)
	req.Header.Set("Tags", notifReq.Tags)

	return http.DefaultClient.Do(req)
}
