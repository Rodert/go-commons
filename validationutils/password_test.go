package validationutils_test

import (
	"testing"

	"github.com/Rodert/go-commons/validationutils"
)

func TestCheckPasswordStrength(t *testing.T) {
	tests := []struct {
		password string
		minScore int
	}{
		{"a", 0},
		{"password", 0},
		{"Password1", 30},
		{"Password123", 40},
		{"Password123!", 60},
		{"P@ssw0rd123!XyZ", 70},
	}

	for _, test := range tests {
		result := validationutils.CheckPasswordStrength(test.password)

		if result.Score < test.minScore {
			t.Errorf("CheckPasswordStrength(%q).Score = %v; want at least %v",
				test.password, result.Score, test.minScore)
		}

		// 确保结果中包含建议字段
		if result.Suggestions == nil {
			t.Errorf("CheckPasswordStrength(%q) should return a non-nil Suggestions slice", test.password)
		}
	}
}

func TestIsPasswordValid(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		minLength      int
		requireUpper   bool
		requireLower   bool
		requireDigit   bool
		requireSpecial bool
		expectedValid  bool
		reasonCount    int
	}{
		{
			name:           "Valid password with all requirements",
			password:       "Password123!",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  true,
			reasonCount:    0,
		},
		{
			name:           "Too short",
			password:       "Pass1!",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    1,
		},
		{
			name:           "No uppercase",
			password:       "password123!",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    1,
		},
		{
			name:           "No lowercase",
			password:       "PASSWORD123!",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    1,
		},
		{
			name:           "No digit",
			password:       "Password!",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    1,
		},
		{
			name:           "No special character",
			password:       "Password123",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    1,
		},
		{
			name:           "Multiple failures",
			password:       "pass",
			minLength:      8,
			requireUpper:   true,
			requireLower:   true,
			requireDigit:   true,
			requireSpecial: true,
			expectedValid:  false,
			reasonCount:    4,
		},
		{
			name:           "Valid with minimal requirements",
			password:       "password",
			minLength:      8,
			requireUpper:   false,
			requireLower:   true,
			requireDigit:   false,
			requireSpecial: false,
			expectedValid:  true,
			reasonCount:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid, reasons := validationutils.IsPasswordValid(
				test.password,
				test.minLength,
				test.requireUpper,
				test.requireLower,
				test.requireDigit,
				test.requireSpecial,
			)

			if valid != test.expectedValid {
				t.Errorf("IsPasswordValid() = %v; want %v", valid, test.expectedValid)
			}

			if len(reasons) != test.reasonCount {
				t.Errorf("IsPasswordValid() returned %d reasons; want %d. Reasons: %v",
					len(reasons), test.reasonCount, reasons)
			}
		})
	}
}
