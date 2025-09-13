package diskutils

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

// TestGetDiskInfo 测试获取磁盘信息
func TestGetDiskInfo(t *testing.T) {
	// 获取当前操作系统的根目录
	rootPath := getRootPath()
	
	// 获取磁盘信息
	diskInfo, err := GetDiskInfo(rootPath)
	if err != nil {
		t.Fatalf("获取磁盘信息失败: %v", err)
	}
	
	// 验证磁盘信息的有效性
	if diskInfo == nil {
		t.Fatal("磁盘信息不应为nil")
	}
	
	// 验证路径
	if diskInfo.Path == "" {
		t.Error("磁盘路径不应为空")
	}
	
	// 验证总空间
	if diskInfo.Total == 0 {
		t.Error("磁盘总空间不应为0")
	}
	
	// 验证已用空间和可用空间
	if diskInfo.Used > diskInfo.Total {
		t.Errorf("已用空间(%d)不应大于总空间(%d)", diskInfo.Used, diskInfo.Total)
	}
	
	if diskInfo.Free > diskInfo.Total {
		t.Errorf("可用空间(%d)不应大于总空间(%d)", diskInfo.Free, diskInfo.Total)
	}
	
	// 验证使用率
	if diskInfo.UsedRatio < 0 || diskInfo.UsedRatio > 100 {
		t.Errorf("使用率应在0-100之间，实际为: %.2f%%", diskInfo.UsedRatio)
	}
	
	// 打印磁盘信息（仅供参考，不是测试的一部分）
	t.Logf("磁盘总量: %d GB, 可用: %d GB, 已用: %d GB, 使用率: %.2f%%",
		diskInfo.Total/1024/1024/1024, diskInfo.Free/1024/1024/1024, 
		diskInfo.Used/1024/1024/1024, diskInfo.UsedRatio)
	t.Logf("磁盘信息: %+v", diskInfo)
}

// TestMultiplePaths 测试多个路径的磁盘信息
func TestMultiplePaths(t *testing.T) {
	// 根据不同操作系统选择测试路径
	var paths []string
	switch runtime.GOOS {
	case "windows":
		paths = []string{"C:\\", "D:\\"}
	case "darwin":
		paths = []string{"/", "/Users", "/tmp"}
	case "linux":
		paths = []string{"/", "/home", "/tmp"}
	default:
		paths = []string{"/"}
	}
	
	// 测试每个路径
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			info, err := GetDiskInfo(path)
			if err != nil {
				t.Logf("跳过 %s: %v", path, err)
				return
			}
			
			// 验证基本信息
			if info.Path == "" {
				t.Errorf("路径 %s: 返回的路径不应为空", path)
			}
			
			if info.Total == 0 {
				t.Errorf("路径 %s: 总空间不应为0", path)
			}
			
			t.Logf("路径=%s, 总空间=%.2fGB, 使用率=%.2f%%",
				info.Path,
				float64(info.Total)/1e9,
				info.UsedRatio,
			)
		})
	}
}

// TestInvalidPath 测试无效路径
func TestInvalidPath(t *testing.T) {
	// 使用一个不太可能存在的路径
	invalidPath := "/path/that/should/not/exist/12345"
	
	// 在Windows上使用不同的无效路径
	if runtime.GOOS == "windows" {
		invalidPath = "X:\\invalid\\path"
	}
	
	// 获取磁盘信息，应该返回错误
	_, err := GetDiskInfo(invalidPath)
	if err == nil {
		t.Errorf("对于无效路径 %s 应返回错误，但没有", invalidPath)
	} else {
		t.Logf("正确地对无效路径返回错误: %v", err)
	}
}

// ExampleGetDiskInfo 展示如何使用GetDiskInfo函数
func ExampleGetDiskInfo() {
	// 获取根目录的磁盘信息
	diskInfo, err := GetDiskInfo("/")
	if err != nil {
		fmt.Printf("获取磁盘信息失败: %v\n", err)
		return
	}
	
	// 打印磁盘信息
	fmt.Printf("路径: %s\n", diskInfo.Path)
	fmt.Printf("总空间: %.2f GB\n", float64(diskInfo.Total)/1e9)
	fmt.Printf("已用空间: %.2f GB\n", float64(diskInfo.Used)/1e9)
	fmt.Printf("可用空间: %.2f GB\n", float64(diskInfo.Free)/1e9)
	fmt.Printf("使用率: %.2f%%\n", diskInfo.UsedRatio)
	
	// 输出不确定，所以不做验证
	// Output:
}

// getRootPath 根据操作系统返回根路径
func getRootPath() string {
	if runtime.GOOS == "windows" {
		// 获取当前工作目录所在的盘符
		wd, err := os.Getwd()
		if err == nil && len(wd) >= 2 && wd[1] == ':' {
			return wd[:2] + "\\"
		}
		return "C:\\"
	}
	return "/"
}
