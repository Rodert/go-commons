package stringutils

import (
	"testing"
)

// 测试 IsEmpty 函数
// Test IsEmpty function
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"whitespace string", " \t\n", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsEmpty(test.input)
			if result != test.expected {
				t.Errorf("IsEmpty(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// 测试 IsNotEmpty 函数
// Test IsNotEmpty function
func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"non-empty string", "hello", true},
		{"whitespace string", " \t\n", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsNotEmpty(test.input)
			if result != test.expected {
				t.Errorf("IsNotEmpty(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// 测试 IsBlank 函数
// Test IsBlank function
func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"whitespace only", " \t\n\r", true},
		{"non-blank string", "hello", false},
		{"string with spaces", "  hello  ", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBlank(test.input)
			if result != test.expected {
				t.Errorf("IsBlank(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// 测试 IsNotBlank 函数
// Test IsNotBlank function
func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"whitespace only", " \t\n\r", false},
		{"non-blank string", "hello", true},
		{"string with spaces", "  hello  ", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsNotBlank(test.input)
			if result != test.expected {
				t.Errorf("IsNotBlank(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// 测试 Trim 函数
// Test Trim function
func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"whitespace only", " \t\n\r", ""},
		{"no whitespace", "hello", "hello"},
		{"leading whitespace", "  hello", "hello"},
		{"trailing whitespace", "hello  ", "hello"},
		{"both sides whitespace", "  hello  ", "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Trim(test.input)
			if result != test.expected {
				t.Errorf("Trim(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 TrimToEmpty 函数
// Test TrimToEmpty function
func TestTrimToEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"whitespace only", " \t\n\r", ""},
		{"no whitespace", "hello", "hello"},
		{"leading whitespace", "  hello", "hello"},
		{"trailing whitespace", "hello  ", "hello"},
		{"both sides whitespace", "  hello  ", "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TrimToEmpty(test.input)
			if result != test.expected {
				t.Errorf("TrimToEmpty(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 Truncate 函数
// Test Truncate function
func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxWidth int
		expected string
	}{
		{"empty string", "", 5, ""},
		{"negative width", "hello", -1, ""},
		{"zero width", "hello", 0, ""},
		{"width less than length", "hello", 3, "hel"},
		{"width equal to length", "hello", 5, "hello"},
		{"width greater than length", "hello", 10, "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Truncate(test.input, test.maxWidth)
			if result != test.expected {
				t.Errorf("Truncate(%q, %d) = %q; want %q", test.input, test.maxWidth, result, test.expected)
			}
		})
	}
}

// 测试 TruncateWithSuffix 函数
// Test TruncateWithSuffix function
func TestTruncateWithSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxWidth int
		suffix   string
		expected string
	}{
		{"empty string", "", 5, "...", ""},
		{"negative width", "hello", -1, "...", ""},
		{"zero width", "hello", 0, "...", ""},
		{"width less than length", "hello world", 8, "...", "hello..."},
		{"width equal to length", "hello", 5, "...", "hello"},
		{"width greater than length", "hello", 10, "...", "hello"},
		{"empty suffix", "hello world", 8, "", "hello wo"},
		{"suffix longer than maxWidth", "hello", 2, "...", ".."},
		{"suffix equal to maxWidth", "hello", 3, "...", "..."},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TruncateWithSuffix(test.input, test.maxWidth, test.suffix)
			if result != test.expected {
				t.Errorf("TruncateWithSuffix(%q, %d, %q) = %q; want %q",
					test.input, test.maxWidth, test.suffix, result, test.expected)
			}
		})
	}
}

// 测试 Capitalize 函数
func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single lowercase", "a", "A"},
		{"single uppercase", "A", "A"},
		{"lowercase word", "hello", "Hello"},
		{"uppercase word", "HELLO", "HELLO"},
		{"mixed case", "hELLo", "HELLo"},
		{"with spaces", "hello world", "Hello world"},
		{"unicode", "你好", "你好"},
		{"number first", "123abc", "123abc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Capitalize(test.input)
			if result != test.expected {
				t.Errorf("Capitalize(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 Uncapitalize 函数
func TestUncapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single lowercase", "a", "a"},
		{"single uppercase", "A", "a"},
		{"lowercase word", "hello", "hello"},
		{"uppercase word", "HELLO", "hELLO"},
		{"mixed case", "HELLo", "hELLo"},
		{"with spaces", "Hello World", "hello World"},
		{"unicode", "你好", "你好"},
		{"number first", "123ABC", "123ABC"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Uncapitalize(test.input)
			if result != test.expected {
				t.Errorf("Uncapitalize(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 ReverseString 函数
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"simple string", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode characters", "你好", "好你"},
		{"mixed unicode", "hello世界", "界世olleh"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReverseString(test.input)
			if result != test.expected {
				t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 ContainsAny 函数
func TestContainsAny(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		searchStrs []string
		expected  bool
	}{
		{"empty string", "", []string{"a"}, false},
		{"empty search strings", "hello", []string{}, false},
		{"contains first", "hello world", []string{"hello", "foo"}, true},
		{"contains second", "hello world", []string{"foo", "world"}, true},
		{"contains none", "hello world", []string{"foo", "bar"}, false},
		{"single match", "test", []string{"es"}, true},
		{"no match", "test", []string{"xyz"}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsAny(test.str, test.searchStrs...)
			if result != test.expected {
				t.Errorf("ContainsAny(%q, %v) = %v; want %v", test.str, test.searchStrs, result, test.expected)
			}
		})
	}
}

// 测试 ContainsAll 函数
func TestContainsAll(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		searchStrs []string
		expected  bool
	}{
		{"empty string", "", []string{"a"}, false},
		{"empty search strings", "hello", []string{}, false},
		{"contains all", "hello world", []string{"hello", "world"}, true},
		{"contains some", "hello world", []string{"hello", "foo"}, false},
		{"contains none", "hello world", []string{"foo", "bar"}, false},
		{"single match", "test", []string{"es"}, true},
		{"multiple matches", "hello world test", []string{"hello", "world", "test"}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsAll(test.str, test.searchStrs...)
			if result != test.expected {
				t.Errorf("ContainsAll(%q, %v) = %v; want %v", test.str, test.searchStrs, result, test.expected)
			}
		})
	}
}

// 测试 SubstringBefore 函数
func TestSubstringBefore(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		separator string
		expected  string
	}{
		{"empty string", "", "/", ""},
		{"empty separator", "hello", "", "hello"},
		{"found separator", "hello/world", "/", "hello"},
		{"not found separator", "hello", "/", "hello"},
		{"multiple separators", "a/b/c", "/", "a"},
		{"separator at start", "/hello", "/", ""},
		{"separator at end", "hello/", "/", "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SubstringBefore(test.str, test.separator)
			if result != test.expected {
				t.Errorf("SubstringBefore(%q, %q) = %q; want %q", test.str, test.separator, result, test.expected)
			}
		})
	}
}

// 测试 SubstringAfter 函数
func TestSubstringAfter(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		separator string
		expected  string
	}{
		{"empty string", "", "/", ""},
		{"empty separator", "hello", "", ""},
		{"found separator", "hello/world", "/", "world"},
		{"not found separator", "hello", "/", ""},
		{"multiple separators", "a/b/c", "/", "b/c"},
		{"separator at start", "/hello", "/", "hello"},
		{"separator at end", "hello/", "/", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SubstringAfter(test.str, test.separator)
			if result != test.expected {
				t.Errorf("SubstringAfter(%q, %q) = %q; want %q", test.str, test.separator, result, test.expected)
			}
		})
	}
}

// 测试 Join 函数
func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		separator string
		elements  []string
		expected  string
	}{
		{"empty elements", ",", []string{}, ""},
		{"single element", ",", []string{"a"}, "a"},
		{"two elements", ",", []string{"a", "b"}, "a,b"},
		{"multiple elements", "-", []string{"a", "b", "c"}, "a-b-c"},
		{"empty separator", "", []string{"a", "b"}, "ab"},
		{"space separator", " ", []string{"hello", "world"}, "hello world"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Join(test.separator, test.elements...)
			if result != test.expected {
				t.Errorf("Join(%q, %v) = %q; want %q", test.separator, test.elements, result, test.expected)
			}
		})
	}
}

// 测试 Split 函数
func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		separator string
		expected  []string
	}{
		{"empty string", "", ",", []string{}},
		{"single element", "a", ",", []string{"a"}},
		{"two elements", "a,b", ",", []string{"a", "b"}},
		{"multiple elements", "a,b,c", ",", []string{"a", "b", "c"}},
		{"empty separator", "ab", "", []string{"a", "b"}},
		{"not found separator", "hello", ",", []string{"hello"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Split(test.str, test.separator)
			if len(result) != len(test.expected) {
				t.Errorf("Split(%q, %q) length = %d; want %d", test.str, test.separator, len(result), len(test.expected))
				return
			}
			for i, v := range result {
				if v != test.expected[i] {
					t.Errorf("Split(%q, %q)[%d] = %q; want %q", test.str, test.separator, i, v, test.expected[i])
				}
			}
		})
	}
}

// 测试 EqualsIgnoreCase 函数
func TestEqualsIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{"both empty", "", "", true},
		{"same case", "hello", "hello", true},
		{"different case", "hello", "HELLO", true},
		{"mixed case", "Hello", "hELLo", true},
		{"different strings", "hello", "world", false},
		{"same length different", "abc", "def", false},
		{"with spaces", "Hello World", "hello world", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EqualsIgnoreCase(test.str1, test.str2)
			if result != test.expected {
				t.Errorf("EqualsIgnoreCase(%q, %q) = %v; want %v", test.str1, test.str2, result, test.expected)
			}
		})
	}
}

// 测试 StartsWith 函数
func TestStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		prefix   string
		expected bool
	}{
		{"both empty", "", "", true},
		{"empty prefix", "hello", "", true},
		{"matches", "hello", "he", true},
		{"doesn't match", "hello", "lo", false},
		{"exact match", "hello", "hello", true},
		{"longer prefix", "hello", "helloworld", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StartsWith(test.str, test.prefix)
			if result != test.expected {
				t.Errorf("StartsWith(%q, %q) = %v; want %v", test.str, test.prefix, result, test.expected)
			}
		})
	}
}

// 测试 EndsWith 函数
func TestEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		suffix   string
		expected bool
	}{
		{"both empty", "", "", true},
		{"empty suffix", "hello", "", true},
		{"matches", "hello", "lo", true},
		{"doesn't match", "hello", "he", false},
		{"exact match", "hello", "hello", true},
		{"longer suffix", "hello", "helloworld", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := EndsWith(test.str, test.suffix)
			if result != test.expected {
				t.Errorf("EndsWith(%q, %q) = %v; want %v", test.str, test.suffix, result, test.expected)
			}
		})
	}
}

// 测试 RemoveStart 函数
func TestRemoveStart(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		remove   string
		expected string
	}{
		{"empty string", "", "pre", ""},
		{"empty remove", "hello", "", "hello"},
		{"matches", "prefixhello", "prefix", "hello"},
		{"doesn't match", "hello", "pre", "hello"},
		{"exact match", "hello", "hello", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveStart(test.str, test.remove)
			if result != test.expected {
				t.Errorf("RemoveStart(%q, %q) = %q; want %q", test.str, test.remove, result, test.expected)
			}
		})
	}
}

// 测试 RemoveEnd 函数
func TestRemoveEnd(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		remove   string
		expected string
	}{
		{"empty string", "", "suf", ""},
		{"empty remove", "hello", "", "hello"},
		{"matches", "hellosuffix", "suffix", "hello"},
		{"doesn't match", "hello", "suf", "hello"},
		{"exact match", "hello", "hello", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveEnd(test.str, test.remove)
			if result != test.expected {
				t.Errorf("RemoveEnd(%q, %q) = %q; want %q", test.str, test.remove, result, test.expected)
			}
		})
	}
}

// 测试 Replace 函数
func TestReplace(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		oldStr   string
		newStr   string
		count    int
		expected string
	}{
		{"empty string", "", "old", "new", -1, ""},
		{"empty oldStr", "hello", "", "new", -1, "hello"},
		{"replace all", "hello hello", "hello", "world", -1, "world world"},
		{"replace once", "hello hello", "hello", "world", 1, "world hello"},
		{"replace twice", "hello hello hello", "hello", "world", 2, "world world hello"},
		{"not found", "hello", "xyz", "world", -1, "hello"},
		{"same old and new", "hello", "hello", "hello", -1, "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Replace(test.str, test.oldStr, test.newStr, test.count)
			if result != test.expected {
				t.Errorf("Replace(%q, %q, %q, %d) = %q; want %q",
					test.str, test.oldStr, test.newStr, test.count, result, test.expected)
			}
		})
	}
}

// 测试 ReplaceAll 函数
func TestReplaceAll(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		oldStr   string
		newStr   string
		expected string
	}{
		{"empty string", "", "old", "new", ""},
		{"empty oldStr", "hello", "", "new", "hello"},
		{"replace all", "hello hello", "hello", "world", "world world"},
		{"not found", "hello", "xyz", "world", "hello"},
		{"same old and new", "hello", "hello", "hello", "hello"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ReplaceAll(test.str, test.oldStr, test.newStr)
			if result != test.expected {
				t.Errorf("ReplaceAll(%q, %q, %q) = %q; want %q",
					test.str, test.oldStr, test.newStr, result, test.expected)
			}
		})
	}
}

// 测试 Repeat 函数
func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		repeat   int
		expected string
	}{
		{"empty string", "", 5, ""},
		{"zero repeat", "hello", 0, ""},
		{"negative repeat", "hello", -1, ""},
		{"single repeat", "hello", 1, "hello"},
		{"multiple repeat", "ab", 3, "ababab"},
		{"large repeat", "x", 5, "xxxxx"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Repeat(test.str, test.repeat)
			if result != test.expected {
				t.Errorf("Repeat(%q, %d) = %q; want %q", test.str, test.repeat, result, test.expected)
			}
		})
	}
}

// 测试 PadLeft 函数
func TestPadLeft(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		size     int
		padChar  rune
		expected string
	}{
		{"empty string", "", 5, '*', "*****"},
		{"size zero", "hello", 0, '*', "hello"},
		{"size negative", "hello", -1, '*', "hello"},
		{"size equal to length", "hello", 5, '*', "hello"},
		{"size less than length", "hello", 3, '*', "hello"},
		{"normal padding", "hi", 5, '*', "***hi"},
		{"unicode", "中文", 5, '*', "***中文"},
		{"emoji", "😀", 5, '*', "****😀"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PadLeft(test.str, test.size, test.padChar)
			if result != test.expected {
				t.Errorf("PadLeft(%q, %d, %c) = %q; want %q",
					test.str, test.size, test.padChar, result, test.expected)
			}
		})
	}
}

// 测试 PadRight 函数
func TestPadRight(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		size     int
		padChar  rune
		expected string
	}{
		{"empty string", "", 5, '*', "*****"},
		{"size zero", "hello", 0, '*', "hello"},
		{"size negative", "hello", -1, '*', "hello"},
		{"size equal to length", "hello", 5, '*', "hello"},
		{"size less than length", "hello", 3, '*', "hello"},
		{"normal padding", "hi", 5, '*', "hi***"},
		{"unicode", "中文", 5, '*', "中文***"},
		{"emoji", "😀", 5, '*', "😀****"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PadRight(test.str, test.size, test.padChar)
			if result != test.expected {
				t.Errorf("PadRight(%q, %d, %c) = %q; want %q",
					test.str, test.size, test.padChar, result, test.expected)
			}
		})
	}
}

// 测试 Center 函数
func TestCenter(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		size     int
		padChar  rune
		expected string
	}{
		{"empty string", "", 5, '*', "*****"},
		{"size zero", "hello", 0, '*', "hello"},
		{"size negative", "hello", -1, '*', "hello"},
		{"size equal to length", "hello", 5, '*', "hello"},
		{"size less than length", "hello", 3, '*', "hello"},
		{"odd padding", "hi", 5, '*', "*hi**"},
		{"even padding", "hi", 6, '*', "**hi**"},
		{"unicode", "中文", 5, '*', "*中文**"},
		{"emoji", "😀", 5, '*', "**😀**"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Center(test.str, test.size, test.padChar)
			if result != test.expected {
				t.Errorf("Center(%q, %d, %c) = %q; want %q",
					test.str, test.size, test.padChar, result, test.expected)
			}
		})
	}
}

// 测试 CountMatches 函数
func TestCountMatches(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		sub      string
		expected int
	}{
		{"empty string", "", "a", 0},
		{"empty sub", "hello", "", 0},
		{"no match", "hello", "xyz", 0},
		{"single match", "hello", "ll", 1},
		{"multiple matches", "hello hello", "hello", 2},
		{"overlapping", "aaa", "aa", 1}, // strings.Count counts non-overlapping matches
		{"exact match", "hello", "hello", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountMatches(test.str, test.sub)
			if result != test.expected {
				t.Errorf("CountMatches(%q, %q) = %d; want %d", test.str, test.sub, result, test.expected)
			}
		})
	}
}

// 测试 ToUpperCase 函数
func TestToUpperCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"all lowercase", "hello", "HELLO"},
		{"all uppercase", "HELLO", "HELLO"},
		{"mixed case", "Hello", "HELLO"},
		{"with numbers", "hello123", "HELLO123"},
		{"with spaces", "hello world", "HELLO WORLD"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToUpperCase(test.input)
			if result != test.expected {
				t.Errorf("ToUpperCase(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 ToLowerCase 函数
func TestToLowerCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"all lowercase", "hello", "hello"},
		{"all uppercase", "HELLO", "hello"},
		{"mixed case", "Hello", "hello"},
		{"with numbers", "HELLO123", "hello123"},
		{"with spaces", "HELLO WORLD", "hello world"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLowerCase(test.input)
			if result != test.expected {
				t.Errorf("ToLowerCase(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 DefaultIfEmpty 函数
func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		defaultStr string
		expected  string
	}{
		{"empty string", "", "default", "default"},
		{"non-empty string", "hello", "default", "hello"},
		{"empty default", "", "", ""},
		{"whitespace", "  ", "default", "  "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DefaultIfEmpty(test.str, test.defaultStr)
			if result != test.expected {
				t.Errorf("DefaultIfEmpty(%q, %q) = %q; want %q",
					test.str, test.defaultStr, result, test.expected)
			}
		})
	}
}

// 测试 DefaultIfBlank 函数
func TestDefaultIfBlank(t *testing.T) {
	tests := []struct {
		name       string
		str        string
		defaultStr string
		expected   string
	}{
		{"empty string", "", "default", "default"},
		{"whitespace only", "   ", "default", "default"},
		{"non-blank string", "hello", "default", "hello"},
		{"string with spaces", "  hello  ", "default", "  hello  "},
		{"tab and newline", "\t\n", "default", "default"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DefaultIfBlank(test.str, test.defaultStr)
			if result != test.expected {
				t.Errorf("DefaultIfBlank(%q, %q) = %q; want %q",
					test.str, test.defaultStr, result, test.expected)
			}
		})
	}
}
