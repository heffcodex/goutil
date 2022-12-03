package goutil

import (
	"io"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
)

var ErrInvalidMIME = errors.New("invalid MIME")

func MIMEValidate(f io.Reader, allowedTypes ...string) (*mimetype.MIME, error) {
	mime, err := mimetype.DetectReader(f)
	if err != nil {
		return nil, err
	}

	for _, t := range allowedTypes {
		if mime.Is(t) {
			return mime, nil
		}
	}

	return nil, ErrInvalidMIME
}

func MIMEReplaceExtension(filename string, mime *mimetype.MIME) string {
	ext := path.Ext(filename)
	validExt := mime.Extension()

	if validExt == ext {
		return filename
	}

	parts := strings.Split(filename, ".")

	if len(parts) < 2 {
		return filename + validExt
	}

	parts[len(parts)-1] = validExt

	return strings.Join(parts, ".")
}
