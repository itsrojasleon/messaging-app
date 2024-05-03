package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/itsrojasleon/messaging-app/common/errors"
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
		errors.HandleError(w, &errors.ValidationError{
			Detail: "Invalid JSON provided",
			Field:  "body",
		})
		return
	}

	dir, _ := os.Getwd()

	schema := jsonschema.NewReferenceLoader("file://" + filepath.Join(dir, "../internal/schemas/signup-schema.json"))
	// TODO: Rename.
	documentLoader := jsonschema.NewStringLoader(`{
		"email": "testuser",
		"password": "securepassword123"
	}`)

	_, isValid := validation.ValidateJSON(schema, documentLoader)

	if !isValid {
		errors.HandleError(w, &errors.ValidationError{})
	}

	// Add logic here to add user to database etc.

	// response := map[string]string{"status": "success", "message": "User registered"}

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)

	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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
