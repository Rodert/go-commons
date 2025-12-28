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
			// 解析格式: "X% user, Y% sys, Z% idle"
			var user, sys, idle float64
			n, err := fmt.Sscanf(usageInfo, "%f%% user, %f%% sys, %f%% idle", &user, &sys, &idle)

			// 如果成功解析了所有三个值，计算总CPU使用率
			if err == nil && n == 3 {
				// 优先使用 100 - idle 来计算总使用率（更准确）
				if idle >= 0 && idle <= 100 {
					return 100.0 - idle, nil
				}
				// 如果idle值异常，使用 user + sys
				if user >= 0 && sys >= 0 {
					return user + sys, nil
				}
			}

			// 如果上面都失败，尝试旧的解析方式（向后兼容）
			usageParts := strings.Split(usageInfo, "%")
			if len(usageParts) >= 1 {
				usageStr := strings.TrimSpace(usageParts[0])
				usage, err := strconv.ParseFloat(usageStr, 64)
				if err == nil {
					return usage, nil
				}
			}
			return 0, fmt.Errorf("unexpected format: %s", usageInfo)
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
