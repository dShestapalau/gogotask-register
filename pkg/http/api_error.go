package utils

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
