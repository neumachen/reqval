package reqval

import "net/http"

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
