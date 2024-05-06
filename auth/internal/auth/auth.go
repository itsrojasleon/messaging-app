package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/itsrojasleon/messaging-app/auth/internal/types"
	httpErrors "github.com/itsrojasleon/messaging-app/common/errors"

	"github.com/itsrojasleon/messaging-app/common/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var credentials types.SignupCreds

	if err := utils.ParseBody(r, &credentials); err != nil {
		httpErrors.HandleError(w, err)
		return
	}

	formattedCreds, err := utils.ToJson(credentials)
	if err != nil {
		httpErrors.HandleError(w, err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		httpErrors.HandleError(w, err)
	}
	schemaPath := filepath.Join(dir, "../internal/schemas/signup.json")
	if valErrs, isValid := utils.ValidateJSON(schemaPath, formattedCreds); !isValid {
		httpErrors.HandleError(w, httpErrors.RequestValidationError{
			ValidationErrors: valErrs,
		})
	}

	// Add logic here to add user to database etc.
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var credentials types.SigninCreds
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "user logged in"})
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Read authorization bearer token.
	// 2. Return JWT data back to user.

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "current user info"})
}
