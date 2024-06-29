package prompt

import (
	"fmt"
	"strings"
)

type Msg struct {
	Text           string
	FixedBlocks    FixedBlocks
	VariableBlocks VariableBlocks
}

func (m Msg) String(kv map[string]string) (string, error) {
	parts := []string{}
	// Fixed Inputs
	fixedInputs := m.FixedBlocks.String()
	if fixedInputs != "" {
		parts = append(parts, fixedInputs)
	}
	// Variable Inputs
	inputBlocks, err := m.VariableBlocks.FromVals(kv)
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

type VariableBlock struct {
	Key   string
	Label string
}

type VariableBlocks []VariableBlock

func (i *VariableBlocks) FromVals(blockOutputs map[string]string) (string, error) {
	var blocks FixedBlocks
	for _, b := range *i {
		if _, ok := blockOutputs[b.Key]; !ok {
			return "", fmt.Errorf("block %s not found in blockOutputs", b.Key)
		}
		blocks = append(blocks, FixedBlock{
			Key:  b.Label,
			Val:  blockOutputs[b.Key],
			Vals: nil,
		})
	}
	return blocks.String(), nil
}

type FixedBlocks []FixedBlock

func (s FixedBlocks) String() string {
	var sb strings.Builder
	for _, block := range s {
		sb.WriteString(block.String(block.Key, block.Val))
	}
	return sb.String()
}

type FixedBlock struct {
	Key  string
	Val  string
	Vals []string
}

func (s *FixedBlock) String(key, value string) string {
	vals := s.Vals
	if s.Val != "" {
		vals = append(s.Vals, s.Val)
	}
	var sb strings.Builder
	for _, v := range vals {
		sb.WriteString(s.buildOne(key, v))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s *FixedBlock) buildOne(key, value string) string {
	key = strings.ToUpper(key)
	return fmt.Sprintf("### %s START###\n%s\n### %s END###\n", key, value, key)
}
