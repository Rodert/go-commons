// Package fileutils 提供文件/IO相关的工具函数
// Package fileutils provides file/IO utility functions
package fileutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile 读取整个文件内容
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - []byte: 文件内容 / file content
//   - error: 如果读取失败则返回错误 / error if reading fails
//
// 示例 / Example:
//   content, err := ReadFile("test.txt")
//
// ReadFile reads the entire file content
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// ReadFileString 读取整个文件内容为字符串
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - string: 文件内容 / file content
//   - error: 如果读取失败则返回错误 / error if reading fails
//
// 示例 / Example:
//   content, err := ReadFileString("test.txt")
//
// ReadFileString reads the entire file content as string
func ReadFileString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile 将内容写入文件
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//   - data: 要写入的数据 / data to write
//   - perm: 文件权限（如果文件不存在） / file permissions (if file doesn't exist)
//
// 返回值 / Returns:
//   - error: 如果写入失败则返回错误 / error if writing fails
//
// 示例 / Example:
//   err := WriteFile("test.txt", []byte("hello"), 0644)
//
// WriteFile writes data to a file
func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}
	return os.WriteFile(filePath, data, perm)
}

// WriteFileString 将字符串内容写入文件
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//   - content: 要写入的内容 / content to write
//   - perm: 文件权限（如果文件不存在） / file permissions (if file doesn't exist)
//
// 返回值 / Returns:
//   - error: 如果写入失败则返回错误 / error if writing fails
//
// 示例 / Example:
//   err := WriteFileString("test.txt", "hello", 0644)
//
// WriteFileString writes string content to a file
func WriteFileString(filePath string, content string, perm os.FileMode) error {
	return WriteFile(filePath, []byte(content), perm)
}

// ReadFileLines 按行读取文件内容
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - []string: 文件行内容（不包含换行符） / file lines (without newline)
//   - error: 如果读取失败则返回错误 / error if reading fails
//
// 示例 / Example:
//   lines, err := ReadFileLines("test.txt")
//
// ReadFileLines reads file content line by line
func ReadFileLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content := string(data)
	// 统一处理不同操作系统的换行符
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(content, "\n")

	// 移除最后一个空行（如果文件以换行符结尾）
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

// Exists 检查文件或目录是否存在
//
// 参数 / Parameters:
//   - path: 文件或目录路径 / file or directory path
//
// 返回值 / Returns:
//   - bool: 如果存在返回true / true if exists
//
// 示例 / Example:
//   exists := Exists("test.txt")
//
// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir 检查路径是否为目录
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - bool: 如果是目录返回true / true if is directory
//
// 示例 / Example:
//   isDir := IsDir("/path/to/dir")
//
// IsDir checks if a path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 检查路径是否为文件
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - bool: 如果是文件返回true / true if is file
//
// 示例 / Example:
//   isFile := IsFile("test.txt")
//
// IsFile checks if a path is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// GetFileSize 获取文件大小（字节）
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - int64: 文件大小（字节） / file size in bytes
//   - error: 如果获取失败则返回错误 / error if fails
//
// 示例 / Example:
//   size, err := GetFileSize("test.txt")
//
// GetFileSize gets file size in bytes
func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// FormatFileSize 格式化文件大小为人类可读的格式
//
// 参数 / Parameters:
//   - size: 文件大小（字节） / file size in bytes
//
// 返回值 / Returns:
//   - string: 格式化后的文件大小 / formatted file size
//
// 示例 / Example:
//   FormatFileSize(1024) // "1.00 KB"
//
// FormatFileSize formats file size to human-readable format
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	return fmt.Sprintf("%.2f %s", float64(size)/float64(div), units[exp+1])
}

// CopyFile 复制文件
//
// 参数 / Parameters:
//   - src: 源文件路径 / source file path
//   - dst: 目标文件路径 / destination file path
//
// 返回值 / Returns:
//   - error: 如果复制失败则返回错误 / error if copying fails
//
// 示例 / Example:
//   err := CopyFile("source.txt", "dest.txt")
//
// CopyFile copies a file
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %w", err)
	}
	defer sourceFile.Close()

	// 确保目标目录存在
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目标目录失败: %w", err)
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	// 复制文件权限
	sourceInfo, err := os.Stat(src)
	if err == nil {
		os.Chmod(dst, sourceInfo.Mode())
	}

	return nil
}

// MoveFile 移动或重命名文件
//
// 参数 / Parameters:
//   - src: 源文件路径 / source file path
//   - dst: 目标文件路径 / destination file path
//
// 返回值 / Returns:
//   - error: 如果移动失败则返回错误 / error if moving fails
//
// 示例 / Example:
//   err := MoveFile("old.txt", "new.txt")
//
// MoveFile moves or renames a file
func MoveFile(src, dst string) error {
	// 确保目标目录存在
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目标目录失败: %w", err)
	}

	return os.Rename(src, dst)
}

// DeleteFile 删除文件
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - error: 如果删除失败则返回错误 / error if deletion fails
//
// 示例 / Example:
//   err := DeleteFile("test.txt")
//
// DeleteFile deletes a file
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// DeleteDir 删除目录（包括所有内容）
//
// 参数 / Parameters:
//   - dirPath: 目录路径 / directory path
//
// 返回值 / Returns:
//   - error: 如果删除失败则返回错误 / error if deletion fails
//
// 示例 / Example:
//   err := DeleteDir("/path/to/dir")
//
// DeleteDir deletes a directory and all its contents
func DeleteDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// CreateDir 创建目录（包括父目录）
//
// 参数 / Parameters:
//   - dirPath: 目录路径 / directory path
//   - perm: 目录权限 / directory permissions
//
// 返回值 / Returns:
//   - error: 如果创建失败则返回错误 / error if creation fails
//
// 示例 / Example:
//   err := CreateDir("/path/to/dir", 0755)
//
// CreateDir creates a directory and all parent directories
func CreateDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

// JoinPath 连接路径组件
//
// 参数 / Parameters:
//   - elements: 路径组件 / path elements
//
// 返回值 / Returns:
//   - string: 连接后的路径 / joined path
//
// 示例 / Example:
//   path := JoinPath("dir", "subdir", "file.txt")
//
// JoinPath joins path elements
func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

// CleanPath 清理路径（移除多余的路径分隔符和"."/".."）
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - string: 清理后的路径 / cleaned path
//
// 示例 / Example:
//   clean := CleanPath("./dir/../file.txt")
//
// CleanPath cleans a path (removes redundant separators and "."/"..")
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// BaseName 获取路径的最后一部分（文件名或目录名）
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - string: 文件名或目录名 / file or directory name
//
// 示例 / Example:
//   name := BaseName("/path/to/file.txt") // "file.txt"
//
// BaseName returns the last element of path
func BaseName(path string) string {
	return filepath.Base(path)
}

// DirName 获取路径的目录部分
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - string: 目录路径 / directory path
//
// 示例 / Example:
//   dir := DirName("/path/to/file.txt") // "/path/to"
//
// DirName returns the directory part of path
func DirName(path string) string {
	return filepath.Dir(path)
}

// ExtName 获取文件扩展名（包含点）
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - string: 文件扩展名 / file extension
//
// 示例 / Example:
//   ext := ExtName("file.txt") // ".txt"
//
// ExtName returns the file extension (with dot)
func ExtName(filePath string) string {
	return filepath.Ext(filePath)
}

// FileNameWithoutExt 获取不带扩展名的文件名
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - string: 不带扩展名的文件名 / file name without extension
//
// 示例 / Example:
//   name := FileNameWithoutExt("file.txt") // "file"
//
// FileNameWithoutExt returns the file name without extension
func FileNameWithoutExt(filePath string) string {
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// AbsPath 获取绝对路径
//
// 参数 / Parameters:
//   - path: 路径 / path
//
// 返回值 / Returns:
//   - string: 绝对路径 / absolute path
//   - error: 如果获取失败则返回错误 / error if fails
//
// 示例 / Example:
//   abs, err := AbsPath("./file.txt")
//
// AbsPath returns the absolute path
func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}

// WalkDir 遍历目录，对每个文件执行回调函数
//
// 参数 / Parameters:
//   - root: 根目录路径 / root directory path
//   - fn: 对每个文件执行的函数 / function to execute for each file
//
// 返回值 / Returns:
//   - error: 如果遍历失败则返回错误 / error if walking fails
//
// 示例 / Example:
//   WalkDir("/path/to/dir", func(path string, info os.FileInfo, err error) error {
//     fmt.Println(path)
//     return nil
//   })
//
// WalkDir walks the directory tree
func WalkDir(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}

// FindFiles 在目录中查找匹配模式的文件
//
// 参数 / Parameters:
//   - root: 根目录路径 / root directory path
//   - pattern: 文件匹配模式（如 "*.txt"） / file pattern (e.g., "*.txt")
//
// 返回值 / Returns:
//   - []string: 匹配的文件路径列表 / matched file paths
//   - error: 如果查找失败则返回错误 / error if searching fails
//
// 示例 / Example:
//   files, err := FindFiles("/path/to/dir", "*.txt")
//
// FindFiles finds files matching the pattern in a directory
func FindFiles(root, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			matched, err := filepath.Match(pattern, info.Name())
			if err != nil {
				return err
			}
			if matched {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}

// GetFileType 获取文件类型描述（基于扩展名）
//
// 参数 / Parameters:
//   - filePath: 文件路径 / file path
//
// 返回值 / Returns:
//   - string: 文件类型描述 / file type description
//
// 示例 / Example:
//   fileType := GetFileType("file.txt") // "text"
//
// GetFileType returns file type description based on extension
func GetFileType(filePath string) string {
	ext := strings.ToLower(ExtName(filePath))

	typeMap := map[string]string{
		".txt":  "text",
		".md":   "markdown",
		".json": "json",
		".xml":  "xml",
		".html": "html",
		".css":  "css",
		".js":   "javascript",
		".go":   "go",
		".py":   "python",
		".java": "java",
		".jpg":  "image",
		".jpeg": "image",
		".png":  "image",
		".gif":  "image",
		".pdf":  "pdf",
		".zip":  "archive",
		".tar":  "archive",
		".gz":   "archive",
	}

	if fileType, ok := typeMap[ext]; ok {
		return fileType
	}
	return "unknown"
}
