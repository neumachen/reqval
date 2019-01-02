package reqval

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredQueryValueValidate_ValidationSuccess(t *testing.T) {
	u := url.Values{}
	u.Add("foo", "moo")

	req := httptest.NewRequest("POST", "http://www.example.com", nil)
	req.URL.RawQuery = u.Encode()
	r := RequiredQueryValue{}

	validationErrors, err := r.Validate(req, "foo")
	assert.NoError(t, err)
	assert.Empty(t, validationErrors)
}

func TestRequiredQueryValueValidate_ValidationFail(t *testing.T) {
	u := url.Values{}
	u.Add("boo", "moo")

	req := httptest.NewRequest("POST", "http://www.example.com", nil)
	req.URL.RawQuery = u.Encode()
	r := RequiredQueryValue{}

	validationErrors, err := r.Validate(req, "foo")
	assert.NoError(t, err)
	assert.NotEmpty(t, validationErrors)
}

func TestRequiredQueryValueValidate_ValidationFail_Multiple(t *testing.T) {
	u := url.Values{}
	u.Add("boo", "moo")
	u.Add("boo", "moo")
	u.Add("boo", "")
	u.Add("boo", "moo")
	u.Add("boo", "")

	req := httptest.NewRequest("POST", "http://www.example.com", nil)
	req.URL.RawQuery = u.Encode()
	r := RequiredQueryValue{}

	validationErrors, err := r.Validate(req, "boo")
	assert.NoError(t, err)
	assert.NotEmpty(t, validationErrors)
}

func TestRequiredQueryValueValidate_ValidationFail_CustomMessage(t *testing.T) {
	u := url.Values{}
	u.Add("boo", "moo")

	req := httptest.NewRequest("POST", "http://www.example.com", nil)
	req.URL.RawQuery = u.Encode()
	r := RequiredQueryValue{Message: "boo"}

	validationErrors, err := r.Validate(req, "foo")
	assert.NoError(t, err)
	assert.NotEmpty(t, validationErrors)
	assert.Equal(t, validationErrors[0].Message(), "boo")
}
