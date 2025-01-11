package internal

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Resource represents an HTTP resource with a specific route and handler.
// Route is the URL path for the resource.
// Handler is the function that handles HTTP requests for the resource.
type Resource struct {
	Route   string
	Handler http.HandlerFunc
}

// Response is a generic struct that encapsulates a response of type Res.
// It includes the actual response data, an HTTP status code, and an error if one occurred.
type Response[Res any] struct {
	Response   Res
	StatusCode int
	Error      error
}

// Handler is a generic type that represents a function which handles HTTP requests.
// It takes an HTTP request and a request body of any type, and returns a pointer to a Response of any type.
//
// Type Parameters:
//   - Req: The type of the request body.
//   - Res: The type of the response body.
//
// Parameters:
//   - r: The HTTP request.
//   - body: The request body of type Req.
//
// Returns:
//   - A pointer to a Response of type Res.
type Handler[Req any, Res any] func(c *Ctx[Req]) *Response[Res]

// NewResource creates a new Resource with the specified route pattern and handler.
// The handler is a generic function that takes an HTTP request and a request body of type Req,
// and returns a response of type Res.
//
// The function decodes the request body into the specified type Req, and calls the handler
// with the decoded request body. If the request body is invalid, it responds with a 400 Bad Request error.
// If the handler returns an error, it responds with the appropriate status code and error message.
//
// The response from the handler is encoded as JSON and written to the response writer with the
// appropriate status code and Content-Type header.
//
// Parameters:
//   - routePattern: The route pattern for the resource.
//   - handler: A generic function that handles the request and returns a response.
//
// Returns:
//
//	A pointer to a Resource with the specified route pattern and handler.
func NewResource[Req any, Res any](routePattern string, handler Handler[Req, Res]) *Resource {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		var reqBody Req

		isEmptyStruct := func() bool {
			_, ok := any(reqBody).(struct{})
			return ok
		}()

		if !isEmptyStruct {
			if r.ContentLength == 0 {
				WriteErrorJSON(w, errors.New("missing request body"), http.StatusBadRequest)
				return
			}

			if err := json.NewDecoder(r.Body).Decode(&reqBody); err == nil {
				WriteErrorJSON(w, errors.New("invalid request body"), http.StatusBadRequest)
				return
			}
		}

		res := handler(&Ctx[Req]{
			Request:  r,
			Response: w,
			Cookies: Cookies{
				Request:  r,
				Response: w,
			},
			Body: reqBody,
		})
		if res.Error != nil {
			WriteErrorJSON(w, res.Error, http.StatusBadRequest)
			return
		}

		WriteJSON(w, res.Response, res.StatusCode)
	}

	return &Resource{
		Route:   routePattern,
		Handler: handlerFunc,
	}
}
