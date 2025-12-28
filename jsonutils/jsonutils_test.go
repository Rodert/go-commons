package jsonutils

import (
	"testing"
)

func TestPrettyJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		shouldErr bool
	}{
		{"valid json string", `{"name":"John","age":30}`, false},
		{"valid json bytes", []byte(`{"name":"John","age":30}`), false},
		{"invalid json", `{name:John}`, true},
		{"nested object", `{"user":{"name":"John","age":30}}`, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := PrettyJSON(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("PrettyJSON(%v) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("PrettyJSON(%v) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result == "" {
				t.Errorf("PrettyJSON(%v) returned empty string", test.input)
			}
		})
	}
}

func TestCompactJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		shouldErr bool
	}{
		{"valid json string", `{ "name": "John", "age": 30 }`, false},
		{"valid json bytes", []byte(`{ "name": "John", "age": 30 }`), false},
		{"invalid json", `{name:John}`, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := CompactJSON(test.input)
			if test.shouldErr && err == nil {
				t.Errorf("CompactJSON(%v) expected error but got none", test.input)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("CompactJSON(%v) unexpected error: %v", test.input, err)
			}
			if !test.shouldErr && result == "" {
				t.Errorf("CompactJSON(%v) returned empty string", test.input)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{Name: "John", Age: 30}
	result, err := StructToMap(user)
	if err != nil {
		t.Errorf("StructToMap(%v) unexpected error: %v", user, err)
	}

	if result["name"] != "John" {
		t.Errorf("StructToMap(%v) name = %v; want John", user, result["name"])
	}
	if age, ok := result["age"].(float64); ok && int(age) != 30 { // JSON numbers are float64
		t.Errorf("StructToMap(%v) age = %v; want 30", user, result["age"])
	} else if !ok {
		t.Errorf("StructToMap(%v) age type assertion failed", user)
	}
}

func TestMapToStruct(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	m := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	var user User
	err := MapToStruct(m, &user)
	if err != nil {
		t.Errorf("MapToStruct(%v, &user) unexpected error: %v", m, err)
	}

	if user.Name != "John" {
		t.Errorf("MapToStruct name = %v; want John", user.Name)
	}
	if user.Age != 30 {
		t.Errorf("MapToStruct age = %v; want 30", user.Age)
	}
}

func TestIsValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"valid json string", `{"name":"John"}`, true},
		{"valid json bytes", []byte(`{"name":"John"}`), true},
		{"invalid json", `{name:John}`, false},
		{"empty string", ``, false},
		{"invalid type", 123, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidJSON(test.input)
			if result != test.expected {
				t.Errorf("IsValidJSON(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestMergeJSON(t *testing.T) {
	obj1 := map[string]interface{}{"a": 1, "b": 2}
	obj2 := map[string]interface{}{"b": 3, "c": 4}

	result, err := MergeJSON(obj1, obj2)
	if err != nil {
		t.Errorf("MergeJSON(%v, %v) unexpected error: %v", obj1, obj2, err)
	}

	if a, ok := result["a"].(int); ok && a != 1 {
		t.Errorf("MergeJSON a = %v; want 1", result["a"])
	} else if !ok {
		t.Errorf("MergeJSON a type assertion failed")
	}
	if b, ok := result["b"].(int); ok && b != 3 { // obj2 should override obj1
		t.Errorf("MergeJSON b = %v; want 3", result["b"])
	} else if !ok {
		t.Errorf("MergeJSON b type assertion failed")
	}
	if c, ok := result["c"].(int); ok && c != 4 {
		t.Errorf("MergeJSON c = %v; want 4", result["c"])
	} else if !ok {
		t.Errorf("MergeJSON c type assertion failed")
	}
}
