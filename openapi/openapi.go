package openapi

import (
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/internal"
)

type Schema struct {
	RequestBody RequestBody
	Response    map[int]Response
}

type ResourceDoc struct {
	Summary     string
	Description string
	Schema      Schema
}

type Resource struct {
	Method  AllowedMethods
	Handler http.HandlerFunc
	Doc     Operation
}

type AllowedMethods struct {
	Method string
}

var GET = AllowedMethods{
	Method: "get",
}

var POST = AllowedMethods{
	Method: "post",
}

var PUT = AllowedMethods{
	Method: "put",
}

var DELETE = AllowedMethods{
	Method: "delete",
}

var PATCH = AllowedMethods{
	Method: "patch",
}

var OPTIONS = AllowedMethods{
	Method: "options",
}

var HEAD = AllowedMethods{
	Method: "head",
}

var CONNECT = AllowedMethods{
	Method: "connect",
}

var TRACE = AllowedMethods{
	Method: "trace",
}

func NewResource[Req any, Res any](routeDoc ResourceDoc, handler func(c *ctx.Request[Req]) *ctx.Response[Res]) Resource {
	operationSpec := NewOperation(routeDoc.Summary, routeDoc.Description, routeDoc.Schema.RequestBody, routeDoc.Schema.Response)

	handlerFunc := internal.NewHandler[Req, Res](handler)

	return Resource{
		Handler: handlerFunc,
		Doc:     *operationSpec,
	}
}
