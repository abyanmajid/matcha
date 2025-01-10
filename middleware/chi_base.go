package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type CorsOptions cors.Options

func AllowContentEncoding() func(next http.Handler) http.Handler {
	return middleware.AllowContentEncoding()
}

func AllowContentType(contentTypes ...string) func(http.Handler) http.Handler {
	return middleware.AllowContentType(contentTypes...)
}

func CleanPath(next http.Handler) http.Handler {
	return middleware.CleanPath(next)
}

func Compress(level int, types ...string) func(next http.Handler) http.Handler {
	return middleware.Compress(level, types...)
}

func ContentCharset(charset ...string) func(next http.Handler) http.Handler {
	return middleware.ContentCharset(charset...)
}

func Cors(opts CorsOptions) func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options(opts))
}

func GetHead(next http.Handler) http.Handler {
	return middleware.GetHead(next)
}

func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func NoCache(h http.Handler) http.Handler {
	return middleware.NoCache(h)
}

func RealIP(h http.Handler) http.Handler {
	return middleware.RealIP(h)
}

func Recoverer(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}
