package openapi

type OpenAPIDocs struct {
	OpenAPI string          `json:"openapi"`
	Info    Info            `json:"info"`
	Paths   map[string]Path `json:"paths,omitempty"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type Path map[string]Operation

type Operation struct {
	Summary     string            `json:"summary,omitempty"`
	Description string            `json:"description,omitempty"`
	Parameters  []*Parameter      `json:"parameters,omitempty"`
	RequestBody *RequestBody      `json:"requestBody,omitempty"`
	Response    *map[int]Response `json:"responses,omitempty"`
}

type Parameter struct {
	In          string `json:"in"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestBody struct {
	Content map[string]*MediaType `json:"content,omitempty"`
}

type Response struct {
	Description string                `json:"description,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"`
}

type MediaType struct {
	ContentSchema *ContentSchema `json:"schema,omitempty"`
}

type ContentSchema struct {
	Type       string                    `json:"type,omitempty"`
	Properties map[string]*ContentSchema `json:"properties,omitempty"`
}
