package utils

type ApplicationError struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}
