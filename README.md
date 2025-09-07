# go-commons

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
- **System utilities (`systemutils`)**: structure prepared for `cpuutils`, `memutils`, `diskutils` (APIs to be added)

## Module

- Module path: `github.com/Rodert/go-commons`
- Go version: `1.24.3`

## Install

```bash
go get github.com/Rodert/go-commons
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
)

func main() {
	fmt.Println(stringutils.IsBlank("  \t\n"))         // true
	fmt.Println(stringutils.Trim("  hello  "))        // "hello"
	fmt.Println(stringutils.TruncateWithSuffix("abcdef", 4, "..")) // "ab.."
	fmt.Println(stringutils.PadLeft("42", 5, '0'))     // "00042"
	fmt.Println(stringutils.ContainsAny("gopher", "go", "java")) // true
}
```

## Examples

- See `stringutils/stringutils_test.go` for a wide range of covered behaviors.
- The `examples/` directory is reserved for runnable samples (currently empty).

## Principles

1. Prefer the standard library over third‑party dependencies
2. Keep APIs small, clear, and well‑tested

## Roadmap

- Flesh out `systemutils/{cpuutils,memutils,diskutils}` packages with basic metrics helpers
- Add more examples under `examples/`

## Contributing

Issues and pull requests are welcome. Please keep code readable and add tests when introducing new functions. 