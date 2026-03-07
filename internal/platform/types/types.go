package types

import "time"

type APIResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    any       `json:"data"`
	Time    time.Time `json:"timestamp"`
}

func NewAPIResponse() APIResponse {
	var response APIResponse
	response.Status = "failed"
	response.Data = nil
	response.Message = "default message"
	response.Time = time.Now()
	return response
}
