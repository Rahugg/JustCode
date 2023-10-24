package main

import (
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func JSON2GoStruct(jsonStr string, typeName string) (string, error) {
	var jsonData interface{}

	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return "", err
	}

	c := NewFile("main")
	c.Add(GenerateStruct(jsonData, typeName))

	var sb strings.Builder
	err = c.Render(&sb)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

// GenerateStruct generates the Go code for the struct.
func GenerateStruct(data interface{}, typeName string) *Statement {
	code := Type().Id(typeName).StructFunc(func(g *Group) {
		if m, ok := data.(map[string]interface{}); ok {
			for key, value := range m {
				fieldType := getGoType(value)
				if customStruct, ok := value.(map[string]interface{}); ok {
					nestedStructName := strings.Title(key)
					g.Id(strings.Title(key)).Add(GenerateStruct(customStruct, nestedStructName))
				} else {
					g.Id(strings.Title(key)).Add(fieldType).Tag(map[string]string{"json": key})
				}
			}
		}
	})

	return code
}

func getGoType(v interface{}) *Statement {
	switch v.(type) {
	case string:
		return String()
	case float64:
		return Float64()
	case bool:
		return Bool()
	case []interface{}:
		return Index().Interface()
	default:
		return Interface()
	}
}

func main() {
	jsonStr := `{"name": "John", "age": 25, "address": {"street": "123 Main St", "city": "Anytown"}}`
	result, err := JSON2GoStruct(jsonStr, "Person")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
