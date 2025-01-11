package openapi

import (
	"reflect"
	"strings"
)

func convertMapToSchema(fieldTypes map[string]string) *Schema {
	properties := make(map[string]*Schema)

	for fieldName, fieldType := range fieldTypes {
		fieldSchema := &Schema{
			Type: getOpenAPIType(fieldType),
		}
		properties[fieldName] = fieldSchema
	}

	return &Schema{
		Type:       "object",
		Properties: properties,
	}
}

func getOpenAPIType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "int64":
		return "integer"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "time.Time":
		return "date-time"
	case "struct":
		return "object"
	case "map":
		return "object"
	case "[]byte":
		return "string"
	default:
		return "object"
	}
}

func convertStructTypeToMap(t interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.TypeOf(t)

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			jsonTag := field.Tag.Get("json")
			fieldName := field.Name

			if jsonTag != "" && jsonTag != "-" {
				tagParts := strings.Split(jsonTag, ",")
				fieldName = tagParts[0]
			}

			result[fieldName] = field.Type.String()
		}
	}
	return result
}
