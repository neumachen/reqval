package reqval

import (
	"net/http"
	"strings"
)

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

	cType := req.Header.Get("Content-Type")
	cTypes := strings.Split(cType, ";")
	if cTypes[0] == "multipart/form-data" {
		err := req.ParseMultipartForm(r.MaxMemory)
		if err != nil {
			return nil, err
		}
	} else {
		err := req.ParseForm()
		if err != nil {
			return nil, err
		}

	}

	formValues := req.PostForm
	if formValues[field] == nil || len(formValues[field]) == 0 {
		validationErrors.Append(NewValidationError(
			SetField(field),
			SetValue(""),
			SetMessage(r.Message),
		))
	}

	for _, formValue := range req.Form[field] {
		if len(formValue) > 0 {
			continue
		}
		validationErrors.Append(NewValidationError(
			SetField(field),
			SetValue(formValue),
			SetMessage(r.Message),
		))
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil

	}
	return validationErrors, nil
}
