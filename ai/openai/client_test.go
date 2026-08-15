package openai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Rodert/go-commons/ai"
)

func TestChatTranslatesRequestAndResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/v1/chat/completions" {
			t.Errorf("path = %s", request.URL.Path)
		}
		if request.Header.Get("Authorization") != "Bearer token" {
			t.Errorf("authorization = %q", request.Header.Get("Authorization"))
		}
		var body map[string]any
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if body["model"] != "test-model" || body["stream"] != nil {
			t.Errorf("request body = %#v", body)
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"id":"chat-1","model":"test-model","choices":[{"message":{"content":"hello","tool_calls":[{"id":"call-1","function":{"name":"weather","arguments":"{\"city\":\"Shanghai\"}"}}]},"finish_reason":"tool_calls"}],"usage":{"prompt_tokens":4,"completion_tokens":2,"total_tokens":6}}`))
	}))
	defer server.Close()

	client := NewClient(Config{APIKey: "token", BaseURL: server.URL + "/v1"})
	response, err := client.Chat(context.Background(), &ai.Request{
		Model:    "test-model",
		Messages: []ai.Message{{Role: ai.RoleUser, Content: "hello"}},
		Tools:    []ai.Tool{{Name: "weather", Parameters: map[string]any{"type": "object"}}},
	})
	if err != nil {
		t.Fatalf("Chat() error = %v", err)
	}
	if response.Content != "hello" || response.Usage.TotalTokens != 6 || response.FinishReason != "tool_calls" {
		t.Errorf("response = %#v", response)
	}
	if len(response.ToolCalls) != 1 || response.ToolCalls[0].Name != "weather" || response.ToolCalls[0].Arguments != `{"city":"Shanghai"}` {
		t.Errorf("tool calls = %#v", response.ToolCalls)
	}
}

func TestStreamParsesSSE(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/event-stream")
		_, _ = writer.Write([]byte("data: {\"id\":\"chat-1\",\"model\":\"test\",\"choices\":[{\"delta\":{\"content\":\"Hel\"}}]}\n\n"))
		_, _ = writer.Write([]byte("data: {\"id\":\"chat-1\",\"model\":\"test\",\"choices\":[{\"delta\":{\"content\":\"lo\"},\"finish_reason\":\"stop\"}],\"usage\":{\"prompt_tokens\":2,\"completion_tokens\":2,\"total_tokens\":4}}\n\n"))
		_, _ = writer.Write([]byte("data: [DONE]\n\n"))
	}))
	defer server.Close()

	client := NewClient(Config{BaseURL: server.URL})
	chunks, err := client.Stream(context.Background(), &ai.Request{Model: "test", Messages: []ai.Message{{Role: ai.RoleUser, Content: "hi"}}})
	if err != nil {
		t.Fatalf("Stream() error = %v", err)
	}

	var content strings.Builder
	var usage *ai.Usage
	for chunk := range chunks {
		if chunk.Err != nil {
			t.Fatalf("stream error = %v", chunk.Err)
		}
		content.WriteString(chunk.Content)
		if chunk.Usage != nil {
			usage = chunk.Usage
		}
	}
	if content.String() != "Hello" || usage == nil || usage.TotalTokens != 4 {
		t.Errorf("content = %q, usage = %#v", content.String(), usage)
	}
}

func TestChatValidatesRequestAndProviderError(t *testing.T) {
	client := NewClient(Config{})
	if _, err := client.Chat(context.Background(), nil); err == nil {
		t.Fatal("Chat(nil) error = nil")
	}
	if _, err := client.Chat(context.Background(), &ai.Request{Model: "model"}); err == nil {
		t.Fatal("Chat() without messages error = nil")
	}

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusTooManyRequests)
		_, _ = writer.Write([]byte(`{"error":"rate limit"}`))
	}))
	defer server.Close()
	client = NewClient(Config{BaseURL: server.URL, HTTPClient: &http.Client{Timeout: time.Second}})
	if _, err := client.Chat(context.Background(), &ai.Request{Model: "model", Messages: []ai.Message{{Role: ai.RoleUser, Content: "hi"}}}); err == nil {
		t.Fatal("Chat() provider error = nil")
	}
}
