package ai

import "context"

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type Message struct {
	Role    Role
	Content string
}

type Response struct {
	Content string
	Model   string
}

type Provider interface {
	Name() string
	Chat(ctx context.Context, messages []Message) (Response, error)
	Stream(ctx context.Context, messages []Message) (<-chan string, error)
}
