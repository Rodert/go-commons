//go:build linux
// +build linux

package memutils

import (
	"os"
	"strconv"
	"strings"
)

// getMemInfo Linux 实现
func getMemInfo() (*MemInfo, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	info := &MemInfo{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSuffix(parts[0], ":")
		val, _ := strconv.ParseUint(parts[1], 10, 64)
		val = val * 1024 // kB -> B
		switch key {
		case "MemTotal":
			info.Total = val
		case "MemAvailable":
			info.Available = val
		}
	}
	info.Used = info.Total - info.Available
	return info, nil
}
