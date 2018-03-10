package reqval

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func multiPartForm(t *testing.T, files map[string]string, params map[string]string) ([]byte, *multipart.Writer) {
	body := &bytes.Buffer{}
	defer body.Reset()

	writer := multipart.NewWriter(body)

	for paramName, filePath := range files {
		file, err := os.Open(filePath)
		assert.NoError(t, err)
		defer file.Close()

		part, err := writer.CreateFormFile(paramName, filepath.Base(filePath))
		assert.NoError(t, err)

		_, err = io.Copy(part, file)
		assert.NoError(t, err)
	}

	for key, val := range params {
		err := writer.WriteField(key, val)
		assert.NoError(t, err)
	}

	err := writer.Close()
	assert.NoError(t, err)

	return body.Bytes(), writer
}

func createTempFile(t *testing.T, fileContent []byte, directory, dirPrefix, fileName string) (string, string) {
	dir, err := ioutil.TempDir(dirPrefix, directory)
	assert.NoError(t, err)

	tmpfn := filepath.Join(dir, fileName)
	err = ioutil.WriteFile(tmpfn, fileContent, 0666)
	assert.NoError(t, err)
	return dir, tmpfn
}
