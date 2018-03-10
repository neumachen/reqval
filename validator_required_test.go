package reqval

import (
	"bytes"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
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

func TestRequiredFormFile_ValidationSuccess(t *testing.T) {
	tmpDir1, tmpFile1 := createTempFile(t, []byte(`test`), "", "", "test.txt")

	defer os.RemoveAll(tmpDir1)

	files := make(map[string]string)
	files["momo"] = tmpFile1

	form, writer := multiPartForm(t, files, nil)

	req := httptest.NewRequest("POST", "http://www.example.com", bytes.NewReader(form))
	req.Header.Add("Content-Type", writer.FormDataContentType())
	r := RequiredFormFile{}

	validationErrors, err := r.Validate(req, "momo")
	assert.NoError(t, err)
	assert.Empty(t, validationErrors)
}

func TestRequiredFormFile_ValidationFail(t *testing.T) {
	tmpDir1, tmpFile1 := createTempFile(t, []byte(`test`), "", "", "test.txt")

	defer os.RemoveAll(tmpDir1)

	files := make(map[string]string)
	files["momo"] = tmpFile1

	form, writer := multiPartForm(t, files, nil)

	req := httptest.NewRequest("POST", "http://www.example.com", bytes.NewReader(form))
	req.Header.Add("Content-Type", writer.FormDataContentType())
	r := RequiredFormFile{}

	validationErrors, err := r.Validate(req, "file")
	assert.NoError(t, err)
	assert.NotEmpty(t, validationErrors)
}
