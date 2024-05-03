package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Description string `json:"description"`
	Field       string `json:"field,omitempty"`
}

type HttpError interface {
	Error() string
	StatusCode() int
	SerializeErrors() []ErrorDetail
}

type ValidationError struct {
	Detail string
	Field  string
}

func (e *ValidationError) Error() string {
	return e.Detail
}

func (e *ValidationError) SerializeErrors() []ErrorDetail {
	return []ErrorDetail{{Description: e.Detail, Field: e.Field}}
}

type UnauthorizedError struct {
	Detail string
}

func (e *UnauthorizedError) Error() string {
	return e.Detail
}

func (e *UnauthorizedError) SerializeErrors() []ErrorDetail {
	return []ErrorDetail{{Description: e.Detail}}
}

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch e := err.(type) {
	case *ValidationError:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: e.SerializeErrors()})
	case *UnauthorizedError:
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: e.SerializeErrors()})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []ErrorDetail{{Description: "Internal server error"}},
		})
	}
}
