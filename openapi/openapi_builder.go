package openapi

import (
	"net/http"

	"github.com/abyanmajid/matcha/internal"
)

type Meta struct {
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
func NewDocs(metadata Meta) OpenAPIDocs {
	return OpenAPIDocs{
		OpenAPI: metadata.OpenAPI,
		Info: Info{
			Title:   metadata.PackageName,
			Version: metadata.PackageVersion,
		},
		Paths: map[string]Path{},
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

func NewPath(pattern string, operation Operation) map[string]Operation {
	return map[string]Operation{
		pattern: operation,
	}
}

func NewSchema(schemaType interface{}) (*ContentSchema, error) {
	err := enforceStructParam(schemaType)
	if err != nil {
		return nil, err
	}
	typeMap := convertStructTypeToMap(schemaType)
	parsedSchema := convertMapToSchema(typeMap)
	return parsedSchema, nil
}

func NewOperation(summary string, description string, requestBody RequestBody, response map[int]Response) *Operation {
	return &Operation{
		Summary:     summary,
		Description: description,
		RequestBody: &requestBody,
		Response:    &response,
	}
}

type paramOptionGetter func(name string, description string) Parameter

type paramOptions struct {
	Query    paramOptionGetter
	Header   paramOptionGetter
	Path     paramOptionGetter
	FormData paramOptionGetter
	Cookie   paramOptionGetter
}

var Param = paramOptions{
	Query: func(name string, description string) Parameter {
		return createOpenAPIParam("query", name, description)
	},
	Header: func(name string, description string) Parameter {
		return createOpenAPIParam("header", name, description)
	},
	Path: func(name string, description string) Parameter {
		return createOpenAPIParam("path", name, description)
	},
	FormData: func(name string, description string) Parameter {
		return createOpenAPIParam("formData", name, description)
	},
	Cookie: func(name string, description string) Parameter {
		return createOpenAPIParam("cookie", name, description)
	},
}

func createOpenAPIParam(in string, name string, description string) Parameter {
	return Parameter{
		In:          in,
		Name:        name,
		Description: description,
	}
}
