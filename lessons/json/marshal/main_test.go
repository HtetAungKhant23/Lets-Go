package main

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type Car struct {
	Brand string `json:"brand"`
	Color string `json:"color"`
}

func Test(t *testing.T) {
	type testCase struct {
		input    []any
		expected [][]byte
	}

	runCases := []testCase{
		{
			input: []any{
				User{"Hak", "123456789", 26},
				Car{"Toyota", "white"},
			},
			expected: [][]byte{
				[]byte(`{"name":"Hak","phone":"123456789","age":26}`),
				[]byte(`{"brand":"Toyota","color":"white"`),
			},
		},
	}

	for _, test := range runCases {
		data, _ := marshalAll(test.input)

		fmt.Println(data)

	}
}
