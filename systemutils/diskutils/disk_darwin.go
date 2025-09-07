//go:build darwin

package diskutils

import (
	"syscall"
)

func getDiskInfo(path string) (*DiskInfo, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return nil, err
	}
	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free
	usedRatio := float64(used) / float64(total) * 100
	return &DiskInfo{
		Path:      path,
		Total:     total,
		Free:      free,
		Used:      used,
		UsedRatio: usedRatio,
	}, nil
}
