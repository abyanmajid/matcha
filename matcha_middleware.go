package matcha

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type MiddlewareType struct {
	BasicAuth            func(realm string, creds map[string]string) func(next http.Handler) http.Handler
	Logger               func(next http.Handler) http.Handler
	CleanPath            func(next http.Handler) http.Handler
	Compress             func(level int, types ...string) func(next http.Handler) http.Handler
	ContentCharset       func(charsets ...string) func(next http.Handler) http.Handler
	AllowContentEncoding func(contentEncoding ...string) func(next http.Handler) http.Handler
	AllowContentType     func(contentTypes ...string) func(http.Handler) http.Handler
	GetHead              func(next http.Handler) http.Handler
	Heartbeat            func(endpoint string) func(http.Handler) http.Handler
	NoCache              func(h http.Handler) http.Handler
	RealIP               func(h http.Handler) http.Handler
	Recoverer            func(next http.Handler) http.Handler
	RequestID            func(next http.Handler) http.Handler
	RequestSize          func(bytes int64) func(http.Handler) http.Handler
	StripSlahses         func(next http.Handler) http.Handler
	RedirectSlashes      func(next http.Handler) http.Handler
	StripPrefix          func(prefix string) func(http.Handler) http.Handler
	Sunset               func(sunsetAt time.Time, links ...string) func(http.Handler) http.Handler
	SupressNotFound      func(router *chi.Mux) func(next http.Handler) http.Handler
	Throttle             func(limit int) func(http.Handler) http.Handler
	Timeout              func(timeout time.Duration) func(next http.Handler) http.Handler
	URLFormat            func(next http.Handler) http.Handler
	Cors                 func(options cors.Options) func(next http.Handler) http.Handler
}

// Middleware is a collection of middleware functions used in the application.
// It includes the following middleware:
// - BasicAuth: middleware for HTTP Basic Authentication.
// - Logger: middleware for logging HTTP requests and responses.
// - CleanPath: middleware for cleaning up the request URL path.
// - Compress: middleware for compressing HTTP responses.
// - ContentCharset: middleware for setting the Content-Type charset.
// - AllowContentEncoding: middleware for allowing specific content encodings.
// - AllowContentType: middleware for allowing specific content types.
// - GetHead: middleware for handling GET and HEAD requests.
// - Heartbeat: middleware for adding a heartbeat endpoint.
// - NoCache: middleware for preventing caching of HTTP responses.
// - Recoverer: middleware for recovering from panics.
// - RequestID: middleware for generating and setting a request ID.
// - RequestSize: middleware for limiting the size of HTTP requests.
// - StripSlashes: middleware for stripping slashes from the URL path.
// - RedirectSlashes: middleware for redirecting requests with trailing slashes.
// - StripPrefix: middleware for stripping a prefix from the URL path.
// - Sunset: middleware for handling the Sunset HTTP header.
// - SupressNotFound: middleware for suppressing 404 Not Found responses.
// - Throttle: middleware for throttling HTTP requests.
// - Timeout: middleware for setting a timeout for HTTP requests.
// - URLFormat: middleware for handling URL format.
// - Cors: middleware for controlling CORS
var Middleware = MiddlewareType{
	BasicAuth:            middleware.BasicAuth,
	Logger:               middleware.Logger,
	CleanPath:            middleware.CleanPath,
	Compress:             middleware.Compress,
	ContentCharset:       middleware.ContentCharset,
	AllowContentEncoding: middleware.AllowContentEncoding,
	AllowContentType:     middleware.AllowContentType,
	GetHead:              middleware.GetHead,
	Heartbeat:            middleware.Heartbeat,
	NoCache:              middleware.NoCache,
	Recoverer:            middleware.Recoverer,
	RequestID:            middleware.RequestID,
	RequestSize:          middleware.RequestSize,
	StripSlahses:         middleware.StripSlashes,
	RedirectSlashes:      middleware.RedirectSlashes,
	StripPrefix:          middleware.StripPrefix,
	Sunset:               middleware.Sunset,
	SupressNotFound:      middleware.SupressNotFound,
	Throttle:             middleware.Throttle,
	Timeout:              middleware.Timeout,
	URLFormat:            middleware.URLFormat,
	Cors:                 cors.Handler,
}
