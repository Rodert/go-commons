// Package netutils 提供网络相关的工具函数
package netutils

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// IsValidIP 检查字符串是否为有效的IPv4或IPv6地址
// 返回布尔值表示是否有效
func IsValidIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

// IsValidIPv4 检查字符串是否为有效的IPv4地址
// 返回布尔值表示是否有效
func IsValidIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}

// IsValidIPv6 检查字符串是否为有效的IPv6地址
// 返回布尔值表示是否有效
func IsValidIPv6(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return ip.To4() == nil
}

// IsValidDomain 检查字符串是否为有效的域名
// 返回布尔值表示是否有效
func IsValidDomain(domain string) bool {
	// 特殊情况：localhost
	if domain == "localhost" {
		return true
	}

	// 支持IDN域名和更广泛的域名格式
	pattern := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9\-]{2,}$`
	valid, _ := regexp.MatchString(pattern, domain)
	return valid
}

// IsPortOpen 检查指定IP和端口是否开放
// 参数:
//   - host: 主机名或IP地址
//   - port: 端口号
//   - timeout: 连接超时时间
//
// 返回:
//   - bool: 端口是否开放
//   - error: 如果发生错误则返回错误信息
func IsPortOpen(host string, port int, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

// ExtractHostPort 从URL中提取主机名和端口号
// 参数:
//   - urlStr: URL字符串
//
// 返回:
//   - string: 主机名
//   - int: 端口号
//   - error: 如果解析失败则返回错误信息
func ExtractHostPort(urlStr string) (string, int, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", 0, err
	}

	host := parsedURL.Hostname()
	port := parsedURL.Port()

	// 如果端口未指定，根据协议设置默认端口
	portNum := 0
	if port != "" {
		portNum, err = strconv.Atoi(port)
		if err != nil {
			return "", 0, err
		}
	} else {
		switch parsedURL.Scheme {
		case "http":
			portNum = 80
		case "https":
			portNum = 443
		case "ftp":
			portNum = 21
		default:
			return "", 0, fmt.Errorf("未知的协议: %s", parsedURL.Scheme)
		}
	}

	return host, portNum, nil
}

// IsURLReachable 检查URL是否可访问
// 参数:
//   - urlStr: 要检查的URL
//   - timeout: 请求超时时间
//
// 返回:
//   - bool: URL是否可访问
//   - int: HTTP状态码 (如果请求成功)
//   - error: 如果发生错误则返回错误信息
func IsURLReachable(urlStr string, timeout time.Duration) (bool, int, error) {
	// 验证URL格式
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false, 0, fmt.Errorf("无效的URL格式: %v", err)
	}

	// 创建带超时的HTTP客户端
	client := &http.Client{
		Timeout: timeout,
	}

	// 发送HEAD请求检查URL可达性
	resp, err := client.Head(urlStr)
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()

	// 2xx 和 3xx 状态码表示URL可达
	return resp.StatusCode >= 200 && resp.StatusCode < 400, resp.StatusCode, nil
}
