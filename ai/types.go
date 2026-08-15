// Package ai defines provider-neutral primitives for LLM applications.
package ai

import "context"

// Role identifies the author of a chat message.
type Role string

const (
	RoleSystem    Role = "system"
	RoleDeveloper Role = "developer"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

// Message is a provider-neutral chat message.
type Message struct {
	Role       Role   `json:"role"`
	Content    string `json:"content,omitempty"`
	Name       string `json:"name,omitempty"`
	ToolCallID string `json:"tool_call_id,omitempty"`
}

// Tool defines a function a model may call. Parameters should be a JSON Schema object.
type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters,omitempty"`
}

// ToolCall is a function call requested by a model. Arguments contains JSON object text.
type ToolCall struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// Request is a provider-neutral chat completion request.
type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature *float64  `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Tools       []Tool    `json:"tools,omitempty"`
}

// Usage describes model token consumption.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Response is a completed model response.
type Response struct {
	ID           string     `json:"id"`
	Model        string     `json:"model"`
	Content      string     `json:"content"`
	FinishReason string     `json:"finish_reason,omitempty"`
	ToolCalls    []ToolCall `json:"tool_calls,omitempty"`
	Usage        Usage      `json:"usage"`
}

// Chunk is one event from a streaming response. Err terminates the stream.
type Chunk struct {
	ID           string     `json:"id,omitempty"`
	Model        string     `json:"model,omitempty"`
	Content      string     `json:"content,omitempty"`
	FinishReason string     `json:"finish_reason,omitempty"`
	ToolCalls    []ToolCall `json:"tool_calls,omitempty"`
	Usage        *Usage     `json:"usage,omitempty"`
	Err          error      `json:"-"`
}

// Client is implemented by AI providers and compatible API clients.
type Client interface {
	Chat(ctx context.Context, request *Request) (*Response, error)
	Stream(ctx context.Context, request *Request) (<-chan Chunk, error)
}
