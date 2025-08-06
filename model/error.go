package model

// ErrorResponse represents a standard error response
// @Description Standard error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request data"`
}

// BadRequestError represents a 400 Bad Request error response
// @Description Bad Request error response
type BadRequestError struct {
	Error string `json:"error" example:"Invalid request data"`
}

// NotFoundError represents a 404 Not Found error response
// @Description Not Found error response
type NotFoundError struct {
	Error string `json:"error" example:"Resource not found"`
}

// InternalServerError represents a 500 Internal Server Error response
// @Description Internal Server Error response
type InternalServerError struct {
	Error string `json:"error" example:"Internal server error occurred"`
}

// DuplicateOrderError represents a duplicate order error response
// @Description Duplicate order error response
type DuplicateOrderError struct {
	Error string `json:"error" example:"Order number already exists"`
}

// ValidationError represents a validation error response
// @Description Validation error response
type ValidationError struct {
	Field   string `json:"field" example:"name"`
	Message string `json:"message" example:"Field is required"`
}

// ValidationErrorResponse represents a validation error response with multiple errors
// @Description Validation error response with multiple field errors
type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors" example:"[{\"field\":\"name\",\"message\":\"Field is required\"}]"`
}
