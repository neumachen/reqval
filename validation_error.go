package reqval

type errorType struct {
	Param   string `json:"param,omitempty"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message,omitempty"`
}

type validationError struct {
	errorType
}

// SetParam ...
func (v *validationError) SetParam(value string) {
	v.errorType.Param = value
}

// SetValue ...
func (v *validationError) SetValue(value string) {
	v.errorType.Value = value
}

// SetMessage ...
func (v *validationError) SetMessage(message string) {
	v.errorType.Message = message
}

// Param ...
func (v *validationError) Param() string {
	return v.errorType.Param
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
	SetParam(param string)
	SetValue(value string)
	SetMessage(messag string)
	Param() string
	Value() string
	Message() string
}

func SetParam(param string) ValidationErrorSetterFunc {
	return func(valError ValidationError) {
		valError.SetParam(param)
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
