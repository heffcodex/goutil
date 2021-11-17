package goutil

import (
	"io"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

var MIMEPDF = []string{"application/pdf"}

func ValidateMIMEType(f io.Reader, allowedTypes ...string) (*mimetype.MIME, error) {
	mt, err := mimetype.DetectReader(f)
	if err != nil {
		return nil, err
	}

	for _, t := range allowedTypes {
		if mt.Is(t) {
			return mt, nil
		}
	}

	return nil, nil
}

func ReplaceMIMEExt(name string, mime *mimetype.MIME) string {
	ext := path.Ext(name)
	validExt := mime.Extension()

	if validExt == ext {
		return name
	}

	parts := strings.Split(name, ".")

	if len(parts) < 2 {
		return name + validExt
	}

	parts[len(parts)-1] = validExt

	return strings.Join(parts, ".")
}
