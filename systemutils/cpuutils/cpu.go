package cpuutils

// CPUInfo 包含 CPU 核心数和负载信息
type CPUInfo struct {
	LogicalCores int
	UsagePercent float64    // 总体使用率，0~100
	LoadAvg      [3]float64 // 1/5/15 分钟平均负载（Linux）
}

// GetCPUInfo 获取当前 CPU 信息
func GetCPUInfo() (*CPUInfo, error) {
	return getCPUInfo()
}
