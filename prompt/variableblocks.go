package prompt

import (
	"fmt"
)

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
