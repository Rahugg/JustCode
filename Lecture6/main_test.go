package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON2GoStruct(t *testing.T) {
	tests := []struct {
		name           string
		jsonInput      string
		expectedOutput string
	}{
		{
			name:      "Simple Fields",
			jsonInput: `{"name":"John", "age":25}`,
			expectedOutput: `package main

type Person struct {
	Name string ` + "`" + `json:"name"` + "`" + `
	Age  float64 ` + "`" + `json:"age"` + "`" + `
}
`,
		},
		{
			name:      "Nested Fields",
			jsonInput: `{"name": "John", "age": 25, "address": {"street": "123 Main St", "city": "Anytown"}}`,
			expectedOutput: `package main


type Person struct{
Name string ` + "`" + `json:"name"` + "`" + `
Age float64 ` + "`" + `json:"age"` + "`" + `
Address type Address struct{
Street string ` + "`" + `json:"street"` + "`" + `
City string ` + "`" + `json:"city"` + "`" + `
}
}
`,
		},
		{
			name:      "Mixed Fields",
			jsonInput: `{"name":"John", "isStudent":true}`,
			expectedOutput: `package main

type Person struct {
	Name      string ` + "`" + `json:"name"` + "`" + `
	IsStudent bool   ` + "`" + `json:"isStudent"` + "`" + `
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := JSON2GoStruct(tt.jsonInput, "Person")
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedOutput, output)

		})
	}
}
