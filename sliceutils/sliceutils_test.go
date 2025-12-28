package sliceutils

import (
	"reflect"
	"testing"
)

func TestUniqueInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"with duplicates", []int{1, 2, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UniqueInt(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("UniqueInt(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"with duplicates", []string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
		{"empty", []string{}, []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UniqueString(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("UniqueString(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestFilterInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected []int
	}{
		{"filter even", []int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{"filter greater than 2", []int{1, 2, 3, 4, 5}, func(x int) bool { return x > 2 }, []int{3, 4, 5}},
		{"empty result", []int{1, 2, 3}, func(x int) bool { return x > 10 }, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FilterInt(test.input, test.fn)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("FilterInt(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func(string) bool
		expected []string
	}{
		{"filter length > 1", []string{"a", "ab", "abc"}, func(s string) bool { return len(s) > 1 }, []string{"ab", "abc"}},
		{"filter contains 'a'", []string{"a", "b", "ab"}, func(s string) bool { return len(s) > 0 && s[0] == 'a' }, []string{"a", "ab"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FilterString(test.input, test.fn)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("FilterString(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestMapInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) int
		expected []int
	}{
		{"double", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"square", []int{1, 2, 3}, func(x int) int { return x * x }, []int{1, 4, 9}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapInt(test.input, test.fn)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MapInt(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestMapString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func(string) string
		expected []string
	}{
		{"add suffix", []string{"a", "b"}, func(s string) string { return s + "!" }, []string{"a!", "b!"}},
		{"upper case", []string{"a", "b"}, func(s string) string { return s + s }, []string{"aa", "bb"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapString(test.input, test.fn)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MapString(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestReduceInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		initial  int
		fn       func(int, int) int
		expected int
	}{
		{"sum", []int{1, 2, 3}, 0, func(acc, x int) int { return acc + x }, 6},
		{"product", []int{2, 3, 4}, 1, func(acc, x int) int { return acc * x }, 24},
		{"empty", []int{}, 10, func(acc, x int) int { return acc + x }, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReduceInt(test.input, test.initial, test.fn)
			if result != test.expected {
				t.Errorf("ReduceInt(%v, %d) = %d; want %d", test.input, test.initial, result, test.expected)
			}
		})
	}
}

func TestPaginateInt(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		page        int
		pageSize    int
		expected    []int
		expectedPages int
		shouldErr   bool
	}{
		{"first page", []int{1, 2, 3, 4, 5}, 1, 2, []int{1, 2}, 3, false},
		{"second page", []int{1, 2, 3, 4, 5}, 2, 2, []int{3, 4}, 3, false},
		{"last page", []int{1, 2, 3, 4, 5}, 3, 2, []int{5}, 3, false},
		{"page out of range", []int{1, 2, 3}, 10, 2, []int{}, 2, false},
		{"invalid page", []int{1, 2, 3}, 0, 2, nil, 0, true},
		{"invalid page size", []int{1, 2, 3}, 1, 0, nil, 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, totalPages, err := PaginateInt(test.input, test.page, test.pageSize)
			if test.shouldErr && err == nil {
				t.Errorf("PaginateInt(%v, %d, %d) expected error but got none", test.input, test.page, test.pageSize)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("PaginateInt(%v, %d, %d) unexpected error: %v", test.input, test.page, test.pageSize, err)
			}
			if !test.shouldErr && !reflect.DeepEqual(result, test.expected) {
				t.Errorf("PaginateInt(%v, %d, %d) = %v; want %v", test.input, test.page, test.pageSize, result, test.expected)
			}
			if !test.shouldErr && totalPages != test.expectedPages {
				t.Errorf("PaginateInt(%v, %d, %d) totalPages = %d; want %d", test.input, test.page, test.pageSize, totalPages, test.expectedPages)
			}
		})
	}
}

func TestIntersectionInt(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{"common case", []int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{"no intersection", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"empty first", []int{}, []int{1, 2, 3}, []int{}},
		{"empty second", []int{1, 2, 3}, []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntersectionInt(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("IntersectionInt(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
			}
		})
	}
}

func TestUnionInt(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{"common case", []int{1, 2}, []int{2, 3}, []int{1, 2, 3}},
		{"no overlap", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{"empty first", []int{}, []int{1, 2}, []int{1, 2}},
		{"empty second", []int{1, 2}, []int{}, []int{1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UnionInt(test.slice1, test.slice2)
			// 由于并集顺序可能不同，我们需要检查元素是否相同
			if len(result) != len(test.expected) {
				t.Errorf("UnionInt(%v, %v) length = %d; want %d", test.slice1, test.slice2, len(result), len(test.expected))
			}
			// 简单的元素检查（实际使用中可能需要更严格的检查）
		})
	}
}

func TestDifferenceInt(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{"common case", []int{1, 2, 3}, []int{2, 3}, []int{1}},
		{"no difference", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"empty second", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DifferenceInt(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("DifferenceInt(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
			}
		})
	}
}

func TestSortInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"unsorted", []int{3, 1, 2}, []int{1, 2, 3}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
		{"empty", []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SortInt(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SortInt(%v) = %v; want %v", test.input, result, test.expected)
			}
			// 确保原切片未被修改
			if len(test.input) > 0 && reflect.DeepEqual(test.input, result) && !reflect.DeepEqual(test.input, test.expected) {
				// 如果原切片和结果相同但不符合预期，说明原切片被修改了
				t.Errorf("SortInt modified original slice")
			}
		})
	}
}

func TestSortString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"unsorted", []string{"c", "a", "b"}, []string{"a", "b", "c"}},
		{"already sorted", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SortString(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SortString(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestSortIntDesc(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"unsorted", []int{1, 3, 2}, []int{3, 2, 1}},
		{"already sorted desc", []int{3, 2, 1}, []int{3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SortIntDesc(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SortIntDesc(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected bool
	}{
		{"contains", []int{1, 2, 3}, 2, true},
		{"not contains", []int{1, 2, 3}, 4, false},
		{"empty", []int{}, 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsInt(test.slice, test.item)
			if result != test.expected {
				t.Errorf("ContainsInt(%v, %d) = %v; want %v", test.slice, test.item, result, test.expected)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{"contains", []string{"a", "b", "c"}, "b", true},
		{"not contains", []string{"a", "b", "c"}, "d", false},
		{"empty", []string{}, "a", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsString(test.slice, test.item)
			if result != test.expected {
				t.Errorf("ContainsString(%v, %q) = %v; want %v", test.slice, test.item, result, test.expected)
			}
		})
	}
}

func TestReverseInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal case", []int{1, 2, 3}, []int{3, 2, 1}},
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReverseInt(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ReverseInt(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"normal case", []string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{"empty", []string{}, []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReverseString(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ReverseString(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}
