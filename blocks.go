package gollm

import (
	"fmt"

	"github.com/lordtatty/gollm/llm"
)

type UserMsg interface {
	String(kv map[string]string) (string, error)
}

type LLMBlock struct {
	Name      string
	LLM       llm.LLM
	SystemMsg string
	UserMsg   UserMsg
}

type BlockResult struct {
	Name string
	Text string
}

func (b *LLMBlock) Run(inputs map[string]string) (*BlockResult, error) {
	userMsg, err := b.UserMsg.String(inputs)
	if err != nil {
		return nil, fmt.Errorf("failed to get user message: %w", err)
	}
	resp, err := b.LLM.Chat(b.SystemMsg, userMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to chat: %w", err)
	}
	return &BlockResult{
		Name: b.Name,
		Text: resp.Text,
	}, nil
}
