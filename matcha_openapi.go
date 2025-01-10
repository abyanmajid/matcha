package matcha

import (
	"net/http"

	"github.com/abyanmajid/matcha/openapi"
	"github.com/go-chi/chi/v5"
)

type MatchaOpenAPI struct {
	matcha *Matcha
	docs   openapi.OpenAPIDocs
}

func NewOpenAPI() *MatchaOpenAPI {
	return &MatchaOpenAPI{
		matcha: New(),
	}
}

func (r *MatchaOpenAPI) Documentation(pattern string, metadata openapi.Metadata) {
	r.docs = openapi.NewDocs(metadata)
	openAPIHandler := openapi.NewHandler(r.docs)

	r.matcha.Get(pattern, openAPIHandler)
}

// Serve mux on a given local address
func (r *MatchaOpenAPI) Serve(addr string) {
	r.matcha.Serve(addr)
}

// Use appends a middleware handler to the Mux middleware stack.
//
// The middleware stack for any Mux will execute before searching for a matching
// route to a specific handler, which provides opportunity to respond early,
// change the course of the request execution, or set request-scoped values for
// the next http.Handler.
func (r *MatchaOpenAPI) Use(middlewares ...func(http.Handler) http.Handler) {
	r.matcha.Use(middlewares...)
}

// Handle adds the route `pattern` that matches any http method to
// execute the `handler` http.Handler.
func (r *MatchaOpenAPI) Handle(pattern string, handler http.Handler) {
	r.matcha.Handle(pattern, handler)
}

// HandleFunc adds the route `pattern` that matches any http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.HandleFunc(pattern, handlerFn)
}

// Method adds the route `pattern` that matches `method` http method to
// execute the `handler` http.Handler.
func (r *MatchaOpenAPI) Method(method, pattern string, handler http.Handler) {
	r.matcha.Method(method, pattern, handler)
}

// MethodFunc adds the route `pattern` that matches `method` http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) MethodFunc(method, pattern string, handlerFn http.HandlerFunc) {
	r.matcha.MethodFunc(method, pattern, handlerFn)
}

// Connect adds the route `pattern` that matches a CONNECT http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Connect(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Connect(pattern, handlerFn)
}

// Delete adds the route `pattern` that matches a DELETE http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Delete(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Delete(pattern, handlerFn)
}

// Get adds the route `pattern` that matches a GET http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Get(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Get(pattern, handlerFn)
}

// Head adds the route `pattern` that matches a HEAD http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Head(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Head(pattern, handlerFn)
}

// Options adds the route `pattern` that matches an OPTIONS http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Options(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Options(pattern, handlerFn)
}

// Patch adds the route `pattern` that matches a PATCH http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Patch(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Patch(pattern, handlerFn)
}

// Post adds the route `pattern` that matches a POST http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Post(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Post(pattern, handlerFn)
}

// Put adds the route `pattern` that matches a PUT http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Put(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Put(pattern, handlerFn)
}

// Trace adds the route `pattern` that matches a TRACE http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *MatchaOpenAPI) Trace(pattern string, handlerFn http.HandlerFunc) {
	r.matcha.Trace(pattern, handlerFn)
}

// NotFound sets a custom http.HandlerFunc for routing paths that could
// not be found. The default 404 handler is `http.NotFound`.
func (r *MatchaOpenAPI) NotFound(handlerFn http.HandlerFunc) {
	r.matcha.NotFound(handlerFn)
}

// NotFoundJSON sets a custom handler for returning JSON responses when paths
// cannot be found. The default response returns a 404 status with a JSON body.
func (r *MatchaOpenAPI) NotFoundJSON() {
	r.matcha.NotFoundJSON()
}

// MethodNotAllowed sets a custom http.HandlerFunc for routing paths where the
// method is unresolved. The default handler returns a 405 with an empty body.
func (r *MatchaOpenAPI) MethodNotAllowed(handlerFn http.HandlerFunc) {
	r.matcha.MethodNotAllowed(handlerFn)
}

// MethodNotAllowedJSON sets a custom handler for returning JSON responses when
// a method is not allowed. The default response returns a 405 status with a JSON body.
func (r *MatchaOpenAPI) MethodNotAllowedJSON() {
	r.matcha.MethodNotAllowedJSON()
}

// With adds inline middlewares for an endpoint handler.
func (r *MatchaOpenAPI) With(middlewares ...func(http.Handler) http.Handler) *MatchaOpenAPI {
	r.matcha.With(middlewares...)
	return r
}

// Group creates a new inline-Mux with a copy of middleware stack. It's useful
// for a group of handlers along the same routing path that use an additional
// set of middlewares. See _examples/.
func (r *MatchaOpenAPI) Group(fn func(r Matcha)) MatchaOpenAPI {
	subRouter := &Matcha{
		mux: chi.NewRouter(),
	}
	fn(*subRouter)
	r.matcha.Mount("/", subRouter.mux)
	return MatchaOpenAPI{
		matcha: subRouter,
		docs:   r.docs,
	}
}

// Mount attaches another http.Handler or chi Router as a subrouter along a routing
// path. It's very useful to split up a large API as many independent routers and
// compose them as a single service using Mount. See _examples/.
//
// Note that Mount() simply sets a wildcard along the `pattern` that will continue
// routing at the `handler`, which in most cases is another chi.Router. As a result,
// if you define two Mount() routes on the exact same pattern the mount will panic.
func (r *MatchaOpenAPI) Mount(pattern string, handler http.Handler) {
	r.matcha.Mount(pattern, handler)
}
