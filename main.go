package main

import (
	"fmt"
	"log"

	"github.com/lordtatty/gollm/llm"
	"github.com/lordtatty/gollm/prompt"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l, err := llm.FromFlag("groq")
	if err != nil {
		return fmt.Errorf("failed to get LLM: %w", err)
	}
	// resp, err := organiserAgent(l.Instance)
	// fmt.Println(resp)
	// if err != nil {
	// 	return fmt.Errorf("failed to planner: %w", err)
	// }

	s, err := sequential(l.Instance)
	if err != nil {
		return fmt.Errorf("failed to get sequential: %w", err)
	}
	err = s.Run()
	if err != nil {
		return fmt.Errorf("failed to run sequential: %w", err)
	}

	return nil
}

type IncludeBlock struct {
	BlockName string
	Label     string
}

type IncludeBlocks []IncludeBlock

func (i *IncludeBlocks) AsStrBlocks(name string, blockOutputs map[string]string) (*prompt.StrBlocks, error) {
	var blocks prompt.StrBlocks
	for _, b := range *i {
		if _, ok := blockOutputs[b.BlockName]; !ok {
			return nil, fmt.Errorf("block %s not found in blockOutputs", b.BlockName)
		}
		blocks = append(blocks, prompt.StrBlock{
			Key:  b.Label,
			Val:  blockOutputs[b.BlockName],
			Vals: nil,
		})
	}
	return &blocks, nil
}

type LLMBlock struct {
	Name          string
	LLM           llm.LLM
	System        string
	Inputs        prompt.StrBlocks
	IncludeBlocks IncludeBlocks
}

type Sequential struct {
	Blocks []LLMBlock
}

func (s *Sequential) Run() error {
	blockOutputs := make(map[string]string)
	for i, block := range s.Blocks {
		prevInputs, err := block.IncludeBlocks.AsStrBlocks(block.Name, blockOutputs)
		if err != nil {
			return fmt.Errorf("failed to get include blocks for block %d: %w", i, err)
		}
		fullBlocks := append(*prevInputs, block.Inputs...)
		fulInput := fullBlocks.Build()
		resp, err := block.LLM.Chat(block.System, fulInput, nil)
		if err != nil {
			return fmt.Errorf("failed to chat block %d: %w", i, err)
		}
		blockOutputs[block.Name] = resp.Text
		fmt.Println(resp.Text)
	}
	return nil
}

func sequential(l llm.LLM) (*Sequential, error) {
	s := &Sequential{
		Blocks: []LLMBlock{
			{
				Name:   "planner",
				LLM:    l,
				System: "You are an expert in planning holidays. Your task is to plan a group holiday. You will be given information on the people involved and what each of them wants to do.  Take into account their preferences and plan a holiday that they will all enjoy.",
				Inputs: prompt.StrBlocks{
					{
						Key:  "People Info",
						Vals: peopleInfoAgent(),
					},
					{
						Key: "Travel Info",
						Val: travelInfoAgent(),
					},
				},
			},
			{
				Name:   "timeBkdown",
				LLM:    l,
				System: "You are an expert in time management. You will be given a plan for a holiday. Now you need to break down the time you have into days and allocate activities to each day. Make sure you take into account the preferences of the people involved",
				IncludeBlocks: IncludeBlocks{
					{
						BlockName: "planner",
						Label:     "Current Plan",
					},
				},
			},
			{
				Name:   "organiser",
				LLM:    l,
				System: "You are an expert organiser. You will been given plans for a holiday. Now you need to work out every task that needs to be done to make the holiday happen, and put them in order. Include things to prepare, things to take, research required, etc.",
				IncludeBlocks: IncludeBlocks{
					{
						BlockName: "planner",
						Label:     "Current Plan",
					},
					{
						BlockName: "timeBkdown",
						Label:     "Time Breakdown",
					},
				},
			},
		},
	}
	return s, nil
}

func peopleInfoAgent() []string {
	return []string{"Name: John Doe\nAge: 30\nOccupation: Software Engineer\nLikes: historical places", "Name: Jane Doe\nAge: 28\nOccupation: Data Scientist\nLikes: architecture"}
}

func travelInfoAgent() string {
	return "Destination: Paris\nDuration: 3 days\nBudget: $1000"
}
