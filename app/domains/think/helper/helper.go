package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func GenerateSchema(v interface{}) string {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	var sb strings.Builder
	sb.WriteString("{")

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		fieldType := field.Type.Name()

		if jsonTag == "" {
			fieldName = toSnakeCase(fieldName)
		} else {
			fieldName = jsonTag
		}

		sb.WriteString(fmt.Sprintf("\"%s\":%s", fieldName, fieldType))
		if i < val.NumField()-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString("}")

	return sb.String()
}

func Encode(v interface{}) string {
	decode, _ := json.Marshal(v)
	return string(decode)
}

func toSnakeCase(str string) string {
	var result []rune
	for i, char := range str {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, char)
	}
	return strings.ToLower(string(result))
}
