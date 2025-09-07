//go:build linux
// +build linux

package cpuutils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// getCPUInfo Linux 实现
func getCPUInfo() (*CPUInfo, error) {
	info := &CPUInfo{
		LogicalCores: runtime.NumCPU(),
	}

	usage, err := getCPUUsageLinux()
	if err != nil {
		return nil, err
	}
	info.UsagePercent = usage

	loadavg, err := getLoadAvgLinux()
	if err == nil {
		info.LoadAvg = loadavg
	}

	return info, nil
}

// 获取 CPU 使用率（读取 /proc/stat 两次计算差值）
func getCPUUsageLinux() (float64, error) {
	total1, idle1, err := readCPUStat()
	if err != nil {
		return 0, err
	}
	time.Sleep(500 * time.Millisecond)
	total2, idle2, err := readCPUStat()
	if err != nil {
		return 0, err
	}

	totalDelta := total2 - total1
	idleDelta := idle2 - idle1
	usage := (float64(totalDelta-idleDelta) / float64(totalDelta)) * 100
	return usage, nil
}

// 读取 /proc/stat 第一行 CPU 数据
func readCPUStat() (total, idle uint64, err error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] != "cpu" {
			err = fmt.Errorf("unexpected format")
			return
		}
		var vals []uint64
		for _, f := range fields[1:] {
			v, e := strconv.ParseUint(f, 10, 64)
			if e != nil {
				err = e
				return
			}
			vals = append(vals, v)
		}
		for _, v := range vals {
			total += v
		}
		idle = vals[3] // idle time
	}
	return
}

// 读取负载平均值
func getLoadAvgLinux() ([3]float64, error) {
	var load [3]float64
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return load, err
	}
	parts := strings.Fields(string(data))
	for i := 0; i < 3; i++ {
		load[i], _ = strconv.ParseFloat(parts[i], 64)
	}
	return load, nil
}
