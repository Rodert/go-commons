// Package openai implements the OpenAI Chat Completions protocol and compatible APIs.
package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Rodert/go-commons/ai"
)

const defaultBaseURL = "https://api.openai.com/v1"

// Config configures an OpenAI-compatible client.
type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Headers    http.Header
}

// Client sends requests using the OpenAI Chat Completions API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	headers    http.Header
}

// NewClient creates an OpenAI-compatible client.
func NewClient(config Config) *Client {
	baseURL := strings.TrimRight(config.BaseURL, "/")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	client := config.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	headers := make(http.Header, len(config.Headers))
	for key, values := range config.Headers {
		headers[key] = append([]string(nil), values...)
	}
	return &Client{apiKey: config.APIKey, baseURL: baseURL, httpClient: client, headers: headers}
}

// Chat sends a non-streaming chat completion request.
func (c *Client) Chat(ctx context.Context, request *ai.Request) (*ai.Response, error) {
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	body, err := json.Marshal(toWireRequest(request, false))
	if err != nil {
		return nil, fmt.Errorf("marshal chat request: %w", err)
	}
	response, err := c.do(ctx, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := checkResponse(response); err != nil {
		return nil, err
	}

	var wire wireResponse
	if err := json.NewDecoder(response.Body).Decode(&wire); err != nil {
		return nil, fmt.Errorf("decode chat response: %w", err)
	}
	return wire.toResponse(), nil
}

// Stream sends a streaming chat completion request and returns SSE chunks.
func (c *Client) Stream(ctx context.Context, request *ai.Request) (<-chan ai.Chunk, error) {
	if err := validateRequest(request); err != nil {
		return nil, err
	}
	body, err := json.Marshal(toWireRequest(request, true))
	if err != nil {
		return nil, fmt.Errorf("marshal stream request: %w", err)
	}
	response, err := c.do(ctx, body)
	if err != nil {
		return nil, err
	}
	if err := checkResponse(response); err != nil {
		response.Body.Close()
		return nil, err
	}

	chunks := make(chan ai.Chunk)
	go consumeSSE(ctx, response.Body, chunks)
	return chunks, nil
}

func validateRequest(request *ai.Request) error {
	if request == nil {
		return fmt.Errorf("chat request must not be nil")
	}
	if strings.TrimSpace(request.Model) == "" {
		return fmt.Errorf("chat request model must not be empty")
	}
	if len(request.Messages) == 0 {
		return fmt.Errorf("chat request must include at least one message")
	}
	return nil
}

func (c *Client) do(ctx context.Context, body []byte) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create chat request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json, text/event-stream")
	if c.apiKey != "" {
		request.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	for key, values := range c.headers {
		request.Header.Del(key)
		request.Header[key] = append([]string(nil), values...)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send chat request: %w", err)
	}
	return response, nil
}

func checkResponse(response *http.Response) error {
	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		return nil
	}
	body, _ := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	return fmt.Errorf("ai provider returned HTTP %d: %s", response.StatusCode, strings.TrimSpace(string(body)))
}

func consumeSSE(ctx context.Context, body io.ReadCloser, chunks chan<- ai.Chunk) {
	defer body.Close()
	defer close(chunks)
	scanner := bufio.NewScanner(body)
	scanner.Buffer(make([]byte, 1024), 1<<20)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if data == "[DONE]" {
			return
		}
		var response wireResponse
		if err := json.Unmarshal([]byte(data), &response); err != nil {
			sendChunk(ctx, chunks, ai.Chunk{Err: fmt.Errorf("decode stream chunk: %w", err)})
			return
		}
		chunk := response.toChunk()
		if !sendChunk(ctx, chunks, chunk) {
			return
		}
	}
	if err := scanner.Err(); err != nil {
		sendChunk(ctx, chunks, ai.Chunk{Err: fmt.Errorf("read stream: %w", err)})
	}
}

func sendChunk(ctx context.Context, chunks chan<- ai.Chunk, chunk ai.Chunk) bool {
	select {
	case <-ctx.Done():
		return false
	case chunks <- chunk:
		return true
	}
}

type wireRequest struct {
	Model       string       `json:"model"`
	Messages    []ai.Message `json:"messages"`
	Temperature *float64     `json:"temperature,omitempty"`
	MaxTokens   int          `json:"max_tokens,omitempty"`
	Tools       []wireTool   `json:"tools,omitempty"`
	Stream      bool         `json:"stream,omitempty"`
}

type wireTool struct {
	Type     string  `json:"type"`
	Function ai.Tool `json:"function"`
}

func toWireRequest(request *ai.Request, stream bool) wireRequest {
	tools := make([]wireTool, 0, len(request.Tools))
	for _, tool := range request.Tools {
		tools = append(tools, wireTool{Type: "function", Function: tool})
	}
	return wireRequest{Model: request.Model, Messages: request.Messages, Temperature: request.Temperature, MaxTokens: request.MaxTokens, Tools: tools, Stream: stream}
}

type wireResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Content   string         `json:"content"`
			ToolCalls []wireToolCall `json:"tool_calls"`
		} `json:"message"`
		Delta struct {
			Content   string         `json:"content"`
			ToolCalls []wireToolCall `json:"tool_calls"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage ai.Usage `json:"usage"`
}

type wireToolCall struct {
	ID       string `json:"id"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

func toToolCalls(calls []wireToolCall) []ai.ToolCall {
	result := make([]ai.ToolCall, 0, len(calls))
	for _, call := range calls {
		result = append(result, ai.ToolCall{ID: call.ID, Name: call.Function.Name, Arguments: call.Function.Arguments})
	}
	return result
}

func (response wireResponse) toResponse() *ai.Response {
	result := &ai.Response{ID: response.ID, Model: response.Model, Usage: response.Usage}
	if len(response.Choices) > 0 {
		choice := response.Choices[0]
		result.Content = choice.Message.Content
		result.ToolCalls = toToolCalls(choice.Message.ToolCalls)
		result.FinishReason = choice.FinishReason
	}
	return result
}

func (response wireResponse) toChunk() ai.Chunk {
	chunk := ai.Chunk{ID: response.ID, Model: response.Model}
	if len(response.Choices) > 0 {
		choice := response.Choices[0]
		chunk.Content = choice.Delta.Content
		chunk.ToolCalls = toToolCalls(choice.Delta.ToolCalls)
		chunk.FinishReason = choice.FinishReason
	}
	if response.Usage.TotalTokens != 0 {
		usage := response.Usage
		chunk.Usage = &usage
	}
	return chunk
}
