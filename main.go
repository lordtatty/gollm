package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lordtatty/dstest/llm"
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
	err = planner(l.Instance)
	if err != nil {
		return fmt.Errorf("failed to planner: %w", err)
	}
	return nil
}

type PlannerState struct {
	PeopleInfo []string
	TravelInfo string
}

func planner(l llm.LLM) error {
	state := &PlannerState{
		PeopleInfo: []string{"Name: John Doe\nAge: 30\nOccupation: Software Engineer", "Name: Jane Doe\nAge: 28\nOccupation: Data Scientist"},
		TravelInfo: "Destination: Paris\nDuration: 3 days\nBudget: $1000",
	}
	inputs := StrBlocks{
		Blocks: []StrBlock{
			{
				Key:  "Person Info",
				Vals: state.PeopleInfo,
			},
			{
				Key: "Travel Info",
				Val: "Destination: Paris\nDuration: 3 days\nBudget: $1000",
			},
		},
	}
	resp, err := l.Chat(
		"You are an expert organiser. You will been given plans for a holiday. Now you need to work out every task that needs to be done to make the holiday happen, and put them in order. Include things to prepare, things to take, research required, etc.",
		inputs.Build(),
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to chat: %w", err)
	}
	fmt.Println(resp.Text)
	return nil
}

type StrBlocks struct {
	Blocks []StrBlock
}

func (s *StrBlocks) Build() string {
	var sb strings.Builder
	for _, block := range s.Blocks {
		sb.WriteString(block.Build(block.Key, block.Val))
		sb.WriteString("\n")
	}
	return sb.String()
}

func MultiStrBlock(key string, vals []string) []StrBlock {
	b := []StrBlock{}
	for _, v := range vals {
		b = append(b, StrBlock{
			Key: key,
			Val: v,
		})
	}
	return b
}

type StrBlock struct {
	Key  string
	Val  string
	Vals []string
}

func (s *StrBlock) Build(key, value string) string {
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

func (s *StrBlock) buildOne(key, value string) string {
	key = strings.ToUpper(key)
	return fmt.Sprintf("### %s START###\n%s\n### %s END###\n", key, value, key)
}
