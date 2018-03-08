package qryval

const requiredValMessage = "This is a required parameter"

// Required checks whether a value is empty or not by determining it's length
func Required(value, message string) (ValidationError, error) {
	if len(value) > 0 {
		return nil, nil
	}
	if message == "" {
		message = requiredValMessage
	}
	return NewValidationError("", value, message), nil
}
