package memutils

import (
	"fmt"
	"testing"
)

func TestGetMemInfo(t *testing.T) {
	mem, err := GetMemInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("总内存: %d MB, 可用: %d MB, 已用: %d MB\n",
		mem.Total/1024/1024, mem.Available/1024/1024, mem.Used/1024/1024)
	fmt.Printf("mem: %+v", mem)
}
