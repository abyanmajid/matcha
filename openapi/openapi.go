package openapi

import (
	"net/http"

	"github.com/abyanmajid/matcha/internal"
)

type Metadata struct {
	OpenAPI        string
	PackageName    string
	PackageVersion string
}

func NewDocs(metadata Metadata) OpenAPIDocs {
	return OpenAPIDocs{
		OpenAPI: metadata.OpenAPI,
		Info: Info{
			Title:   metadata.PackageName,
			Version: metadata.PackageVersion,
		},
	}
}

func NewHandler(docs OpenAPIDocs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal.WriteJSON(w, docs, http.StatusOK)
	}
}
