package gollm

import (
	"fmt"
)

type RunnableBlock interface {
	Name() string
	Run(map[string]string) (*BlockResult, error)
}

type RunnableBlocks []RunnableBlock

func (r *RunnableBlocks) Names() []string {
	names := []string{}
	for _, block := range *r {
		names = append(names, block.Name())
	}
	return names
}

func (r *RunnableBlocks) Valid() error {
	// ensure all blocks have unique names
	names := map[string]struct{}{}
	for _, block := range *r {
		name := block.Name()
		if _, ok := names[name]; ok {
			return fmt.Errorf("block name %s is not unique", name)
		}
		names[name] = struct{}{}
	}
	return nil
}

type FlowResp struct {
	Run     []string
	NotRun  []string
	Outputs map[string]BlockResult
}

func (r *FlowResp) markRun(name string) {
	r.Run = append(r.Run, name)
	// remove from NotRun
	for i, n := range r.NotRun {
		if n == name {
			r.NotRun = append(r.NotRun[:i], r.NotRun[i+1:]...)
			break
		}
	}
}

func (r *FlowResp) setOutput(name string, output BlockResult) {
	r.Outputs[name] = output
}

func (r *FlowResp) OutputStrs() map[string]string {
	result := make(map[string]string)
	for name, output := range r.Outputs {
		result[name] = output.Text
	}
	return result
}

func Flow(blocks RunnableBlocks) (*FlowResp, error) {
	if err := blocks.Valid(); err != nil {
		return nil, fmt.Errorf("invalid blocks: %w", err)
	}
	result := &FlowResp{
		NotRun:  blocks.Names(),
		Outputs: make(map[string]BlockResult),
	}
	for i, block := range blocks {
		result.markRun(block.Name())
		resp, err := block.Run(result.OutputStrs())
		if err != nil {
			return nil, fmt.Errorf("failed to run block %d, %s: %w", i, block.Name(), err)
		}
		result.setOutput(block.Name(), *resp)
	}
	return result, nil
}
