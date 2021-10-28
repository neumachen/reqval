package reqval

type errorType struct {
	Field   string `json:"field,omitempty"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message,omitempty"`
}

type validationError struct {
	errorType
}

// SetField ...
func (v *validationError) SetField(field string) {
	v.errorType.Field = field
}

// SetValue ...
func (v *validationError) SetValue(value string) {
	v.errorType.Value = value
}

// SetMessage ...
func (v *validationError) SetMessage(message string) {
	v.errorType.Message = message
}

// Field ...
func (v *validationError) Field() string {
	return v.errorType.Field
}

// Value ...
func (v *validationError) Value() string {
	return v.errorType.Value
}

// Message ...
func (v *validationError) Message() string {
	return v.errorType.Message
}

var _ ValidationError = (*validationError)(nil)

// ValidationError ...
type ValidationError interface {
	SetField(field string)
	SetValue(value string)
	SetMessage(message string)
	Field() string
	Value() string
	Message() string
}

func SetField(field string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetField(field)
	}
}

func SetValue(value string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetValue(value)
	}
}

func SetMessage(message string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetMessage(message)
	}
}

type ValidationErrorSetterFunc func(validationError ValidationError)

// NewValidationError ...
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

func insertToValidationErrors(
	original []ValidationError,
	position int,
	value ValidationError,
) []ValidationError {
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

func NewValidationErrors() ValidationErrors {
	return make([]ValidationError, 0)
}

// ValidationErrors ...
type ValidationErrors []ValidationError

func (v ValidationErrors) GetLength() int {
	return len(v)
}

func (v *ValidationErrors) Append(valErrors ...ValidationError) {
	if len(valErrors) == 0 {
		return
	}

	for i := range valErrors {
		*v = insertToValidationErrors(*v, v.GetLength(), valErrors[i])
	}
}
