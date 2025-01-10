package matcha

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/abyanmajid/matcha/internal"
	"github.com/abyanmajid/matcha/logger"
	"github.com/go-chi/chi/v5"
)

type Matcha struct {
	mux *chi.Mux
}

// New creates and returns a new instance of .Mux by initializing
// and configuring the  router.
func New() *Matcha {
	return &Matcha{
		mux: chi.NewRouter(),
	}
}

// Serve mux on a given local address
func (r *Matcha) Serve(addr string) {
	port, err := internal.GetPortFromAddr(addr)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	internal.PrintIntro()

	logger.Info("A Matcha web server is now running on port %s... 🌿", port)

	http.ListenAndServe(addr, r.mux)
}

// Use appends a middleware handler to the Mux middleware stack.
//
// The middleware stack for any Mux will execute before searching for a matching
// route to a specific handler, which provides opportunity to respond early,
// change the course of the request execution, or set request-scoped values for
// the next http.Handler.
func (r *Matcha) Use(middlewares ...func(http.Handler) http.Handler) {
	r.mux.Use(middlewares...)
}

// Handle adds the route `pattern` that matches any http method to
// execute the `handler` http.Handler.
func (r *Matcha) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

// HandleFunc adds the route `pattern` that matches any http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handlerFn)
}

// Method adds the route `pattern` that matches `method` http method to
// execute the `handler` http.Handler.
func (r *Matcha) Method(method, pattern string, handler http.Handler) {
	r.mux.Method(method, pattern, handler)
}

// MethodFunc adds the route `pattern` that matches `method` http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) MethodFunc(method, pattern string, handlerFn http.HandlerFunc) {
	r.mux.MethodFunc(method, pattern, handlerFn)
}

// Connect adds the route `pattern` that matches a CONNECT http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Connect(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Connect(pattern, handlerFn)
}

// Delete adds the route `pattern` that matches a DELETE http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Delete(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Delete(pattern, handlerFn)
}

// Get adds the route `pattern` that matches a GET http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Get(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Get(pattern, handlerFn)
}

// Head adds the route `pattern` that matches a HEAD http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Head(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Head(pattern, handlerFn)
}

// Options adds the route `pattern` that matches an OPTIONS http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Options(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Options(pattern, handlerFn)
}

// Patch adds the route `pattern` that matches a PATCH http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Patch(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Patch(pattern, handlerFn)
}

// Post adds the route `pattern` that matches a POST http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Post(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Post(pattern, handlerFn)
}

// Put adds the route `pattern` that matches a PUT http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Put(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Put(pattern, handlerFn)
}

// Trace adds the route `pattern` that matches a TRACE http method to
// execute the `handlerFn` http.HandlerFunc.
func (r *Matcha) Trace(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Trace(pattern, handlerFn)
}

// NotFound sets a custom http.HandlerFunc for routing paths that could
// not be found. The default 404 handler is `http.NotFound`.
func (r *Matcha) NotFound(handlerFn http.HandlerFunc) {
	r.mux.NotFound(handlerFn)
}

// MethodNotAllowed sets a custom http.HandlerFunc for routing paths where the
// method is unresolved. The default handler returns a 405 with an empty body.
func (r *Matcha) MethodNotAllowed(handlerFn http.HandlerFunc) {
	r.mux.MethodNotAllowed(handlerFn)
}

// With adds inline middlewares for an endpoint handler.
func (r *Matcha) With(middlewares ...func(http.Handler) http.Handler) Matcha {
	return Matcha{
		mux: r.mux.With(middlewares...).(*chi.Mux),
	}
}

// ErrorJSON sets custom handlers for returning JSON responses when paths
// cannot be found or method not allowed.
func (r *Matcha) ErrorJSON() {
	r.mux.NotFound(func(w http.ResponseWriter, req *http.Request) {
		internal.WriteErrorJSON(w, errors.New("resource not found"), http.StatusNotFound)
	})

	r.mux.MethodNotAllowed(func(w http.ResponseWriter, req *http.Request) {
		internal.WriteErrorJSON(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
	})
}

// Group creates a new inline-Mux with a copy of middleware stack. It's useful
// for a group of handlers along the same routing path that use an additional
// set of middlewares. See _examples/.
func (r *Matcha) Group(fn func(r Matcha)) Matcha {
	subRouter := &Matcha{
		mux: chi.NewRouter(),
	}
	fn(*subRouter)
	r.mux.Mount("/", subRouter.mux)
	return *subRouter
}

// Route creates a new Mux and mounts it along the `pattern` as a subrouter.
// Effectively, this is a short-hand call to Mount. See _examples/.
func (r *Matcha) Route(pattern string, fn func(r Matcha)) Matcha {
	if fn == nil {
		panic(fmt.Sprintf("chi: attempting to Route() a nil subrouter on '%s'", pattern))
	}
	subRouter := New()
	fn(*subRouter)
	r.mux.Mount(pattern, subRouter.mux)
	return Matcha{
		mux: subRouter.mux,
	}
}

// Mount attaches another http.Handler or chi Router as a subrouter along a routing
// path. It's very useful to split up a large API as many independent routers and
// compose them as a single service using Mount. See _examples/.
//
// Note that Mount() simply sets a wildcard along the `pattern` that will continue
// routing at the `handler`, which in most cases is another chi.Router. As a result,
// if you define two Mount() routes on the exact same pattern the mount will panic.
func (r *Matcha) Mount(pattern string, handler http.Handler) {
	r.mux.Mount(pattern, handler)
}
