package umime

import (
	"strings"
	"testing"

	"github.com/gabriel-vasile/mimetype"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		mime, err := Validate(strings.NewReader("test"), "text/plain")
		require.NoError(t, err)
		require.Equal(t, "text/plain; charset=utf-8", mime.String())
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := Validate(strings.NewReader("test"), "image/jpeg")
		require.Error(t, err)
		require.Equal(t, ErrInvalidMIME, err)
	})
}

func TestReplaceExtension(t *testing.T) {
	require.Equal(t, "test.jpg", ReplaceExt("test.png", mimetype.Lookup("image/jpeg")))   // replace
	require.Equal(t, "test.jpg", ReplaceExt("test", mimetype.Lookup("image/jpeg")))       // append
	require.Equal(t, "test", ReplaceExt("test.png", mimetype.Lookup("application/tzif"))) // trim
}
