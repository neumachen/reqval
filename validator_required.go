package reqval

import (
	"net/http"
)

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

	if fieldValues == nil || len(fieldValues) == 0 {
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

const requiredFormValueMessage = "No form value set"

// RequiredFormValue ...
type RequiredFormValue struct {
	Message   string
	MaxMemory int64
}

// Validate ...
func (r *RequiredFormValue) Validate(req *http.Request, field string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)

	if r.Message == "" {
		r.Message = requiredFormValueMessage
	}

	if r.MaxMemory == 0 {
		r.MaxMemory = (10 * 1024 * 1024)
	}
	err := req.ParseMultipartForm(r.MaxMemory)
	if err != nil {
		return nil, err
	}

	formValues := req.PostForm
	if formValues[field] == nil || len(formValues[field]) == 0 {
		validationErrors = append(validationErrors, NewValidationError(field, "", r.Message))
	}

	for _, formValue := range formValues[field] {
		if len(formValue) > 0 {
			continue
		}
		validationErrors = append(validationErrors, NewValidationError(field, formValue, r.Message))
	}

	if len(validationErrors) == 0 {
		return nil, nil

	}
	return validationErrors, nil
}

const requiredFormFileMessage = "No file sent"

// RequiredFormFile ...
type RequiredFormFile struct {
	Message string
}

// Validate ...
func (r *RequiredFormFile) Validate(req *http.Request, field string) (ValidationErrors, error) {
	validationErrors := make(ValidationErrors, 0)

	if r.Message == "" {
		r.Message = requiredFormFileMessage
	}

	_, _, err := req.FormFile(field)
	switch err {
	case nil:
		// do nothing
	case http.ErrMissingFile:
		validationErrors = append(validationErrors, NewValidationError(field, "", r.Message))
	default:
		return nil, err
	}
	if len(validationErrors) == 0 {
		return nil, nil
	}
	return validationErrors, nil
}
