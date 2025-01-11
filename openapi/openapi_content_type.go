package openapi

func Json(schema *ContentSchema) map[string]*MediaType {
	return map[string]*MediaType{
		"application/json": {
			ContentSchema: schema,
		},
	}
}

func SimpleErrorSchema() *ContentSchema {
	errorSchema := convertMapToSchema(map[string]string{
		"error": "string",
	})
	return errorSchema
}
