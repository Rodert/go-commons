package validationutils_test

import (
	"testing"

	"github.com/Rodert/go-commons/validationutils"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"test.name@example.co.uk", true},
		{"test+label@example.com", true},
		{"test@localhost", false},
		{"test@", false},
		{"@example.com", false},
		{"test", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsEmail(test.email)
		if result != test.expected {
			t.Errorf("IsEmail(%q) = %v; want %v", test.email, result, test.expected)
		}
	}
}

func TestIsCNMobile(t *testing.T) {
	tests := []struct {
		mobile   string
		expected bool
	}{
		{"13812345678", true},
		{"15912345678", true},
		{"17012345678", true},
		{"19912345678", true},
		{"12345678901", false},
		{"1381234567", false},   // 少一位
		{"138123456789", false}, // 多一位
		{"23812345678", false},  // 不是1开头
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsCNMobile(test.mobile)
		if result != test.expected {
			t.Errorf("IsCNMobile(%q) = %v; want %v", test.mobile, result, test.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://example.com", true},
		{"http://example.com/path", true},
		{"https://example.com/path?query=value", true},
		{"example.com", false},
		{"http://", false},
		{"http:/example.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsURL(test.url)
		if result != test.expected {
			t.Errorf("IsURL(%q) = %v; want %v", test.url, result, test.expected)
		}
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"255.255.255.255", true},
		{"0.0.0.0", true},
		{"256.0.0.1", false},
		{"192.168.1", false},
		{"192.168.1.1.1", false},
		{"192.168.1.a", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsIPv4(test.ip)
		if result != test.expected {
			t.Errorf("IsIPv4(%q) = %v; want %v", test.ip, result, test.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"123abc", false},
		{"abc123", false},
		{"abc", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsNumeric(test.str)
		if result != test.expected {
			t.Errorf("IsNumeric(%q) = %v; want %v", test.str, result, test.expected)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"abc", true},
		{"ABC", true},
		{"abcABC", true},
		{"abc123", false},
		{"123abc", false},
		{"123", false},
		{"abc-def", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsAlpha(test.str)
		if result != test.expected {
			t.Errorf("IsAlpha(%q) = %v; want %v", test.str, result, test.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"abc123", true},
		{"123abc", true},
		{"abc", true},
		{"123", true},
		{"abc-123", false},
		{"abc 123", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.IsAlphaNumeric(test.str)
		if result != test.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v; want %v", test.str, result, test.expected)
		}
	}
}

func TestHasMinLength(t *testing.T) {
	tests := []struct {
		str       string
		minLength int
		expected  bool
	}{
		{"abc", 3, true},
		{"abc", 2, true},
		{"abc", 4, false},
		{"", 1, false},
		{"", 0, true},
	}

	for _, test := range tests {
		result := validationutils.HasMinLength(test.str, test.minLength)
		if result != test.expected {
			t.Errorf("HasMinLength(%q, %d) = %v; want %v", test.str, test.minLength, result, test.expected)
		}
	}
}

func TestHasMaxLength(t *testing.T) {
	tests := []struct {
		str       string
		maxLength int
		expected  bool
	}{
		{"abc", 3, true},
		{"abc", 4, true},
		{"abc", 2, false},
		{"", 0, true},
	}

	for _, test := range tests {
		result := validationutils.HasMaxLength(test.str, test.maxLength)
		if result != test.expected {
			t.Errorf("HasMaxLength(%q, %d) = %v; want %v", test.str, test.maxLength, result, test.expected)
		}
	}
}

func TestContainsUppercase(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"ABC", true},
		{"abcD", true},
		{"Abc", true},
		{"abc", false},
		{"123", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.ContainsUppercase(test.str)
		if result != test.expected {
			t.Errorf("ContainsUppercase(%q) = %v; want %v", test.str, result, test.expected)
		}
	}
}

func TestContainsLowercase(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"abc", true},
		{"ABCd", true},
		{"aBC", true},
		{"ABC", false},
		{"123", false},
		{"", false},
	}

	for _, test := range tests {
		result := validationutils.ContainsLowercase(test.str)
		if result != test.expected {
			t.Errorf("ContainsLowercase(%q) = %v; want %v", test.str, result, test.expected)
		}
	}
}

func TestContainsSpecialChar(t *testing.T) {
	tests := []struct {
		str          string
		specialChars string
		expected     bool
	}{
		{"abc!", "", true},
		{"abc@123", "", true},
		{"abc#def", "", true},
		{"abc123", "", false},
		{"abcABC", "", false},
		{"abc123", "123", true},
		{"abcXYZ", "XYZ", true},
		{"", "", false},
	}

	for _, test := range tests {
		result := validationutils.ContainsSpecialChar(test.str, test.specialChars)
		if result != test.expected {
			t.Errorf("ContainsSpecialChar(%q, %q) = %v; want %v", test.str, test.specialChars, result, test.expected)
		}
	}
}
