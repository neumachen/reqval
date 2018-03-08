package qryval

// ValidationError ...
type ValidationError interface {
	SetField(value string)
	Field() string
	Value() string
	Message() string
}

type vError struct {
	Field   string `json:"field"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

type validationError struct {
	VErr *vError `json:"validation_error"`
}

// SetValue ...
func (v *validationError) SetField(value string) {
	v.VErr.Field = value
}

// ValField ...
func (v *validationError) Field() string {
	return v.VErr.Field
}

// ValValue ...
func (v *validationError) Value() string {
	return v.VErr.Field
}

// ValidationMessage ...
func (v *validationError) Message() string {
	return v.VErr.Field
}

// NewValidationError ...
func NewValidationError(field, value, message string) ValidationError {
	return &validationError{
		&vError{
			Field: field, Value: value, Message: message,
		},
	}
}

// ValidationErrors ...
type ValidationErrors []ValidationError

// Count ...
func (v ValidationErrors) Count() int {
	return len(v)
}
