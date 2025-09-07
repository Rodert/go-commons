package memutils

// MemInfo 包含系统内存信息
type MemInfo struct {
	Total     uint64 // 总内存（字节）
	Available uint64 // 可用内存（字节）
	Used      uint64 // 已用内存（字节）
}

// GetMemInfo 获取系统内存信息
func GetMemInfo() (*MemInfo, error) {
	return getMemInfo()
}
