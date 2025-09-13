//go:build darwin
// +build darwin

package cpuutils

import (
	"fmt"
	"math"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// getCPUInfo macOS 实现
func getCPUInfo() (*CPUInfo, error) {
	info := &CPUInfo{
		LogicalCores: runtime.NumCPU(),
	}

	// 获取CPU使用率
	usage, err := getCPUUsageDarwin()
	if err != nil {
		return nil, err
	}
	info.UsagePercent = usage

	// 获取负载平均值
	loadavg, err := getLoadAvgDarwin()
	if err == nil {
		info.LoadAvg = loadavg
	}

	return info, nil
}

// 使用top命令获取CPU使用率
func getCPUUsageDarwin() (float64, error) {
	cmd := exec.Command("top", "-l", "1", "-n", "0")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "CPU usage") {
			parts := strings.Split(line, ": ")
			if len(parts) < 2 {
				return 0, fmt.Errorf("unexpected format")
			}
			usageInfo := parts[1]
			usageParts := strings.Split(usageInfo, "%")
			if len(usageParts) < 1 {
				return 0, fmt.Errorf("unexpected format")
			}
			usageStr := strings.TrimSpace(usageParts[0])
			usage, err := strconv.ParseFloat(usageStr, 64)
			if err != nil {
				return 0, err
			}
			return usage, nil
		}
	}

	return 0, fmt.Errorf("could not find CPU usage information")
}

// 获取负载平均值
func getLoadAvgDarwin() ([3]float64, error) {
	cmd := exec.Command("sysctl", "-n", "vm.loadavg")
	output, err := cmd.Output()
	if err != nil {
		return [3]float64{math.NaN(), math.NaN(), math.NaN()}, err
	}

	// 输出格式类似于 "{ 1.23 0.45 0.67 }"
	outputStr := string(output)
	outputStr = strings.Trim(outputStr, "{ }\n")
	parts := strings.Fields(outputStr)

	var loadavg [3]float64
	for i := 0; i < 3 && i < len(parts); i++ {
		loadavg[i], _ = strconv.ParseFloat(parts[i], 64)
	}

	return loadavg, nil
}