package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Rodert/go-commons/fileutils"
)

func main() {
	fmt.Println("=== Go Commons File Utils Demo ===")
	fmt.Println()

	// 创建临时目录用于演示
	tmpDir, err := os.MkdirTemp("", "fileutils_demo_*")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	fmt.Println("临时目录:", tmpDir)
	fmt.Println()

	// 文件读写
	fmt.Println("文件读写 / File Read/Write:")
	testFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!\n这是第二行"
	err = fileutils.WriteFileString(testFile, content, 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Println("写入文件:", testFile)

	readContent, err := fileutils.ReadFileString(testFile)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}
	fmt.Println("读取内容:", readContent)

	lines, err := fileutils.ReadFileLines(testFile)
	if err != nil {
		fmt.Printf("按行读取失败: %v\n", err)
		return
	}
	fmt.Println("文件行数:", len(lines))
	for i, line := range lines {
		fmt.Printf("  第%d行: %s\n", i+1, line)
	}
	fmt.Println()

	// 文件存在性检查
	fmt.Println("文件存在性检查 / File Existence:")
	fmt.Println("文件存在:", fileutils.Exists(testFile))
	fmt.Println("目录存在:", fileutils.IsDir(tmpDir))
	fmt.Println("是文件:", fileutils.IsFile(testFile))
	fmt.Println("是目录:", fileutils.IsDir(testFile))
	fmt.Println()

	// 文件大小
	fmt.Println("文件大小 / File Size:")
	size, err := fileutils.GetFileSize(testFile)
	if err != nil {
		fmt.Printf("获取文件大小失败: %v\n", err)
	} else {
		fmt.Println("文件大小 (字节):", size)
		fmt.Println("格式化大小:", fileutils.FormatFileSize(size))
		fmt.Println("1KB:", fileutils.FormatFileSize(1024))
		fmt.Println("1MB:", fileutils.FormatFileSize(1024*1024))
		fmt.Println("1GB:", fileutils.FormatFileSize(1024*1024*1024))
	}
	fmt.Println()

	// 路径处理
	fmt.Println("路径处理 / Path Operations:")
	testPath := "/path/to/file.txt"
	fmt.Println("原始路径:", testPath)
	fmt.Println("目录名:", fileutils.DirName(testPath))
	fmt.Println("文件名:", fileutils.BaseName(testPath))
	fmt.Println("扩展名:", fileutils.ExtName(testPath))
	fmt.Println("无扩展名:", fileutils.FileNameWithoutExt(testPath))
	fmt.Println("连接路径:", fileutils.JoinPath("dir", "subdir", "file.txt"))
	fmt.Println("清理路径:", fileutils.CleanPath("./dir/../file.txt"))
	absPath, err := fileutils.AbsPath(".")
	if err == nil {
		fmt.Println("绝对路径:", absPath)
	}
	fmt.Println()

	// 文件操作
	fmt.Println("文件操作 / File Operations:")
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")
	fileutils.WriteFileString(srcFile, "这是源文件", 0644)
	fmt.Println("复制文件:", srcFile, "->", dstFile)
	err = fileutils.CopyFile(srcFile, dstFile)
	if err != nil {
		fmt.Printf("复制文件失败: %v\n", err)
	} else {
		copiedContent, _ := fileutils.ReadFileString(dstFile)
		fmt.Println("复制后内容:", copiedContent)
	}

	moveFile := filepath.Join(tmpDir, "moved.txt")
	fmt.Println("移动文件:", dstFile, "->", moveFile)
	err = fileutils.MoveFile(dstFile, moveFile)
	if err != nil {
		fmt.Printf("移动文件失败: %v\n", err)
	} else {
		fmt.Println("移动成功，原文件存在:", fileutils.Exists(dstFile))
		fmt.Println("目标文件存在:", fileutils.Exists(moveFile))
	}
	fmt.Println()

	// 目录操作
	fmt.Println("目录操作 / Directory Operations:")
	newDir := filepath.Join(tmpDir, "new", "sub", "dir")
	err = fileutils.CreateDir(newDir, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
	} else {
		fmt.Println("创建目录:", newDir)
		fmt.Println("目录存在:", fileutils.IsDir(newDir))
	}
	fmt.Println()

	// 文件查找
	fmt.Println("文件查找 / File Finding:")
	// 创建多个测试文件
	testFiles := []string{"file1.txt", "file2.txt", "file3.log"}
	for _, f := range testFiles {
		testPath := filepath.Join(tmpDir, f)
		fileutils.WriteFileString(testPath, "test", 0644)
	}
	foundFiles, err := fileutils.FindFiles(tmpDir, "*.txt")
	if err != nil {
		fmt.Printf("查找文件失败: %v\n", err)
	} else {
		fmt.Println("找到的 .txt 文件:")
		for _, f := range foundFiles {
			fmt.Println("  -", f)
		}
	}
	fmt.Println()

	// 文件类型
	fmt.Println("文件类型 / File Type:")
	testTypes := []string{"file.txt", "file.json", "file.go", "file.jpg", "file.pdf", "file.xyz"}
	for _, f := range testTypes {
		fmt.Printf("  %s -> %s\n", f, fileutils.GetFileType(f))
	}
	fmt.Println()

	// 目录遍历
	fmt.Println("目录遍历 / Directory Walking:")
	fmt.Println("遍历目录:", tmpDir)
	err = fileutils.WalkDir(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("  [DIR]  %s\n", path)
		} else {
			size, _ := fileutils.GetFileSize(path)
			fmt.Printf("  [FILE] %s (%s)\n", path, fileutils.FormatFileSize(size))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}
}
