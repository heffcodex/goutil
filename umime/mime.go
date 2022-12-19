package umime

import (
	"io"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
)

var ErrInvalidMIME = errors.New("invalid MIME")

func Validate(f io.Reader, allowedTypes ...string) (*mimetype.MIME, error) {
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

func ReplaceExt(filename string, mime *mimetype.MIME) string {
	ext := path.Ext(filename)
	validExt := mime.Extension()

	if validExt == ext {
		return filename
	}

	return strings.TrimSuffix(filename, ext) + validExt
}
