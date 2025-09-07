//go:build windows

package diskutils

import (
	"syscall"
	"unsafe"
)

// getDiskInfo Windows 实现
func getDiskInfo(path string) (*DiskInfo, error) {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}

	// 加载 kernel32.dll
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getDiskFreeSpaceEx := kernel32.NewProc("GetDiskFreeSpaceExW")

	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64

	r1, _, err := getDiskFreeSpaceEx.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
	)
	if r1 == 0 {
		return nil, err
	}

	used := totalNumberOfBytes - totalNumberOfFreeBytes
	usedRatio := float64(used) / float64(totalNumberOfBytes) * 100

	return &DiskInfo{
		Path:      path,
		Total:     totalNumberOfBytes,
		Free:      totalNumberOfFreeBytes,
		Used:      used,
		UsedRatio: usedRatio,
	}, nil
}
