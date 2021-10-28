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

func (r RequestValidations) GetLength() int {
	return len(r)
}

// Validate ...
func Validate(req *http.Request, validations RequestValidations) (ValidationErrors, error) {
	validationErrors := NewValidationErrors()
	for param, validators := range validations {
		for _, validator := range validators {
			valErrs, err := validator.Validate(req, param)
			if err != nil {
				return nil, err
			}
			if valErrs.GetLength() > 0 {
				validationErrors.Append(
					valErrs...,
				)
			}
		}
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil
	}
	return validationErrors, nil
}
