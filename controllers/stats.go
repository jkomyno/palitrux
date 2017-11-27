package controllers

import (
	"net/http"

	"github.com/jkomyno/palitrux/stats"
)

func StatsController(w http.ResponseWriter, r *http.Request) {
	res := stats.GetRuntimeStats()
	ReplyWithJSON(w, res)
}
