package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ValidationError struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type ErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

type CustomError interface {
	Error() string
	StatusCode() int
	SerializeErrors() ErrorResponse
}

type RequestValidationError struct {
	ValidationErrors []ValidationError
}

func (e RequestValidationError) Error() string {
	return "Error validating incoming properties"
}

func (e RequestValidationError) StatusCode() int {
	return http.StatusBadRequest
}

func (e RequestValidationError) SerializeErrors() ErrorResponse {
	return ErrorResponse{
		Errors: e.ValidationErrors,
	}
}

type UnauthorizedError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return "Lacks valid authentication credentials for the requested resource"
}

func (e UnauthorizedError) SerializeErrors() ErrorResponse {
	return ErrorResponse{
		Errors: []ValidationError{{Message: e.Message}},
	}
}

func (e UnauthorizedError) StatusCode() int {
	return http.StatusUnauthorized
}

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if customErr, ok := err.(CustomError); ok {
		w.WriteHeader(customErr.StatusCode())
		json.NewEncoder(w).Encode(customErr.SerializeErrors())
	} else {
		log.Printf("something went wrong: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []ValidationError{{Message: "Internal server error"}},
		})
	}
}
