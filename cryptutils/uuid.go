package cryptutils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// GenerateUUID 生成一个随机的UUID (版本4)
// 返回:
//   - string: 生成的UUID字符串
//   - error: 如果生成失败则返回错误信息
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return "", err
	}

	// 设置版本 (4) 和变体位
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // 版本 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // 变体 RFC4122

	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16]), nil
}

// GenerateUUIDWithoutHyphens 生成一个没有连字符的UUID
// 返回:
//   - string: 生成的UUID字符串（无连字符）
//   - error: 如果生成失败则返回错误信息
func GenerateUUIDWithoutHyphens() (string, error) {
	uuid, err := GenerateUUID()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(uuid, "-", ""), nil
}

// GenerateRandomHex 生成指定长度的随机十六进制字符串
// 参数:
//   - length: 要生成的十六进制字符串长度（字节数的两倍）
//
// 返回:
//   - string: 生成的十六进制字符串
//   - error: 如果生成失败则返回错误信息
func GenerateRandomHex(length int) (string, error) {
	// 确保长度是偶数
	if length%2 != 0 {
		length++
	}

	// 计算需要的字节数
	bytes := length / 2
	randomBytes, err := GenerateRandomBytes(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}

// GenerateRandomString 生成指定长度的随机字符串
// 参数:
//   - length: 要生成的字符串长度
//   - charset: 字符集（如果为空，则使用默认字符集）
//
// 返回:
//   - string: 生成的随机字符串
//   - error: 如果生成失败则返回错误信息
func GenerateRandomString(length int, charset string) (string, error) {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	charsetLength := len(charset)
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	result := make([]byte, length)
	for i, b := range randomBytes {
		result[i] = charset[int(b)%charsetLength]
	}

	return string(result), nil
}
