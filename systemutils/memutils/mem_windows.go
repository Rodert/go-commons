//go:build windows

package memutils

import (
	"syscall"
	"unsafe"
)

// MEMORYSTATUSEX 定义
type memoryStatusEx struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

// getMemInfo Windows 实现
func getMemInfo() (*MemInfo, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	globalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

	var mem memoryStatusEx
	mem.Length = uint32(unsafe.Sizeof(mem))

	r1, _, err := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&mem)))
	if r1 == 0 {
		return nil, err
	}

	info := &MemInfo{
		Total:     mem.TotalPhys,
		Available: mem.AvailPhys,
		Used:      mem.TotalPhys - mem.AvailPhys,
	}
	return info, nil
}
