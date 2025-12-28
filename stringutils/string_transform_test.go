package stringutils

import (
	"testing"
)

// 测试 Reverse 函数
// Test Reverse function
func TestReverse(t *testing.T) {
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
		{"emoji", "😀😁😂", "😂😁😀"},
		{"mixed content", "a1b2c3", "3c2b1a"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Reverse(test.input)
			if result != test.expected {
				t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 SwapCase 函数
// Test SwapCase function
func TestSwapCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"all lowercase", "hello", "HELLO"},
		{"all uppercase", "HELLO", "hello"},
		{"mixed case", "Hello World", "hELLO wORLD"},
		{"with numbers", "Hello123", "hELLO123"},
		{"with special chars", "Hello!@#", "hELLO!@#"},
		{"single char lowercase", "a", "A"},
		{"single char uppercase", "A", "a"},
		{"mixed unicode", "Hello世界", "hELLO世界"},
		{"numbers only", "123", "123"},
		{"special chars only", "!@#", "!@#"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SwapCase(test.input)
			if result != test.expected {
				t.Errorf("SwapCase(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

// 测试 PadCenter 函数
// Test PadCenter function
func TestPadCenter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
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
		{"single char", "a", 5, '*', "**a**"},
		{"with unicode", "中文", 5, '*', "*中文**"},
		{"with emoji", "😀", 5, '*', "**😀**"},
		{"large padding", "x", 10, '-', "----x-----"},
		{"space padding", "test", 10, ' ', "   test   "},
		{"number padding", "abc", 7, '0', "00abc00"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PadCenter(test.input, test.size, test.padChar)
			if result != test.expected {
				t.Errorf("PadCenter(%q, %d, %c) = %q; want %q",
					test.input, test.size, test.padChar, result, test.expected)
			}
			// 验证结果长度（字符数）应该等于size（如果size > 0且input长度 <= size）
			if test.size > 0 {
				actualRuneCount := len([]rune(result))
				inputRuneCount := len([]rune(test.input))
				if inputRuneCount <= test.size && actualRuneCount != test.size {
					t.Errorf("PadCenter(%q, %d, %c) result length = %d; want %d",
						test.input, test.size, test.padChar, actualRuneCount, test.size)
				}
			}
		})
	}
}

