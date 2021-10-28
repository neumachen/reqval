package reqval

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequiredQueryValue_Validate(t *testing.T) {
	t.Run("validation success", func(t *testing.T) {
		u := url.Values{}
		u.Add("foo", "moo")

		req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		r := RequiredQueryValue{}

		validationErrors, err := r.Validate(req, "foo")
		require.NoError(t, err)
		require.Empty(t, validationErrors)
	})

	t.Run("validation failed", func(t *testing.T) {
		u := url.Values{}
		u.Add("boo", "moo")

		req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		r := RequiredQueryValue{}

		validationErrors, err := r.Validate(req, "foo")
		require.NoError(t, err)
		require.NotEmpty(t, validationErrors)
	})

	t.Run("validation failed - multiple", func(t *testing.T) {
		u := url.Values{}
		u.Add("boo", "moo")
		u.Add("boo", "moo")
		u.Add("boo", "")
		u.Add("boo", "moo")
		u.Add("boo", "")

		req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		r := RequiredQueryValue{}

		validationErrors, err := r.Validate(req, "boo")
		require.NoError(t, err)
		require.NotEmpty(t, validationErrors)
	})

	t.Run("validation failed - custom message", func(t *testing.T) {
		u := url.Values{}
		u.Add("boo", "moo")

		req := httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
		req.URL.RawQuery = u.Encode()
		r := RequiredQueryValue{MessageFunc: func(field string) string {
			return "boo"
		}}

		validationErrors, err := r.Validate(req, "foo")
		require.NoError(t, err)
		require.NotEmpty(t, validationErrors)
		require.Equal(t, validationErrors[0].Message(), "boo")
	})
}
