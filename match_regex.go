package reqval

import (
	"net/http"
	"regexp"
)

// MatchRegexPattern ...
type MatchRegexPattern struct {
	RegexPattern *regexp.Regexp
	Message      string
}

const matchRegexPatternMessage = "Must be int matching regex: ^([0-9]|[1-9][0-9])$"

// Validate checks whether a value is empty or not by determining it's length
func (m *MatchRegexPattern) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	validationErrors := make(ValidationErrors, 0)

	if m.Message == "" {
		m.Message = matchRegexPatternMessage
	}

	// we do not check for presence of the fields since this validator is
	// meant for validating that a value is an int. If presence is
	// required, the required validator should be used in conjuction with
	// this one.
	if fieldValues == nil || len(fieldValues) == 0 {
		return nil, nil
	}

	for _, fieldValue := range fieldValues {
		if m.RegexPattern.MatchString(fieldValue) {
			continue
		}
		validationErrors = append(validationErrors, NewValidationError(field, fieldValue, m.Message))
	}

	if len(validationErrors) == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
