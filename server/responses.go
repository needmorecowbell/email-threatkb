package main

// GenericSuccessResponse represents a generic success response.
type GenericSuccessResponse struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// GenericErrorResponse represents a generic error response.
type GenericErrorResponse struct {
	Success bool        `json:"success" default:"false"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
