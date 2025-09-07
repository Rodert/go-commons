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