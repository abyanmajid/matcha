package matcha

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Ctx is a generic context struct that holds the HTTP request and response,
// along with a body of type Req. It is used to pass around the HTTP request
// and response objects, as well as any additional data needed for processing
// the request.
//
// Type Parameters:
//
//	Req - The type of the body data.
//
// Fields:
//
//	Request  - The HTTP request object.
//	Response - The HTTP response writer.
//	Body     - The body data of type Req.
type Ctx[Req any] struct {
	Request  *http.Request
	Response http.ResponseWriter
	Cookies  Cookies
	Body     Req
}

// Cookies abstraction
type Cookies struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// CookieOptions represents the options for setting a cookie.
// Path specifies the URL path that must exist in the requested URL for the browser to send the Cookie header.
// Domain specifies the domain for which the cookie is valid.
// MaxAge defines the maximum age of the cookie in seconds.
// Secure indicates whether the cookie should only be transmitted over secure protocols such as HTTPS.
// HttpOnly indicates whether the cookie is accessible only through the HTTP protocol and not through client-side scripts.
// SameSite controls whether a cookie is sent with cross-site requests, providing some protection against cross-site request forgery attacks.
// Expires specifies the expiration date of the cookie.
type CookieOptions struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
	Expires  time.Time
}

// GetHeader retrieves the value of a specific HTTP header.
func (c *Ctx[Req]) GetHeader(header string) string {
	return c.Request.Header.Get(header)
}

// GetQueryParam retrieves the value of a specific query parameter.
func (c *Ctx[Req]) GetQueryParam(queryParam string) string {
	return c.Request.URL.Query().Get(queryParam)
}

// GetPathParam retrieves a placeholder value from the URL path.
func (c *Ctx[Req]) GetPathParam(pathParam string) string {
	return chi.URLParam(c.Request, pathParam)
}

// GetQueryParamDefault retrieves a query parameter value or returns a default value if not found.
func (c *Ctx[Req]) GetQueryParamDefault(queryParam, defaultValue string) string {
	value := c.Request.URL.Query().Get(queryParam)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetPathParamDefault retrieves a path parameter value or returns a default value if not found.
func (c *Ctx[Req]) GetPathParamDefault(pathParam, defaultValue string) string {
	value := chi.URLParam(c.Request, pathParam)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetIP retrieves the client's IP address from the request.
func (c *Ctx[Req]) GetIP() string {
	xff := c.GetHeader("X-Forwarded-For")
	if xff != "" {
		return xff
	}
	return c.Request.RemoteAddr
}

// GetForm retrieves a form via their name
func (c *Ctx[Req]) GetForm(formName string) string {
	return c.Request.FormValue(formName)
}

// GetFile retrieves a file from a multipart form request.
func (c *Ctx[Req]) GetFile(fieldName string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(fieldName)
}

// Redirect redirects the client to a given URL
func (c *Ctx[Req]) Redirect(statusCode int, url string) {
	http.Redirect(c.Response, c.Request, url, statusCode)
}

// Json writes a JSON response to the client with the specified status code.
func (c *Ctx[Req]) Json(statusCode int, data any) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(statusCode)
	return json.NewEncoder(c.Response).Encode(data)
}

// Text writes a plain text response to the client with the specified status code.
func (c *Ctx[Req]) Text(statusCode int, message string) {
	c.Response.Header().Set("Content-Type", "text/plain")
	c.Response.WriteHeader(statusCode)
	c.Response.Write([]byte(message))
}

// GetCookie retrieves the value of a specific cookie.
func (co *Cookies) GetCookie(name string) (string, error) {
	cookie, err := co.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// SetCookie sets an HTTP cookie in the response using flexible options.
func (co *Cookies) SetCookie(name, value string, options *CookieOptions) {
	cookie := &http.Cookie{
		Name:  name,
		Value: value,
	}

	if options != nil {
		if options.Path != "" {
			cookie.Path = options.Path
		}
		if options.Domain != "" {
			cookie.Domain = options.Domain
		}
		if options.MaxAge != 0 {
			cookie.MaxAge = options.MaxAge
		}
		cookie.Secure = options.Secure
		cookie.HttpOnly = options.HttpOnly
		if !options.Expires.IsZero() {
			cookie.Expires = options.Expires
		}
		cookie.SameSite = options.SameSite
	}

	http.SetCookie(co.Response, cookie)
}
