# go-commons

<sub><sup>[English README](README.md) | 中文</sup></sub>

[![Go Reference](https://pkg.go.dev/badge/github.com/Rodert/go-commons.svg)](https://pkg.go.dev/github.com/Rodert/go-commons)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](LICENSE)

一组精简的 Go 实用工具包，包含字符串工具与基础系统工具，尽量不依赖第三方库。

## 特性

- **尽量不使用第三方依赖**：优先使用 Go 标准库
- **字符串工具（`stringutils`）**：
  - 空与空白：`IsEmpty`、`IsNotEmpty`、`IsBlank`、`IsNotBlank`、`Trim`、`TrimToEmpty`
  - 子串与判断：`ContainsAny`、`ContainsAll`、`SubstringBefore`、`SubstringAfter`、`StartsWith`、`EndsWith`
  - 转换：`Capitalize`、`Uncapitalize`、`ReverseString`、`ToUpperCase`、`ToLowerCase`
  - 替换与连接：`Join`、`Split`、`Replace`、`ReplaceAll`、`Repeat`
  - 填充与居中：`PadLeft`、`PadRight`、`Center`
  - 其他：`Truncate`、`TruncateWithSuffix`、`CountMatches`、`DefaultIfEmpty`、`DefaultIfBlank`
- **系统工具（`systemutils`）**：已建立 `cpuutils`、`memutils`、`diskutils` 目录（功能待补充）

## 模块

- 模块路径：`github.com/Rodert/go-commons`
- Go 版本：`1.24.3`

## 安装

```bash
go get github.com/Rodert/go-commons
```

## 使用示例

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

## 示例

- 参考 `stringutils/stringutils_test.go` 获取更多覆盖的行为示例。
- `examples/` 目录预留为可运行示例（目前为空）。

## 原则

1. 优先使用标准库，尽量避免第三方依赖
2. 保持 API 简洁、清晰并配套测试

## 规划

- 充实 `systemutils/{cpuutils,memutils,diskutils}` 包的基础指标能力
- 在 `examples/` 中补充可运行示例

## 贡献

欢迎提交 Issue 与 PR。请保持代码可读性，并在新增函数时补充测试。 