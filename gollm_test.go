package gollm_test

import (
	"testing"

	"github.com/lordtatty/gollm"
	"github.com/lordtatty/gollm/llm"
	"github.com/lordtatty/gollm/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLLMBlock_Run(t *testing.T) {
	assert := assert.New(t)

	chatResp := &llm.ChatResp{
		Text: "response message",
	}
	mLLM := mocks.NewLLM(t)
	var ch chan string
	mLLM.EXPECT().Chat("system message", "### GIVEN KEY 1 START###\nValue1\n### GIVEN KEY 1 END###\n\n\n\nuser message", ch).Return(chatResp, nil)

	sut := gollm.LLMBlock{
		LLM:       mLLM,
		Name:      "block1",
		SystemMsg: "system message",
		UserMsg:   gollm.Msg{Text: "user message"},
		ExpectKeys: gollm.ExpectKeys{
			{
				Key:   "Key1",
				Label: "given key 1",
			},
		},
	}

	result, err := sut.Run(map[string]string{"Key1": "Value1"})
	assert.Nil(err)

	want := &gollm.BlockResult{
		Name: "block1",
		Text: "response message",
	}

	assert.Equal(want, result)
}
