package main

import (
	"fmt"
	"log"

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
	resp, err := l.Instance.Chat("You are a 1980s skteboarding expert", "Hello, world!", nil)
	if err != nil {
		return fmt.Errorf("failed to chat: %w", err)
	}
	fmt.Println(resp)
	return nil
}
