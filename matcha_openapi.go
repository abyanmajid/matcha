package matcha

import (
	"sync"
)

type OpenAPIHandler struct {
	mu   sync.Mutex
	Docs OpenAPIDocumentation
}

type OpenAPIDocumentation struct {
	OpenAPI string                 `json:"openapi"`
	Info    OpenAPIInfo            `json:"info"`
	Paths   map[string]OpenAPIPath `json:"paths,omitempty"`
}

type OpenAPIInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type OpenAPIPath map[string]OpenAPIOperation

type OpenAPIOperation struct {
	Summary     string                  `json:"summary,omitempty"`
	Description string                  `json:"description,omitempty"`
	Parameters  []OpenAPIParameter      `json:"parameters,omitempty"`
	RequestBody OpenAPIRequestBody      `json:"requestBody,omitempty"`
	Response    map[int]OpenAPIResponse `json:"response,omitempty"`
}

type OpenAPIParameter struct {
	In          string `json:"in"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required"`
}

type OpenAPIRequestBody struct {
	Required bool                         `json:"required"`
	Content  map[string]*OpenAPIMediaType `json:"content,omitempty"`
}

type OpenAPIResponse struct {
	Description string                       `json:"description,omitempty"`
	Content     map[string]*OpenAPIMediaType `json:"content,omitempty"`
}

// MediaType represents the media type of a request or response body.
type OpenAPIMediaType struct {
	Schema *OpenAPISchema `json:"schema,omitempty"`
}

type OpenAPISchema struct {
	Type       string                    `json:"type,omitempty"`
	Properties map[string]*OpenAPISchema `json:"properties,omitempty"`
	Required   []string                  `json:"required,omitempty"`
}
