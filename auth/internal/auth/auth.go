package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	httpErrors "github.com/itsrojasleon/messaging-app/common/errors"
	"github.com/itsrojasleon/messaging-app/common/validation"
	jsonschema "github.com/xeipuuv/gojsonschema"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpErrors.HandleError(w, httpErrors.RequestValidationError{
			ValidationErrors: []httpErrors.ValidationError{{
				Message: "Invalid JSON provided",
				Field:   "body",
			}},
		})
		return
	}

	dir, _ := os.Getwd()

	refLoader := jsonschema.NewReferenceLoader("file://" + filepath.Join(dir, "../internal/schemas/signup-schema.json"))
	strLoader := jsonschema.NewStringLoader(`{
		"email": "testuser",
		"password": "securepassword123"
	}`)

	if valErrs, isValid := validation.ValidateJSON(refLoader, strLoader); !isValid {
		httpErrors.HandleError(w, httpErrors.RequestValidationError{
			ValidationErrors: valErrs,
		})
	}

	// Add logic here to add user to database etc.
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "User logged in"})
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Read authorization bearer token.
	// 2. Return JWT data back to user.

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Current user info"})
}
