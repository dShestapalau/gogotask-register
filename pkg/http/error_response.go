package utils

type ErrorResponse struct {
	Errors []ApiError `json:"errors"`
}
