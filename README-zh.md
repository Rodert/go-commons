# go-commons

<sub><sup>[English README](README.md) | 中文</sup></sub>

[![Go Reference](https://pkg.go.dev/badge/github.com/Rodert/go-commons.svg)](https://pkg.go.dev/github.com/Rodert/go-commons)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](LICENSE)
[![Go Tests](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml)
[![Go Lint](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml)
[![codecov](https://codecov.io/gh/Rodert/go-commons/branch/main/graph/badge.svg)](https://codecov.io/gh/Rodert/go-commons)

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
- **系统工具（`systemutils`）**：
  - CPU工具（`cpuutils`）：`GetCPUInfo` - 获取CPU核心数、使用率百分比和负载平均值
  - 内存工具（`memutils`）：`GetMemInfo` - 获取总内存、可用内存和已用内存
  - 磁盘工具（`diskutils`）：`GetDiskInfo` - 获取磁盘空间信息，包括总空间、可用空间、已用空间和使用率

## 模块

- 模块路径：`github.com/Rodert/go-commons`
- Go 版本：`1.24.3`

## 安装

```bash
go get github.com/Rodert/go-commons
```

## 开发

### 自动格式化

本项目使用Git钩子在每次提交前自动格式化Go代码。

安装pre-commit钩子：

```bash
make hooks
```

手动格式化所有Go文件：

```bash
make fmt
```

## 使用示例

### 字符串工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
)

func main() {
	// 基本字符串操作
	fmt.Println(stringutils.IsBlank("  \t\n"))         // true
	fmt.Println(stringutils.Trim("  hello  "))        // "hello"
	fmt.Println(stringutils.TruncateWithSuffix("abcdef", 4, "..")) // "ab.."
	fmt.Println(stringutils.PadLeft("42", 5, '0'))     // "00042"
	fmt.Println(stringutils.ContainsAny("gopher", "go", "java")) // true
	
	// 字符串转换
	fmt.Println(stringutils.Reverse("hello"))         // "olleh"
	fmt.Println(stringutils.SwapCase("Hello World"))  // "hELLO wORLD"
	fmt.Println(stringutils.PadCenter("hello", 9, '*')) // "**hello**"
}
```

### 系统工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/systemutils/cpuutils"
	"github.com/Rodert/go-commons/systemutils/memutils"
	"github.com/Rodert/go-commons/systemutils/diskutils"
)

func main() {
	// 获取CPU信息
	cpuInfo, err := cpuutils.GetCPUInfo()
	if err == nil {
		fmt.Printf("CPU核心数: %d\n", cpuInfo.LogicalCores)
		fmt.Printf("CPU使用率: %.2f%%\n", cpuInfo.UsagePercent)
		fmt.Printf("负载平均值: %.2f, %.2f, %.2f\n", 
			cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
	}
	
	// 获取内存信息
	memInfo, err := memutils.GetMemInfo()
	if err == nil {
		fmt.Printf("总内存: %d 字节\n", memInfo.Total)
		fmt.Printf("可用内存: %d 字节\n", memInfo.Available)
		fmt.Printf("已用内存: %d 字节\n", memInfo.Used)
	}
	
	// 获取磁盘信息
	diskInfo, err := diskutils.GetDiskInfo("/")
	if err == nil {
		fmt.Printf("磁盘路径: %s\n", diskInfo.Path)
		fmt.Printf("总空间: %d 字节\n", diskInfo.Total)
		fmt.Printf("可用空间: %d 字节\n", diskInfo.Free)
		fmt.Printf("已用空间: %d 字节\n", diskInfo.Used)
		fmt.Printf("使用率: %.2f%%\n", diskInfo.UsedRatio)
	}
}
```

## 示例

- 参考 `stringutils/stringutils_test.go` 获取更多覆盖的行为示例。
- 查看 `examples/` 目录获取可运行示例。

## 测试

本项目包含一个Makefile，用于简化测试和其他开发任务：

```bash
# 运行所有测试
make test

# 运行特定包的测试
make test-pkg PKG=./stringutils

# 运行测试并生成覆盖率报告
make cover

# 运行基准测试
make bench

# 格式化代码并运行测试
make

# 显示所有可用命令
make help
```

## 原则

1. 优先使用标准库，尽量避免第三方依赖
2. 保持 API 简洁、清晰并配套测试

## 规划

- 增强 `systemutils` 包的详细指标和监控能力
- 在 `examples/` 中补充可运行示例
- 改进跨平台兼容性和测试
- 添加更多字符串操作工具

## 开发时间线

- **2025-09-07**: 项目初始化，创建基础README和LICENSE
- **2025-09-08**: 
  - 添加`stringutils`包中的核心字符串工具函数
  - 实现CPU、内存和磁盘监控的系统工具
  - 添加跨平台支持（Linux、macOS、Windows）
  - 创建示例和完善文档
  - 添加字符串转换函数（`Reverse`、`SwapCase`、`PadCenter`）

## 贡献

欢迎提交 Issue 与 PR。请保持代码可读性，并在新增函数时补充测试。