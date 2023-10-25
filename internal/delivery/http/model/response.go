package model

type SuccessResponse struct {
	Status string      `json:"status"`
	Code   int64       `json:"code"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
