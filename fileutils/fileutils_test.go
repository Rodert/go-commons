package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "test_read_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	testContent := []byte("Hello, World!")
	if _, err := tmpFile.Write(testContent); err != nil {
		t.Fatalf("写入测试内容失败: %v", err)
	}
	tmpFile.Close()

	// 测试读取
	content, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadFile(%q) 返回错误: %v", tmpFile.Name(), err)
	}
	if string(content) != string(testContent) {
		t.Errorf("ReadFile(%q) = %q; want %q", tmpFile.Name(), content, testContent)
	}
}

func TestReadFileString(t *testing.T) {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "test_read_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	testContent := "Hello, World!"
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("写入测试内容失败: %v", err)
	}
	tmpFile.Close()

	// 测试读取
	content, err := ReadFileString(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadFileString(%q) 返回错误: %v", tmpFile.Name(), err)
	}
	if content != testContent {
		t.Errorf("ReadFileString(%q) = %q; want %q", tmpFile.Name(), content, testContent)
	}
}

func TestWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_write.txt")
	testContent := []byte("Hello, World!")

	err := WriteFile(testFile, testContent, 0644)
	if err != nil {
		t.Errorf("WriteFile(%q) 返回错误: %v", testFile, err)
	}

	// 验证文件内容
	content, err := ReadFile(testFile)
	if err != nil {
		t.Errorf("读取写入的文件失败: %v", err)
	}
	if string(content) != string(testContent) {
		t.Errorf("WriteFile 写入的内容不匹配: got %q, want %q", content, testContent)
	}
}

func TestWriteFileString(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_write_string.txt")
	testContent := "Hello, World!"

	err := WriteFileString(testFile, testContent, 0644)
	if err != nil {
		t.Errorf("WriteFileString(%q) 返回错误: %v", testFile, err)
	}

	// 验证文件内容
	content, err := ReadFileString(testFile)
	if err != nil {
		t.Errorf("读取写入的文件失败: %v", err)
	}
	if content != testContent {
		t.Errorf("WriteFileString 写入的内容不匹配: got %q, want %q", content, testContent)
	}
}

func TestReadFileLines(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_lines_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testContent := "line1\nline2\nline3\n"
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("写入测试内容失败: %v", err)
	}
	tmpFile.Close()

	lines, err := ReadFileLines(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadFileLines(%q) 返回错误: %v", tmpFile.Name(), err)
	}

	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Errorf("ReadFileLines(%q) 返回 %d 行; want %d 行", tmpFile.Name(), len(lines), len(expected))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("ReadFileLines(%q)[%d] = %q; want %q", tmpFile.Name(), i, line, expected[i])
		}
	}
}

func TestExists(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_exists_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	if !Exists(tmpFile.Name()) {
		t.Errorf("Exists(%q) = false; want true", tmpFile.Name())
	}

	if Exists("/nonexistent/path/that/does/not/exist") {
		t.Error("Exists(不存在的路径) = true; want false")
	}
}

func TestIsDir(t *testing.T) {
	tmpDir := t.TempDir()

	if !IsDir(tmpDir) {
		t.Errorf("IsDir(%q) = false; want true", tmpDir)
	}

	tmpFile, err := os.CreateTemp(tmpDir, "test_file_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer tmpFile.Close()

	if IsDir(tmpFile.Name()) {
		t.Errorf("IsDir(%q) = true; want false", tmpFile.Name())
	}
}

func TestIsFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "test_file_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer tmpFile.Close()

	if !IsFile(tmpFile.Name()) {
		t.Errorf("IsFile(%q) = false; want true", tmpFile.Name())
	}

	if IsFile(tmpDir) {
		t.Errorf("IsFile(%q) = true; want false", tmpDir)
	}
}

func TestGetFileSize(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_size_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testContent := "Hello, World!"
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("写入测试内容失败: %v", err)
	}
	tmpFile.Close()

	size, err := GetFileSize(tmpFile.Name())
	if err != nil {
		t.Errorf("GetFileSize(%q) 返回错误: %v", tmpFile.Name(), err)
	}
	if size != int64(len(testContent)) {
		t.Errorf("GetFileSize(%q) = %d; want %d", tmpFile.Name(), size, len(testContent))
	}
}

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"bytes", 512, "512 B"},
		{"kilobytes", 1024, "1.00 KB"},
		{"megabytes", 1024 * 1024, "1.00 MB"},
		{"gigabytes", 1024 * 1024 * 1024, "1.00 GB"},
		{"terabytes", 1024 * 1024 * 1024 * 1024, "1.00 TB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FormatFileSize(test.size)
			if result != test.expected {
				t.Errorf("FormatFileSize(%d) = %q; want %q", test.size, result, test.expected)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	testContent := []byte("Hello, World!")
	if err := WriteFile(srcFile, testContent, 0644); err != nil {
		t.Fatalf("创建源文件失败: %v", err)
	}

	err := CopyFile(srcFile, dstFile)
	if err != nil {
		t.Errorf("CopyFile(%q, %q) 返回错误: %v", srcFile, dstFile, err)
	}

	// 验证目标文件内容
	content, err := ReadFile(dstFile)
	if err != nil {
		t.Errorf("读取复制的文件失败: %v", err)
	}
	if string(content) != string(testContent) {
		t.Errorf("CopyFile 复制的内容不匹配: got %q, want %q", content, testContent)
	}
}

func TestMoveFile(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	testContent := []byte("Hello, World!")
	if err := WriteFile(srcFile, testContent, 0644); err != nil {
		t.Fatalf("创建源文件失败: %v", err)
	}

	err := MoveFile(srcFile, dstFile)
	if err != nil {
		t.Errorf("MoveFile(%q, %q) 返回错误: %v", srcFile, dstFile, err)
	}

	// 验证源文件不存在
	if Exists(srcFile) {
		t.Errorf("MoveFile 后源文件仍存在: %q", srcFile)
	}

	// 验证目标文件存在且内容正确
	content, err := ReadFile(dstFile)
	if err != nil {
		t.Errorf("读取移动的文件失败: %v", err)
	}
	if string(content) != string(testContent) {
		t.Errorf("MoveFile 移动的内容不匹配: got %q, want %q", content, testContent)
	}
}

func TestDeleteFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_delete_*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	tmpFile.Close()

	if !Exists(tmpFile.Name()) {
		t.Fatalf("临时文件不存在: %q", tmpFile.Name())
	}

	err = DeleteFile(tmpFile.Name())
	if err != nil {
		t.Errorf("DeleteFile(%q) 返回错误: %v", tmpFile.Name(), err)
	}

	if Exists(tmpFile.Name()) {
		t.Errorf("DeleteFile 后文件仍存在: %q", tmpFile.Name())
	}
}

func TestDeleteDir(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	if err := WriteFileString(testFile, "test", 0644); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	if !Exists(tmpDir) {
		t.Fatalf("临时目录不存在: %q", tmpDir)
	}

	err := DeleteDir(tmpDir)
	if err != nil {
		t.Errorf("DeleteDir(%q) 返回错误: %v", tmpDir, err)
	}

	if Exists(tmpDir) {
		t.Errorf("DeleteDir 后目录仍存在: %q", tmpDir)
	}
}

func TestCreateDir(t *testing.T) {
	tmpDir := t.TempDir()
	newDir := filepath.Join(tmpDir, "new", "sub", "dir")

	err := CreateDir(newDir, 0755)
	if err != nil {
		t.Errorf("CreateDir(%q) 返回错误: %v", newDir, err)
	}

	if !IsDir(newDir) {
		t.Errorf("CreateDir 后目录不存在: %q", newDir)
	}
}

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		expected string
	}{
		{"simple", []string{"dir", "file.txt"}, filepath.Join("dir", "file.txt")},
		{"multiple", []string{"dir", "sub", "file.txt"}, filepath.Join("dir", "sub", "file.txt")},
		{"absolute", []string{"/", "path", "to", "file.txt"}, filepath.Join("/", "path", "to", "file.txt")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JoinPath(test.elements...)
			if result != test.expected {
				t.Errorf("JoinPath(%v) = %q; want %q", test.elements, result, test.expected)
			}
		})
	}
}

func TestCleanPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"simple", "./dir/../file.txt", filepath.Clean("./dir/../file.txt")},
		{"dots", ".././file.txt", filepath.Clean(".././file.txt")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CleanPath(test.path)
			if result != test.expected {
				t.Errorf("CleanPath(%q) = %q; want %q", test.path, result, test.expected)
			}
		})
	}
}

func TestBaseName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"simple", "/path/to/file.txt", "file.txt"},
		{"no dir", "file.txt", "file.txt"},
		{"directory", "/path/to/dir", "dir"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BaseName(test.path)
			if result != test.expected {
				t.Errorf("BaseName(%q) = %q; want %q", test.path, result, test.expected)
			}
		})
	}
}

func TestDirName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"simple", "/path/to/file.txt", filepath.Dir("/path/to/file.txt")},
		{"no dir", "file.txt", filepath.Dir("file.txt")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DirName(test.path)
			if result != test.expected {
				t.Errorf("DirName(%q) = %q; want %q", test.path, result, test.expected)
			}
		})
	}
}

func TestExtName(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{"with ext", "file.txt", ".txt"},
		{"no ext", "file", ""},
		{"multiple dots", "file.backup.txt", ".txt"},
		{"hidden file", ".gitignore", ".gitignore"}, // filepath.Ext returns ".gitignore" for hidden files
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ExtName(test.filePath)
			if result != test.expected {
				t.Errorf("ExtName(%q) = %q; want %q", test.filePath, result, test.expected)
			}
		})
	}
}

func TestFileNameWithoutExt(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{"with ext", "file.txt", "file"},
		{"no ext", "file", "file"},
		{"multiple dots", "file.backup.txt", "file.backup"},
		{"path with ext", "/path/to/file.txt", "file"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FileNameWithoutExt(test.filePath)
			if result != test.expected {
				t.Errorf("FileNameWithoutExt(%q) = %q; want %q", test.filePath, result, test.expected)
			}
		})
	}
}

func TestAbsPath(t *testing.T) {
	result, err := AbsPath(".")
	if err != nil {
		t.Errorf("AbsPath(.) 返回错误: %v", err)
	}
	if result == "" {
		t.Error("AbsPath(.) 返回空字符串")
	}
}

func TestFindFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// 创建测试文件
	files := []string{"file1.txt", "file2.txt", "file3.log"}
	for _, f := range files {
		testFile := filepath.Join(tmpDir, f)
		if err := WriteFileString(testFile, "test", 0644); err != nil {
			t.Fatalf("创建测试文件失败: %v", err)
		}
	}

	// 查找 .txt 文件
	foundFiles, err := FindFiles(tmpDir, "*.txt")
	if err != nil {
		t.Errorf("FindFiles(%q, *.txt) 返回错误: %v", tmpDir, err)
	}

	if len(foundFiles) != 2 {
		t.Errorf("FindFiles(%q, *.txt) 返回 %d 个文件; want 2", tmpDir, len(foundFiles))
	}
}

func TestGetFileType(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{"txt", "file.txt", "text"},
		{"json", "file.json", "json"},
		{"go", "file.go", "go"},
		{"image", "file.jpg", "image"},
		{"unknown", "file.xyz", "unknown"},
		{"no ext", "file", "unknown"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetFileType(test.filePath)
			if result != test.expected {
				t.Errorf("GetFileType(%q) = %q; want %q", test.filePath, result, test.expected)
			}
		})
	}
}
