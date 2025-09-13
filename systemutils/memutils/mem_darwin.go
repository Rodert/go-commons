//go:build darwin
// +build darwin

package memutils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// getMemInfo macOS 实现
func getMemInfo() (*MemInfo, error) {
	// 使用vm_stat命令获取内存信息
	cmd := exec.Command("vm_stat")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// 解析输出
	lines := strings.Split(string(output), "\n")
	pageSize := uint64(4096) // 默认页大小为4KB
	var free, inactive uint64

	for _, line := range lines {
		if strings.Contains(line, "page size of") {
			parts := strings.Split(line, "page size of")
			if len(parts) >= 2 {
				sizeStr := strings.TrimSpace(parts[1])
				sizeStr = strings.TrimSuffix(sizeStr, " bytes")
				size, err := strconv.ParseUint(sizeStr, 10, 64)
				if err == nil {
					pageSize = size
				}
			}
		} else if strings.Contains(line, "Pages free") {
			free = parseVMStatLine(line)
		} else if strings.Contains(line, "Pages inactive") {
			inactive = parseVMStatLine(line)
		}
	}

	// 获取总内存
	cmd = exec.Command("sysctl", "-n", "hw.memsize")
	output, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	totalStr := strings.TrimSpace(string(output))
	total, err := strconv.ParseUint(totalStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse total memory: %v", err)
	}

	// 计算可用内存和已用内存
	available := (free + inactive) * pageSize
	used := total - available

	return &MemInfo{
		Total:     total,
		Available: available,
		Used:      used,
	}, nil
}

// 解析vm_stat输出行，提取数值
func parseVMStatLine(line string) uint64 {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return 0
	}
	valStr := strings.TrimSpace(parts[1])
	valStr = strings.TrimSuffix(valStr, ".")
	val, _ := strconv.ParseUint(valStr, 10, 64)
	return val
}
