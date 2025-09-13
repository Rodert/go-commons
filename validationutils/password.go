package validationutils

import (
	"strings"
	"unicode"
)

// PasswordStrengthLevel 密码强度级别
type PasswordStrengthLevel int

const (
	// WeakPassword 弱密码
	WeakPassword PasswordStrengthLevel = iota
	// MediumPassword 中等强度密码
	MediumPassword
	// StrongPassword 强密码
	StrongPassword
	// VeryStrongPassword 非常强的密码
	VeryStrongPassword
)

// PasswordStrengthResult 密码强度检查结果
type PasswordStrengthResult struct {
	// Level 密码强度级别
	Level PasswordStrengthLevel
	// Score 密码强度得分（0-100）
	Score int
	// Suggestions 改进建议
	Suggestions []string
}

// CheckPasswordStrength 检查密码强度
// 参数:
//   - password: 要检查的密码
//
// 返回:
//   - PasswordStrengthResult: 密码强度检查结果
func CheckPasswordStrength(password string) PasswordStrengthResult {
	result := PasswordStrengthResult{}
	suggestions := []string{}

	// 基础分数：密码长度
	passwordLength := len(password)
	lengthScore := 0
	switch {
	case passwordLength >= 12:
		lengthScore = 25
	case passwordLength >= 8:
		lengthScore = 15
		suggestions = append(suggestions, "考虑使用更长的密码（至少12个字符）")
	case passwordLength >= 6:
		lengthScore = 10
		suggestions = append(suggestions, "密码太短，建议使用至少8个字符")
	default:
		lengthScore = 5
		suggestions = append(suggestions, "密码太短，存在安全风险，建议使用至少8个字符")
	}

	// 字符多样性分数
	hasLower := ContainsLowercase(password)
	hasUpper := ContainsUppercase(password)
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}

	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"
	hasSpecial = ContainsSpecialChar(password, specialChars)

	charTypeScore := 0
	if hasLower {
		charTypeScore += 10
	} else {
		suggestions = append(suggestions, "添加小写字母可以提高密码强度")
	}

	if hasUpper {
		charTypeScore += 10
	} else {
		suggestions = append(suggestions, "添加大写字母可以提高密码强度")
	}

	if hasDigit {
		charTypeScore += 10
	} else {
		suggestions = append(suggestions, "添加数字可以提高密码强度")
	}

	if hasSpecial {
		charTypeScore += 15
	} else {
		suggestions = append(suggestions, "添加特殊字符可以提高密码强度")
	}

	// 复杂性分数
	complexityScore := 0

	// 检查是否有连续的相同字符
	hasConsecutiveSame := false
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i] == password[i+2] {
			hasConsecutiveSame = true
			break
		}
	}

	if !hasConsecutiveSame {
		complexityScore += 10
	} else {
		suggestions = append(suggestions, "避免使用连续的相同字符")
	}

	// 检查是否有连续的数字或字母
	hasSequential := false
	for i := 0; i < len(password)-2; i++ {
		// 检查数字序列
		if unicode.IsDigit(rune(password[i])) &&
			unicode.IsDigit(rune(password[i+1])) &&
			unicode.IsDigit(rune(password[i+2])) {
			if int(password[i+1])-int(password[i]) == 1 &&
				int(password[i+2])-int(password[i+1]) == 1 {
				hasSequential = true
				break
			}
		}

		// 检查字母序列
		if unicode.IsLetter(rune(password[i])) &&
			unicode.IsLetter(rune(password[i+1])) &&
			unicode.IsLetter(rune(password[i+2])) {
			if strings.ToLower(string(password[i+1]))[0]-strings.ToLower(string(password[i]))[0] == 1 &&
				strings.ToLower(string(password[i+2]))[0]-strings.ToLower(string(password[i+1]))[0] == 1 {
				hasSequential = true
				break
			}
		}
	}

	if !hasSequential {
		complexityScore += 10
	} else {
		suggestions = append(suggestions, "避免使用连续的数字或字母序列")
	}

	// 检查常见密码模式
	commonPasswords := []string{"password", "123456", "qwerty", "admin", "welcome"}
	isCommon := false
	for _, common := range commonPasswords {
		if strings.Contains(strings.ToLower(password), common) {
			isCommon = true
			break
		}
	}

	if !isCommon {
		complexityScore += 10
	} else {
		suggestions = append(suggestions, "避免使用常见的密码或密码片段")
	}

	// 计算总分
	totalScore := lengthScore + charTypeScore + complexityScore

	// 确定密码强度级别
	switch {
	case totalScore >= 80:
		result.Level = VeryStrongPassword
	case totalScore >= 60:
		result.Level = StrongPassword
	case totalScore >= 40:
		result.Level = MediumPassword
	default:
		result.Level = WeakPassword
	}

	result.Score = totalScore
	result.Suggestions = suggestions

	return result
}

// IsPasswordValid 根据指定的规则验证密码是否有效
// 参数:
//   - password: 要验证的密码
//   - minLength: 最小长度
//   - requireUpper: 是否要求包含大写字母
//   - requireLower: 是否要求包含小写字母
//   - requireDigit: 是否要求包含数字
//   - requireSpecial: 是否要求包含特殊字符
//
// 返回:
//   - bool: 密码是否有效
//   - []string: 不满足的规则列表
func IsPasswordValid(password string, minLength int, requireUpper, requireLower, requireDigit, requireSpecial bool) (bool, []string) {
	invalidReasons := []string{}

	// 检查长度
	if len(password) < minLength {
		invalidReasons = append(invalidReasons, "密码长度不足")
	}

	// 检查是否包含大写字母
	if requireUpper && !ContainsUppercase(password) {
		invalidReasons = append(invalidReasons, "密码必须包含至少一个大写字母")
	}

	// 检查是否包含小写字母
	if requireLower && !ContainsLowercase(password) {
		invalidReasons = append(invalidReasons, "密码必须包含至少一个小写字母")
	}

	// 检查是否包含数字
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}

	if requireDigit && !hasDigit {
		invalidReasons = append(invalidReasons, "密码必须包含至少一个数字")
	}

	// 检查是否包含特殊字符
	if requireSpecial && !ContainsSpecialChar(password, "") {
		invalidReasons = append(invalidReasons, "密码必须包含至少一个特殊字符")
	}

	return len(invalidReasons) == 0, invalidReasons
}
