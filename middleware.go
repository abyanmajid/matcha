package matcha

import (
	"net/http"

	"github.com/abyanmajid/matcha/internal/middleware"
)

type MiddlewareType struct {
	BasicAuth func(realm string, creds map[string]string) func(next http.Handler) http.Handler
}

var Middleware = MiddlewareType{
	BasicAuth: middleware.BasicAuth,
}
