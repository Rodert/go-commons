// Package errorutils 提供错误处理相关的工具函数
// Package errorutils provides error handling utility functions
package errorutils

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// ErrorType 错误类型
// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeUnknown 未知错误类型
	// ErrorTypeUnknown represents unknown error type
	ErrorTypeUnknown ErrorType = "unknown"
	// ErrorTypeValidation 验证错误类型
	// ErrorTypeValidation represents validation error type
	ErrorTypeValidation ErrorType = "validation"
	// ErrorTypeNotFound 未找到错误类型
	// ErrorTypeNotFound represents not found error type
	ErrorTypeNotFound ErrorType = "not_found"
	// ErrorTypePermission 权限错误类型
	// ErrorTypePermission represents permission error type
	ErrorTypePermission ErrorType = "permission"
	// ErrorTypeNetwork 网络错误类型
	// ErrorTypeNetwork represents network error type
	ErrorTypeNetwork ErrorType = "network"
	// ErrorTypeTimeout 超时错误类型
	// ErrorTypeTimeout represents timeout error type
	ErrorTypeTimeout ErrorType = "timeout"
	// ErrorTypeInternal 内部错误类型
	// ErrorTypeInternal represents internal error type
	ErrorTypeInternal ErrorType = "internal"
)

// ErrorCode 错误代码
// ErrorCode represents error code
type ErrorCode int

const (
	// ErrorCodeUnknown 未知错误代码
	// ErrorCodeUnknown represents unknown error code
	ErrorCodeUnknown ErrorCode = 0
	// ErrorCodeInvalidInput 无效输入错误代码
	// ErrorCodeInvalidInput represents invalid input error code
	ErrorCodeInvalidInput ErrorCode = 1000
	// ErrorCodeNotFound 未找到错误代码
	// ErrorCodeNotFound represents not found error code
	ErrorCodeNotFound ErrorCode = 1001
	// ErrorCodeUnauthorized 未授权错误代码
	// ErrorCodeUnauthorized represents unauthorized error code
	ErrorCodeUnauthorized ErrorCode = 1002
	// ErrorCodeForbidden 禁止访问错误代码
	// ErrorCodeForbidden represents forbidden error code
	ErrorCodeForbidden ErrorCode = 1003
	// ErrorCodeTimeout 超时错误代码
	// ErrorCodeTimeout represents timeout error code
	ErrorCodeTimeout ErrorCode = 1004
	// ErrorCodeInternal 内部错误代码
	// ErrorCodeInternal represents internal error code
	ErrorCodeInternal ErrorCode = 5000
)

// WrappedError 包装的错误，包含原始错误、消息和堆栈信息
// WrappedError wraps an error with message and stack trace
type WrappedError struct {
	// Err 原始错误 / original error
	Err error
	// Message 错误消息 / error message
	Message string
	// Stack 堆栈跟踪 / stack trace
	Stack []string
	// Type 错误类型 / error type
	Type ErrorType
	// Code 错误代码 / error code
	Code ErrorCode
}

// Error 实现 error 接口
// Error implements the error interface
func (e *WrappedError) Error() string {
	if e.Message != "" {
		if e.Err != nil {
			return fmt.Sprintf("%s: %v", e.Message, e.Err)
		}
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return "unknown error"
}

// Unwrap 返回原始错误，用于 errors.Unwrap
// Unwrap returns the original error for errors.Unwrap
func (e *WrappedError) Unwrap() error {
	return e.Err
}

// Wrap 包装一个错误，添加消息和堆栈跟踪
//
// 参数 / Parameters:
//   - err: 要包装的错误，如果为nil则返回nil / error to wrap, returns nil if err is nil
//   - message: 错误消息 / error message
//
// 返回值 / Returns:
//   - error: 包装后的错误 / wrapped error
//
// 示例 / Example:
//
//	err := errors.New("original error")
//	wrapped := Wrap(err, "failed to process")
//
// Wrap wraps an error with a message and stack trace
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &WrappedError{
		Err:     err,
		Message: message,
		Stack:   captureStack(2),
		Type:    ErrorTypeUnknown,
		Code:    ErrorCodeUnknown,
	}
}

// Wrapf 使用格式化字符串包装错误
//
// 参数 / Parameters:
//   - err: 要包装的错误，如果为nil则返回nil / error to wrap, returns nil if err is nil
//   - format: 格式化字符串 / format string
//   - args: 格式化参数 / format arguments
//
// 返回值 / Returns:
//   - error: 包装后的错误 / wrapped error
//
// 示例 / Example:
//
//	err := errors.New("original error")
//	wrapped := Wrapf(err, "failed to process %s", "file.txt")
//
// Wrapf wraps an error with a formatted message
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &WrappedError{
		Err:     err,
		Message: fmt.Sprintf(format, args...),
		Stack:   captureStack(2),
		Type:    ErrorTypeUnknown,
		Code:    ErrorCodeUnknown,
	}
}

// WithStack 为错误添加堆栈跟踪
//
// 参数 / Parameters:
//   - err: 要添加堆栈的错误，如果为nil则返回nil / error to add stack to, returns nil if err is nil
//
// 返回值 / Returns:
//   - error: 带有堆栈跟踪的错误 / error with stack trace
//
// 示例 / Example:
//
//	err := errors.New("some error")
//	errWithStack := WithStack(err)
//
// WithStack adds stack trace to an error
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	// 如果已经是 WrappedError，直接返回
	// If already a WrappedError, return as is
	if wrapped, ok := err.(*WrappedError); ok {
		return wrapped
	}
	return &WrappedError{
		Err:     err,
		Message: "",
		Stack:   captureStack(2),
		Type:    ErrorTypeUnknown,
		Code:    ErrorCodeUnknown,
	}
}

// NewWithType 创建一个带类型的错误
//
// 参数 / Parameters:
//   - errType: 错误类型 / error type
//   - code: 错误代码 / error code
//   - message: 错误消息 / error message
//
// 返回值 / Returns:
//   - error: 新创建的错误 / newly created error
//
// 示例 / Example:
//
//	err := NewWithType(ErrorTypeValidation, ErrorCodeInvalidInput, "invalid input")
//
// NewWithType creates a new error with type and code
func NewWithType(errType ErrorType, code ErrorCode, message string) error {
	return &WrappedError{
		Err:     nil,
		Message: message,
		Stack:   captureStack(2),
		Type:    errType,
		Code:    code,
	}
}

// WrapWithType 包装错误并指定类型和代码
//
// 参数 / Parameters:
//   - err: 要包装的错误 / error to wrap
//   - errType: 错误类型 / error type
//   - code: 错误代码 / error code
//   - message: 错误消息 / error message
//
// 返回值 / Returns:
//   - error: 包装后的错误 / wrapped error
//
// 示例 / Example:
//
//	err := errors.New("original error")
//	wrapped := WrapWithType(err, ErrorTypeValidation, ErrorCodeInvalidInput, "validation failed")
//
// WrapWithType wraps an error with type and code
func WrapWithType(err error, errType ErrorType, code ErrorCode, message string) error {
	if err == nil {
		return nil
	}
	return &WrappedError{
		Err:     err,
		Message: message,
		Stack:   captureStack(2),
		Type:    errType,
		Code:    code,
	}
}

// StackTrace 获取错误的堆栈跟踪
//
// 参数 / Parameters:
//   - err: 错误 / error
//
// 返回值 / Returns:
//   - []string: 堆栈跟踪信息，如果没有则返回空切片 / stack trace, empty slice if not available
//
// 示例 / Example:
//
//	stack := StackTrace(err)
//
// StackTrace returns the stack trace of an error
func StackTrace(err error) []string {
	if err == nil {
		return nil
	}
	if wrapped, ok := err.(*WrappedError); ok {
		return wrapped.Stack
	}
	return nil
}

// IsType 检查错误是否为指定类型
//
// 参数 / Parameters:
//   - err: 错误 / error
//   - errType: 要检查的错误类型 / error type to check
//
// 返回值 / Returns:
//   - bool: 如果错误是指定类型则返回true / true if error is of the specified type
//
// 示例 / Example:
//
//	if IsType(err, ErrorTypeValidation) { ... }
//
// IsType checks if an error is of a specific type
func IsType(err error, errType ErrorType) bool {
	if err == nil {
		return false
	}
	var wrapped *WrappedError
	if errors.As(err, &wrapped) {
		return wrapped.Type == errType
	}
	return false
}

// IsCode 检查错误是否为指定代码
//
// 参数 / Parameters:
//   - err: 错误 / error
//   - code: 要检查的错误代码 / error code to check
//
// 返回值 / Returns:
//   - bool: 如果错误是指定代码则返回true / true if error is of the specified code
//
// 示例 / Example:
//
//	if IsCode(err, ErrorCodeInvalidInput) { ... }
//
// IsCode checks if an error is of a specific code
func IsCode(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}
	var wrapped *WrappedError
	if errors.As(err, &wrapped) {
		return wrapped.Code == code
	}
	return false
}

// GetType 获取错误的类型
//
// 参数 / Parameters:
//   - err: 错误 / error
//
// 返回值 / Returns:
//   - ErrorType: 错误类型，如果不是WrappedError则返回ErrorTypeUnknown / error type, ErrorTypeUnknown if not WrappedError
//
// 示例 / Example:
//
//	errType := GetType(err)
//
// GetType returns the type of an error
func GetType(err error) ErrorType {
	if err == nil {
		return ErrorTypeUnknown
	}
	var wrapped *WrappedError
	if errors.As(err, &wrapped) {
		return wrapped.Type
	}
	return ErrorTypeUnknown
}

// GetCode 获取错误的代码
//
// 参数 / Parameters:
//   - err: 错误 / error
//
// 返回值 / Returns:
//   - ErrorCode: 错误代码，如果不是WrappedError则返回ErrorCodeUnknown / error code, ErrorCodeUnknown if not WrappedError
//
// 示例 / Example:
//
//	code := GetCode(err)
//
// GetCode returns the code of an error
func GetCode(err error) ErrorCode {
	if err == nil {
		return ErrorCodeUnknown
	}
	var wrapped *WrappedError
	if errors.As(err, &wrapped) {
		return wrapped.Code
	}
	return ErrorCodeUnknown
}

// FormatError 格式化错误消息，包含堆栈跟踪
//
// 参数 / Parameters:
//   - err: 错误 / error
//   - includeStack: 是否包含堆栈跟踪 / whether to include stack trace
//
// 返回值 / Returns:
//   - string: 格式化后的错误消息 / formatted error message
//
// 示例 / Example:
//
//	formatted := FormatError(err, true)
//
// FormatError formats an error message with optional stack trace
func FormatError(err error, includeStack bool) string {
	if err == nil {
		return ""
	}

	var sb strings.Builder
	var wrapped *WrappedError

	if errors.As(err, &wrapped) {
		// 构建错误消息
		// Build error message
		if wrapped.Message != "" {
			sb.WriteString(wrapped.Message)
			if wrapped.Err != nil {
				sb.WriteString(": ")
			}
		}
		if wrapped.Err != nil {
			sb.WriteString(wrapped.Err.Error())
		}

		// 添加类型和代码信息
		// Add type and code information
		if wrapped.Type != ErrorTypeUnknown || wrapped.Code != ErrorCodeUnknown {
			sb.WriteString(" [")
			if wrapped.Type != ErrorTypeUnknown {
				sb.WriteString("type: ")
				sb.WriteString(string(wrapped.Type))
			}
			if wrapped.Type != ErrorTypeUnknown && wrapped.Code != ErrorCodeUnknown {
				sb.WriteString(", ")
			}
			if wrapped.Code != ErrorCodeUnknown {
				sb.WriteString("code: ")
				sb.WriteString(fmt.Sprintf("%d", wrapped.Code))
			}
			sb.WriteString("]")
		}

		// 添加堆栈跟踪
		// Add stack trace
		if includeStack && len(wrapped.Stack) > 0 {
			sb.WriteString("\nStack trace:\n")
			for i, frame := range wrapped.Stack {
				sb.WriteString(fmt.Sprintf("  %d. %s\n", i+1, frame))
			}
		}
	} else {
		// 普通错误
		// Regular error
		sb.WriteString(err.Error())
	}

	return sb.String()
}

// captureStack 捕获堆栈跟踪
// captureStack captures the stack trace
func captureStack(skip int) []string {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	var stack []string
	for {
		frame, more := frames.Next()
		stack = append(stack, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	return stack
}
