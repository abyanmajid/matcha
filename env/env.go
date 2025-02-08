package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/abyanmajid/matcha/internal"
)

func Dotenv(filenames ...string) (err error) {
	return internal.Dotenv(filenames...)
}

func Load(config interface{}) error {
	configValue := reflect.ValueOf(config).Elem()
	configType := configValue.Type()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fieldValue := configValue.Field(i)

		envKey := field.Tag.Get("name")
		if envKey == "" {
			return fmt.Errorf("field %s is missing a `name` tag", field.Name)
		}

		required := field.Tag.Get("required") == "true"

		defaultValue := field.Tag.Get("default")

		envValue, exists := os.LookupEnv(envKey)
		if !exists {
			if required && defaultValue == "" {
				return fmt.Errorf("required environment variable %s is missing", envKey)
			}
			envValue = defaultValue
		}

		if err := setField(fieldValue, envValue, field.Type.Kind()); err != nil {
			return fmt.Errorf("failed to set field %s: %v", field.Name, err)
		}
	}

	return nil
}

func setField(field reflect.Value, value string, kind reflect.Kind) error {
	switch kind {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid integer value: %v", err)
		}
		field.SetInt(int64(intValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean value: %v", err)
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported type: %v", kind)
	}
	return nil
}
