package errorutils

import (
	"errors"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message string
		wantNil bool
	}{
		{"nil error", nil, "message", true},
		{"with error", errors.New("original"), "wrapped", false},
		{"empty message", errors.New("original"), "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.err, tt.message)
			if tt.wantNil && result != nil {
				t.Errorf("Wrap() = %v, want nil", result)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("Wrap() = nil, want error")
			}
			if result != nil {
				if !strings.Contains(result.Error(), tt.message) && tt.message != "" {
					t.Errorf("Wrap() error = %v, should contain %v", result.Error(), tt.message)
				}
			}
		})
	}
}

func TestWrapf(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		format  string
		args    []interface{}
		wantNil bool
	}{
		{"nil error", nil, "message %s", []interface{}{"test"}, true},
		{"with error", errors.New("original"), "wrapped %s", []interface{}{"test"}, false},
		{"no args", errors.New("original"), "wrapped", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result error
			if tt.args != nil {
				result = Wrapf(tt.err, tt.format, tt.args...)
			} else {
				// Use a constant format string when no args
				result = Wrapf(tt.err, "%s", tt.format)
			}
			if tt.wantNil && result != nil {
				t.Errorf("Wrapf() = %v, want nil", result)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("Wrapf() = nil, want error")
			}
			if result != nil && tt.args != nil {
				expected := "test"
				if !strings.Contains(result.Error(), expected) {
					t.Errorf("Wrapf() error = %v, should contain %v", result.Error(), expected)
				}
			}
		})
	}
}

func TestWithStack(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantNil bool
	}{
		{"nil error", nil, true},
		{"with error", errors.New("original"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WithStack(tt.err)
			if tt.wantNil && result != nil {
				t.Errorf("WithStack() = %v, want nil", result)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("WithStack() = nil, want error")
			}
			if result != nil {
				stack := StackTrace(result)
				if len(stack) == 0 {
					t.Errorf("WithStack() should have stack trace")
				}
			}
		})
	}
}

func TestNewWithType(t *testing.T) {
	tests := []struct {
		name    string
		errType ErrorType
		code    ErrorCode
		message string
	}{
		{"validation error", ErrorTypeValidation, ErrorCodeInvalidInput, "invalid input"},
		{"not found error", ErrorTypeNotFound, ErrorCodeNotFound, "not found"},
		{"empty message", ErrorTypeInternal, ErrorCodeInternal, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewWithType(tt.errType, tt.code, tt.message)
			if err == nil {
				t.Errorf("NewWithType() = nil, want error")
			}
			if GetType(err) != tt.errType {
				t.Errorf("GetType() = %v, want %v", GetType(err), tt.errType)
			}
			if GetCode(err) != tt.code {
				t.Errorf("GetCode() = %v, want %v", GetCode(err), tt.code)
			}
		})
	}
}

func TestWrapWithType(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		errType ErrorType
		code    ErrorCode
		message string
		wantNil bool
	}{
		{"nil error", nil, ErrorTypeValidation, ErrorCodeInvalidInput, "message", true},
		{"with error", errors.New("original"), ErrorTypeValidation, ErrorCodeInvalidInput, "wrapped", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapWithType(tt.err, tt.errType, tt.code, tt.message)
			if tt.wantNil && result != nil {
				t.Errorf("WrapWithType() = %v, want nil", result)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("WrapWithType() = nil, want error")
			}
			if result != nil {
				if GetType(result) != tt.errType {
					t.Errorf("GetType() = %v, want %v", GetType(result), tt.errType)
				}
				if GetCode(result) != tt.code {
					t.Errorf("GetCode() = %v, want %v", GetCode(result), tt.code)
				}
			}
		})
	}
}

func TestStackTrace(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"nil error", nil, 0},
		{"wrapped error", Wrap(errors.New("test"), "wrapped"), 1},
		{"regular error", errors.New("test"), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := StackTrace(tt.err)
			if len(stack) != tt.want {
				if tt.want > 0 && len(stack) == 0 {
					t.Errorf("StackTrace() = %v, want non-empty stack", stack)
				} else if tt.want == 0 && len(stack) > 0 {
					t.Errorf("StackTrace() = %v, want empty stack", stack)
				}
			}
		})
	}
}

func TestIsType(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		errType ErrorType
		want    bool
	}{
		{"nil error", nil, ErrorTypeValidation, false},
		{"matching type", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorTypeValidation, true},
		{"non-matching type", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorTypeNotFound, false},
		{"regular error", errors.New("test"), ErrorTypeValidation, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsType(tt.err, tt.errType)
			if result != tt.want {
				t.Errorf("IsType() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestIsCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code ErrorCode
		want bool
	}{
		{"nil error", nil, ErrorCodeInvalidInput, false},
		{"matching code", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorCodeInvalidInput, true},
		{"non-matching code", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorCodeNotFound, false},
		{"regular error", errors.New("test"), ErrorCodeInvalidInput, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCode(tt.err, tt.code)
			if result != tt.want {
				t.Errorf("IsCode() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want ErrorType
	}{
		{"nil error", nil, ErrorTypeUnknown},
		{"validation error", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorTypeValidation},
		{"regular error", errors.New("test"), ErrorTypeUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetType(tt.err)
			if result != tt.want {
				t.Errorf("GetType() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestGetCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want ErrorCode
	}{
		{"nil error", nil, ErrorCodeUnknown},
		{"with code", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "test"), ErrorCodeInvalidInput},
		{"regular error", errors.New("test"), ErrorCodeUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCode(tt.err)
			if result != tt.want {
				t.Errorf("GetCode() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestFormatError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		includeStack bool
		wantEmpty    bool
		wantStack    bool
	}{
		{"nil error", nil, false, true, false},
		{"wrapped error no stack", Wrap(errors.New("original"), "wrapped"), false, false, false},
		{"wrapped error with stack", Wrap(errors.New("original"), "wrapped"), true, false, true},
		{"regular error", errors.New("test"), false, false, false},
		{"typed error", NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "validation failed"), true, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatError(tt.err, tt.includeStack)
			if tt.wantEmpty && result != "" {
				t.Errorf("FormatError() = %v, want empty", result)
			}
			if !tt.wantEmpty && result == "" {
				t.Errorf("FormatError() = empty, want non-empty")
			}
			if tt.wantStack && !strings.Contains(result, "Stack trace:") {
				t.Errorf("FormatError() should contain stack trace")
			}
			if !tt.wantStack && strings.Contains(result, "Stack trace:") {
				t.Errorf("FormatError() should not contain stack trace")
			}
		})
	}
}

func TestErrorChain(t *testing.T) {
	// 测试错误链
	// Test error chain
	original := errors.New("original error")
	wrapped1 := Wrap(original, "first wrap")
	wrapped2 := Wrapf(wrapped1, "second wrap: %s", "test")

	// 测试 Unwrap
	// Test Unwrap
	if errors.Unwrap(wrapped2) != wrapped1 {
		t.Errorf("Unwrap() should return wrapped1")
	}
	if errors.Unwrap(wrapped1) != original {
		t.Errorf("Unwrap() should return original")
	}

	// 测试 Is
	// Test Is
	if !errors.Is(wrapped2, original) {
		t.Errorf("errors.Is() should find original error")
	}
	if !errors.Is(wrapped2, wrapped1) {
		t.Errorf("errors.Is() should find wrapped1")
	}
}

func TestWrappedError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *WrappedError
		contains []string
	}{
		{"with message and err", &WrappedError{Err: errors.New("original"), Message: "wrapped"}, []string{"wrapped", "original"}},
		{"only message", &WrappedError{Message: "wrapped"}, []string{"wrapped"}},
		{"only err", &WrappedError{Err: errors.New("original")}, []string{"original"}},
		{"empty", &WrappedError{}, []string{"unknown error"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("Error() = %v, should contain %v", result, want)
				}
			}
		})
	}
}

