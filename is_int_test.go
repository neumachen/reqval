package reqval

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInt_ValidationSuccess(t *testing.T) {
	u := url.Values{}
	u.Add("int", "1")

	req := httptest.NewRequest("POST", "http://www.example.com", nil)
	req.URL.RawQuery = u.Encode()
	r := IsInt{}

	validationErrors, err := r.Validate(req, "int")
	assert.NoError(t, err)
	assert.Empty(t, validationErrors)
}
