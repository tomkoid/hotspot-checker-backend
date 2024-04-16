package models

type Response struct {
	Count                   int     `json:"count"`
	LastUpdateTimeSeconds   float64 `json:"lastUpdateTime"`
	Stopped                 bool    `json:"stopped"`
	SentConsecutiveMessages int     `json:"sentConsecutiveMessages"`
}
