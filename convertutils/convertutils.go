// Package convertutils 提供类型转换相关的工具函数
// Package convertutils provides type conversion utility functions
package convertutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// StringToInt 将字符串转换为整数
//
// 参数 / Parameters:
//   - s: 字符串 / string
//
// 返回值 / Returns:
//   - int: 转换后的整数 / converted integer
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   StringToInt("123") // 123, nil
//
// StringToInt converts a string to int
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToInt64 将字符串转换为int64
//
// 参数 / Parameters:
//   - s: 字符串 / string
//   - base: 进制（10表示十进制） / base (10 for decimal)
//
// 返回值 / Returns:
//   - int64: 转换后的整数 / converted integer
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   StringToInt64("123", 10) // 123, nil
//
// StringToInt64 converts a string to int64
func StringToInt64(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

// StringToFloat64 将字符串转换为float64
//
// 参数 / Parameters:
//   - s: 字符串 / string
//
// 返回值 / Returns:
//   - float64: 转换后的浮点数 / converted float
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   StringToFloat64("123.45") // 123.45, nil
//
// StringToFloat64 converts a string to float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// StringToBool 将字符串转换为bool
//
// 参数 / Parameters:
//   - s: 字符串 / string
//
// 返回值 / Returns:
//   - bool: 转换后的布尔值 / converted boolean
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   StringToBool("true") // true, nil
//
// StringToBool converts a string to bool
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// IntToString 将整数转换为字符串
//
// 参数 / Parameters:
//   - i: 整数 / integer
//
// 返回值 / Returns:
//   - string: 转换后的字符串 / converted string
//
// 示例 / Example:
//   IntToString(123) // "123"
//
// IntToString converts an int to string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString 将int64转换为字符串
//
// 参数 / Parameters:
//   - i: int64整数 / int64 integer
//   - base: 进制（10表示十进制） / base (10 for decimal)
//
// 返回值 / Returns:
//   - string: 转换后的字符串 / converted string
//
// 示例 / Example:
//   Int64ToString(123, 10) // "123"
//
// Int64ToString converts an int64 to string
func Int64ToString(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// Float64ToString 将float64转换为字符串
//
// 参数 / Parameters:
//   - f: 浮点数 / float
//   - prec: 精度（小数位数） / precision (decimal places)
//
// 返回值 / Returns:
//   - string: 转换后的字符串 / converted string
//
// 示例 / Example:
//   Float64ToString(123.456, 2) // "123.46"
//
// Float64ToString converts a float64 to string
func Float64ToString(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

// BoolToString 将bool转换为字符串
//
// 参数 / Parameters:
//   - b: 布尔值 / boolean
//
// 返回值 / Returns:
//   - string: 转换后的字符串 / converted string
//
// 示例 / Example:
//   BoolToString(true) // "true"
//
// BoolToString converts a bool to string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// ToInt 将任意类型转换为int（支持string, int, int64, float64）
//
// 参数 / Parameters:
//   - v: 要转换的值 / value to convert
//
// 返回值 / Returns:
//   - int: 转换后的整数 / converted integer
//   - error: 如果转换失败则返回错误 / error if conversion fails
//
// 示例 / Example:
//   ToInt("123") // 123, nil
//   ToInt(123)   // 123, nil
//
// ToInt converts any type to int (supports string, int, int64, float64)
func ToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case float32:
		return int(val), nil
	case float64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	default:
		return 0, fmt.Errorf("无法转换为int: %T", v)
	}
}

// ToString 将任意类型转换为string
//
// 参数 / Parameters:
//   - v: 要转换的值 / value to convert
//
// 返回值 / Returns:
//   - string: 转换后的字符串 / converted string
//
// 示例 / Example:
//   ToString(123)    // "123"
//   ToString(true)   // "true"
//
// ToString converts any type to string
func ToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case bool:
		return strconv.FormatBool(val)
	case []byte:
		return string(val)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// DeepCopy 深拷贝任意对象（通过JSON序列化/反序列化）
//
// 参数 / Parameters:
//   - src: 源对象（必须是指针） / source object (must be a pointer)
//   - dst: 目标对象（必须是指针） / destination object (must be a pointer)
//
// 返回值 / Returns:
//   - error: 如果拷贝失败则返回错误 / error if copy fails
//
// 示例 / Example:
//   type User struct { Name string; Age int }
//   src := &User{Name: "John", Age: 30}
//   var dst User
//   DeepCopy(src, &dst)
//
// DeepCopy performs deep copy of any object via JSON serialization
func DeepCopy(src, dst interface{}) error {
	// 使用JSON进行深拷贝
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Ptr {
		return fmt.Errorf("源对象必须是指针类型")
	}
	if dstValue.Kind() != reflect.Ptr {
		return fmt.Errorf("目标对象必须是指针类型")
	}

	// 序列化源对象
	jsonBytes, err := toJSONBytes(src)
	if err != nil {
		return fmt.Errorf("序列化源对象失败: %w", err)
	}

	// 反序列化到目标对象
	if err := fromJSONBytes(jsonBytes, dst); err != nil {
		return fmt.Errorf("反序列化到目标对象失败: %w", err)
	}

	return nil
}

// toJSONBytes 将对象转换为JSON字节数组
func toJSONBytes(v interface{}) ([]byte, error) {
	// 如果是字符串或字节数组，直接返回
	switch val := v.(type) {
	case string:
		return []byte(val), nil
	case []byte:
		return val, nil
	}

	// 使用标准库json包进行序列化
	return json.Marshal(v)
}

// fromJSONBytes 将JSON字节数组转换为对象
func fromJSONBytes(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
