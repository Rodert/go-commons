package cpuutils

import (
	"fmt"
	"testing"
)

func TestGetCPUInfo(t *testing.T) {
	cpuInfo, err := GetCPUInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("CPU 核心数: %d\n", cpuInfo.LogicalCores)
	fmt.Printf("CPU 使用率: %.2f%%\n", cpuInfo.UsagePercent)
	fmt.Printf("负载平均: %.2f, %.2f, %.2f\n", cpuInfo.LoadAvg[0], cpuInfo.LoadAvg[1], cpuInfo.LoadAvg[2])
	fmt.Printf("cpuInfo: %+v", cpuInfo)
}
