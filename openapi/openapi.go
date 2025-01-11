package openapi

import (
	"fmt"
	"net/http"

	"github.com/abyanmajid/matcha/internal"
)

type Metadata struct {
	OpenAPI        string
	PackageName    string
	PackageVersion string
}

// NewDocs creates a new OpenAPIDocs instance using the provided metadata.
// It initializes the OpenAPI version and the Info section with the package name and version.
//
// Parameters:
//   - metadata: Metadata containing OpenAPI version, package name, and package version.
//
// Returns:
//   - OpenAPIDocs: A new instance of OpenAPIDocs populated with the provided metadata.
func NewDocs(metadata Metadata) OpenAPIDocs {
	return OpenAPIDocs{
		OpenAPI: metadata.OpenAPI,
		Info: Info{
			Title:   metadata.PackageName,
			Version: metadata.PackageVersion,
		},
	}
}

// NewHandler creates a new HTTP handler function that serves the provided OpenAPI documentation.
// It writes the OpenAPI documentation as a JSON response with an HTTP status code of 200 (OK).
//
// Parameters:
//   - docs: OpenAPIDocs containing the OpenAPI documentation to be served.
//
// Returns:
//   - http.HandlerFunc: An HTTP handler function that serves the OpenAPI documentation.
func NewHandler(docs OpenAPIDocs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal.WriteJSON(w, docs, http.StatusOK)
	}
}

type CustomResponse struct {
	Code    int
	Message string
}

func NewOperation(schema interface{}) {
	x := convertStructTypeToMap(schema)
	s := convertMapToSchema(x)
	fmt.Println(x)
	fmt.Println(s.Type)
}
