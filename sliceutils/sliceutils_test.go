package sliceutils

import (
	"reflect"
	"testing"
)

// TestUnique 测试泛型 Unique 函数
func TestUnique(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
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
				result := Unique(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Unique(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
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
				result := Unique(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Unique(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})
}

// TestFilter 测试泛型 Filter 函数
func TestFilter(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
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
				result := Filter(test.input, test.fn)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Filter(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
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
				result := Filter(test.input, test.fn)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Filter(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})
}

// TestMap 测试泛型 Map 函数
func TestMap(t *testing.T) {
	t.Run("int to int", func(t *testing.T) {
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
				result := Map(test.input, test.fn)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Map(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("string to string", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []string
			fn       func(string) string
			expected []string
		}{
			{"add suffix", []string{"a", "b"}, func(s string) string { return s + "!" }, []string{"a!", "b!"}},
			{"duplicate", []string{"a", "b"}, func(s string) string { return s + s }, []string{"aa", "bb"}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := Map(test.input, test.fn)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Map(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})
}

// TestReduce 测试泛型 Reduce 函数
func TestReduce(t *testing.T) {
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
			result := Reduce(test.input, test.initial, test.fn)
			if result != test.expected {
				t.Errorf("Reduce(%v, %d) = %d; want %d", test.input, test.initial, result, test.expected)
			}
		})
	}
}

// TestPaginate 测试泛型 Paginate 函数
func TestPaginate(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		page          int
		pageSize      int
		expected      []int
		expectedPages int
		shouldErr     bool
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
			result, totalPages, err := Paginate(test.input, test.page, test.pageSize)
			if test.shouldErr && err == nil {
				t.Errorf("Paginate(%v, %d, %d) expected error but got none", test.input, test.page, test.pageSize)
			}
			if !test.shouldErr && err != nil {
				t.Errorf("Paginate(%v, %d, %d) unexpected error: %v", test.input, test.page, test.pageSize, err)
			}
			if !test.shouldErr && !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Paginate(%v, %d, %d) = %v; want %v", test.input, test.page, test.pageSize, result, test.expected)
			}
			if !test.shouldErr && totalPages != test.expectedPages {
				t.Errorf("Paginate(%v, %d, %d) totalPages = %d; want %d", test.input, test.page, test.pageSize, totalPages, test.expectedPages)
			}
		})
	}
}

// TestIntersection 测试泛型 Intersection 函数
func TestIntersection(t *testing.T) {
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
			result := Intersection(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Intersection(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
			}
		})
	}
}

// TestUnion 测试泛型 Union 函数
func TestUnion(t *testing.T) {
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
			result := Union(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Union(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
			}
		})
	}
}

// TestDifference 测试泛型 Difference 函数
func TestDifference(t *testing.T) {
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
			result := Difference(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Difference(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
			}
		})
	}
}

// TestSort 测试泛型 Sort 函数
func TestSort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
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
				result := Sort(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Sort(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
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
				result := Sort(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Sort(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})
}

// TestSortDesc 测试泛型 SortDesc 函数
func TestSortDesc(t *testing.T) {
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
			result := SortDesc(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SortDesc(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// TestContains 测试泛型 Contains 函数
func TestContains(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
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
				result := Contains(test.slice, test.item)
				if result != test.expected {
					t.Errorf("Contains(%v, %d) = %v; want %v", test.slice, test.item, result, test.expected)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
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
				result := Contains(test.slice, test.item)
				if result != test.expected {
					t.Errorf("Contains(%v, %q) = %v; want %v", test.slice, test.item, result, test.expected)
				}
			})
		}
	})
}

// TestReverse 测试泛型 Reverse 函数
func TestReverse(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
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
				result := Reverse(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Reverse(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
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
				result := Reverse(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Reverse(%v) = %v; want %v", test.input, result, test.expected)
				}
			})
		}
	})
}
