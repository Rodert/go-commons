package stringutils

import (
	"strings"
	"unicode"
)

// Reverse 反转字符串
// Reverse reverses a string
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SwapCase 交换字符串中字母的大小写
// SwapCase swaps the case of all alphabetic characters in a string
func SwapCase(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

// PadCenter 在字符串两侧填充字符，使其居中
// PadCenter pads a string on both sides with a specified character
func PadCenter(str string, size int, padChar rune) string {
	if len(str) >= size {
		return str
	}
	
	padsNeeded := size - len(str)
	padLeft := padsNeeded / 2
	padRight := padsNeeded - padLeft
	
	return strings.Repeat(string(padChar), padLeft) + str + strings.Repeat(string(padChar), padRight)
}