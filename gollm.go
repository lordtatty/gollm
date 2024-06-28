package gollm

import (
	"fmt"

	"github.com/lordtatty/gollm/llm"
	"github.com/lordtatty/gollm/prompt"
)

type Msg struct {
	Text string
}

func (m Msg) String() string {
	return m.Text
}

type ExpectKey struct {
	Key   string
	Label string
}

type ExpectKeys []ExpectKey

func (i *ExpectKeys) FromVals(name string, blockOutputs map[string]string) (string, error) {
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
	String() string
}

type LLMBlock struct {
	Name       string
	LLM        llm.LLM
	SystemMsg  string
	UserMsg    UserMsg
	ExpectKeys ExpectKeys
}

type BlockResult struct {
	Name string
	Text string
}

func (b *LLMBlock) Run(inputs map[string]string) (*BlockResult, error) {
	inputBlocks, err := b.ExpectKeys.FromVals(b.Name, inputs)
	if err != nil {
		return nil, fmt.Errorf("missing values for input keys: %w", err)
	}
	fulInput := inputBlocks + "\n\n" + b.UserMsg.String()
	resp, err := b.LLM.Chat(b.SystemMsg, fulInput, nil)
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

// func (s *Sequential) Run() error {
// 	blockOutputs := make(map[string]string)
// 	for i, block := range s.Blocks {
// 		prevInputs, err := block.IncludeBlocks.AsStrBlocks(block.Name, blockOutputs)
// 		if err != nil {
// 			return fmt.Errorf("failed to get include blocks for block %d: %w", i, err)
// 		}
// 		fulInput := prevInputs.String() + "\n\n" + block.UserMsg.String()
// 		resp, err := block.LLM.Chat(block.SystemMsg, fulInput, nil)
// 		if err != nil {
// 			return fmt.Errorf("failed to chat block %d: %w", i, err)
// 		}
// 		blockOutputs[block.Name] = resp.Text
// 		fmt.Println(resp.Text)
// 	}
// 	return nil
// }
