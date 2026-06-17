package llm

import (
	"context"
)

type Client interface {
	Chat(ctx context.Context, prompt string) (string, error)
}
