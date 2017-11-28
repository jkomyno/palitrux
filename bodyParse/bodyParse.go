package bodyParse

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/jkomyno/palitrux/errors"
)

const maxMemory int64 = 5 << 20 // 5 MB

// isMultipart checks if Content-Type is multipart/*
func isMultipart(r *http.Request) bool {
	return strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/")
}

// readFileFromBody extracts a file from a multipart request
func readFileFromBody(r *http.Request) (multipart.File, *errors.Error) {
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return nil, &errors.ErrorFileSizeExceeded
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, &errors.ErrorMissingParam
	}
	defer file.Close()

	return file, nil
}

// ReadImageFromBody reads an image from a request
func ReadImageFromBody(r *http.Request) (image.Image, *errors.Error) {
	if !isMultipart(r) {
		return nil, &errors.ErrorContentType
	}

	file, errFile := readFileFromBody(r)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	img, _, errImg := image.Decode(file)
	if errImg != nil {
		return nil, &errors.ErrorMIMETypeUnsupported
	}

	return img, nil
}
