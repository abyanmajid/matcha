package openapi

import (
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/internal"
)

type Schema struct {
	Parameters  []Parameter
	RequestBody RequestBody
	Responses   map[int]Response
}

type ResourceDoc struct {
	Summary     string
	Description string
	Schema      Schema
}

type Resource struct {
	Name    string
	Handler http.HandlerFunc
	Doc     Operation
}

func NewResource[Req any, Res any](name string, routeDoc ResourceDoc, handler func(c *ctx.Request[Req]) *ctx.Response[Res]) Resource {
	operationSpec := NewOperation(routeDoc.Summary, routeDoc.Description, routeDoc.Schema.Parameters, routeDoc.Schema.RequestBody, routeDoc.Schema.Responses)
	handlerFunc := internal.NewHandler[Req, Res](handler)

	return Resource{
		Name:    name,
		Handler: handlerFunc,
		Doc:     *operationSpec,
	}
}
