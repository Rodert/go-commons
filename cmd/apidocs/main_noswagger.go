//go:build noswagger
// +build noswagger

package main

// 当使用 -tags=noswagger 构建时，这个空的main函数将被使用
// 这样就不会生成apidocs可执行文件
func main() {
	// 空实现，不做任何事情
}
