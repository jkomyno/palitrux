package controllers

import (
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"

	"github.com/jkomyno/palitra"
	"github.com/jkomyno/palitrux/bodyParse"
	"github.com/jkomyno/palitrux/errors"
)

func DominantColorsController(w http.ResponseWriter, r *http.Request) {
	img, errImg := bodyParse.ReadImageFromBody(r)
	if errImg != nil {
		errors.ReplyWithError(w, *errImg)
		return
	}

	limit, errLimit := strconv.Atoi(r.FormValue("limit"))
	if limit == 0 || errLimit != nil {
		errors.ReplyWithError(w, errors.ErrorMissingParam)
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
