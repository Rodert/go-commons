package memutils

import (
	"fmt"
	"testing"
	"time"
)

// TestGetMemInfo 测试获取内存信息
func TestGetMemInfo(t *testing.T) {
	memInfo, err := GetMemInfo()
	if err != nil {
		t.Fatalf("获取内存信息失败: %v", err)
	}
	
	// 验证内存信息的有效性
	if memInfo == nil {
		t.Fatal("内存信息不应为nil")
	}
	
	// 验证总内存大小
	if memInfo.Total == 0 {
		t.Error("总内存不应为0")
	}
	
	// 验证已用内存和可用内存
	if memInfo.Used > memInfo.Total {
		t.Errorf("已用内存(%d)不应大于总内存(%d)", memInfo.Used, memInfo.Total)
	}
	
	if memInfo.Available > memInfo.Total {
		t.Errorf("可用内存(%d)不应大于总内存(%d)", memInfo.Available, memInfo.Total)
	}
	
	// 打印内存信息（仅供参考，不是测试的一部分）
	t.Logf("总内存: %d MB, 可用: %d MB, 已用: %d MB",
		memInfo.Total/1024/1024, memInfo.Available/1024/1024, memInfo.Used/1024/1024)
	t.Logf("内存信息: %+v", memInfo)
}

// TestMemInfoConsistency 测试内存信息的一致性
func TestMemInfoConsistency(t *testing.T) {
	// 获取第一次内存信息
	memInfo1, err := GetMemInfo()
	if err != nil {
		t.Fatalf("第一次获取内存信息失败: %v", err)
	}
	
	// 短暂等待
	time.Sleep(100 * time.Millisecond)
	
	// 获取第二次内存信息
	memInfo2, err := GetMemInfo()
	if err != nil {
		t.Fatalf("第二次获取内存信息失败: %v", err)
	}
	
	// 验证总内存应该保持一致
	if memInfo1.Total != memInfo2.Total {
		t.Errorf("两次获取的总内存不一致: %d != %d", memInfo1.Total, memInfo2.Total)
	}
	
	// 打印两次内存使用情况的差异（仅供参考）
	usedDiff := int64(memInfo2.Used) - int64(memInfo1.Used)
	availableDiff := int64(memInfo2.Available) - int64(memInfo1.Available)
	
	t.Logf("内存使用变化: %+d KB, 可用内存变化: %+d KB", 
		usedDiff/1024, availableDiff/1024)
}

// ExampleGetMemInfo 展示如何使用GetMemInfo函数
func ExampleGetMemInfo() {
	memInfo, err := GetMemInfo()
	if err != nil {
		fmt.Printf("获取内存信息失败: %v\n", err)
		return
	}
	
	// 计算内存使用率
	usagePercent := float64(memInfo.Used) / float64(memInfo.Total) * 100
	
	fmt.Printf("总内存: %d MB\n", memInfo.Total/1024/1024)
	fmt.Printf("已用内存: %d MB\n", memInfo.Used/1024/1024)
	fmt.Printf("可用内存: %d MB\n", memInfo.Available/1024/1024)
	fmt.Printf("内存使用率: %.2f%%\n", usagePercent)
	
	// 输出不确定，所以不做验证
	// Output:
}
