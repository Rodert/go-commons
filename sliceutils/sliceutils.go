// Package sliceutils 提供切片/集合相关的工具函数
// Package sliceutils provides slice/collection utility functions
package sliceutils

import (
	"fmt"
	"reflect"
	"sort"
)

// Unique 去除切片中的重复元素，保持原有顺序
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//
// 返回值 / Returns:
//   - []interface{}: 去重后的切片 / deduplicated slice
//
// 示例 / Example:
//   Unique([]interface{}{1, 2, 2, 3, 3, 3}) // []interface{}{1, 2, 3}
//
// Unique removes duplicate elements from a slice while preserving order
func Unique(slice interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	seen := make(map[interface{}]bool)
	result := make([]interface{}, 0, v.Len())

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// UniqueInt 去除整数切片中的重复元素
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//
// 返回值 / Returns:
//   - []int: 去重后的切片 / deduplicated slice
//
// 示例 / Example:
//   UniqueInt([]int{1, 2, 2, 3, 3, 3}) // []int{1, 2, 3}
//
// UniqueInt removes duplicate elements from an int slice
func UniqueInt(slice []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// UniqueString 去除字符串切片中的重复元素
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//
// 返回值 / Returns:
//   - []string: 去重后的切片 / deduplicated slice
//
// 示例 / Example:
//   UniqueString([]string{"a", "b", "b", "c"}) // []string{"a", "b", "c"}
//
// UniqueString removes duplicate elements from a string slice
func UniqueString(slice []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Filter 过滤切片，返回满足条件的元素
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//   - fn: 过滤函数，返回true表示保留该元素 / filter function, returns true to keep element
//
// 返回值 / Returns:
//   - []interface{}: 过滤后的切片 / filtered slice
//
// 示例 / Example:
//   Filter([]interface{}{1, 2, 3, 4, 5}, func(x interface{}) bool {
//     return x.(int) > 2
//   }) // []interface{}{3, 4, 5}
//
// Filter filters a slice based on a predicate function
func Filter(slice interface{}, fn func(interface{}) bool) []interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]interface{}, 0)
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

// FilterInt 过滤整数切片
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//   - fn: 过滤函数 / filter function
//
// 返回值 / Returns:
//   - []int: 过滤后的切片 / filtered slice
//
// 示例 / Example:
//   FilterInt([]int{1, 2, 3, 4, 5}, func(x int) bool { return x > 2 })
//
// FilterInt filters an int slice based on a predicate function
func FilterInt(slice []int, fn func(int) bool) []int {
	result := make([]int, 0, len(slice))
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// FilterString 过滤字符串切片
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//   - fn: 过滤函数 / filter function
//
// 返回值 / Returns:
//   - []string: 过滤后的切片 / filtered slice
//
// 示例 / Example:
//   FilterString([]string{"a", "ab", "abc"}, func(s string) bool { return len(s) > 1 })
//
// FilterString filters a string slice based on a predicate function
func FilterString(slice []string, fn func(string) bool) []string {
	result := make([]string, 0, len(slice))
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 对切片中的每个元素应用函数，返回新切片
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//   - fn: 映射函数 / map function
//
// 返回值 / Returns:
//   - []interface{}: 映射后的切片 / mapped slice
//
// 示例 / Example:
//   Map([]interface{}{1, 2, 3}, func(x interface{}) interface{} {
//     return x.(int) * 2
//   }) // []interface{}{2, 4, 6}
//
// Map applies a function to each element of a slice
func Map(slice interface{}, fn func(interface{}) interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		result[i] = fn(item)
	}

	return result
}

// MapInt 对整数切片中的每个元素应用函数
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//   - fn: 映射函数 / map function
//
// 返回值 / Returns:
//   - []int: 映射后的切片 / mapped slice
//
// 示例 / Example:
//   MapInt([]int{1, 2, 3}, func(x int) int { return x * 2 })
//
// MapInt applies a function to each element of an int slice
func MapInt(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// MapString 对字符串切片中的每个元素应用函数
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//   - fn: 映射函数 / map function
//
// 返回值 / Returns:
//   - []string: 映射后的切片 / mapped slice
//
// 示例 / Example:
//   MapString([]string{"a", "b"}, func(s string) string { return s + "!" })
//
// MapString applies a function to each element of a string slice
func MapString(slice []string, fn func(string) string) []string {
	result := make([]string, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Reduce 归约切片，从左到右累积值
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//   - initial: 初始值 / initial value
//   - fn: 归约函数，参数为(累积值, 当前元素) / reduce function with (accumulator, current) params
//
// 返回值 / Returns:
//   - interface{}: 归约后的值 / reduced value
//
// 示例 / Example:
//   Reduce([]interface{}{1, 2, 3}, 0, func(acc, x interface{}) interface{} {
//     return acc.(int) + x.(int)
//   }) // 6
//
// Reduce reduces a slice to a single value by applying a function
func Reduce(slice interface{}, initial interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return initial
	}

	accumulator := initial
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		accumulator = fn(accumulator, item)
	}

	return accumulator
}

// ReduceInt 归约整数切片
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//   - initial: 初始值 / initial value
//   - fn: 归约函数 / reduce function
//
// 返回值 / Returns:
//   - int: 归约后的值 / reduced value
//
// 示例 / Example:
//   ReduceInt([]int{1, 2, 3}, 0, func(acc, x int) int { return acc + x })
//
// ReduceInt reduces an int slice to a single value
func ReduceInt(slice []int, initial int, fn func(int, int) int) int {
	accumulator := initial
	for _, item := range slice {
		accumulator = fn(accumulator, item)
	}
	return accumulator
}

// Paginate 对切片进行分页
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//   - page: 页码（从1开始） / page number (starts from 1)
//   - pageSize: 每页大小 / page size
//
// 返回值 / Returns:
//   - []interface{}: 分页后的切片 / paginated slice
//   - int: 总页数 / total pages
//   - error: 如果参数无效则返回错误 / error if params invalid
//
// 示例 / Example:
//   result, totalPages, _ := Paginate([]interface{}{1, 2, 3, 4, 5}, 1, 2)
//
// Paginate paginates a slice
func Paginate(slice interface{}, page, pageSize int) ([]interface{}, int, error) {
	if page < 1 {
		return nil, 0, fmt.Errorf("页码必须大于0")
	}
	if pageSize < 1 {
		return nil, 0, fmt.Errorf("每页大小必须大于0")
	}

	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, 0, fmt.Errorf("输入必须是切片类型")
	}

	total := v.Len()
	totalPages := (total + pageSize - 1) / pageSize // 向上取整

	if page > totalPages {
		return []interface{}{}, totalPages, nil
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}

	result := make([]interface{}, end-start)
	for i := start; i < end; i++ {
		result[i-start] = v.Index(i).Interface()
	}

	return result, totalPages, nil
}

// PaginateInt 对整数切片进行分页
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//   - page: 页码（从1开始） / page number (starts from 1)
//   - pageSize: 每页大小 / page size
//
// 返回值 / Returns:
//   - []int: 分页后的切片 / paginated slice
//   - int: 总页数 / total pages
//   - error: 如果参数无效则返回错误 / error if params invalid
//
// 示例 / Example:
//   result, totalPages, _ := PaginateInt([]int{1, 2, 3, 4, 5}, 1, 2)
//
// PaginateInt paginates an int slice
func PaginateInt(slice []int, page, pageSize int) ([]int, int, error) {
	if page < 1 {
		return nil, 0, fmt.Errorf("页码必须大于0")
	}
	if pageSize < 1 {
		return nil, 0, fmt.Errorf("每页大小必须大于0")
	}

	total := len(slice)
	totalPages := (total + pageSize - 1) / pageSize

	if page > totalPages {
		return []int{}, totalPages, nil
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}

	return slice[start:end], totalPages, nil
}

// Intersection 求两个切片的交集
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []interface{}: 交集切片 / intersection slice
//
// 示例 / Example:
//   Intersection([]interface{}{1, 2, 3}, []interface{}{2, 3, 4}) // []interface{}{2, 3}
//
// Intersection returns the intersection of two slices
func Intersection(slice1, slice2 interface{}) []interface{} {
	v1 := reflect.ValueOf(slice1)
	v2 := reflect.ValueOf(slice2)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		return nil
	}

	// 将第二个切片转换为map以便快速查找
	set := make(map[interface{}]bool)
	for i := 0; i < v2.Len(); i++ {
		set[v2.Index(i).Interface()] = true
	}

	result := make([]interface{}, 0)
	seen := make(map[interface{}]bool)
	for i := 0; i < v1.Len(); i++ {
		item := v1.Index(i).Interface()
		if set[item] && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// IntersectionInt 求两个整数切片的交集
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []int: 交集切片 / intersection slice
//
// 示例 / Example:
//   IntersectionInt([]int{1, 2, 3}, []int{2, 3, 4}) // []int{2, 3}
//
// IntersectionInt returns the intersection of two int slices
func IntersectionInt(slice1, slice2 []int) []int {
	set := make(map[int]bool)
	for _, item := range slice2 {
		set[item] = true
	}

	result := make([]int, 0)
	seen := make(map[int]bool)
	for _, item := range slice1 {
		if set[item] && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Union 求两个切片的并集
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []interface{}: 并集切片 / union slice
//
// 示例 / Example:
//   Union([]interface{}{1, 2}, []interface{}{2, 3}) // []interface{}{1, 2, 3}
//
// Union returns the union of two slices
func Union(slice1, slice2 interface{}) []interface{} {
	seen := make(map[interface{}]bool)
	result := make([]interface{}, 0)

	v1 := reflect.ValueOf(slice1)
	if v1.Kind() == reflect.Slice {
		for i := 0; i < v1.Len(); i++ {
			item := v1.Index(i).Interface()
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}
	}

	v2 := reflect.ValueOf(slice2)
	if v2.Kind() == reflect.Slice {
		for i := 0; i < v2.Len(); i++ {
			item := v2.Index(i).Interface()
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}
	}

	return result
}

// UnionInt 求两个整数切片的并集
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []int: 并集切片 / union slice
//
// 示例 / Example:
//   UnionInt([]int{1, 2}, []int{2, 3}) // []int{1, 2, 3}
//
// UnionInt returns the union of two int slices
func UnionInt(slice1, slice2 []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, item := range slice1 {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	for _, item := range slice2 {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Difference 求两个切片的差集（slice1 - slice2）
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []interface{}: 差集切片 / difference slice
//
// 示例 / Example:
//   Difference([]interface{}{1, 2, 3}, []interface{}{2, 3}) // []interface{}{1}
//
// Difference returns the difference of two slices (slice1 - slice2)
func Difference(slice1, slice2 interface{}) []interface{} {
	v1 := reflect.ValueOf(slice1)
	v2 := reflect.ValueOf(slice2)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		return nil
	}

	// 将第二个切片转换为map以便快速查找
	set := make(map[interface{}]bool)
	for i := 0; i < v2.Len(); i++ {
		set[v2.Index(i).Interface()] = true
	}

	result := make([]interface{}, 0)
	seen := make(map[interface{}]bool)
	for i := 0; i < v1.Len(); i++ {
		item := v1.Index(i).Interface()
		if !set[item] && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// DifferenceInt 求两个整数切片的差集
//
// 参数 / Parameters:
//   - slice1: 第一个切片 / first slice
//   - slice2: 第二个切片 / second slice
//
// 返回值 / Returns:
//   - []int: 差集切片 / difference slice
//
// 示例 / Example:
//   DifferenceInt([]int{1, 2, 3}, []int{2, 3}) // []int{1}
//
// DifferenceInt returns the difference of two int slices
func DifferenceInt(slice1, slice2 []int) []int {
	set := make(map[int]bool)
	for _, item := range slice2 {
		set[item] = true
	}

	result := make([]int, 0)
	seen := make(map[int]bool)
	for _, item := range slice1 {
		if !set[item] && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// SortInt 对整数切片进行排序（升序）
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//
// 返回值 / Returns:
//   - []int: 排序后的切片 / sorted slice
//
// 示例 / Example:
//   SortInt([]int{3, 1, 2}) // []int{1, 2, 3}
//
// SortInt sorts an int slice in ascending order
func SortInt(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	sort.Ints(result)
	return result
}

// SortString 对字符串切片进行排序（升序）
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//
// 返回值 / Returns:
//   - []string: 排序后的切片 / sorted slice
//
// 示例 / Example:
//   SortString([]string{"c", "a", "b"}) // []string{"a", "b", "c"}
//
// SortString sorts a string slice in ascending order
func SortString(slice []string) []string {
	result := make([]string, len(slice))
	copy(result, slice)
	sort.Strings(result)
	return result
}

// SortIntDesc 对整数切片进行排序（降序）
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//
// 返回值 / Returns:
//   - []int: 排序后的切片 / sorted slice
//
// 示例 / Example:
//   SortIntDesc([]int{1, 3, 2}) // []int{3, 2, 1}
//
// SortIntDesc sorts an int slice in descending order
func SortIntDesc(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}

// SortStringDesc 对字符串切片进行排序（降序）
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//
// 返回值 / Returns:
//   - []string: 排序后的切片 / sorted slice
//
// 示例 / Example:
//   SortStringDesc([]string{"a", "c", "b"}) // []string{"c", "b", "a"}
//
// SortStringDesc sorts a string slice in descending order
func SortStringDesc(slice []string) []string {
	result := make([]string, len(slice))
	copy(result, slice)
	sort.Sort(sort.Reverse(sort.StringSlice(result)))
	return result
}

// Contains 检查切片是否包含指定元素
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//   - item: 要查找的元素 / item to find
//
// 返回值 / Returns:
//   - bool: 如果包含返回true / true if contains
//
// 示例 / Example:
//   Contains([]interface{}{1, 2, 3}, 2) // true
//
// Contains checks if a slice contains an item
func Contains(slice interface{}, item interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

// ContainsInt 检查整数切片是否包含指定元素
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//   - item: 要查找的元素 / item to find
//
// 返回值 / Returns:
//   - bool: 如果包含返回true / true if contains
//
// 示例 / Example:
//   ContainsInt([]int{1, 2, 3}, 2) // true
//
// ContainsInt checks if an int slice contains an item
func ContainsInt(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsString 检查字符串切片是否包含指定元素
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//   - item: 要查找的元素 / item to find
//
// 返回值 / Returns:
//   - bool: 如果包含返回true / true if contains
//
// 示例 / Example:
//   ContainsString([]string{"a", "b", "c"}, "b") // true
//
// ContainsString checks if a string slice contains an item
func ContainsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Reverse 反转切片
//
// 参数 / Parameters:
//   - slice: 输入切片 / input slice
//
// 返回值 / Returns:
//   - []interface{}: 反转后的切片 / reversed slice
//
// 示例 / Example:
//   Reverse([]interface{}{1, 2, 3}) // []interface{}{3, 2, 1}
//
// Reverse reverses a slice
func Reverse(slice interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil
	}

	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[v.Len()-1-i] = v.Index(i).Interface()
	}
	return result
}

// ReverseInt 反转整数切片
//
// 参数 / Parameters:
//   - slice: 输入整数切片 / input int slice
//
// 返回值 / Returns:
//   - []int: 反转后的切片 / reversed slice
//
// 示例 / Example:
//   ReverseInt([]int{1, 2, 3}) // []int{3, 2, 1}
//
// ReverseInt reverses an int slice
func ReverseInt(slice []int) []int {
	result := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		result[len(slice)-1-i] = slice[i]
	}
	return result
}

// ReverseString 反转字符串切片
//
// 参数 / Parameters:
//   - slice: 输入字符串切片 / input string slice
//
// 返回值 / Returns:
//   - []string: 反转后的切片 / reversed slice
//
// 示例 / Example:
//   ReverseString([]string{"a", "b", "c"}) // []string{"c", "b", "a"}
//
// ReverseString reverses a string slice
func ReverseString(slice []string) []string {
	result := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		result[len(slice)-1-i] = slice[i]
	}
	return result
}
