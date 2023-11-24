package main

type ConsumerRes struct {
	Message   interface{} `json:"message"`
	MessageID int         `json:"message_id"`
	Error     string      `json:"error"`
}

type ConsumerReq struct {
	MessageID int  `json:"message_id"`
	Destroy   bool `json:"destroy"`
}
