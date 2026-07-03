// Package sliceutils 提供切片/集合相关的工具函数
// Package sliceutils provides slice/collection utility functions
package sliceutils

import (
	"cmp"
	"fmt"
	"slices"
)

// Unique 去除切片中的重复元素，保持原有顺序
//
// 示例 / Example:
//
//	Unique([]int{1, 2, 2, 3}) // []int{1, 2, 3}
//
// Unique removes duplicate elements from a slice while preserving order
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{}, len(slice))
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Filter 过滤切片，返回满足条件的元素
//
// 示例 / Example:
//
//	Filter([]int{1, 2, 3, 4, 5}, func(x int) bool { return x > 2 }) // []int{3, 4, 5}
//
// Filter returns elements from slice for which fn returns true
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 对切片中的每个元素应用函数，返回新切片
//
// 示例 / Example:
//
//	Map([]int{1, 2, 3}, func(x int) int { return x * 2 }) // []int{2, 4, 6}
//
// Map applies fn to each element and returns the resulting slice
func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Reduce 归约切片，从左到右累积值
//
// 示例 / Example:
//
//	Reduce([]int{1, 2, 3}, 0, func(acc, x int) int { return acc + x }) // 6
//
// Reduce reduces a slice to a single value by applying fn left-to-right
func Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R {
	acc := initial
	for _, item := range slice {
		acc = fn(acc, item)
	}
	return acc
}

// Contains 检查切片是否包含指定元素
//
// 示例 / Example:
//
//	Contains([]int{1, 2, 3}, 2) // true
//
// Contains reports whether item is present in slice
func Contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// Intersection 求两个切片的交集，保持 slice1 中的元素顺序
//
// 示例 / Example:
//
//	Intersection([]int{1, 2, 3}, []int{2, 3, 4}) // []int{2, 3}
//
// Intersection returns elements present in both slices
func Intersection[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]struct{}, len(slice2))
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	seen := make(map[T]struct{})
	result := make([]T, 0)
	for _, item := range slice1 {
		if _, inSet := set[item]; inSet {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// Union 求两个切片的并集，保持先后顺序
//
// 示例 / Example:
//
//	Union([]int{1, 2}, []int{2, 3}) // []int{1, 2, 3}
//
// Union returns all distinct elements from both slices
func Union[T comparable](slice1, slice2 []T) []T {
	seen := make(map[T]struct{}, len(slice1)+len(slice2))
	result := make([]T, 0, len(slice1)+len(slice2))
	for _, item := range append(slice1, slice2...) {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Difference 求差集（属于 slice1 但不属于 slice2 的元素）
//
// 示例 / Example:
//
//	Difference([]int{1, 2, 3}, []int{2, 3}) // []int{1}
//
// Difference returns elements in slice1 that are not in slice2
func Difference[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]struct{}, len(slice2))
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	seen := make(map[T]struct{})
	result := make([]T, 0)
	for _, item := range slice1 {
		if _, inSet := set[item]; !inSet {
			if _, alreadySeen := seen[item]; !alreadySeen {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// Reverse 返回切片的反转副本，不修改原切片
//
// 示例 / Example:
//
//	Reverse([]int{1, 2, 3}) // []int{3, 2, 1}
//
// Reverse returns a reversed copy of the slice
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	slices.Reverse(result)
	return result
}

// Sort 返回排序后的副本（升序），不修改原切片
//
// 示例 / Example:
//
//	Sort([]int{3, 1, 2}) // []int{1, 2, 3}
//
// Sort returns a sorted copy of the slice in ascending order
func Sort[T cmp.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	slices.Sort(result)
	return result
}

// SortDesc 返回降序排序后的副本，不修改原切片
//
// 示例 / Example:
//
//	SortDesc([]int{1, 3, 2}) // []int{3, 2, 1}
//
// SortDesc returns a sorted copy of the slice in descending order
func SortDesc[T cmp.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	slices.SortFunc(result, func(a, b T) int { return cmp.Compare(b, a) })
	return result
}

// Paginate 对切片进行分页
//
// 参数 / Parameters:
//   - page: 页码（从1开始） / page number (starts from 1)
//   - pageSize: 每页大小 / items per page
//
// 返回值 / Returns:
//   - []T: 当前页的元素 / elements on the current page
//   - int: 总页数 / total pages
//   - error: 参数无效时返回错误 / error if params invalid
//
// 示例 / Example:
//
//	result, totalPages, _ := Paginate([]int{1, 2, 3, 4, 5}, 1, 2) // [1,2], 3
//
// Paginate returns the items for the requested page
func Paginate[T any](slice []T, page, pageSize int) ([]T, int, error) {
	if page < 1 {
		return nil, 0, fmt.Errorf("页码必须大于0 / page must be >= 1")
	}
	if pageSize < 1 {
		return nil, 0, fmt.Errorf("每页大小必须大于0 / pageSize must be >= 1")
	}

	total := len(slice)
	totalPages := (total + pageSize - 1) / pageSize

	if page > totalPages {
		return []T{}, totalPages, nil
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}

	return slice[start:end], totalPages, nil
}
