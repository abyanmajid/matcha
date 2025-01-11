package openapi

type Schema struct {
	Parameters  []Parameter
	RequestBody RequestBody
	Response    map[int]Response
}

type RouteDoc struct {
	Method      string
	Summary     string
	Description string
	Schema      Schema
}

type Route struct {
	Path string
	Doc  Operation
}

func NewRoute(routePattern string, routeDoc RouteDoc) Route {
	operationSpec := NewOperation(routeDoc.Summary, routeDoc.Description, routeDoc.Schema.RequestBody, routeDoc.Schema.Response, routeDoc.Schema.Parameters)

	return Route{
		Path: routePattern,
		Doc:  *operationSpec,
	}
}
