package reqval

import (
	"fmt"
	"net/http"
)

// ArrayContainsStr ...
func ArrayContainsStr(strSlice []string, searchString string) bool {
	for _, value := range strSlice {
		if value == searchString {
			return true
		}
	}
	return false
}

var isOneOfStringsMessageFunc = func(field string, strs []string) string {
	return fmt.Sprintf("query param: %s is not one of %v", field, strs)
}

type IsOneOfStrings struct {
	MessageFunc    func(field string, strs []string) string
	AllowedStrings []string
}

func (i *IsOneOfStrings) Validate(req *http.Request, field string) (ValidationErrors, error) {
	if len(i.AllowedStrings) == 0 {
		return nil, fmt.Errorf("failed to validate query field: %s, allowed strings is empty", field)
	}

	fieldValues := req.URL.Query()[field]
	if len(fieldValues) < 1 {
		return nil, nil
	}
	if IsNil(i.MessageFunc) {
		i.MessageFunc = isOneOfStringsMessageFunc
	}

	validationErrors := make(ValidationErrors, 0)
	for _, fieldValue := range fieldValues {
		if !ArrayContainsStr(i.AllowedStrings, fieldValue) {
			validationErrors.Append(NewValidationError(
				SetField(field),
				SetMessage(i.MessageFunc(field, i.AllowedStrings)),
			))
		}
	}

	if validationErrors.GetLength() == 0 {
		return nil, nil
	}
	return validationErrors, nil
}
