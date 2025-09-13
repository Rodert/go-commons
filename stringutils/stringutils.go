// Package stringutils 提供了常用的字符串操作工具函数
// Package stringutils provides common string utility functions
package stringutils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty 检查字符串是否为空
//
// 参数 / Parameters:
//
//	str - 要检查的字符串 / string to check
//
// 返回值 / Returns:
//
//	bool - 如果字符串为空返回true，否则返回false / true if string is empty, false otherwise
//
// 示例 / Example:
//
//	IsEmpty("")     // true
//	IsEmpty("abc")  // false
//
// IsEmpty checks if a string is empty
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty 检查字符串是否非空
// IsNotEmpty checks if a string is not empty
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank 检查字符串是否为空白（空或只包含空白字符）
// IsBlank checks if a string is blank (empty or contains only whitespace)
func IsBlank(str string) bool {
	if IsEmpty(str) {
		return true
	}
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsNotBlank 检查字符串是否非空白
// IsNotBlank checks if a string is not blank
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// Trim 去除字符串两端的空白字符
// Trim removes whitespace from both ends of a string
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// TrimToEmpty 去除字符串两端的空白字符，如果结果为nil则返回空字符串
// TrimToEmpty removes whitespace from both ends of a string, returns empty string if result is nil
func TrimToEmpty(str string) string {
	return Trim(str)
}

// Truncate 截断字符串到指定长度
// Truncate truncates a string to a specified length
func Truncate(str string, maxWidth int) string {
	if maxWidth < 0 {
		return ""
	}
	if len(str) <= maxWidth {
		return str
	}
	return str[0:maxWidth]
}

// TruncateWithSuffix 截断字符串到指定长度并添加后缀
// TruncateWithSuffix truncates a string to a specified length and adds a suffix
func TruncateWithSuffix(str string, maxWidth int, suffix string) string {
	// 处理特殊情况
	if maxWidth < 0 {
		return ""
	}

	// 如果原字符串长度小于等于最大宽度，直接返回原字符串
	if len(str) <= maxWidth {
		return str
	}

	// 如果后缀长度大于等于最大宽度，返回截断的后缀
	if len(suffix) >= maxWidth {
		return suffix[:maxWidth]
	}

	// 正常情况：截断字符串并添加后缀
	return str[0:maxWidth-len(suffix)] + suffix
}

// Capitalize 将字符串的第一个字符转为大写
// Capitalize converts the first character of a string to uppercase
func Capitalize(str string) string {
	if IsEmpty(str) {
		return str
	}
	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[size:]
}

// Uncapitalize 将字符串的第一个字符转为小写
// Uncapitalize converts the first character of a string to lowercase
func Uncapitalize(str string) string {
	if IsEmpty(str) {
		return str
	}
	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToLower(r)) + str[size:]
}

// ReverseString 反转字符串
// ReverseString reverses a string
func ReverseString(str string) string {
	if IsEmpty(str) {
		return str
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsAny 检查字符串是否包含任意一个子字符串
// ContainsAny checks if a string contains any of the specified substrings
func ContainsAny(str string, searchStrs ...string) bool {
	if IsEmpty(str) || len(searchStrs) == 0 {
		return false
	}
	for _, searchStr := range searchStrs {
		if strings.Contains(str, searchStr) {
			return true
		}
	}
	return false
}

// ContainsAll 检查字符串是否包含所有子字符串
// ContainsAll checks if a string contains all of the specified substrings
func ContainsAll(str string, searchStrs ...string) bool {
	if IsEmpty(str) || len(searchStrs) == 0 {
		return false
	}
	for _, searchStr := range searchStrs {
		if !strings.Contains(str, searchStr) {
			return false
		}
	}
	return true
}

// SubstringBefore 返回字符串中指定分隔符之前的子字符串
// SubstringBefore returns the substring before the first occurrence of a separator
func SubstringBefore(str string, separator string) string {
	if IsEmpty(str) || IsEmpty(separator) {
		return str
	}
	pos := strings.Index(str, separator)
	if pos == -1 {
		return str
	}
	return str[0:pos]
}

// SubstringAfter 返回字符串中指定分隔符之后的子字符串
// SubstringAfter returns the substring after the first occurrence of a separator
func SubstringAfter(str string, separator string) string {
	if IsEmpty(str) {
		return str
	}
	if IsEmpty(separator) {
		return ""
	}
	pos := strings.Index(str, separator)
	if pos == -1 {
		return ""
	}
	return str[pos+len(separator):]
}

// Join 使用分隔符连接字符串数组
// Join concatenates the elements of a string array into a single string with a separator
func Join(separator string, elements ...string) string {
	return strings.Join(elements, separator)
}

// Split 使用分隔符分割字符串
// Split divides a string into substrings using a separator
func Split(str string, separator string) []string {
	if IsEmpty(str) {
		return []string{}
	}
	return strings.Split(str, separator)
}

// EqualsIgnoreCase 忽略大小写比较两个字符串
// EqualsIgnoreCase compares two strings ignoring case
func EqualsIgnoreCase(str1 string, str2 string) bool {
	return strings.EqualFold(str1, str2)
}

// StartsWith 检查字符串是否以指定前缀开始
// StartsWith checks if a string starts with a specified prefix
func StartsWith(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// EndsWith 检查字符串是否以指定后缀结束
// EndsWith checks if a string ends with a specified suffix
func EndsWith(str string, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// RemoveStart 移除字符串开头的指定前缀
// RemoveStart removes a prefix from the start of a string
func RemoveStart(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if StartsWith(str, remove) {
		return str[len(remove):]
	}
	return str
}

// RemoveEnd 移除字符串结尾的指定后缀
// RemoveEnd removes a suffix from the end of a string
func RemoveEnd(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if EndsWith(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// Replace 替换字符串中的指定子字符串
// Replace replaces a specified substring with another substring in a string
func Replace(str string, oldStr string, newStr string, count int) string {
	if IsEmpty(str) || IsEmpty(oldStr) || oldStr == newStr {
		return str
	}
	return strings.Replace(str, oldStr, newStr, count)
}

// ReplaceAll 替换字符串中所有的指定子字符串
// ReplaceAll replaces all occurrences of a specified substring with another substring in a string
func ReplaceAll(str string, oldStr string, newStr string) string {
	return Replace(str, oldStr, newStr, -1)
}

// Repeat 重复字符串指定次数
// Repeat returns a string consisting of a specified number of copies of the original string
func Repeat(str string, repeat int) string {
	if IsEmpty(str) || repeat <= 0 {
		return ""
	}
	return strings.Repeat(str, repeat)
}

// PadLeft 在字符串左侧填充指定字符到指定长度
// PadLeft pads the left side of a string with a specified character to a specified length
func PadLeft(str string, size int, padChar rune) string {
	if size <= 0 {
		return str
	}
	pLen := size - utf8.RuneCountInString(str)
	if pLen <= 0 {
		return str
	}
	return Repeat(string(padChar), pLen) + str
}

// PadRight 在字符串右侧填充指定字符到指定长度
// PadRight pads the right side of a string with a specified character to a specified length
func PadRight(str string, size int, padChar rune) string {
	if size <= 0 {
		return str
	}
	pLen := size - utf8.RuneCountInString(str)
	if pLen <= 0 {
		return str
	}
	return str + Repeat(string(padChar), pLen)
}

// Center 在字符串两侧填充指定字符使其居中
// Center pads both sides of a string with a specified character to center it within a specified length
func Center(str string, size int, padChar rune) string {
	if size <= 0 {
		return str
	}
	pLen := size - utf8.RuneCountInString(str)
	if pLen <= 0 {
		return str
	}
	pLenLeft := pLen / 2
	pLenRight := pLen - pLenLeft
	return Repeat(string(padChar), pLenLeft) + str + Repeat(string(padChar), pLenRight)
}

// CountMatches 计算字符串中指定子字符串出现的次数
// CountMatches counts the number of occurrences of a substring in a string
func CountMatches(str string, sub string) int {
	if IsEmpty(str) || IsEmpty(sub) {
		return 0
	}
	return strings.Count(str, sub)
}

// ToUpperCase 将字符串转为大写
// ToUpperCase converts a string to uppercase
func ToUpperCase(str string) string {
	if IsEmpty(str) {
		return str
	}
	return strings.ToUpper(str)
}

// ToLowerCase 将字符串转为小写
// ToLowerCase converts a string to lowercase
func ToLowerCase(str string) string {
	if IsEmpty(str) {
		return str
	}
	return strings.ToLower(str)
}

// DefaultIfEmpty 如果字符串为空则返回默认值
// DefaultIfEmpty returns a default value if a string is empty
func DefaultIfEmpty(str string, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	}
	return str
}

// DefaultIfBlank 如果字符串为空白则返回默认值
// DefaultIfBlank returns a default value if a string is blank
func DefaultIfBlank(str string, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}
	return str
}
