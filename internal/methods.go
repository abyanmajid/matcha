package internal

type HTTPMethod string

const (
	Get     HTTPMethod = "GET"
	Post    HTTPMethod = "POST"
	Put     HTTPMethod = "PUT"
	Delete  HTTPMethod = "DELETE"
	Patch   HTTPMethod = "PATCH"
	Options HTTPMethod = "OPTIONS"
	Head    HTTPMethod = "HEAD"
	Trace   HTTPMethod = "TRACE"
	Connect HTTPMethod = "CONNECT"
)

func IsValidHTTPMethod(method HTTPMethod) bool {
	switch method {
	case Get, Post, Put, Delete, Patch, Options, Head, Trace, Connect:
		return true
	default:
		return false
	}
}
