package cryptutils_test

import (
	"regexp"
	"testing"

	"github.com/Rodert/go-commons/cryptutils"
)

func TestGenerateUUID(t *testing.T) {
	// UUID格式的正则表达式
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	// 生成多个UUID并验证格式
	for i := 0; i < 5; i++ {
		uuid, err := cryptutils.GenerateUUID()
		if err != nil {
			t.Errorf("GenerateUUID() unexpected error: %v", err)
			continue
		}

		// 验证UUID格式
		if !uuidRegex.MatchString(uuid) {
			t.Errorf("GenerateUUID() = %q; does not match UUID format", uuid)
		}

		// 验证UUID长度
		if len(uuid) != 36 {
			t.Errorf("GenerateUUID() = %q; length = %d, want 36", uuid, len(uuid))
		}
	}

	// 验证多次生成的UUID是否不同
	uuid1, _ := cryptutils.GenerateUUID()
	uuid2, _ := cryptutils.GenerateUUID()
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUID() returned identical UUIDs: %q", uuid1)
	}
}

func TestGenerateUUIDWithoutHyphens(t *testing.T) {
	// 无连字符UUID格式的正则表达式
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{32}$`)

	// 生成多个无连字符UUID并验证格式
	for i := 0; i < 5; i++ {
		uuid, err := cryptutils.GenerateUUIDWithoutHyphens()
		if err != nil {
			t.Errorf("GenerateUUIDWithoutHyphens() unexpected error: %v", err)
			continue
		}

		// 验证UUID格式
		if !uuidRegex.MatchString(uuid) {
			t.Errorf("GenerateUUIDWithoutHyphens() = %q; does not match UUID format", uuid)
		}

		// 验证UUID长度
		if len(uuid) != 32 {
			t.Errorf("GenerateUUIDWithoutHyphens() = %q; length = %d, want 32", uuid, len(uuid))
		}

		// 验证不包含连字符
		if regexp.MustCompile(`-`).MatchString(uuid) {
			t.Errorf("GenerateUUIDWithoutHyphens() = %q; contains hyphens", uuid)
		}
	}

	// 验证多次生成的UUID是否不同
	uuid1, _ := cryptutils.GenerateUUIDWithoutHyphens()
	uuid2, _ := cryptutils.GenerateUUIDWithoutHyphens()
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUIDWithoutHyphens() returned identical UUIDs: %q", uuid1)
	}
}

func TestGenerateRandomHex(t *testing.T) {
	lengths := []int{0, 2, 10, 16, 32, 64}

	for _, length := range lengths {
		// 生成随机十六进制字符串
		hex, err := cryptutils.GenerateRandomHex(length)
		if err != nil {
			t.Errorf("GenerateRandomHex(%d) unexpected error: %v", length, err)
			continue
		}

		// 如果长度是奇数，确保结果长度是偶数（因为函数会自动调整）
		expectedLength := length
		if length%2 != 0 {
			expectedLength = length + 1
		}

		// 检查长度
		if len(hex) != expectedLength {
			t.Errorf("GenerateRandomHex(%d) returned string of length %d; want %d",
				length, len(hex), expectedLength)
		}

		// 验证是否只包含十六进制字符
		if !regexp.MustCompile(`^[0-9a-f]*$`).MatchString(hex) {
			t.Errorf("GenerateRandomHex(%d) = %q; contains non-hex characters", length, hex)
		}

		// 对于长度大于0的情况，生成第二个随机字符串并确保它们不同
		if length > 0 {
			hex2, err := cryptutils.GenerateRandomHex(length)
			if err != nil {
				t.Errorf("Second GenerateRandomHex(%d) unexpected error: %v", length, err)
				continue
			}

			// 检查两次生成的随机字符串是否相同（概率极低）
			if hex == hex2 {
				t.Errorf("GenerateRandomHex(%d) returned identical values in two calls: %q", length, hex)
			}
		}
	}
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		length  int
		charset string
	}{
		{0, ""},
		{10, ""},
		{20, ""},
		{10, "ABC"},
		{10, "123"},
		{10, "!@#"},
	}

	for _, test := range tests {
		// 生成随机字符串
		str, err := cryptutils.GenerateRandomString(test.length, test.charset)
		if err != nil {
			t.Errorf("GenerateRandomString(%d, %q) unexpected error: %v",
				test.length, test.charset, err)
			continue
		}

		// 检查长度
		if len(str) != test.length {
			t.Errorf("GenerateRandomString(%d, %q) returned string of length %d; want %d",
				test.length, test.charset, len(str), test.length)
		}

		// 如果指定了字符集，验证生成的字符串只包含字符集中的字符
		if test.charset != "" {
			for _, char := range str {
				if !regexp.MustCompile(`[` + regexp.QuoteMeta(test.charset) + `]`).MatchString(string(char)) {
					t.Errorf("GenerateRandomString(%d, %q) = %q; contains character %q not in charset",
						test.length, test.charset, str, string(char))
					break
				}
			}
		}

		// 对于长度大于0的情况，生成第二个随机字符串并确保它们不同
		if test.length > 0 {
			str2, err := cryptutils.GenerateRandomString(test.length, test.charset)
			if err != nil {
				t.Errorf("Second GenerateRandomString(%d, %q) unexpected error: %v",
					test.length, test.charset, err)
				continue
			}

			// 如果字符集非常小且长度短，可能会生成相同的字符串，所以只在字符集足够大时检查
			if (test.charset == "" || len(test.charset) > 5) && test.length > 5 {
				if str == str2 {
					t.Errorf("GenerateRandomString(%d, %q) returned identical values in two calls: %q",
						test.length, test.charset, str)
				}
			}
		}
	}
}
