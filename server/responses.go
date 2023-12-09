package main

// GenericSuccessResponse represents a generic success response.
type GenericSuccessResponse struct {
	Success bool        `json:"success,omitempty" default:"true"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// GenericErrorResponse represents a generic error response.
type GenericErrorResponse struct {
	Status  string      `json:"success,omitempty" default:"false"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
