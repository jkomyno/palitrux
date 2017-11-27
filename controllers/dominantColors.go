package controllers

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"
	"strings"

	"github.com/jkomyno/palitra"
	"github.com/jkomyno/palitrux/errors"
)

func DominantColorsController(w http.ResponseWriter, r *http.Request) {
	if !IsMultipart(r) {
		errors.ReplyWithError(w, errors.ErrorContentType)
		return
	}

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		fmt.Println("err", err)
		errors.ReplyWithError(w, errors.ErrorFileSizeExceeded)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		errors.ReplyWithError(w, errors.ErrorMissingParam)
		return
	}
	defer file.Close()

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if limit == 0 || err != nil {
		errors.ReplyWithError(w, errors.ErrorMissingParam)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		errors.ReplyWithError(w, errors.ErrorMIMETypeUnsupported)
		return
	}

	colorPercentageChan := make(chan []palitra.ColorPercentageT, 1)

	go func(colorPercentageChan chan []palitra.ColorPercentageT) {
		palette := palitra.GetPalette(img, limit, 75)
		colorPercentageChan <- palette
	}(colorPercentageChan)

	palette := <-colorPercentageChan

	ReplyWithJSON(w, palette)
}

// IsMultipart checks if Content-Type is multipart/form-data
func IsMultipart(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data")
}
