package router

import (
	"net/http"
	"path"

	"github.com/jkomyno/palitrux/config"
	"github.com/jkomyno/palitrux/controllers"
)

func NewServerMux(c *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(join("/stats", c), Middleware(controllers.StatsController, c))
	mux.Handle(join("/dominantColors", c), Middleware(controllers.DominantColorsController, c))

	return mux
}

func join(route string, c *config.Config) string {
	return path.Join(c.PathPrefix, route)
}
