package reqval

import "net/http"

const requiredMessage = "This is a required parameter"

// RequiredQueryValue ...
type RequiredQueryValue struct {
	Message string
}

// Validate checks whether a value is empty or not by determining it's length
func (r *RequiredQueryValue) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	validationErrors := make(ValidationErrors, 0)

	if r.Message == "" {
		r.Message = requiredMessage
	}

	if len(fieldValues) == 0 {
		validationErrors = append(validationErrors, NewValidationError(field, "", r.Message))
	}

	for _, fieldValue := range fieldValues {
		if len(fieldValue) > 0 {
			continue
		}
		validationErrors = append(validationErrors, NewValidationError(field, fieldValue, r.Message))

	}

	if len(validationErrors) == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
