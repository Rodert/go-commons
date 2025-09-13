// Package cryptutils 提供加密解密相关的工具函数
package cryptutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// MD5Hash 计算字符串的MD5哈希值
// 参数:
//   - data: 要计算哈希的数据
//
// 返回:
//   - string: 十六进制格式的MD5哈希值
func MD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// SHA1Hash 计算字符串的SHA1哈希值
// 参数:
//   - data: 要计算哈希的数据
//
// 返回:
//   - string: 十六进制格式的SHA1哈希值
func SHA1Hash(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

// SHA256Hash 计算字符串的SHA256哈希值
// 参数:
//   - data: 要计算哈希的数据
//
// 返回:
//   - string: 十六进制格式的SHA256哈希值
func SHA256Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// SHA512Hash 计算字符串的SHA512哈希值
// 参数:
//   - data: 要计算哈希的数据
//
// 返回:
//   - string: 十六进制格式的SHA512哈希值
func SHA512Hash(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// Base64Encode 将数据编码为Base64字符串
// 参数:
//   - data: 要编码的数据
//
// 返回:
//   - string: Base64编码的字符串
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode 解码Base64字符串
// 参数:
//   - encoded: Base64编码的字符串
//
// 返回:
//   - []byte: 解码后的数据
//   - error: 如果解码失败则返回错误信息
func Base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// AESEncrypt 使用AES-GCM模式加密数据
// 参数:
//   - plaintext: 要加密的明文数据
//   - key: 加密密钥 (必须是16, 24或32字节长，对应AES-128, AES-192或AES-256)
//
// 返回:
//   - []byte: 加密后的数据 (包含随机生成的nonce)
//   - error: 如果加密失败则返回错误信息
func AESEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 创建随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 加密数据
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AESDecrypt 使用AES-GCM模式解密数据
// 参数:
//   - ciphertext: 要解密的密文数据 (包含nonce)
//   - key: 解密密钥 (必须是16, 24或32字节长，对应AES-128, AES-192或AES-256)
//
// 返回:
//   - []byte: 解密后的明文数据
//   - error: 如果解密失败则返回错误信息
func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 检查密文长度
	if len(ciphertext) < gcm.NonceSize() {
		return nil, fmt.Errorf("密文太短")
	}

	// 提取nonce
	nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	// 解密数据
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// GenerateRandomBytes 生成指定长度的随机字节
// 参数:
//   - length: 要生成的随机字节长度
//
// 返回:
//   - []byte: 生成的随机字节
//   - error: 如果生成失败则返回错误信息
func GenerateRandomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
