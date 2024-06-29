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
