package main

import (
	"fmt"

	"github.com/Rodert/go-commons/jsonutils"
)

func main() {
	fmt.Println("=== Go Commons JSON Utils Demo ===")
	fmt.Println()

	// JSON美化
	fmt.Println("JSON美化 / Pretty JSON:")
	compactJSON := `{"name":"John","age":30,"city":"New York"}`
	pretty, err := jsonutils.PrettyJSON(compactJSON)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Println("原始JSON:", compactJSON)
		fmt.Println("美化后:")
		fmt.Println(pretty)
	}
	fmt.Println()

	// JSON压缩
	fmt.Println("JSON压缩 / Compact JSON:")
	prettyJSON := `{
  "name": "John",
  "age": 30,
  "city": "New York"
}`
	compact, err := jsonutils.CompactJSON(prettyJSON)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Println("原始JSON:")
		fmt.Println(prettyJSON)
		fmt.Println("压缩后:", compact)
	}
	fmt.Println()

	// Struct转Map
	fmt.Println("Struct转Map / Struct to Map:")
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}
	user := &User{Name: "John", Age: 30, City: "New York"}
	userMap, err := jsonutils.StructToMap(user)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("Struct: %+v\n", user)
		fmt.Printf("Map: %v\n", userMap)
	}
	fmt.Println()

	// Map转Struct
	fmt.Println("Map转Struct / Map to Struct:")
	m := map[string]interface{}{
		"name": "Jane",
		"age":  25,
		"city": "London",
	}
	var newUser User
	err = jsonutils.MapToStruct(m, &newUser)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("Map: %v\n", m)
		fmt.Printf("Struct: %+v\n", newUser)
	}
	fmt.Println()

	// JSON验证
	fmt.Println("JSON验证 / JSON Validation:")
	validJSON := `{"name":"John"}`
	invalidJSON := `{name:John}`
	fmt.Printf("'%s' 是否有效: %v\n", validJSON, jsonutils.IsValidJSON(validJSON))
	fmt.Printf("'%s' 是否有效: %v\n", invalidJSON, jsonutils.IsValidJSON(invalidJSON))
	fmt.Println()

	// JSON合并
	fmt.Println("JSON合并 / Merge JSON:")
	obj1 := map[string]interface{}{"a": 1, "b": 2}
	obj2 := map[string]interface{}{"b": 3, "c": 4}
	merged, err := jsonutils.MergeJSON(obj1, obj2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("对象1: %v\n", obj1)
		fmt.Printf("对象2: %v\n", obj2)
		fmt.Printf("合并后: %v\n", merged)
	}
}
