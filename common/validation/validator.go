package validation

import (
	"github.com/itsrojasleon/messaging-app/common/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON(schemaLoader, documentLoader gojsonschema.JSONLoader) ([]errors.ValidationError, bool) {
	res, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return []errors.ValidationError{
			{
				Message: "Internal server error during validation: " + err.Error(),
			},
		}, false
	}

	if res.Valid() {
		return nil, true
	}

	var valErrs []errors.ValidationError
	for _, desc := range res.Errors() {
		valErrs = append(valErrs, errors.ValidationError{
			Message: desc.Description(),
			Field:   desc.Field(),
		})
	}
	return valErrs, false
}
