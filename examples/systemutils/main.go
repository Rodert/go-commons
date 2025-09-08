package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Rodert/go-commons/systemutils/cpuutils"
	"github.com/Rodert/go-commons/systemutils/diskutils"
	"github.com/Rodert/go-commons/systemutils/memutils"
)

func main() {
	fmt.Println("=== Go Commons System Utils Demo ===")
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
	fmt.Println()

	// 系统监控示例
	fmt.Println("📊 System Monitoring (5 seconds):")
	monitorSystem()
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
