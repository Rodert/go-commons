package main

import (
	"fmt"
	"time"

	"github.com/Rodert/go-commons/netutils"
)

func main() {
	// 演示IP地址验证
	ipv4 := "192.168.1.1"
	ipv6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	invalidIP := "256.256.256.256"

	fmt.Printf("%s 是有效的IP地址: %t\n", ipv4, netutils.IsValidIP(ipv4))
	fmt.Printf("%s 是有效的IPv4地址: %t\n", ipv4, netutils.IsValidIPv4(ipv4))
	fmt.Printf("%s 是有效的IPv6地址: %t\n", ipv4, netutils.IsValidIPv6(ipv4))

	fmt.Printf("%s 是有效的IP地址: %t\n", ipv6, netutils.IsValidIP(ipv6))
	fmt.Printf("%s 是有效的IPv4地址: %t\n", ipv6, netutils.IsValidIPv4(ipv6))
	fmt.Printf("%s 是有效的IPv6地址: %t\n", ipv6, netutils.IsValidIPv6(ipv6))

	fmt.Printf("%s 是有效的IP地址: %t\n", invalidIP, netutils.IsValidIP(invalidIP))

	// 演示域名验证
	domains := []string{"example.com", "sub.example.com", "localhost", "invalid-domain"}
	for _, domain := range domains {
		fmt.Printf("%s 是有效的域名: %t\n", domain, netutils.IsValidDomain(domain))
	}

	// 演示从URL提取主机名和端口
	urls := []string{
		"http://example.com",
		"https://example.com:8443/path",
		"ftp://files.example.com:2121",
	}

	for _, urlStr := range urls {
		host, port, err := netutils.ExtractHostPort(urlStr)
		if err != nil {
			fmt.Printf("从 %s 提取主机名和端口失败: %v\n", urlStr, err)
		} else {
			fmt.Printf("URL: %s -> 主机名: %s, 端口: %d\n", urlStr, host, port)
		}
	}

	// 演示检查端口是否开放
	// 注意：这些检查可能会因网络环境而有不同结果
	fmt.Println("\n检查端口是否开放 (可能需要几秒钟)...")

	// 检查Google的80端口（通常是开放的）
	isOpen, err := netutils.IsPortOpen("google.com", 80, 2*time.Second)
	if err != nil {
		fmt.Printf("检查google.com:80失败: %v\n", err)
	} else {
		fmt.Printf("google.com:80 是否开放: %t\n", isOpen)
	}

	// 检查一个不太可能开放的端口
	isOpen, err = netutils.IsPortOpen("example.com", 9999, 1*time.Second)
	if err != nil {
		fmt.Printf("检查example.com:9999失败: %v\n", err)
	} else {
		fmt.Printf("example.com:9999 是否开放: %t\n", isOpen)
	}

	// 演示检查URL是否可访问
	fmt.Println("\n检查URL是否可访问 (可能需要几秒钟)...")

	// 检查Google的可访问性
	isReachable, statusCode, err := netutils.IsURLReachable("https://www.google.com", 3*time.Second)
	if err != nil {
		fmt.Printf("检查https://www.google.com失败: %v\n", err)
	} else {
		fmt.Printf("https://www.google.com 是否可访问: %t, 状态码: %d\n", isReachable, statusCode)
	}

	// 检查一个不存在的URL
	isReachable, statusCode, err = netutils.IsURLReachable("https://this-domain-does-not-exist-123456789.com", 3*time.Second)
	if err != nil {
		fmt.Printf("检查不存在的域名失败: %v\n", err)
	} else {
		fmt.Printf("不存在的域名 是否可访问: %t, 状态码: %d\n", isReachable, statusCode)
	}
}
