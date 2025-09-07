package diskutils

import (
	"fmt"
	"testing"
)

func TestGetDiskInfo(t *testing.T) {
	// 磁盘
	disk, _ := GetDiskInfo("/")
	fmt.Printf("磁盘总量: %d GB, 可用: %d GB, 已用: %d GB, 使用率: %.2f%%\n",
		disk.Total/1024/1024/1024, disk.Free/1024/1024/1024, disk.Used/1024/1024/1024, disk.UsedRatio)
	fmt.Printf("disk: %+v", disk)
}

func TestDiskInfo(t *testing.T) {
	paths := []string{"/", "C:\\"}
	for _, path := range paths {
		info, err := GetDiskInfo(path)
		if err != nil {
			t.Logf("skip %s: %v", path, err)
			continue
		}
		fmt.Printf("Path=%s, Total=%.2fGB, UsedRatio=%.2f%%\n",
			info.Path,
			float64(info.Total)/1e9,
			info.UsedRatio,
		)
	}
}
