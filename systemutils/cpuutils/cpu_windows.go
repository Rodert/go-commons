//go:build windows
// +build windows

package cpuutils

import (
	"math"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

// Windows 实现 CPU 信息
func getCPUInfo() (*CPUInfo, error) {
	usage, err := getCPUUsageWindows()
	if err != nil {
		usage = math.NaN()
	}

	info := &CPUInfo{
		LogicalCores: runtime.NumCPU(),
		UsagePercent: usage,
		LoadAvg:      [3]float64{math.NaN(), math.NaN(), math.NaN()}, // Windows 无负载平均
	}
	return info, nil
}

// 调用 Windows API 获取 CPU 使用率
func getCPUUsageWindows() (float64, error) {
	idle1, kernel1, user1, err := getSystemTimes()
	if err != nil {
		return 0, err
	}

	time.Sleep(500 * time.Millisecond) // 采样间隔

	idle2, kernel2, user2, err := getSystemTimes()
	if err != nil {
		return 0, err
	}

	idle := float64(idle2 - idle1)
	kernel := float64(kernel2 - kernel1)
	user := float64(user2 - user1)

	usage := 100 * (kernel + user - idle) / (kernel + user)
	if usage < 0 {
		usage = 0
	} else if usage > 100 {
		usage = 100
	}

	return usage, nil
}

// 系统调用
var (
	modkernel32        = syscall.NewLazyDLL("kernel32.dll")
	procGetSystemTimes = modkernel32.NewProc("GetSystemTimes")
)

type filetime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

func filetimeToUint64(ft filetime) uint64 {
	return (uint64(ft.HighDateTime) << 32) + uint64(ft.LowDateTime)
}

func getSystemTimes() (idle, kernel, user uint64, err error) {
	var idleTime, kernelTime, userTime filetime
	r1, _, e1 := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&idleTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)
	if r1 == 0 {
		err = e1
		return
	}

	idle = filetimeToUint64(idleTime)
	kernel = filetimeToUint64(kernelTime)
	user = filetimeToUint64(userTime)
	return
}
