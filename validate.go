package qryval

import "net/http"

// Validator ...
type Validator func(value, message string) (ValidationError, error)

// Validators ...
type Validators []Validator

// Validations ...
type Validations map[string]Validators

// Validate ...
func Validate(req *http.Request, validations Validations) (ValidationErrors, error) {
	qry := req.URL.Query()
	valErrs := make(ValidationErrors, 0)
	for field, validators := range validations {
		vv := qry[field]
		for _, validator := range validators {
			if vv == nil || len(vv) == 0 {
				valErr, err := validator("", "")
				if err != nil {
					return nil, err
				}
				if valErr != nil {
					valErrs = append(valErrs, valErr)
				}
				continue
			}
			if len(vv) == 1 {
				valErr, err := validator(vv[0], "")
				if err != nil {
					return nil, err
				}
				if valErr != nil {
					valErrs = append(valErrs, valErr)
				}
				continue
			}
			for _, v := range vv {
				valErr, err := validator(v, "")
				if err != nil {
					return nil, err
				}
				if valErr != nil {
					valErr.SetField(field)
					valErrs = append(valErrs, valErr)
				}
			}
		}
	}
	if len(valErrs) == 0 {
		return nil, nil
	}
	return valErrs, nil
}
