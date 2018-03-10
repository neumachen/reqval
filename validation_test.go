package reqval

import (
	"bytes"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	validations := RequestValidations{
		"foo_id":      RequestValidators{&RequiredFormValue{}},
		"test_file_1": RequestValidators{&RequiredFormFile{}},
		"test_file_2": RequestValidators{&RequiredFormFile{}},
	}
	tmpDir1, tmpFile1 := createTempFile(t, []byte(`test`), "", "", "test.txt")
	tmpDir1, tmpFile2 := createTempFile(t, []byte(`test`), "", "", "test.txt")

	defer os.RemoveAll(tmpDir1)

	files := make(map[string]string)
	files["test_file_1"] = tmpFile1
	files["test_file_2"] = tmpFile2

	formFields := map[string]string{
		"foo_id": "1",
		"boo_id": "2",
	}

	form, writer := multiPartForm(t, files, formFields)

	req := httptest.NewRequest("POST", "http://www.example.com", bytes.NewReader(form))
	req.Header.Add("Content-Type", writer.FormDataContentType())
	valErrs, err := Validate(req, validations)
	assert.NoError(t, err)
	assert.Empty(t, valErrs)

}
