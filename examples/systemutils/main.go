package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/Rodert/go-commons/systemutils/cpuutils"
	"github.com/Rodert/go-commons/systemutils/diskutils"
	"github.com/Rodert/go-commons/systemutils/memutils"
)

func main() {
	fmt.Println("=== Go Commons System Utils Demo ===")
	fmt.Println()

	// 基本系统信息
	fmt.Println("🖥️  Basic System Information:")
	fmt.Printf("  OS: %s\n", runtime.GOOS)
	fmt.Printf("  Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("  Go Version: %s\n", runtime.Version())
	fmt.Printf("  Hostname: %s\n", getHostname())
	fmt.Println()

	// CPU 信息示例
	fmt.Println("🖥️  CPU Information:")
	cpuInfo, err := cpuutils.GetCPUInfo()
	if err != nil {
		log.Printf("Failed to get CPU info: %v", err)
	} else {
		fmt.Printf("  Logical Cores: %d\n", cpuInfo.LogicalCores)
		fmt.Printf("  Usage: %.2f%%\n", cpuInfo.UsagePercent)
		fmt.Printf("  Load Average: %.2f %.2f %.2f\n", 
			cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
		fmt.Printf("  GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	}
	fmt.Println()

	// 内存信息示例
	fmt.Println("💾 Memory Information:")
	memInfo, err := memutils.GetMemInfo()
	if err != nil {
		log.Printf("Failed to get memory info: %v", err)
	} else {
		fmt.Printf("  Total: %s\n", formatBytes(memInfo.Total))
		fmt.Printf("  Used: %s\n", formatBytes(memInfo.Used))
		fmt.Printf("  Available: %s\n", formatBytes(memInfo.Available))
		fmt.Printf("  Usage: %.2f%%\n", float64(memInfo.Used)/float64(memInfo.Total)*100)
	}

	// Go 运行时内存信息
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Println("  Go Runtime Memory:")
	fmt.Printf("    Allocated: %s\n", formatBytes(rtm.Alloc))
	fmt.Printf("    Total Allocated: %s\n", formatBytes(rtm.TotalAlloc))
	fmt.Printf("    System: %s\n", formatBytes(rtm.Sys))
	fmt.Printf("    GC Cycles: %d\n", rtm.NumGC)
	fmt.Println()

	// 磁盘信息示例
	fmt.Println("💿 Disk Information:")
	diskInfo, err := diskutils.GetDiskInfo("/")
	if err != nil {
		log.Printf("Failed to get disk info: %v", err)
	} else {
		fmt.Printf("  Path: %s\n", diskInfo.Path)
		fmt.Printf("  Total: %s\n", formatBytes(diskInfo.Total))
		fmt.Printf("  Used: %s\n", formatBytes(diskInfo.Used))
		fmt.Printf("  Free: %s\n", formatBytes(diskInfo.Free))
		fmt.Printf("  Usage: %.2f%%\n", diskInfo.UsedRatio)
	}

	// 检查多个路径
	paths := []string{"/", "/home", "/tmp"}
	fmt.Println("  Multiple Disk Paths:")
	for _, path := range paths {
		info, err := diskutils.GetDiskInfo(path)
		if err != nil {
			continue
		}
		fmt.Printf("    %s: %.1f%% used (%.1f GB free)\n", 
			info.Path, info.UsedRatio, float64(info.Free)/1e9)
	}
	fmt.Println()

	// 系统监控示例
	fmt.Println("📊 System Monitoring (5 seconds):")
	monitorSystem()

	// 资源使用率图表示例
	fmt.Println("📈 Resource Usage Chart:")
	showResourceChart()
}

// getHostname 获取主机名
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

// formatBytes 格式化字节数为可读格式
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// monitorSystem 系统监控示例
func monitorSystem() {
	for i := 0; i < 5; i++ {
		fmt.Printf("  [%d/5] ", i+1)
		
		// CPU 使用率
		if cpuInfo, err := cpuutils.GetCPUInfo(); err == nil {
			fmt.Printf("CPU: %.1f%% ", cpuInfo.UsagePercent)
		}
		
		// 内存使用率
		if memInfo, err := memutils.GetMemInfo(); err == nil {
			usage := float64(memInfo.Used) / float64(memInfo.Total) * 100
			fmt.Printf("Memory: %.1f%% ", usage)
		}
		
		// 磁盘使用率
		if diskInfo, err := diskutils.GetDiskInfo("/"); err == nil {
			fmt.Printf("Disk: %.1f%%", diskInfo.UsedRatio)
		}
		
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}

// showResourceChart 显示资源使用率图表
func showResourceChart() {
	// 获取CPU和内存使用率
	cpuInfo, err1 := cpuutils.GetCPUInfo()
	memInfo, err2 := memutils.GetMemInfo()
	diskInfo, err3 := diskutils.GetDiskInfo("/")
	
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("  Error getting system information")
		return
	}
	
	cpuPercent := int(cpuInfo.UsagePercent / 10)
	memPercent := int(float64(memInfo.Used) / float64(memInfo.Total) * 10)
	diskPercent := int(diskInfo.UsedRatio / 10)
	
	// 显示图表
	fmt.Println("  Resource Usage (each █ = 10%):")
	
	fmt.Printf("  CPU  [%s%s] %.1f%%\n", 
		repeatChar('█', cpuPercent), 
		repeatChar('░', 10-cpuPercent),
		cpuInfo.UsagePercent)
	
	fmt.Printf("  MEM  [%s%s] %.1f%%\n", 
		repeatChar('█', memPercent), 
		repeatChar('░', 10-memPercent),
		float64(memInfo.Used)/float64(memInfo.Total)*100)
	
	fmt.Printf("  DISK [%s%s] %.1f%%\n", 
		repeatChar('█', diskPercent), 
		repeatChar('░', 10-diskPercent),
		diskInfo.UsedRatio)
}

// repeatChar 重复字符
func repeatChar(char rune, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += string(char)
	}
	return result
}
