package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// CorsOptions is a type alias for cors.Options.
type CorsOptions cors.Options

// AllowContentEncoding is a middleware that adds support for handling specific content encoding types
// in HTTP requests. It ensures the allowed encodings are respected in the request.
func AllowContentEncoding() func(next http.Handler) http.Handler {
	return middleware.AllowContentEncoding()
}

// AllowContentType is a middleware that allows specific content types for HTTP requests.
// It ensures that only the specified content types are accepted by the server.
func AllowContentType(contentTypes ...string) func(http.Handler) http.Handler {
	return middleware.AllowContentType(contentTypes...)
}

// CleanPath is a middleware that cleans the path of the HTTP request by removing redundant slashes,
// which could be helpful for routing and consistency.
func CleanPath(next http.Handler) http.Handler {
	return middleware.CleanPath(next)
}

// Compress is a middleware that compresses HTTP responses.
// The `level` determines the compression level (higher is better compression but slower),
// and `types` specifies the content types to apply compression to.
func Compress(level int, types ...string) func(next http.Handler) http.Handler {
	return middleware.Compress(level, types...)
}

// ContentCharset is a middleware that sets the content charset for the response.
// It allows specifying the charsets that the server will accept for content encoding.
func ContentCharset(charset ...string) func(next http.Handler) http.Handler {
	return middleware.ContentCharset(charset...)
}

// Cors is a middleware that handles Cross-Origin Resource Sharing (CORS) for HTTP requests.
// It configures the allowed origins, methods, and other CORS-related headers.
func Cors(opts CorsOptions) func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options(opts))
}

// GetHead is a middleware that ensures HTTP GET and HEAD requests are handled correctly.
func GetHead(next http.Handler) http.Handler {
	return middleware.GetHead(next)
}

// Logger is a middleware that logs HTTP requests and their associated responses, including status codes,
// the request method, and request duration. It is useful for debugging and monitoring.
func Logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

// NoCache is a middleware that prevents HTTP responses from being cached. It ensures that responses
// are always fetched fresh from the server without using a cached version.
func NoCache(h http.Handler) http.Handler {
	return middleware.NoCache(h)
}

// RealIP is a middleware that extracts the real IP address of the client from the request headers.
// It is particularly useful when your server is behind a reverse proxy or load balancer.
func RealIP(h http.Handler) http.Handler {
	return middleware.RealIP(h)
}

// Recoverer is a middleware that recovers from any panics during the request handling process,
// returning a 500 Internal Server Error and logging the panic. It ensures that the application doesn't
// crash due to unexpected errors.
func Recoverer(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}

// RedirectSlashes is a middleware that automatically redirects requests with a trailing slash to the
// same URL without the slash. This ensures consistent routing and avoids unnecessary duplicate paths.
func RedirectSlashes(next http.Handler) http.Handler {
	return middleware.RedirectSlashes(next)
}

// StripSlashes is a middleware that removes a specified prefix from the URL path of incoming requests.
// This is useful when you want to remove a common prefix for routing.
func StripSlashes(prefix string) func(http.Handler) http.Handler {
	return middleware.StripPrefix(prefix)
}

// Throttle is a middleware that limits the number of requests that can be handled per second.
// This helps prevent abuse and protects the server from being overwhelmed by too many requests.
func Throttle(limit int) func(http.Handler) http.Handler {
	return middleware.Throttle(limit)
}

// ThrottleBacklog is a middleware that implements throttling with a backlog of requests.
// It limits the number of requests that can be handled in a given time window, and if the limit is reached,
// it will queue the requests for a backlog time before rejecting them.
func ThrottleBacklog(limit int, backlogLimit int, backlogTimeout time.Duration) func(http.Handler) http.Handler {
	return middleware.ThrottleBacklog(limit, backlogLimit, backlogTimeout)
}

// Timeout is a middleware that sets a timeout for HTTP requests. If a request takes longer than the
// specified `timeout` duration, the server will respond with a 408 Request Timeout status.
func Timeout(timeout time.Duration) func(next http.Handler) http.Handler {
	return middleware.Timeout(timeout)
}

// RouteHeaders is a middleware that allows for setting custom HTTP headers on the route level.
// This can be used for setting specific headers for certain routes, like security or caching headers.
func RouteHeaders() middleware.HeaderRouter {
	return middleware.RouteHeaders()
}

// All is a method on Ratelimit that applies global rate-limiting to all incoming requests. It limits
// the number of requests that can be made by any client in a specific time window.
func RateLimitAll(requestLimit int, windowLength time.Duration) func(next http.Handler) http.Handler {
	return httprate.LimitAll(requestLimit, windowLength)
}

// ByIP is a method on Ratelimit that applies rate-limiting based on the client's IP address. It limits
// the number of requests a single IP can make in a specific time window.
func RateLimitByIP(requestLimit int, windowLength time.Duration) func(next http.Handler) http.Handler {
	return httprate.LimitByIP(requestLimit, windowLength)
}

// ByRealIP is a method on Ratelimit that applies rate-limiting based on the real client IP (e.g., after
// handling reverse proxies). It limits the number of requests made from a specific real IP address.
func RateLimitByRealIP(requestLimit int, windowLength time.Duration) func(next http.Handler) http.Handler {
	return httprate.LimitByRealIP(requestLimit, windowLength)
}
