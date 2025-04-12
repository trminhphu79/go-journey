package utils

import (
	"reflect"
	"strings"
	"unicode"
)

func ToCamelCase(input string) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			result += strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
		}
	}
	return result
}

func FieldToCamelCase(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)

	// If pointer, get the underlying element
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		// Convert field name from TitleCase to camelCase
		key := field.Name
		camelKey := string(unicode.ToLower(rune(key[0]))) + key[1:]
		value := v.Field(i).Interface()

		result[camelKey] = value
	}

	return result
}
