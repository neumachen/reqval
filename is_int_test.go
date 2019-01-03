package reqval

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInt_Validate_Success(t *testing.T) {
	tests := []string{
		"01",
		"1",
		"99999999",
		"099999999",
	}

	for _, test := range tests {
		u := url.Values{}
		u.Add("int", test)

		req := httptest.NewRequest("POST", "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		r := IsInt{}

		validationErrors, err := r.Validate(req, "int")
		assert.NoError(t, err)
		assert.Empty(t, validationErrors)
	}
}
