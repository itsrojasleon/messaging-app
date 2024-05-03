package validation

import (
	"github.com/itsrojasleon/messaging-app/common/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON(schemaLoader, documentLoader gojsonschema.JSONLoader) ([]errors.ErrorDetail, bool) {
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return []errors.ErrorDetail{
			{
				Description: "Internal server error during validation: " + err.Error(),
			},
		}, false
	}

	if result.Valid() {
		return nil, true
	}

	var validationErrors []errors.ErrorDetail
	for _, desc := range result.Errors() {
		validationErrors = append(validationErrors, errors.ErrorDetail{
			Description: desc.Description(),
			Field:       desc.Field(),
		})
	}
	return validationErrors, false
}
