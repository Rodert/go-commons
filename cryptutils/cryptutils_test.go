package cryptutils_test

import (
	"testing"

	"github.com/Rodert/go-commons/cryptutils"
)

func TestMD5Hash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"Hello, World!", "65a8e27d8879283831b664bd8b7f0ad4"},
	}

	for _, test := range tests {
		result := cryptutils.MD5Hash([]byte(test.input))
		if result != test.expected {
			t.Errorf("MD5Hash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestSHA1Hash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"Hello, World!", "0a0a9f2a6772942557ab5355d76af442f8f65e01"},
	}

	for _, test := range tests {
		result := cryptutils.SHA1Hash([]byte(test.input))
		if result != test.expected {
			t.Errorf("SHA1Hash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestSHA256Hash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"Hello, World!", "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"},
	}

	for _, test := range tests {
		result := cryptutils.SHA256Hash([]byte(test.input))
		if result != test.expected {
			t.Errorf("SHA256Hash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestSHA512Hash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
		{"Hello, World!", "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387"},
	}

	for _, test := range tests {
		result := cryptutils.SHA512Hash([]byte(test.input))
		if result != test.expected {
			t.Errorf("SHA512Hash(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "aGVsbG8="},
		{"Hello, World!", "SGVsbG8sIFdvcmxkIQ=="},
	}

	for _, test := range tests {
		result := cryptutils.Base64Encode([]byte(test.input))
		if result != test.expected {
			t.Errorf("Base64Encode(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		{"", "", false},
		{"aGVsbG8=", "hello", false},
		{"SGVsbG8sIFdvcmxkIQ==", "Hello, World!", false},
		{"invalid base64", "", true},
	}

	for _, test := range tests {
		result, err := cryptutils.Base64Decode(test.input)

		if test.expectError && err == nil {
			t.Errorf("Base64Decode(%q) expected error, got nil", test.input)
			continue
		}

		if !test.expectError && err != nil {
			t.Errorf("Base64Decode(%q) unexpected error: %v", test.input, err)
			continue
		}

		// 只有在没有预期错误的情况下才比较结果
		if !test.expectError && string(result) != test.expected {
			t.Errorf("Base64Decode(%q) = %v; want %v", test.input, string(result), test.expected)
		}
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
	}{
		{"hello", "0123456789abcdef0123456789abcdef"},
		{"Hello, World!", "0123456789abcdef0123456789abcdef"},
		{"This is a longer text to test AES encryption and decryption", "0123456789abcdef0123456789abcdef"},
	}

	for _, test := range tests {
		// 加密
		ciphertext, err := cryptutils.AESEncrypt([]byte(test.plaintext), []byte(test.key))
		if err != nil {
			t.Errorf("AESEncrypt(%q, %q) unexpected error: %v", test.plaintext, test.key, err)
			continue
		}

		// 确保加密后的文本不同于原文
		if string([]byte(test.plaintext)) == string(ciphertext) {
			t.Errorf("AESEncrypt(%q, %q) = %q; encryption did not change the text",
				test.plaintext, test.key, ciphertext)
			continue
		}

		// 解密
		decrypted, err := cryptutils.AESDecrypt(ciphertext, []byte(test.key))
		if err != nil {
			t.Errorf("AESDecrypt(%q, %q) unexpected error: %v", ciphertext, test.key, err)
			continue
		}

		// 确保解密后的文本与原文相同
		if string(decrypted) != test.plaintext {
			t.Errorf("AESDecrypt(AESEncrypt(%q)) = %q; want %q",
				test.plaintext, string(decrypted), test.plaintext)
		}
	}

	// 测试无效密钥
	_, err := cryptutils.AESEncrypt([]byte("hello"), []byte("short key"))
	if err == nil {
		t.Error("AESEncrypt with invalid key length did not return error")
	}

	// 测试无效的密文
	_, err = cryptutils.AESDecrypt([]byte("invalid ciphertext"), []byte("0123456789abcdef0123456789abcdef"))
	if err == nil {
		t.Error("AESDecrypt with invalid ciphertext did not return error")
	}
}

func TestGenerateRandomBytes(t *testing.T) {
	lengths := []int{0, 1, 10, 16, 32, 64}

	for _, length := range lengths {
		// 生成随机字节
		bytes, err := cryptutils.GenerateRandomBytes(length)
		if err != nil {
			t.Errorf("GenerateRandomBytes(%d) unexpected error: %v", length, err)
			continue
		}

		// 检查长度
		if len(bytes) != length {
			t.Errorf("GenerateRandomBytes(%d) returned %d bytes; want %d",
				length, len(bytes), length)
		}

		// 对于长度大于0的情况，生成第二个随机字节数组并确保它们不同
		if length > 0 {
			bytes2, err := cryptutils.GenerateRandomBytes(length)
			if err != nil {
				t.Errorf("Second GenerateRandomBytes(%d) unexpected error: %v", length, err)
				continue
			}

			// 检查两次生成的随机字节是否相同（概率极低）
			if string(bytes) == string(bytes2) {
				t.Errorf("GenerateRandomBytes(%d) returned identical values in two calls", length)
			}
		}
	}
}
