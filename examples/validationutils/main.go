package main

import (
	"fmt"

	"github.com/Rodert/go-commons/validationutils"
)

func main() {
	// 演示邮箱验证
	emails := []string{
		"user@example.com",
		"admin@subdomain.example.co.uk",
		"invalid-email",
		"missing@domain",
	}

	fmt.Println("邮箱验证:")
	for _, email := range emails {
		fmt.Printf("%s 是有效的邮箱: %t\n", email, validationutils.IsEmail(email))
	}

	// 演示中国手机号验证
	mobiles := []string{
		"13812345678",
		"15912345678",
		"19912345678",
		"1381234567",  // 少一位
		"23812345678", // 不是1开头
	}

	fmt.Println("\n中国手机号验证:")
	for _, mobile := range mobiles {
		fmt.Printf("%s 是有效的中国手机号: %t\n", mobile, validationutils.IsCNMobile(mobile))
	}

	// 演示URL验证
	urls := []string{
		"http://example.com",
		"https://subdomain.example.com/path?query=value",
		"ftp://files.example.org",
		"invalid-url",
		"http://", // 不完整
	}

	fmt.Println("\nURL验证:")
	for _, url := range urls {
		fmt.Printf("%s 是有效的URL: %t\n", url, validationutils.IsURL(url))
	}

	// 演示IPv4验证
	ips := []string{
		"192.168.1.1",
		"10.0.0.1",
		"255.255.255.255",
		"256.0.0.1", // 超出范围
		"192.168.1", // 不完整
	}

	fmt.Println("\nIPv4验证:")
	for _, ip := range ips {
		fmt.Printf("%s 是有效的IPv4地址: %t\n", ip, validationutils.IsIPv4(ip))
	}

	// 演示中国身份证号验证
	idCards := []string{
		"110101199001011234", // 示例18位号码
		"11010119900101123X", // 示例18位号码，X结尾
		"110101900101123",    // 示例15位号码
		"1101011990010",      // 不完整
	}

	fmt.Println("\n中国身份证号验证:")
	for _, idCard := range idCards {
		fmt.Printf("%s 是有效的中国身份证号: %t\n", idCard, validationutils.IsCNIDCard(idCard))
	}

	// 演示密码强度验证
	passwords := []string{
		"abc123",                  // 简单密码
		"Password123",             // 中等强度
		"P@ssw0rd123!",            // 强密码
		"SuperStr0ng!P@ssw0rd123", // 非常强的密码
	}

	fmt.Println("\n密码强度验证:")
	for _, password := range passwords {
		result := validationutils.CheckPasswordStrength(password)

		// 将密码强度级别转换为文字描述
		var levelDesc string
		switch result.Level {
		case validationutils.WeakPassword:
			levelDesc = "弱"
		case validationutils.MediumPassword:
			levelDesc = "中等"
		case validationutils.StrongPassword:
			levelDesc = "强"
		case validationutils.VeryStrongPassword:
			levelDesc = "非常强"
		}

		fmt.Printf("%s 的密码强度: %s (得分: %d)\n", password, levelDesc, result.Score)

		// 显示改进建议
		if len(result.Suggestions) > 0 {
			fmt.Println("  改进建议:")
			for _, suggestion := range result.Suggestions {
				fmt.Printf("  - %s\n", suggestion)
			}
		}
	}
}
