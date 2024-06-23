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
	inputs := prompt.StrBlocks{
		{
			Key:  "People Info",
			Vals: state.PeopleInfo,
		},
		{
			Key: "Travel Info",
			Val: state.TravelInfo,
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
