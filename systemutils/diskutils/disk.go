package diskutils

// DiskInfo 磁盘信息
type DiskInfo struct {
	Path      string  // 挂载点路径
	Total     uint64  // 总大小（字节）
	Free      uint64  // 可用空间（字节）
	Used      uint64  // 已用空间（字节）
	UsedRatio float64 // 使用率百分比
}

// GetDiskInfo 获取指定路径磁盘信息
func GetDiskInfo(path string) (*DiskInfo, error) {
	return getDiskInfo(path)
}
