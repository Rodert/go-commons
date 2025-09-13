# go-commons

<sub><sup>English | [中文 README](README-zh.md)</sup></sub>

[![Go Reference](https://pkg.go.dev/badge/github.com/Rodert/go-commons.svg)](https://pkg.go.dev/github.com/Rodert/go-commons)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](LICENSE)
[![Go Tests](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml)
[![Go Lint](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml)
[![codecov](https://codecov.io/gh/Rodert/go-commons/branch/main/graph/badge.svg)](https://codecov.io/gh/Rodert/go-commons)

A small collection of Go utility packages focused on string helpers and basic system utilities, with minimal third‑party dependencies.

## Features

- **No third‑party deps**: Prefer using the Go standard library where possible
- **String utilities (`stringutils`)**:
  - Emptiness and whitespace: `IsEmpty`, `IsNotEmpty`, `IsBlank`, `IsNotBlank`, `Trim`, `TrimToEmpty`
  - Substrings and checks: `ContainsAny`, `ContainsAll`, `SubstringBefore`, `SubstringAfter`, `StartsWith`, `EndsWith`
  - Transformations: `Capitalize`, `Uncapitalize`, `ReverseString`, `ToUpperCase`, `ToLowerCase`
  - Replace and join: `Join`, `Split`, `Replace`, `ReplaceAll`, `Repeat`
  - Padding and centering: `PadLeft`, `PadRight`, `Center`
  - Misc: `Truncate`, `TruncateWithSuffix`, `CountMatches`, `DefaultIfEmpty`, `DefaultIfBlank`
- **System utilities (`systemutils`)**:
  - CPU utilities (`cpuutils`): `GetCPUInfo` - retrieve CPU cores, usage percentage, and load averages
  - Memory utilities (`memutils`): `GetMemInfo` - get total, available, and used memory
  - Disk utilities (`diskutils`): `GetDiskInfo` - get disk space information including total, free, used space and usage ratio

## Module

- Module path: `github.com/Rodert/go-commons`
- Go version: `1.24.3`

## Install

```bash
go get github.com/Rodert/go-commons
```

## Development

### Auto-formatting

This project uses Git hooks to automatically format Go code before each commit.

To install the pre-commit hook:

```bash
make hooks
```

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

- See `stringutils/stringutils_test.go` for a wide range of covered behaviors.
- Check the `examples/` directory for runnable samples.

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

1. Prefer the standard library over third‑party dependencies
2. Keep APIs small, clear, and well‑tested

## Roadmap

- Enhance `systemutils` packages with more detailed metrics and monitoring capabilities
- Add more examples under `examples/`
- Improve cross-platform compatibility and testing
- Add more string manipulation utilities

## Development Timeline

- **2025-09-07**: Initial project setup, basic README and LICENSE
- **2025-09-08**: 
  - Added core string utilities in `stringutils` package
  - Implemented system utilities for CPU, memory, and disk monitoring
  - Added cross-platform support (Linux, macOS, Windows)
  - Created examples and comprehensive documentation
  - Added string transformation functions (`Reverse`, `SwapCase`, `PadCenter`)

## Contributing

Issues and pull requests are welcome. Please keep code readable and add tests when introducing new functions.