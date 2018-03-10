package reqval

import "net/http"

// RequestValidator ...
type RequestValidator interface {
	Validate(req *http.Request, field string) (ValidationErrors, error)
}

// RequestValidators ...
type RequestValidators []RequestValidator

// RequestValidations ...
type RequestValidations map[string]RequestValidators

// Validate ...
func Validate(req *http.Request, validations RequestValidations) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)
	for field, validators := range validations {
		for _, validator := range validators {
			valErrs, err := validator.Validate(req, field)
			if err != nil {
				return nil, err
			}
			if len(valErrs) > 0 {
				validationErrors = append(validationErrors, valErrs...)
			}
		}
	}
	if len(validationErrors) == 0 {
		return nil, nil
	}
	return validationErrors, nil
}
