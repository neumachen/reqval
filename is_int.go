package reqval

import (
	"net/http"
	"regexp"
)

// IsInt ...
type IsInt struct {
	Message string
}

const isIntMessage = "Must be int matching regex: ^[0-9]+$"

var isIntRegex = regexp.MustCompile("^[0-9]+$")

// Validate checks whether a value is empty or not by determining it's length
func (i *IsInt) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	if i.Message == "" {
		i.Message = isIntMessage
	}

	// we do not check for presence of the fields since this validator is
	// meant for validating that a value is an int. If presence is
	// required, the required validator should be used in conjuction with
	// this one.
	if len(fieldValues) == 0 {
		return nil, nil
	}

	validationErrors := make(ValidationErrors, 0)

	for _, fieldValue := range fieldValues {
		if isIntRegex.MatchString(fieldValue) {
			continue
		}
		validationErrors.Append(NewValidationError(
			SetParam(field),
			SetValue(fieldValue),
			SetMessage(i.Message),
		))
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
