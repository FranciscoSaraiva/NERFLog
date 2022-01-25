package models

type ApiResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    string `json:"data,omitempty"`
}
