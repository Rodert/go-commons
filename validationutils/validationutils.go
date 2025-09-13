// Package validationutils 提供数据验证相关的工具函数
package validationutils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 常用的正则表达式模式
const (
	// 邮箱地址正则表达式
	emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 中国大陆手机号正则表达式
	cnMobilePattern = `^1[3-9]\d{9}$`

	// URL正则表达式
	urlPattern = `^(https?|ftp)://[^\s/$.?#].[^\s]*$`

	// IPv4地址正则表达式
	ipv4Pattern = `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

	// 中国身份证号码正则表达式（支持15位和18位）
	cnIDCardPattern = `^(\d{15}|\d{17}[0-9Xx])$`

	// 邮政编码正则表达式
	postalCodePattern = `^\d{6}$`
)

// IsEmail 验证字符串是否为有效的电子邮件地址
// 参数:
//   - email: 要验证的电子邮件地址
//
// 返回:
//   - bool: 是否为有效的电子邮件地址
func IsEmail(email string) bool {
	regex := regexp.MustCompile(emailPattern)
	return regex.MatchString(email)
}

// IsCNMobile 验证字符串是否为有效的中国大陆手机号
// 参数:
//   - mobile: 要验证的手机号
//
// 返回:
//   - bool: 是否为有效的中国大陆手机号
func IsCNMobile(mobile string) bool {
	regex := regexp.MustCompile(cnMobilePattern)
	return regex.MatchString(mobile)
}

// IsURL 验证字符串是否为有效的URL
// 参数:
//   - url: 要验证的URL
//
// 返回:
//   - bool: 是否为有效的URL
func IsURL(url string) bool {
	regex := regexp.MustCompile(urlPattern)
	return regex.MatchString(url)
}

// IsIPv4 验证字符串是否为有效的IPv4地址
// 参数:
//   - ip: 要验证的IP地址
//
// 返回:
//   - bool: 是否为有效的IPv4地址
func IsIPv4(ip string) bool {
	regex := regexp.MustCompile(ipv4Pattern)
	return regex.MatchString(ip)
}

// IsCNIDCard 验证字符串是否为有效的中国身份证号码
// 参数:
//   - idCard: 要验证的身份证号码
//
// 返回:
//   - bool: 是否为有效的中国身份证号码
func IsCNIDCard(idCard string) bool {
	// 基本格式验证
	regex := regexp.MustCompile(cnIDCardPattern)
	if !regex.MatchString(idCard) {
		return false
	}

	// 对18位身份证进行校验码验证
	if len(idCard) == 18 {
		return validateCNIDCard18(idCard)
	}

	// 15位身份证只做格式验证
	return true
}

// validateCNIDCard18 验证18位中国身份证号码的校验码
// 参数:
//   - idCard: 18位身份证号码
//
// 返回:
//   - bool: 校验码是否正确
func validateCNIDCard18(idCard string) bool {
	// 加权因子
	weight := [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	// 校验码对应值
	validateCode := [...]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	// 计算校验码
	sum := 0
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(idCard[i]))
		sum += num * weight[i]
	}

	// 获取校验位
	mod := sum % 11
	expectedCode := validateCode[mod]

	// 比较校验码
	actualCode := byte(unicode.ToUpper(rune(idCard[17])))
	return actualCode == expectedCode
}

// IsPostalCode 验证字符串是否为有效的中国邮政编码
// 参数:
//   - code: 要验证的邮政编码
//
// 返回:
//   - bool: 是否为有效的中国邮政编码
func IsPostalCode(code string) bool {
	regex := regexp.MustCompile(postalCodePattern)
	return regex.MatchString(code)
}

// IsNumeric 验证字符串是否只包含数字
// 参数:
//   - str: 要验证的字符串
//
// 返回:
//   - bool: 是否只包含数字
func IsNumeric(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return len(str) > 0
}

// IsAlpha 验证字符串是否只包含字母
// 参数:
//   - str: 要验证的字符串
//
// 返回:
//   - bool: 是否只包含字母
func IsAlpha(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return len(str) > 0
}

// IsAlphaNumeric 验证字符串是否只包含字母和数字
// 参数:
//   - str: 要验证的字符串
//
// 返回:
//   - bool: 是否只包含字母和数字
func IsAlphaNumeric(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return len(str) > 0
}

// HasMinLength 验证字符串是否达到最小长度
// 参数:
//   - str: 要验证的字符串
//   - minLength: 最小长度
//
// 返回:
//   - bool: 是否达到最小长度
func HasMinLength(str string, minLength int) bool {
	return len(str) >= minLength
}

// HasMaxLength 验证字符串是否不超过最大长度
// 参数:
//   - str: 要验证的字符串
//   - maxLength: 最大长度
//
// 返回:
//   - bool: 是否不超过最大长度
func HasMaxLength(str string, maxLength int) bool {
	return len(str) <= maxLength
}

// IsInRange 验证数值是否在指定范围内
// 参数:
//   - value: 要验证的数值
//   - min: 最小值
//   - max: 最大值
//
// 返回:
//   - bool: 是否在指定范围内
func IsInRange(value, min, max int) bool {
	return value >= min && value <= max
}

// ContainsUppercase 验证字符串是否包含大写字母
// 参数:
//   - str: 要验证的字符串
//
// 返回:
//   - bool: 是否包含大写字母
func ContainsUppercase(str string) bool {
	for _, char := range str {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// ContainsLowercase 验证字符串是否包含小写字母
// 参数:
//   - str: 要验证的字符串
//
// 返回:
//   - bool: 是否包含小写字母
func ContainsLowercase(str string) bool {
	for _, char := range str {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

// ContainsSpecialChar 验证字符串是否包含特殊字符
// 参数:
//   - str: 要验证的字符串
//   - specialChars: 特殊字符集（如果为空，则使用默认特殊字符集）
//
// 返回:
//   - bool: 是否包含特殊字符
func ContainsSpecialChar(str string, specialChars string) bool {
	if specialChars == "" {
		specialChars = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"
	}

	for _, char := range str {
		if strings.ContainsRune(specialChars, char) {
			return true
		}
	}
	return false
}
