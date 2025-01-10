package matcha

import (
	"net/http"
	"time"

	"github.com/abyanmajid/matcha/internal"
	"github.com/abyanmajid/matcha/internal/middleware"
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
	SupressNotFound      func(router *internal.Mux) func(next http.Handler) http.Handler
	Throttle             func(limit int) func(http.Handler) http.Handler
	Timeout              func(timeout time.Duration) func(next http.Handler) http.Handler
	URLFormat            func(next http.Handler) http.Handler
}

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
}
