package reqval

import (
	"bytes"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
