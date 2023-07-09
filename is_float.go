package reqval

import (
	"net/http"
	"regexp"
)

// IsFloat validates the given value is of int type by matching the string value
// of the queyr parameter to the given regex
type IsFloat struct {
	// Message the validation faiure detail returned
	Message string
	// Extremum if given is set as both the Min and the Max, even if
	// either are given
	Extremum *float64
	// Min if given, it will check if the value is equal to or greater
	// than the preset value
	Min *float64
	// Max if given, it will check if the value is equal to or less
	// than the preset value
	Max *float64
}

const isFloatMessage = "failed to recognize value as float"

var floatRegex = regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)|\.\d+)(?:\d[eE][+\-]?\d+)?$`)

// Validate checks whether a value is empty or not by determining it's length
func (i *IsFloat) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	if i.Message == "" {
		i.Message = isFloatMessage
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
		if isFloat := floatRegex.MatchString(fieldValue); !isFloat {
			validationErrors.Append(NewValidationError(
				SetField(field),
				SetValue(fieldValue),
				SetMessage(i.Message),
			))
		}
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
