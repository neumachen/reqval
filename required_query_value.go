package reqval

import (
	"fmt"
	"net/http"
)

var requiredQueryParamMessageFunc = func(field string) string {
	return fmt.Sprintf("query param: %s is required but was not provided", field)
}

// RequiredQueryValue ...
type RequiredQueryValue struct {
	MessageFunc func(field string) string
}

// Validate checks whether a value is empty or not by determining it's length
func (r *RequiredQueryValue) Validate(req *http.Request, field string) (ValidationErrors, error) {
	fieldValues := req.URL.Query()[field]

	if IsNil(r.MessageFunc) {
		r.MessageFunc = requiredQueryParamMessageFunc
	}

	validationErrors := make(ValidationErrors, 0)
	if len(fieldValues) == 0 {
		validationErrors.Append(NewValidationError(
			SetField(field),
			SetMessage(r.MessageFunc(field)),
		))
	} else {
		for _, fieldValue := range fieldValues {
			if len(fieldValue) > 0 {
				continue
			}
			validationErrors.Append(NewValidationError(
				SetField(field),
				SetValue(fieldValue),
				SetMessage(r.MessageFunc(field)),
			))
		}
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
