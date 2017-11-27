package router

import (
	"net/http"

	"github.com/jkomyno/palitrux/config"
	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
)

func handleLimiterError(err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "rate limiter error: "+err.Error(), http.StatusInternalServerError)
	})
}

func throttleLimiter(next http.Handler, c *config.Config) http.Handler {
	store, err := memstore.New(65536) // 32 bytes store
	if err != nil {
		return handleLimiterError(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerSec(c.HTTPRateLimit),
		MaxBurst: c.HTTPBurst,
	}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		return handleLimiterError(err)
	}

	// called for each request
	httpRateLimiter := throttled.HTTPRateLimiter{
		// cdetermines whether the request is permitted
		// and updates internal state.
		RateLimiter: rateLimiter,
		// generates a key for the limiter
		VaryBy: &throttled.VaryBy{
			Method: true,
		},
	}

	return httpRateLimiter.RateLimit(next)
}
