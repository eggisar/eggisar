package model

type Event struct {
	Subject string      `json:"subject"`
	Action  string      `json:"action"`
	Data    interface{} `json:"data"`
}
