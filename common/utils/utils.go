package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	httpErrors "github.com/itsrojasleon/messaging-app/common/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ParseBody(r *http.Request, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr || reflect.TypeOf(v).Elem().Kind() != reflect.Struct {
		return errors.New("provided variable is not a pointer to struct")
	}

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return httpErrors.RequestValidationError{
			ValidationErrors: []httpErrors.ValidationError{{
				Message: "invalid JSON provided",
				Field:   "body",
			}},
		}
	}
	return nil
}

func ValidateJSON(schemaPath, schemaString string) ([]httpErrors.ValidationError, bool) {
	refLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	strLoader := gojsonschema.NewStringLoader(schemaString)

	res, err := gojsonschema.Validate(refLoader, strLoader)
	if err != nil {
		return []httpErrors.ValidationError{
			{
				Message: "internal server error during validation: " + err.Error(),
			},
		}, false
	}
	if res.Valid() {
		return nil, true
	}

	var valErrs []httpErrors.ValidationError
	for _, desc := range res.Errors() {
		valErrs = append(valErrs, httpErrors.ValidationError{
			Message: desc.Description(),
			Field:   desc.Field(),
		})
	}
	return valErrs, false
}

func ToJson(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("cannot encode value into JSON")
	}
	return string(jsonData), nil
}
