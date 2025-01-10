package middleware

import (
	"net/http"
	"path"

	"github.com/abyanmajid/matcha/internal"
)

func CleanPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := internal.RouteContext(r.Context())

		routePath := rctx.RoutePath
		if routePath == "" {
			if r.URL.RawPath != "" {
				routePath = r.URL.RawPath
			} else {
				routePath = r.URL.Path
			}
			rctx.RoutePath = path.Clean(routePath)
		}

		next.ServeHTTP(w, r)
	})
}
