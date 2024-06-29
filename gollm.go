package gollm

import (
	"fmt"
	"strings"

	"github.com/lordtatty/gollm/llm"
	"github.com/lordtatty/gollm/prompt"
)

type Msg struct {
	Text        string
	FixedKVs    prompt.StrBlocks
	VariableKVs VariableKVs
}

func (m Msg) String(kv map[string]string) (string, error) {
	parts := []string{}
	// Fixed Inputs
	fixedInputs := m.FixedKVs.String()
	if fixedInputs != "" {
		parts = append(parts, fixedInputs)
	}
	// Variable Inputs
	inputBlocks, err := m.VariableKVs.FromVals(kv)
	if err != nil {
		return "", fmt.Errorf("missing values for input keys: %w", err)
	}
	if inputBlocks != "" {
		parts = append(parts, inputBlocks)
	}
	// Prompt Text
	if m.Text != "" {
		parts = append(parts, m.Text)
	}
	result := strings.Join(parts, "\n\n")
	return result, nil
}

type VariableKV struct {
	Key   string
	Label string
}

type VariableKVs []VariableKV

func (i *VariableKVs) FromVals(blockOutputs map[string]string) (string, error) {
	var blocks prompt.StrBlocks
	for _, b := range *i {
		if _, ok := blockOutputs[b.Key]; !ok {
			return "", fmt.Errorf("block %s not found in blockOutputs", b.Key)
		}
		blocks = append(blocks, prompt.StrBlock{
			Key:  b.Label,
			Val:  blockOutputs[b.Key],
			Vals: nil,
		})
	}
	return blocks.String(), nil
}

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

type Sequential struct {
	Blocks []LLMBlock
}

func (s *Sequential) Run() error {
	blockOutputs := make(map[string]string)
	for i, block := range s.Blocks {
		resp, err := block.Run(blockOutputs)
		if err != nil {
			return fmt.Errorf("failed to run block %d: %w", i, err)
		}
		blockOutputs[block.Name] = resp.Text
		fmt.Println(resp.Text)
	}
	return nil
}
