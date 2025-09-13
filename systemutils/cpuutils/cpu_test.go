package cpuutils

import (
	"fmt"
	"testing"
	"time"
)

// TestGetCPUInfo 测试获取CPU信息
func TestGetCPUInfo(t *testing.T) {
	cpuInfo, err := GetCPUInfo()
	if err != nil {
		t.Fatalf("获取CPU信息失败: %v", err)
	}
	
	// 验证CPU信息的有效性
	if cpuInfo == nil {
		t.Fatal("CPU信息不应为nil")
	}
	
	// 验证逻辑核心数
	if cpuInfo.LogicalCores <= 0 {
		t.Errorf("逻辑核心数应大于0，实际为: %d", cpuInfo.LogicalCores)
	}
	
	// 验证CPU使用率范围
	if cpuInfo.UsagePercent < 0 || cpuInfo.UsagePercent > 100 {
		t.Errorf("CPU使用率应在0-100之间，实际为: %.2f%%", cpuInfo.UsagePercent)
	}
	
	// 打印CPU信息（仅供参考，不是测试的一部分）
	t.Logf("CPU 核心数: %d", cpuInfo.LogicalCores)
	t.Logf("CPU 使用率: %.2f%%", cpuInfo.UsagePercent)
	t.Logf("负载平均: %.2f, %.2f, %.2f", cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
	t.Logf("CPU信息: %+v", cpuInfo)
}

// TestCPUInfoStability 测试多次获取CPU信息的稳定性
func TestCPUInfoStability(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过稳定性测试")
	}
	
	// 多次获取CPU信息，验证稳定性
	for i := 0; i < 3; i++ {
		cpuInfo, err := GetCPUInfo()
		if err != nil {
			t.Fatalf("第%d次获取CPU信息失败: %v", i+1, err)
		}
		
		// 验证基本信息
		if cpuInfo.LogicalCores <= 0 {
			t.Errorf("第%d次：逻辑核心数应大于0，实际为: %d", i+1, cpuInfo.LogicalCores)
		}
		
		t.Logf("第%d次：CPU使用率: %.2f%%", i+1, cpuInfo.UsagePercent)
		
		// 短暂等待，让CPU使用率有变化
		time.Sleep(500 * time.Millisecond)
	}
}

// ExampleGetCPUInfo 展示如何使用GetCPUInfo函数
func ExampleGetCPUInfo() {
	cpuInfo, err := GetCPUInfo()
	if err != nil {
		fmt.Printf("获取CPU信息失败: %v\n", err)
		return
	}
	
	fmt.Printf("CPU 核心数: %d\n", cpuInfo.LogicalCores)
	fmt.Printf("CPU 使用率: %.2f%%\n", cpuInfo.UsagePercent)
	fmt.Printf("负载平均: %.2f, %.2f, %.2f\n", cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
	
	// 输出不确定，所以不做验证
	// Output:
}
