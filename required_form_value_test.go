package reqval

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredFormValue_ValidationSuccess(t *testing.T) {
	u := url.Values{}
	u.Add("foo", "moo")

	req := httptest.NewRequest("POST", "http://www.example.com", strings.NewReader(u.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = u.Encode()
	r := RequiredFormValue{}

	validationErrors, err := r.Validate(req, "foo")
	assert.NoError(t, err)
	assert.Empty(t, validationErrors)
}

func TestRequiredFormValue_ValidationFail(t *testing.T) {
	u := url.Values{}
	u.Add("foo", "moo")

	req := httptest.NewRequest("POST", "http://www.example.com", strings.NewReader(u.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = u.Encode()
	r := RequiredFormValue{}

	validationErrors, err := r.Validate(req, "zoo")
	assert.NoError(t, err)
	assert.NotEmpty(t, validationErrors)
}
