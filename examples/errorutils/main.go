package main

import (
	"errors"
	"fmt"
	"github.com/Rodert/go-commons/errorutils"
)

func main() {
	fmt.Println("=== Error Utils Examples ===\n")

	// 示例1: 基本错误包装
	// Example 1: Basic error wrapping
	fmt.Println("1. Basic Error Wrapping:")
	originalErr := errors.New("file not found")
	wrappedErr := errorutils.Wrap(originalErr, "failed to read config")
	fmt.Printf("   Original: %v\n", originalErr)
	fmt.Printf("   Wrapped:  %v\n\n", wrappedErr)

	// 示例2: 格式化错误包装
	// Example 2: Formatted error wrapping
	fmt.Println("2. Formatted Error Wrapping:")
	fileErr := errors.New("permission denied")
	formattedErr := errorutils.Wrapf(fileErr, "failed to open file: %s", "config.json")
	fmt.Printf("   Formatted: %v\n\n", formattedErr)

	// 示例3: 添加堆栈跟踪
	// Example 3: Adding stack trace
	fmt.Println("3. Error with Stack Trace:")
	stackErr := errorutils.WithStack(errors.New("something went wrong"))
	stack := errorutils.StackTrace(stackErr)
	if len(stack) > 0 {
		fmt.Printf("   Error: %v\n", stackErr)
		fmt.Printf("   Stack (first 3 frames):\n")
		for i, frame := range stack {
			if i >= 3 {
				break
			}
			fmt.Printf("     %d. %s\n", i+1, frame)
		}
	}
	fmt.Println()

	// 示例4: 带类型的错误
	// Example 4: Error with type
	fmt.Println("4. Typed Errors:")
	validationErr := errorutils.NewWithType(
		errorutils.ErrorTypeValidation,
		errorutils.ErrorCodeInvalidInput,
		"email format is invalid",
	)
	fmt.Printf("   Error: %v\n", validationErr)
	fmt.Printf("   Type: %v\n", errorutils.GetType(validationErr))
	fmt.Printf("   Code: %v\n\n", errorutils.GetCode(validationErr))

	// 示例5: 包装错误并指定类型
	// Example 5: Wrap error with type
	fmt.Println("5. Wrap Error with Type:")
	notFoundErr := errors.New("resource not found")
	typedErr := errorutils.WrapWithType(
		notFoundErr,
		errorutils.ErrorTypeNotFound,
		errorutils.ErrorCodeNotFound,
		"user not found",
	)
	fmt.Printf("   Error: %v\n", typedErr)
	fmt.Printf("   Type: %v\n", errorutils.GetType(typedErr))
	fmt.Printf("   Code: %v\n\n", errorutils.GetCode(typedErr))

	// 示例6: 错误类型检查
	// Example 6: Error type checking
	fmt.Println("6. Error Type Checking:")
	testErr := errorutils.NewWithType(
		errorutils.ErrorTypeValidation,
		errorutils.ErrorCodeInvalidInput,
		"validation failed",
	)
	if errorutils.IsType(testErr, errorutils.ErrorTypeValidation) {
		fmt.Printf("   Error is of type Validation\n")
	}
	if errorutils.IsCode(testErr, errorutils.ErrorCodeInvalidInput) {
		fmt.Printf("   Error has code InvalidInput\n")
	}
	fmt.Println()

	// 示例7: 错误格式化
	// Example 7: Error formatting
	fmt.Println("7. Error Formatting:")
	complexErr := errorutils.WrapWithType(
		errors.New("database connection failed"),
		errorutils.ErrorTypeNetwork,
		errorutils.ErrorCodeTimeout,
		"failed to connect to database",
	)
	fmt.Println("   Without stack:")
	fmt.Printf("   %s\n\n", errorutils.FormatError(complexErr, false))
	fmt.Println("   With stack (first 100 chars):")
	formatted := errorutils.FormatError(complexErr, true)
	if len(formatted) > 100 {
		fmt.Printf("   %s...\n\n", formatted[:100])
	} else {
		fmt.Printf("   %s\n\n", formatted)
	}

	// 示例8: 错误链
	// Example 8: Error chain
	fmt.Println("8. Error Chain:")
	baseErr := errors.New("base error")
	chain1 := errorutils.Wrap(baseErr, "level 1")
	chain2 := errorutils.Wrapf(chain1, "level 2: %s", "context")
	fmt.Printf("   Chain: %v\n", chain2)
	
	// 使用标准库的 errors.Is 检查
	// Use standard library's errors.Is to check
	if errors.Is(chain2, baseErr) {
		fmt.Printf("   ✓ Can find base error in chain\n")
	}
	if errors.Is(chain2, chain1) {
		fmt.Printf("   ✓ Can find intermediate error in chain\n")
	}
}

