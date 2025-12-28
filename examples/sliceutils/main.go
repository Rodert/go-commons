package main

import (
	"fmt"

	"github.com/Rodert/go-commons/sliceutils"
)

func main() {
	fmt.Println("=== Go Commons Slice Utils Demo ===")
	fmt.Println()

	// 去重
	fmt.Println("去重 / Unique:")
	numbers := []int{1, 2, 2, 3, 3, 3, 4}
	uniqueNumbers := sliceutils.UniqueInt(numbers)
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("去重: %v\n", uniqueNumbers)

	strings := []string{"a", "b", "b", "c", "c", "c"}
	uniqueStrings := sliceutils.UniqueString(strings)
	fmt.Printf("原始: %v\n", strings)
	fmt.Printf("去重: %v\n", uniqueStrings)
	fmt.Println()

	// 过滤
	fmt.Println("过滤 / Filter:")
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := sliceutils.FilterInt(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("偶数: %v\n", evenNumbers)

	longStrings := sliceutils.FilterString([]string{"a", "ab", "abc", "abcd"}, func(s string) bool {
		return len(s) > 2
	})
	fmt.Printf("长度>2: %v\n", longStrings)
	fmt.Println()

	// 映射
	fmt.Println("映射 / Map:")
	numbers = []int{1, 2, 3, 4, 5}
	doubled := sliceutils.MapInt(numbers, func(x int) int {
		return x * 2
	})
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("加倍: %v\n", doubled)

	upperStrings := sliceutils.MapString([]string{"hello", "world"}, func(s string) string {
		return s + "!"
	})
	fmt.Printf("添加!: %v\n", upperStrings)
	fmt.Println()

	// 归约
	fmt.Println("归约 / Reduce:")
	numbers = []int{1, 2, 3, 4, 5}
	sum := sliceutils.ReduceInt(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("求和: %d\n", sum)

	product := sliceutils.ReduceInt([]int{2, 3, 4}, 1, func(acc, x int) int {
		return acc * x
	})
	fmt.Printf("求积: %d\n", product)
	fmt.Println()

	// 分页
	fmt.Println("分页 / Paginate:")
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	page1, totalPages, _ := sliceutils.PaginateInt(numbers, 1, 3)
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("第1页 (每页3条): %v (共%d页)\n", page1, totalPages)
	page2, _, _ := sliceutils.PaginateInt(numbers, 2, 3)
	fmt.Printf("第2页: %v\n", page2)
	page3, _, _ := sliceutils.PaginateInt(numbers, 3, 3)
	fmt.Printf("第3页: %v\n", page3)
	fmt.Println()

	// 集合操作
	fmt.Println("集合操作 / Set Operations:")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	fmt.Printf("切片1: %v\n", slice1)
	fmt.Printf("切片2: %v\n", slice2)

	intersection := sliceutils.IntersectionInt(slice1, slice2)
	fmt.Printf("交集: %v\n", intersection)

	union := sliceutils.UnionInt(slice1, slice2)
	fmt.Printf("并集: %v\n", union)

	difference := sliceutils.DifferenceInt(slice1, slice2)
	fmt.Printf("差集 (slice1 - slice2): %v\n", difference)
	fmt.Println()

	// 排序
	fmt.Println("排序 / Sort:")
	numbers = []int{3, 1, 4, 1, 5, 9, 2, 6}
	sorted := sliceutils.SortInt(numbers)
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("升序: %v\n", sorted)

	descSorted := sliceutils.SortIntDesc(numbers)
	fmt.Printf("降序: %v\n", descSorted)

	strings = []string{"banana", "apple", "cherry", "date"}
	sortedStrings := sliceutils.SortString(strings)
	fmt.Printf("原始: %v\n", strings)
	fmt.Printf("升序: %v\n", sortedStrings)
	fmt.Println()

	// 包含检查
	fmt.Println("包含检查 / Contains:")
	numbers = []int{1, 2, 3, 4, 5}
	fmt.Printf("切片: %v\n", numbers)
	fmt.Printf("包含 3: %v\n", sliceutils.ContainsInt(numbers, 3))
	fmt.Printf("包含 6: %v\n", sliceutils.ContainsInt(numbers, 6))

	strings = []string{"apple", "banana", "cherry"}
	fmt.Printf("切片: %v\n", strings)
	fmt.Printf("包含 \"banana\": %v\n", sliceutils.ContainsString(strings, "banana"))
	fmt.Printf("包含 \"orange\": %v\n", sliceutils.ContainsString(strings, "orange"))
	fmt.Println()

	// 反转
	fmt.Println("反转 / Reverse:")
	numbers = []int{1, 2, 3, 4, 5}
	reversed := sliceutils.ReverseInt(numbers)
	fmt.Printf("原始: %v\n", numbers)
	fmt.Printf("反转: %v\n", reversed)

	strings = []string{"a", "b", "c", "d"}
	reversedStrings := sliceutils.ReverseString(strings)
	fmt.Printf("原始: %v\n", strings)
	fmt.Printf("反转: %v\n", reversedStrings)
}
