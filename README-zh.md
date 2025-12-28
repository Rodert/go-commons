# go-commons

<sub><sup>[English README](README.md) | 中文</sup></sub>

[![Go Reference](https://pkg.go.dev/badge/github.com/Rodert/go-commons.svg)](https://pkg.go.dev/github.com/Rodert/go-commons)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](LICENSE)
[![Go Tests](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-test.yml)
[![Go Lint](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml/badge.svg)](https://github.com/Rodert/go-commons/actions/workflows/go-lint.yml)
[![codecov](https://codecov.io/gh/Rodert/go-commons/branch/main/graph/badge.svg)](https://codecov.io/gh/Rodert/go-commons)

一组全面的 Go 实用工具包，尽量不依赖第三方库，为常见开发任务提供必要的工具。

## 特性

- **尽量不使用第三方依赖**：优先使用 Go 标准库
- **字符串工具（`stringutils`）**：
  - 空与空白：`IsEmpty`、`IsNotEmpty`、`IsBlank`、`IsNotBlank`、`Trim`、`TrimToEmpty`
  - 子串与判断：`ContainsAny`、`ContainsAll`、`SubstringBefore`、`SubstringAfter`、`StartsWith`、`EndsWith`
  - 转换：`Capitalize`、`Uncapitalize`、`ReverseString`、`ToUpperCase`、`ToLowerCase`
  - 替换与连接：`Join`、`Split`、`Replace`、`ReplaceAll`、`Repeat`
  - 填充与居中：`PadLeft`、`PadRight`、`Center`
  - 其他：`Truncate`、`TruncateWithSuffix`、`CountMatches`、`DefaultIfEmpty`、`DefaultIfBlank`
- **时间工具（`timeutils`）**：
  - 时间格式化与解析：`FormatTime`、`ParseTime`
  - 时间计算：`AddDays`、`AddMonths`、`AddYears`、`DaysBetween`、`HoursBetween`、`MinutesBetween`
  - 相对时间：`TimeAgo`、`TimeAgoEn`
  - 时间范围：`Today`、`ThisWeek`、`ThisMonth`、`ThisYear`
  - 时区转换：`ToTimezone`、`ToUTC`
  - 时间判断：`IsToday`、`IsWeekend`、`IsWeekday`
- **文件工具（`fileutils`）**：
  - 文件读写：`ReadFile`、`WriteFile`、`ReadFileLines`
  - 目录操作：`WalkDir`、`FindFiles`
  - 文件操作：`Copy`、`Move`、`Delete`、`Exists`
  - 路径工具：`JoinPath`、`CleanPath`、`BaseName`、`DirName`
  - 文件类型检测：`GetFileType`、`IsDir`、`IsFile`
- **切片工具（`sliceutils`）**：
  - 去重：`Unique`、`UniqueInt`、`UniqueString`
  - 函数式操作：`Filter`、`Map`、`Reduce`
  - 分页：`Paginate`、`PaginateInt`
  - 集合操作：`Intersection`、`Union`、`Difference`
  - 排序：`Sort`、`SortInt`、`SortString`、`SortIntDesc`、`SortStringDesc`
- **JSON/转换工具（`jsonutils`、`convertutils`）**：
  - JSON格式化：`PrettyJSON`、`CompactJSON`
  - 类型转换：`MapToStruct`、`StructToMap`、`StringToInt`、`IntToString`、`FloatToString`
  - 深拷贝：`DeepCopy`
  - JSON验证与合并：`IsValidJSON`、`MergeJSON`
- **错误处理工具（`errorutils`）**：
  - 错误包装：`Wrap`、`Wrapf`、`WithStack`
  - 堆栈跟踪：`StackTrace`
  - 错误分类：`IsType`、`IsCode`、`GetType`、`GetCode`
  - 错误格式化：`FormatError`
- **配置工具（`configutils`）**：
  - 配置加载：`LoadFromJSON`、`LoadFromJSONString`、`LoadFromEnv`
  - 类型安全访问：`GetString`、`GetInt`、`GetFloat`、`GetBool`、`GetStringSlice`
  - 配置管理：`Set`、`Get`、`Has`、`Merge`、`SetDefaults`
  - 配置验证：`Validate`
  - 结构体解析：`Unmarshal`
- **并发工具（`concurrentutils`）**：
  - 工作池：`WorkerPool` - 管理并发任务执行
  - 限流器：`RateLimiter` - 使用令牌桶算法控制请求速率
  - 安全计数器：`SafeCounter` - 使用原子操作的线程安全计数器
  - 安全缓存：`SafeCache` - 支持懒加载的线程安全内存缓存
- **系统工具（`systemutils`）**：
  - CPU工具（`cpuutils`）：`GetCPUInfo` - 获取CPU核心数、使用率百分比和负载平均值
  - 内存工具（`memutils`）：`GetMemInfo` - 获取总内存、可用内存和已用内存
  - 磁盘工具（`diskutils`）：`GetDiskInfo` - 获取磁盘空间信息，包括总空间、可用空间、已用空间和使用率

## 模块

- 模块路径：`github.com/Rodert/go-commons`
- Go 版本：`1.24.7`

## 安装

```bash
go get github.com/Rodert/go-commons
```

## 快速开始

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/stringutils"
	"github.com/Rodert/go-commons/timeutils"
	"github.com/Rodert/go-commons/configutils"
)

func main() {
	// 字符串工具
	fmt.Println(stringutils.IsBlank("  "))  // true
	fmt.Println(stringutils.Trim("  hello  "))  // "hello"
	
	// 时间工具
	now := timeutils.Now()
	fmt.Println(timeutils.FormatTime(now, timeutils.DefaultDateTimeFormat))
	
	// 配置工具
	config := configutils.NewConfig()
	config.Set("app.name", "MyApp")
	fmt.Println(config.GetString("app.name", ""))  // "MyApp"
}
```

## 包概览

本库包含以下包：

- **`stringutils`** - 字符串操作和验证工具
- **`timeutils`** - 时间和日期操作、格式化和计算
- **`fileutils`** - 文件和目录操作、路径工具
- **`sliceutils`** - 切片操作：去重、过滤、分页、排序
- **`jsonutils`** - JSON格式化和验证
- **`convertutils`** - 类型转换和深拷贝
- **`errorutils`** - 错误包装、堆栈跟踪和错误分类
- **`configutils`** - 配置管理，支持JSON和环境变量
- **`concurrentutils`** - 并发工具：工作池、限流器、安全计数器和缓存
- **`systemutils`** - 系统监控：CPU、内存和磁盘工具
  - `cpuutils` - CPU信息和使用率
  - `memutils` - 内存信息
  - `diskutils` - 磁盘空间信息

## 开发

### 自动格式化

本项目使用Git钩子在每次提交前自动格式化Go代码。

安装pre-commit钩子：

```bash
make hooks
```

### API文档

本项目包含一个基于Swagger UI的交互式API文档界面。这使您可以通过Web界面探索和测试库中的函数。

#### 📌 在线API文档

**访问我们的在线API文档：[https://rodert.github.io/go-commons](https://rodert.github.io/go-commons)**

在线文档从main分支自动部署，提供最新的API参考。

![API文档界面](images/api-img.png)

#### 本地开发

在本地启动API文档服务器：

```bash
./run_apidocs.sh
```

然后在浏览器中访问 [http://localhost:8080](http://localhost:8080) 查看交互式API文档。

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

### 错误处理工具

```go
package main

import (
	"errors"
	"fmt"
	"github.com/Rodert/go-commons/errorutils"
)

func main() {
	// 包装错误并添加上下文
	err := errors.New("file not found")
	wrapped := errorutils.Wrap(err, "failed to read config")
	
	// 检查错误类型
	if errorutils.IsType(wrapped, errorutils.ErrorTypeInternal) {
		fmt.Println("内部错误")
	}
	
	// 格式化错误（包含堆栈跟踪）
	fmt.Println(errorutils.FormatError(wrapped, true))
}
```

### 配置工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/configutils"
)

func main() {
	// 从JSON加载配置
	config, _ := configutils.LoadConfigFromJSON("config.json")
	
	// 获取值（带默认值）
	host := config.GetString("database.host", "localhost")
	port := config.GetInt("database.port", 3306)
	debug := config.GetBool("app.debug", false)
	
	// 从环境变量加载配置
	envConfig := configutils.LoadConfigFromEnv("APP_")
	fmt.Println(envConfig.GetString("name", "default"))
}
```

### 并发工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/concurrentutils"
)

func main() {
	// 工作池
	pool := concurrentutils.NewWorkerPool(10)
	pool.Start()
	defer pool.Stop()
	
	pool.Submit(func() {
		fmt.Println("任务执行")
	})
	
	// 限流器
	limiter := concurrentutils.NewRateLimiter(100) // 每秒100个请求
	if limiter.Allow() {
		// 处理请求
	}
	
	// 安全计数器
	counter := concurrentutils.NewSafeCounter(0)
	counter.Increment(1)
	fmt.Println(counter.Get())
	
	// 安全缓存
	cache := concurrentutils.NewSafeCache()
	cache.Set("key", "value")
	val, _ := cache.Get("key")
	fmt.Println(val)
}
```

### 时间工具

```go
package main

import (
	"fmt"
	"time"
	"github.com/Rodert/go-commons/timeutils"
)

func main() {
	now := time.Now()
	
	// 格式化
	fmt.Println(timeutils.FormatTime(now, timeutils.DefaultDateTimeFormat))
	
	// 计算
	tomorrow := timeutils.AddDays(now, 1)
	nextMonth := timeutils.AddMonths(now, 1)
	
	// 相对时间
	fmt.Println(timeutils.TimeAgo(now.Add(-2 * time.Hour)))  // "2小时前"
	
	// 时间判断
	fmt.Println(timeutils.IsToday(now))  // true
	fmt.Println(timeutils.IsWeekend(now))  // 取决于日期
}
```

### 文件工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/fileutils"
)

func main() {
	// 读取文件
	content, _ := fileutils.ReadFile("config.json")
	
	// 写入文件
	fileutils.WriteFile("output.txt", []byte("Hello World"))
	
	// 文件操作
	if fileutils.Exists("file.txt") {
		fileutils.Copy("file.txt", "file_copy.txt")
	}
	
	// 路径工具
	base := fileutils.BaseName("/path/to/file.txt")  // "file.txt"
	dir := fileutils.DirName("/path/to/file.txt")    // "/path/to"
}
```

### 切片工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/sliceutils"
)

func main() {
	// 去重
	nums := []int{1, 2, 2, 3, 3, 3}
	unique := sliceutils.UniqueInt(nums)  // [1, 2, 3]
	
	// 过滤
	even := sliceutils.Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	
	// 分页
	page := sliceutils.PaginateInt(nums, 1, 2)  // 第1页，每页2条
	
	// 集合操作
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	intersection := sliceutils.Intersection(a, b)  // [2, 3]
}
```

### JSON/转换工具

```go
package main

import (
	"fmt"
	"github.com/Rodert/go-commons/jsonutils"
	"github.com/Rodert/go-commons/convertutils"
)

func main() {
	// JSON格式化
	jsonStr := `{"name":"John","age":30}`
	pretty, _ := jsonutils.PrettyJSON(jsonStr)
	fmt.Println(pretty)
	
	// 类型转换
	num := convertutils.StringToInt("123", 0)  // 123
	str := convertutils.IntToString(456)       // "456"
	
	// 深拷贝
	original := map[string]interface{}{"key": "value"}
	copied := convertutils.DeepCopy(original)
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

`examples/` 目录提供了全面的示例：

- `examples/stringutils/` - 字符串操作示例
- `examples/timeutils/` - 时间和日期操作
- `examples/fileutils/` - 文件和目录操作
- `examples/sliceutils/` - 切片操作和函数式编程
- `examples/jsonutils/` - JSON处理示例
- `examples/configutils/` - 配置管理
- `examples/errorutils/` - 错误处理模式
- `examples/concurrentutils/` - 并发工具
- `examples/systemutils/` - 系统监控

您也可以查看测试文件（如 `*_test.go`）获取更多使用示例。

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

1. **最小依赖**：优先使用标准库，尽量避免第三方依赖
2. **简洁API**：保持 API 简洁、清晰并配套测试
3. **跨平台**：支持 Linux、macOS 和 Windows
4. **完善文档**：提供全面的文档和示例
5. **生产就绪**：经过充分测试，代码覆盖率高

## 性能

所有工具都针对性能进行了优化：
- 尽可能减少或避免内存分配
- 高效的算法（大多数操作为 O(n)）
- 使用原子操作和同步原语的线程安全实现
- 在热点路径中避免反射开销

## 许可证

本项目采用 [Unlicense](LICENSE) 许可证 - 详见 LICENSE 文件。

## 规划

- [ ] HTTP工具增强（URL构建器、查询参数解析、重试机制）
- [ ] 编码/解码工具（URL、HTML、Hex）
- [ ] 数学工具（精确浮点计算、随机数、百分比）
- [ ] 反射工具（结构体字段操作、标签解析）
- [ ] 日志工具（结构化日志、日志轮转、彩色输出）
- [ ] 增强 `systemutils` 包的详细指标
- [ ] 添加更多示例和使用场景
- [ ] 改进跨平台兼容性和测试

## 开发时间线

- **2025-09-07**: 项目初始化，创建基础README和LICENSE
- **2025-09-08**: 
  - 添加`stringutils`包中的核心字符串工具函数
  - 实现CPU、内存和磁盘监控的系统工具
  - 添加跨平台支持（Linux、macOS、Windows）
  - 创建示例和完善文档
  - 添加字符串转换函数（`Reverse`、`SwapCase`、`PadCenter`）
- **2025-01-XX**: 
  - 添加时间工具（`timeutils`）- 时间格式化、计算、时区转换
  - 添加文件工具（`fileutils`）- 文件I/O、目录操作、路径工具
  - 添加切片工具（`sliceutils`）- 去重、函数式操作、分页、排序
  - 添加JSON/转换工具（`jsonutils`、`convertutils`）- JSON格式化、类型转换、深拷贝
  - 添加错误处理工具（`errorutils`）- 错误包装、堆栈跟踪、错误分类
  - 添加配置工具（`configutils`）- 配置加载、验证、类型安全访问
  - 添加并发工具（`concurrentutils`）- 工作池、限流器、安全计数器、安全缓存

## 贡献

欢迎提交 Issue 与 PR。请保持代码可读性，并在新增函数时补充测试。