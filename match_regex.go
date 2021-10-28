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

const matchRegexPatternMessage = "Must regex pattern"

// Validate checks whether a value is empty or not by determining it's length
func (m *MatchRegexPattern) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	if m.Message == "" {
		m.Message = matchRegexPatternMessage
	}

	// we do not check for presence of the fields since this validator is
	// meant for validating that a value matches a regex pattern. If
	// presence is required, the required validator should be used in
	// conjuction with this one.
	if len(fieldValues) == 0 {
		return nil, nil
	}

	validationErrors := NewValidationErrors()
	for _, fieldValue := range fieldValues {
		if m.RegexPattern.MatchString(fieldValue) {
			continue
		}
		validationErrors.Append(NewValidationError(
			SetParam(field),
			SetValue(fieldValue),
			SetMessage(m.Message),
		))
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
