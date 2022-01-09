package model

type ErrorBody struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
