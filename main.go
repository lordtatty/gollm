package main

import (
	"fmt"
	"log"

	"github.com/lordtatty/dstest/llm"
	"github.com/lordtatty/dstest/prompt"
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
	err = organiserAgent(l.Instance)
	if err != nil {
		return fmt.Errorf("failed to planner: %w", err)
	}
	return nil
}

func organiserAgent(l llm.LLM) error {
	bkdwn, err := timeBreakdownAgent(l)
	if err != nil {
		return fmt.Errorf("failed to get time breakdown: %w", err)
	}
	inputs := prompt.StrBlocks{
		{
			Key:  "People Info",
			Vals: peopleInfoAgent(),
		},
		{
			Key: "Travel Info",
			Val: travelInfoAgent(),
		},
		{
			Key: "Time Breakdown",
			Val: bkdwn,
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

func peopleInfoAgent() []string {
	return []string{"Name: John Doe\nAge: 30\nOccupation: Software Engineer", "Name: Jane Doe\nAge: 28\nOccupation: Data Scientist"}
}

func travelInfoAgent() string {
	return "Destination: Paris\nDuration: 3 days\nBudget: $1000"
}

func timeBreakdownAgent(l llm.LLM) (string, error) {
	plans := "Day 1: Visit Eiffel Tower for John\nDay 2: Visit Louvre Museum, but Jane may not want to go\nDay 3: Visit Notre Dame Cathedral"
	inputs := prompt.StrBlocks{
		{
			Key: "Plan Info",
			Val: plans,
		},
	}
	resp, err := l.Chat(
		"You are an expert in time management. You will be given a plan for a holiday. Now you need to break down the time you have into days and allocate activities to each day. Make sure you take into account the preferences of the people involved",
		inputs.Build(),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to chat: %w", err)
	}
	return resp.Text, nil
}
