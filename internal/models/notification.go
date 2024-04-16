package models

type NotificationRequest struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Tags     string `json:"tags"`
	Priority string `json:"priority"`
}
