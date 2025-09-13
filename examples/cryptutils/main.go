package main

import (
	"fmt"

	"github.com/Rodert/go-commons/cryptutils"
)

func main() {
	// 演示MD5哈希
	text := "Hello, World!"
	md5Hash := cryptutils.MD5Hash([]byte(text))
	fmt.Printf("MD5(\"%s\") = %s\n", text, md5Hash)

	// 演示SHA1哈希
	sha1Hash := cryptutils.SHA1Hash([]byte(text))
	fmt.Printf("SHA1(\"%s\") = %s\n", text, sha1Hash)

	// 演示SHA256哈希
	sha256Hash := cryptutils.SHA256Hash([]byte(text))
	fmt.Printf("SHA256(\"%s\") = %s\n", text, sha256Hash)

	// 演示SHA512哈希
	sha512Hash := cryptutils.SHA512Hash([]byte(text))
	fmt.Printf("SHA512(\"%s\") = %s\n", text, sha512Hash)

	// 演示Base64编码和解码
	encoded := cryptutils.Base64Encode([]byte(text))
	fmt.Printf("Base64Encode(\"%s\") = %s\n", text, encoded)

	decoded, err := cryptutils.Base64Decode(encoded)
	if err != nil {
		fmt.Printf("Base64Decode error: %v\n", err)
	} else {
		fmt.Printf("Base64Decode(\"%s\") = %s\n", encoded, string(decoded))
	}

	// 演示AES加密和解密
	key := "0123456789abcdef0123456789abcdef" // 32字节的AES-256密钥
	ciphertext, err := cryptutils.AESEncrypt([]byte(text), []byte(key))
	if err != nil {
		fmt.Printf("AESEncrypt error: %v\n", err)
	} else {
		fmt.Printf("AESEncrypt(\"%s\") = %x\n", text, ciphertext)

		decrypted, err := cryptutils.AESDecrypt(ciphertext, []byte(key))
		if err != nil {
			fmt.Printf("AESDecrypt error: %v\n", err)
		} else {
			fmt.Printf("AESDecrypt(ciphertext) = %s\n", string(decrypted))
		}
	}

	// 演示生成随机字节
	randomBytes, err := cryptutils.GenerateRandomBytes(16)
	if err != nil {
		fmt.Printf("GenerateRandomBytes error: %v\n", err)
	} else {
		fmt.Printf("GenerateRandomBytes(16) = %x\n", randomBytes)
	}
}
