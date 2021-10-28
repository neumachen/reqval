package reqval

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsOneOfStrings_Validate(t *testing.T) {
	allowedStrings := []string{"a", "b", "c"}
	t.Run("validation success", func(t *testing.T) {
		tests := allowedStrings

		allowedStrings := []string{"a", "b", "c"}

		for _, test := range tests {
			u := url.Values{}
			u.Add("str", test)

			req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			validator := IsOneOfStrings{
				AllowedStrings: allowedStrings,
			}

			validationErrors, err := validator.Validate(req, "str")
			require.NoError(t, err)
			require.Empty(t, validationErrors)
		}
	})

	t.Run("validation error - no allowed strings", func(t *testing.T) {
		tests := allowedStrings

		for _, test := range tests {
			u := url.Values{}
			u.Add("str", test)

			req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			validator := IsOneOfStrings{}

			validationErrors, err := validator.Validate(req, "str")
			require.Error(t, err)
			require.Empty(t, validationErrors)
		}
	})

	t.Run("validation failed", func(t *testing.T) {
		tests := []string{
			"z",
			"moo",
			"too",
		}

		for _, test := range tests {
			u := url.Values{}
			u.Add("str", test)

			req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			validator := IsOneOfStrings{
				AllowedStrings: allowedStrings,
			}

			validationErrors, err := validator.Validate(req, "str")
			require.NoError(t, err)
			require.NotEmpty(t, validationErrors)
		}
	})

	t.Run("validation failed - multiple", func(t *testing.T) {
		tests := allowedStrings
		tests = append(tests, "zoo", "loo", "koo")

		u := url.Values{}
		for _, test := range tests {
			u.Add("str", test)
		}

		req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		validator := IsOneOfStrings{
			AllowedStrings: allowedStrings,
		}

		validationErrors, err := validator.Validate(req, "str")
		require.NoError(t, err)
		require.NotEmpty(t, validationErrors)
	})
}
