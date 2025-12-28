# go-commons

<sub><sup>English | [中文 README](README-zh.md)</sup></sub>

[![Go Reference](https://pkg.go.dev/badge/github.com/Rodert/go-commons.svg)](https://pkg.go.dev/github.com/Rodert/go-commons)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](LICENSE)
[![Go Tests](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml)
[![Go Lint](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml)
[![codecov](https://codecov.io/gh/Rodert/go-commons/branch/main/graph/badge.svg)](https://codecov.io/gh/Rodert/go-commons)

A comprehensive collection of Go utility packages with minimal third‑party dependencies, providing essential tools for common development tasks.

## Features

- **No third‑party deps**: Prefer using the Go standard library where possible
- **String utilities (`stringutils`)**:
  - Emptiness and whitespace: `IsEmpty`, `IsNotEmpty`, `IsBlank`, `IsNotBlank`, `Trim`, `TrimToEmpty`
  - Substrings and checks: `ContainsAny`, `ContainsAll`, `SubstringBefore`, `SubstringAfter`, `StartsWith`, `EndsWith`
  - Transformations: `Capitalize`, `Uncapitalize`, `ReverseString`, `ToUpperCase`, `ToLowerCase`
  - Replace and join: `Join`, `Split`, `Replace`, `ReplaceAll`, `Repeat`
  - Padding and centering: `PadLeft`, `PadRight`, `Center`
  - Misc: `Truncate`, `TruncateWithSuffix`, `CountMatches`, `DefaultIfEmpty`, `DefaultIfBlank`
- **Time utilities (`timeutils`)**:
  - Time formatting and parsing: `FormatTime`, `ParseTime`
  - Time calculations: `AddDays`, `AddMonths`, `AddYears`, `DaysBetween`, `HoursBetween`, `MinutesBetween`
  - Relative time: `TimeAgo`, `TimeAgoEn`
  - Time ranges: `Today`, `ThisWeek`, `ThisMonth`, `ThisYear`
  - Timezone conversion: `ToTimezone`, `ToUTC`
  - Time checks: `IsToday`, `IsWeekend`, `IsWeekday`
- **File utilities (`fileutils`)**:
  - File I/O: `ReadFile`, `WriteFile`, `ReadFileLines`
  - Directory operations: `WalkDir`, `FindFiles`
  - File operations: `Copy`, `Move`, `Delete`, `Exists`
  - Path utilities: `JoinPath`, `CleanPath`, `BaseName`, `DirName`
  - File type detection: `GetFileType`, `IsDir`, `IsFile`
- **Slice utilities (`sliceutils`)**:
  - Deduplication: `Unique`, `UniqueInt`, `UniqueString`
  - Functional operations: `Filter`, `Map`, `Reduce`
  - Pagination: `Paginate`, `PaginateInt`
  - Set operations: `Intersection`, `Union`, `Difference`
  - Sorting: `Sort`, `SortInt`, `SortString`, `SortIntDesc`, `SortStringDesc`
- **JSON/Convert utilities (`jsonutils`, `convertutils`)**:
  - JSON formatting: `PrettyJSON`, `CompactJSON`
  - Type conversion: `MapToStruct`, `StructToMap`, `StringToInt`, `IntToString`, `FloatToString`
  - Deep copy: `DeepCopy`
  - JSON validation and merging: `IsValidJSON`, `MergeJSON`
- **Error utilities (`errorutils`)**:
  - Error wrapping: `Wrap`, `Wrapf`, `WithStack`
  - Stack trace: `StackTrace`
  - Error classification: `IsType`, `IsCode`, `GetType`, `GetCode`
  - Error formatting: `FormatError`
- **Config utilities (`configutils`)**:
  - Configuration loading: `LoadFromJSON`, `LoadFromJSONString`, `LoadFromEnv`
  - Type-safe getters: `GetString`, `GetInt`, `GetFloat`, `GetBool`, `GetStringSlice`
  - Configuration management: `Set`, `Get`, `Has`, `Merge`, `SetDefaults`
  - Validation: `Validate`
  - Struct unmarshaling: `Unmarshal`
- **Concurrent utilities (`concurrentutils`)**:
  - Worker pool: `WorkerPool` - manage concurrent task execution
  - Rate limiter: `RateLimiter` - control request rate with token bucket algorithm
  - Safe counter: `SafeCounter` - thread-safe counter with atomic operations
  - Safe cache: `SafeCache` - thread-safe in-memory cache with lazy loading
- **System utilities (`systemutils`)**:
  - CPU utilities (`cpuutils`): `GetCPUInfo` - retrieve CPU cores, usage percentage, and load averages
  - Memory utilities (`memutils`): `GetMemInfo` - get total, available, and used memory
  - Disk utilities (`diskutils`): `GetDiskInfo` - get disk space information including total, free, used space and usage ratio

## Module

- Module path: `github.com/Rodert/go-commons`
- Go version: `1.24.7`

## Install

```bash
go get github.com/Rodert/go-commons
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
	"github.com/Rodert/go-commons/timeutils"
	"github.com/Rodert/go-commons/configutils"
)

func main() {
	// String utilities
	fmt.Println(stringutils.IsBlank("  "))  // true
	fmt.Println(stringutils.Trim("  hello  "))  // "hello"
	
	// Time utilities
	now := timeutils.Now()
	fmt.Println(timeutils.FormatTime(now, timeutils.DefaultDateTimeFormat))
	
	// Config utilities
	config := configutils.NewConfig()
	config.Set("app.name", "MyApp")
	fmt.Println(config.GetString("app.name", ""))  // "MyApp"
}
```

## Package Overview

This library is organized into the following packages:

- **`stringutils`** - String manipulation and validation utilities
- **`timeutils`** - Time and date operations, formatting, and calculations
- **`fileutils`** - File and directory operations, path utilities
- **`sliceutils`** - Slice operations: deduplication, filtering, pagination, sorting
- **`jsonutils`** - JSON formatting and validation
- **`convertutils`** - Type conversion and deep copying
- **`errorutils`** - Error wrapping, stack traces, and error classification
- **`configutils`** - Configuration management with JSON and environment variable support
- **`concurrentutils`** - Concurrency utilities: worker pools, rate limiting, safe counters and caches
- **`systemutils`** - System monitoring: CPU, memory, and disk utilities
  - `cpuutils` - CPU information and usage
  - `memutils` - Memory information
  - `diskutils` - Disk space information

## Development

### Auto-formatting

This project uses Git hooks to automatically format Go code before each commit.

To install the pre-commit hook:

```bash
make hooks
```

### API Documentation

This project includes an interactive API documentation interface using Swagger UI. This allows you to explore and test the library's functions through a web interface.

#### 📌 Online API Documentation

**Visit our API documentation online at: [https://rodert.github.io/go-commons](https://rodert.github.io/go-commons)**

The online documentation is automatically deployed from the main branch and provides the most up-to-date API reference.

![API Documentation Interface](images/api-img.png)

#### Local Development

To start the API documentation server locally:

```bash
./run_apidocs.sh
```

Then open your browser and navigate to [http://localhost:8080](http://localhost:8080) to view the interactive API documentation.

To manually format all Go files:

```bash
make fmt
```

## Usage

### String Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
)

func main() {
	// Basic string operations
	fmt.Println(stringutils.IsBlank("  \t\n"))         // true
	fmt.Println(stringutils.Trim("  hello  "))        // "hello"
	fmt.Println(stringutils.TruncateWithSuffix("abcdef", 4, "..")) // "ab.."
	fmt.Println(stringutils.PadLeft("42", 5, '0'))     // "00042"
	fmt.Println(stringutils.ContainsAny("gopher", "go", "java")) // true
	
	// String transformations
	fmt.Println(stringutils.Reverse("hello"))         // "olleh"
	fmt.Println(stringutils.SwapCase("Hello World"))  // "hELLO wORLD"
	fmt.Println(stringutils.PadCenter("hello", 9, '*')) // "**hello**"
}
```

### Error Utilities

```go
package main

import (
	"errors"
	"fmt"
	"github.com/Rodert/go-commons/errorutils"
)

func main() {
	// Wrap errors with context
	err := errors.New("file not found")
	wrapped := errorutils.Wrap(err, "failed to read config")
	
	// Check error type
	if errorutils.IsType(wrapped, errorutils.ErrorTypeInternal) {
		fmt.Println("Internal error")
	}
	
	// Format error with stack trace
	fmt.Println(errorutils.FormatError(wrapped, true))
}
```

### Config Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/configutils"
)

func main() {
	// Load from JSON
	config, _ := configutils.LoadConfigFromJSON("config.json")
	
	// Get values with defaults
	host := config.GetString("database.host", "localhost")
	port := config.GetInt("database.port", 3306)
	debug := config.GetBool("app.debug", false)
	
	// Load from environment variables
	envConfig := configutils.LoadConfigFromEnv("APP_")
	fmt.Println(envConfig.GetString("name", "default"))
}
```

### Concurrent Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/concurrentutils"
)

func main() {
	// Worker pool
	pool := concurrentutils.NewWorkerPool(10)
	pool.Start()
	defer pool.Stop()
	
	pool.Submit(func() {
		fmt.Println("Task executed")
	})
	
	// Rate limiter
	limiter := concurrentutils.NewRateLimiter(100) // 100 req/s
	if limiter.Allow() {
		// Process request
	}
	
	// Safe counter
	counter := concurrentutils.NewSafeCounter(0)
	counter.Increment(1)
	fmt.Println(counter.Get())
	
	// Safe cache
	cache := concurrentutils.NewSafeCache()
	cache.Set("key", "value")
	val, _ := cache.Get("key")
	fmt.Println(val)
}
```

### Time Utilities

```go
package main

import (
	"fmt"
	"time"
	"github.com/Rodert/go-commons/timeutils"
)

func main() {
	now := time.Now()
	
	// Formatting
	fmt.Println(timeutils.FormatTime(now, timeutils.DefaultDateTimeFormat))
	
	// Calculations
	tomorrow := timeutils.AddDays(now, 1)
	nextMonth := timeutils.AddMonths(now, 1)
	
	// Relative time
	fmt.Println(timeutils.TimeAgo(now.Add(-2 * time.Hour)))  // "2小时前"
	
	// Time checks
	fmt.Println(timeutils.IsToday(now))  // true
	fmt.Println(timeutils.IsWeekend(now))  // depends on day
}
```

### File Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/fileutils"
)

func main() {
	// Read file
	content, _ := fileutils.ReadFile("config.json")
	
	// Write file
	fileutils.WriteFile("output.txt", []byte("Hello World"))
	
	// File operations
	if fileutils.Exists("file.txt") {
		fileutils.Copy("file.txt", "file_copy.txt")
	}
	
	// Path utilities
	base := fileutils.BaseName("/path/to/file.txt")  // "file.txt"
	dir := fileutils.DirName("/path/to/file.txt")    // "/path/to"
}
```

### Slice Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/sliceutils"
)

func main() {
	// Deduplication
	nums := []int{1, 2, 2, 3, 3, 3}
	unique := sliceutils.UniqueInt(nums)  // [1, 2, 3]
	
	// Filter
	even := sliceutils.Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	
	// Pagination
	page := sliceutils.PaginateInt(nums, 1, 2)  // page 1, size 2
	
	// Set operations
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	intersection := sliceutils.Intersection(a, b)  // [2, 3]
}
```

### JSON/Convert Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/jsonutils"
	"github.com/Rodert/go-commons/convertutils"
)

func main() {
	// JSON formatting
	jsonStr := `{"name":"John","age":30}`
	pretty, _ := jsonutils.PrettyJSON(jsonStr)
	fmt.Println(pretty)
	
	// Type conversion
	num := convertutils.StringToInt("123", 0)  // 123
	str := convertutils.IntToString(456)       // "456"
	
	// Deep copy
	original := map[string]interface{}{"key": "value"}
	copied := convertutils.DeepCopy(original)
}
```

### System Utilities

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/systemutils/cpuutils"
	"github.com/Rodert/go-commons/systemutils/memutils"
	"github.com/Rodert/go-commons/systemutils/diskutils"
)

func main() {
	// Get CPU information
	cpuInfo, err := cpuutils.GetCPUInfo()
	if err == nil {
		fmt.Printf("CPU Cores: %d\n", cpuInfo.LogicalCores)
		fmt.Printf("CPU Usage: %.2f%%\n", cpuInfo.UsagePercent)
		fmt.Printf("Load Average: %.2f, %.2f, %.2f\n", 
			cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
	}
	
	// Get memory information
	memInfo, err := memutils.GetMemInfo()
	if err == nil {
		fmt.Printf("Total Memory: %d bytes\n", memInfo.Total)
		fmt.Printf("Available Memory: %d bytes\n", memInfo.Available)
		fmt.Printf("Used Memory: %d bytes\n", memInfo.Used)
	}
	
	// Get disk information
	diskInfo, err := diskutils.GetDiskInfo("/")
	if err == nil {
		fmt.Printf("Disk Path: %s\n", diskInfo.Path)
		fmt.Printf("Total Space: %d bytes\n", diskInfo.Total)
		fmt.Printf("Free Space: %d bytes\n", diskInfo.Free)
		fmt.Printf("Used Space: %d bytes\n", diskInfo.Used)
		fmt.Printf("Usage Ratio: %.2f%%\n", diskInfo.UsedRatio)
	}
}
```

## Examples

Comprehensive examples are available in the `examples/` directory:

- `examples/stringutils/` - String manipulation examples
- `examples/timeutils/` - Time and date operations
- `examples/fileutils/` - File and directory operations
- `examples/sliceutils/` - Slice operations and functional programming
- `examples/jsonutils/` - JSON processing examples
- `examples/configutils/` - Configuration management
- `examples/errorutils/` - Error handling patterns
- `examples/concurrentutils/` - Concurrency utilities
- `examples/systemutils/` - System monitoring

You can also check the test files (e.g., `*_test.go`) for more usage examples.

## Testing

This project includes a Makefile to simplify running tests and other development tasks:

```bash
# Run all tests
make test

# Run tests for a specific package
make test-pkg PKG=./stringutils

# Run tests with coverage report
make cover

# Run benchmarks
make bench

# Format code and run tests
make

# Show all available commands
make help
```

## Principles

1. **Minimal dependencies**: Prefer the standard library over third‑party dependencies
2. **Simple APIs**: Keep APIs small, clear, and well‑tested
3. **Cross-platform**: Support Linux, macOS, and Windows
4. **Well-documented**: Comprehensive documentation with examples
5. **Production-ready**: Thoroughly tested with high code coverage

## Performance

All utilities are designed for performance:
- Zero or minimal allocations where possible
- Efficient algorithms (e.g., O(n) for most operations)
- Thread-safe implementations using atomic operations and sync primitives
- No reflection overhead in hot paths

## License

This project is licensed under the [Unlicense](LICENSE) - see the LICENSE file for details.

## Roadmap

- [ ] HTTP utilities enhancement (URL builder, query parsing, retry mechanism)
- [ ] Encoding/decoding utilities (URL, HTML, Hex)
- [ ] Math utilities (precise float calculations, random numbers, percentage)
- [ ] Reflection utilities (struct field manipulation, tag parsing)
- [ ] Logging utilities (structured logging, log rotation, colored output)
- [ ] Enhance `systemutils` packages with more detailed metrics
- [ ] Add more examples and use cases
- [ ] Improve cross-platform compatibility and testing

## Development Timeline

- **2025-09-07**: Initial project setup, basic README and LICENSE
- **2025-09-08**: 
  - Added core string utilities in `stringutils` package
  - Implemented system utilities for CPU, memory, and disk monitoring
  - Added cross-platform support (Linux, macOS, Windows)
  - Created examples and comprehensive documentation
  - Added string transformation functions (`Reverse`, `SwapCase`, `PadCenter`)
- **2025-01-XX**: 
  - Added time utilities (`timeutils`) - time formatting, calculations, timezone conversion
  - Added file utilities (`fileutils`) - file I/O, directory operations, path utilities
  - Added slice utilities (`sliceutils`) - deduplication, functional operations, pagination, sorting
  - Added JSON/Convert utilities (`jsonutils`, `convertutils`) - JSON formatting, type conversion, deep copy
  - Added error utilities (`errorutils`) - error wrapping, stack trace, error classification
  - Added config utilities (`configutils`) - configuration loading, validation, type-safe access
  - Added concurrent utilities (`concurrentutils`) - worker pool, rate limiter, safe counter, safe cache

## Contributing

Issues and pull requests are welcome. Please keep code readable and add tests when introducing new functions.