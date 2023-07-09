package reqval

// validationError represents an error with validation information.
type validationError struct {
	// Field represents the field associated with the validation error.
	Field string `json:"field,omitempty"`
	// Value represents the value associated with the validation error.
	Value string `json:"value,omitempty"`
	// Message represents the validation error message.
	Message string `json:"message,omitempty"`
}

// SetField sets the field value of the validationError.
func (v *validationError) SetField(field string) {
	v.Field = field
}

// SetValue sets the value value of the validationError.
func (v *validationError) SetValue(value string) {
	v.Value = value
}

// SetMessage sets the message value of the validationError.
func (v *validationError) SetMessage(message string) {
	v.Message = message
}

// Field returns the field associated with the validationError.
func (v validationError) GetField() string {
	return v.Field
}

// Value returns the value associated with the validationError.
func (v validationError) GetValue() string {
	return v.Value
}

// Message returns the error message of the validationError.
func (v validationError) GetMessage() string {
	return v.Message
}

// ValidationError is an interface for errors with validation information.
type ValidationError interface {
	SetField(field string)
	SetValue(value string)
	SetMessage(message string)
	GetField() string
	GetValue() string
	GetMessage() string
}

// SetField creates a ValidationErrorSetterFunc that sets the field value of a ValidationError.
func SetField(field string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetField(field)
	}
}

// SetValue creates a ValidationErrorSetterFunc that sets the value value of a ValidationError.
func SetValue(value string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetValue(value)
	}
}

// SetMessage creates a ValidationErrorSetterFunc that sets the message value of a ValidationError.
func SetMessage(message string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetMessage(message)
	}
}

// ValidationErrorSetterFunc is a function type that sets values of a ValidationError.
type ValidationErrorSetterFunc func(validationError ValidationError)

// NewValidationError creates a new ValidationError based on the provided ValidationErrorSetterFuncs.
func NewValidationError(setterFuncs ...ValidationErrorSetterFunc) ValidationError {
	if len(setterFuncs) == 0 {
		return nil
	}

	valErr := &validationError{}

	for i := range setterFuncs {
		setterFuncs[i](valErr)
	}

	return valErr
}

// insertToValidationErrors inserts a ValidationError into an existing slice of ValidationErrors at a specific position.
// It handles the resizing and copying of the slice to accommodate the new element.
func insertToValidationErrors(original []ValidationError, position int, value ValidationError) []ValidationError {
	l := len(original)
	target := original
	if cap(original) == l {
		target = make([]ValidationError, l+1, l+10)
		copy(target, original[:position])
	} else {
		target = append(target, &validationError{})
	}
	copy(target[position+1:], original[position:])
	target[position] = value
	return target
}

// NewValidationErrors creates a new slice of ValidationErrors with the provided ValidationError instances.
func NewValidationErrors(validationErrors ...ValidationError) ValidationErrors {
	newValErrors := make([]ValidationError, 0)
	if count := len(validationErrors); count > 0 {
		newValErrors = append(newValErrors, validationErrors...)
	}

	return newValErrors
}

// ValidationErrors represents a slice of ValidationError instances.
type ValidationErrors []ValidationError

// GetLength returns the length of the ValidationErrors slice.
func (v ValidationErrors) GetLength() int {
	return len(v)
}

// Append adds new ValidationErrors to the existing ValidationErrors slice.
func (v *ValidationErrors) Append(valErrors ...ValidationError) {
	if len(valErrors) == 0 {
		return
	}

	for i := range valErrors {
		*v = insertToValidationErrors(*v, v.GetLength(), valErrors[i])
	}
}

