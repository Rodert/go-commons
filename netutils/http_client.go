package netutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient 是一个简化的HTTP客户端
type HTTPClient struct {
	client *http.Client
}

// HTTPResponse 表示HTTP响应
type HTTPResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// NewHTTPClient 创建一个新的HTTP客户端
// 参数:
//   - timeout: 请求超时时间
//
// 返回:
//   - *HTTPClient: 新创建的HTTP客户端
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get 发送GET请求
// 参数:
//   - url: 请求URL
//   - headers: 请求头
//
// 返回:
//   - *HTTPResponse: HTTP响应
//   - error: 如果发生错误则返回错误信息
func (c *HTTPClient) Get(url string, headers map[string]string) (*HTTPResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return c.doRequest(req)
}

// Post 发送POST请求
// 参数:
//   - url: 请求URL
//   - headers: 请求头
//   - body: 请求体
//
// 返回:
//   - *HTTPResponse: HTTP响应
//   - error: 如果发生错误则返回错误信息
func (c *HTTPClient) Post(url string, headers map[string]string, body []byte) (*HTTPResponse, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return c.doRequest(req)
}

// PostJSON 发送JSON格式的POST请求
// 参数:
//   - url: 请求URL
//   - headers: 请求头
//   - data: 要发送的数据对象，将被转换为JSON
//
// 返回:
//   - *HTTPResponse: HTTP响应
//   - error: 如果发生错误则返回错误信息
func (c *HTTPClient) PostJSON(url string, headers map[string]string, data interface{}) (*HTTPResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON编码失败: %v", err)
	}

	// 确保设置了正确的Content-Type
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"

	return c.Post(url, headers, jsonData)
}

// doRequest 执行HTTP请求
// 参数:
//   - req: HTTP请求
//
// 返回:
//   - *HTTPResponse: HTTP响应
//   - error: 如果发生错误则返回错误信息
func (c *HTTPClient) doRequest(req *http.Request) (*HTTPResponse, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}

// GetJSON 发送GET请求并将响应解析为JSON
// 参数:
//   - url: 请求URL
//   - headers: 请求头
//   - result: 用于存储解析结果的对象指针
//
// 返回:
//   - error: 如果发生错误则返回错误信息
func (c *HTTPClient) GetJSON(url string, headers map[string]string, result interface{}) error {
	resp, err := c.Get(url, headers)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}

	return json.Unmarshal(resp.Body, result)
}
