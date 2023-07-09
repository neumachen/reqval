package reqval

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsFloat_Validate(t *testing.T) {
	t.Run("validation successful", func(t *testing.T) {
		tests := []string{
			"1.1",
			"-2.3",
			"1.5555555",
			"-9.1233444",
			"132.3423345",
			"24.12941",
			"35.84260",
		}

		for _, test := range tests {
			u := url.Values{}
			u.Add("int", test)

			req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			r := IsFloat{}

			validationErrors, err := r.Validate(req, "int")
			require.NoError(t, err)
			require.Empty(t, validationErrors)
		}
	})

	t.Run("validation failed", func(t *testing.T) {
		tests := []string{
			"01",
			"1",
			"99999999",
			"099999999",
			"a",
			"c",
			"d",
			"0A3",
		}

		for _, test := range tests {
			u := url.Values{}
			u.Add("int", test)

			req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			r := IsFloat{}

			validationErrors, err := r.Validate(req, "int")
			require.NoError(t, err)
			require.NotEmpty(t, validationErrors)
		}
	})
}
