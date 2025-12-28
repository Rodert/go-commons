// Package jsonutils 提供JSON处理相关的工具函数
// Package jsonutils provides JSON processing utility functions
package jsonutils

import (
	"encoding/json"
	"fmt"
)

// PrettyJSON 美化JSON字符串（添加缩进和换行）
//
// 参数 / Parameters:
//   - data: JSON字符串或字节数组 / JSON string or bytes
//
// 返回值 / Returns:
//   - string: 美化后的JSON字符串 / pretty JSON string
//   - error: 如果JSON无效则返回错误 / error if JSON is invalid
//
// 示例 / Example:
//   PrettyJSON(`{"name":"John","age":30}`)
//
// PrettyJSON formats JSON string with indentation and newlines
func PrettyJSON(data interface{}) (string, error) {
	var jsonBytes []byte
	var err error

	switch v := data.(type) {
	case string:
		jsonBytes = []byte(v)
	case []byte:
		jsonBytes = v
	default:
		return "", fmt.Errorf("不支持的数据类型: %T", data)
	}

	// 首先解析JSON以确保格式正确
	var jsonObj interface{}
	if err := json.Unmarshal(jsonBytes, &jsonObj); err != nil {
		return "", fmt.Errorf("无效的JSON: %w", err)
	}

	// 重新编码并格式化
	prettyBytes, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("格式化JSON失败: %w", err)
	}

	return string(prettyBytes), nil
}

// CompactJSON 压缩JSON字符串（移除所有不必要的空白字符）
//
// 参数 / Parameters:
//   - data: JSON字符串或字节数组 / JSON string or bytes
//
// 返回值 / Returns:
//   - string: 压缩后的JSON字符串 / compact JSON string
//   - error: 如果JSON无效则返回错误 / error if JSON is invalid
//
// 示例 / Example:
//   CompactJSON(`{ "name": "John", "age": 30 }`)
//
// CompactJSON compacts JSON string by removing unnecessary whitespace
func CompactJSON(data interface{}) (string, error) {
	var jsonBytes []byte
	var err error

	switch v := data.(type) {
	case string:
		jsonBytes = []byte(v)
	case []byte:
		jsonBytes = v
	default:
		return "", fmt.Errorf("不支持的数据类型: %T", data)
	}

	// 解析JSON
	var jsonObj interface{}
	if err := json.Unmarshal(jsonBytes, &jsonObj); err != nil {
		return "", fmt.Errorf("无效的JSON: %w", err)
	}

	// 重新编码并压缩
	compactBytes, err := json.Marshal(jsonObj)
	if err != nil {
		return "", fmt.Errorf("压缩JSON失败: %w", err)
	}

	return string(compactBytes), nil
}

// StructToMap 将结构体转换为map
//
// 参数 / Parameters:
//   - v: 要转换的结构体（必须是指针） / struct to convert (must be a pointer)
//
// 返回值 / Returns:
//   - map[string]interface{}: 转换后的map / converted map
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   type User struct { Name string; Age int }
//   user := &User{Name: "John", Age: 30}
//   m, _ := StructToMap(user)
//
// StructToMap converts a struct to map[string]interface{}
func StructToMap(v interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("序列化失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, fmt.Errorf("反序列化失败: %w", err)
	}

	return result, nil
}

// MapToStruct 将map转换为结构体
//
// 参数 / Parameters:
//   - m: 要转换的map / map to convert
//   - v: 目标结构体（必须是指针） / target struct (must be a pointer)
//
// 返回值 / Returns:
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   type User struct { Name string; Age int }
//   m := map[string]interface{}{"name": "John", "age": 30}
//   var user User
//   MapToStruct(m, &user)
//
// MapToStruct converts a map to struct
func MapToStruct(m map[string]interface{}, v interface{}) error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("序列化失败: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, v); err != nil {
		return fmt.Errorf("反序列化失败: %w", err)
	}

	return nil
}

// IsValidJSON 检查字符串是否为有效的JSON
//
// 参数 / Parameters:
//   - data: JSON字符串或字节数组 / JSON string or bytes
//
// 返回值 / Returns:
//   - bool: 如果有效返回true / true if valid
//
// 示例 / Example:
//   IsValidJSON(`{"name":"John"}`) // true
//
// IsValidJSON checks if a string is valid JSON
func IsValidJSON(data interface{}) bool {
	var jsonBytes []byte

	switch v := data.(type) {
	case string:
		jsonBytes = []byte(v)
	case []byte:
		jsonBytes = v
	default:
		return false
	}

	var jsonObj interface{}
	return json.Unmarshal(jsonBytes, &jsonObj) == nil
}

// MergeJSON 合并多个JSON对象
//
// 参数 / Parameters:
//   - jsonObjects: JSON对象列表（后面的会覆盖前面的） / list of JSON objects (later ones override earlier ones)
//
// 返回值 / Returns:
//   - map[string]interface{}: 合并后的JSON对象 / merged JSON object
//   - error: 如果合并失败则返回错误 / error if merge fails
//
// 示例 / Example:
//   obj1 := map[string]interface{}{"a": 1, "b": 2}
//   obj2 := map[string]interface{}{"b": 3, "c": 4}
//   merged, _ := MergeJSON(obj1, obj2) // {"a":1, "b":3, "c":4}
//
// MergeJSON merges multiple JSON objects
func MergeJSON(jsonObjects ...map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, obj := range jsonObjects {
		for k, v := range obj {
			result[k] = v
		}
	}

	return result, nil
}
