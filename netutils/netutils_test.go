package netutils_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Rodert/go-commons/netutils"
)

func TestHTTPClientResponseBodyLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("12345"))
	}))
	defer server.Close()

	client := netutils.NewHTTPClientWithMaxResponseBody(time.Second, 4)
	if _, err := client.Get(server.URL, nil); err == nil {
		t.Fatal("Get() error = nil, want response size error")
	}
}

func TestHTTPClientGetWithContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	client := netutils.NewHTTPClient(time.Second)
	if _, err := client.GetWithContext(ctx, server.URL, nil); err == nil {
		t.Fatal("GetWithContext() error = nil, want context cancellation error")
	}
}

func TestHTTPClientPostAndJSONHelpers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusTeapot)
			return
		}
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.Header.Get("X-Request-ID") != "test" {
			t.Errorf("X-Request-ID = %q", r.Header.Get("X-Request-ID"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	client := netutils.NewHTTPClient(time.Second)
	response, err := client.PostJSON(server.URL, map[string]string{"X-Request-ID": "test"}, map[string]string{"name": "commons"})
	if err != nil {
		t.Fatalf("PostJSON() error = %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", response.StatusCode)
	}

	var decoded struct {
		OK bool `json:"ok"`
	}
	if err := client.GetJSON(server.URL, nil, &decoded); err == nil {
		t.Fatal("GetJSON() error = nil, want non-2xx error")
	}
}

func TestIsValidIPv4(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"255.255.255.255", true},
		{"0.0.0.0", true},
		{"256.0.0.1", false},
		{"192.168.1", false},
		{"192.168.1.1.1", false},
		{"192.168.1.a", false},
		{"", false},
	}

	for _, test := range tests {
		result := netutils.IsValidIPv4(test.ip)
		if result != test.expected {
			t.Errorf("IsValidIPv4(%q) = %v; want %v", test.ip, result, test.expected)
		}
	}
}

func TestIsValidDomain(t *testing.T) {
	tests := []struct {
		domain   string
		expected bool
	}{
		{"example.com", true},
		{"sub.example.com", true},
		{"sub-domain.example.co.uk", true},
		{"xn--p1ai.xn--p1ai", true}, // IDN domains
		{"localhost", true},
		{"example", false},      // Missing TLD
		{"example..com", false}, // Double dot
		{"example-.com", false}, // Hyphen at end of label
		{"-example.com", false}, // Hyphen at start of label
		{"exam_ple.com", false}, // Underscore in domain
		{"example.com.", false}, // Trailing dot
		{"", false},
	}

	for _, test := range tests {
		result := netutils.IsValidDomain(test.domain)
		if result != test.expected {
			t.Errorf("IsValidDomain(%q) = %v; want %v", test.domain, result, test.expected)
		}
	}
}

func TestIsPortOpen(t *testing.T) {
	// 启动一个测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 从服务器地址中提取主机和端口
	host, port, err := netutils.ExtractHostPort(server.URL)
	if err != nil {
		t.Fatalf("Failed to extract host and port from URL: %v", err)
	}

	// 测试端口是否开放
	isOpen, _ := netutils.IsPortOpen(host, port, 2*time.Second)
	if !isOpen {
		t.Errorf("IsPortOpen(%s, %d) = false; want true", host, port)
	}

	// 测试一个不太可能开放的端口
	isOpen, _ = netutils.IsPortOpen(host, 54321, 1*time.Second)
	if isOpen {
		t.Errorf("IsPortOpen(%s, %d) = true; want false", host, 54321)
	}
}

func TestIsURLReachable(t *testing.T) {
	// 启动一个测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 测试可达的URL
	isReachable, _, _ := netutils.IsURLReachable(server.URL, 2*time.Second)
	if !isReachable {
		t.Errorf("IsURLReachable(%s) = false; want true", server.URL)
	}

	// 测试不可达的URL
	isReachable, _, _ = netutils.IsURLReachable("http://localhost:54321/nonexistent", 1*time.Second)
	if isReachable {
		t.Errorf("IsURLReachable(%s) = true; want false", "http://localhost:54321/nonexistent")
	}
}

func TestExtractHostPort(t *testing.T) {
	tests := []struct {
		url          string
		expectedHost string
		expectedPort int
		expectError  bool
	}{
		{"http://example.com", "example.com", 80, false},
		{"https://example.com", "example.com", 443, false},
		{"http://example.com:8080", "example.com", 8080, false},
		{"https://example.com:8443", "example.com", 8443, false},
		{"ftp://example.com", "example.com", 21, false},
		{"not-a-url", "", 0, true},
		{"", "", 0, true},
	}

	for _, test := range tests {
		host, port, err := netutils.ExtractHostPort(test.url)

		if test.expectError && err == nil {
			t.Errorf("ExtractHostPort(%q) expected error, got nil", test.url)
			continue
		}

		if !test.expectError && err != nil {
			t.Errorf("ExtractHostPort(%q) unexpected error: %v", test.url, err)
			continue
		}

		if host != test.expectedHost || port != test.expectedPort {
			t.Errorf("ExtractHostPort(%q) = (%q, %d); want (%q, %d)",
				test.url, host, port, test.expectedHost, test.expectedPort)
		}
	}
}
