package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
