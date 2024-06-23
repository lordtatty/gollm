package llm

import (
	"fmt"
	"os"
)

func boolptr(b bool) *bool {
	return &b
}

type ChatOpts struct {
	StreamCh chan string
}
type ChatResp struct {
	Text string
}

type LLM interface {
	Chat(systemMsg, userMsg string, streamCh chan string) (*ChatResp, error)
}

type LLMInstance struct {
	Instance LLM
}

func FromFlag(llmFlag string) (*LLMInstance, error) {
	llm := &LLMInstance{}
	switch llmFlag {
	case "ollama":
		llm.Instance = &Ollama{}
	case "groq":
		g := &Groq{
			APIKey: os.Getenv("GROQ_API_KEY"),
			// Model:  llm.GROQ_MODEL_LLAMA3_70B_8192,
		}
		llm.Instance = g
	default:
		return nil, fmt.Errorf("unknown llm: %s", llmFlag)
	}
	return llm, nil
}
