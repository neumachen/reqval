package reqval

import (
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchRegex_Validate(t *testing.T) {
	t.Run("does not match regex pattern", func(t *testing.T) {
		tests := []string{
			"01",
			"1",
			"99999999",
			"099999999",
		}

		pattern, err := regexp.Compile("^[a-z]$")
		require.NoError(t, err)

		for _, test := range tests {
			u := url.Values{}
			u.Add("field", test)

			req := httptest.NewRequest("POST", "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			r := MatchRegexPattern{
				RegexPattern: pattern,
			}

			validationErrors, err := r.Validate(req, "field")
			require.NoError(t, err)
			require.NotEmpty(t, validationErrors)
		}
	})

	t.Run("match regex pattern", func(t *testing.T) {
		tests := []string{
			"a",
			"aaaaa",
			"aaaaaa",
			"manbearpig",
		}

		pattern, err := regexp.Compile("^[a-z]+$")
		require.NoError(t, err)

		for _, test := range tests {
			u := url.Values{}
			u.Add("field", test)

			req := httptest.NewRequest("POST", "http://www.example.com", nil)
			req.URL.RawQuery = u.Encode()
			r := MatchRegexPattern{
				RegexPattern: pattern,
			}

			validationErrors, err := r.Validate(req, "field")
			require.NoError(t, err)
			require.Empty(t, validationErrors)
		}
	})
}
