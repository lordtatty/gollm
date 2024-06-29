package gollm_test

import (
	"testing"

	"github.com/lordtatty/gollm"
	"github.com/lordtatty/gollm/llm"
	"github.com/lordtatty/gollm/mocks"
	"github.com/lordtatty/gollm/prompt"
	"github.com/stretchr/testify/assert"
)

func TestMsg_String(t *testing.T) {
	assert := assert.New(t)

	sut := gollm.Msg{
		Text: "user message",
		VariableKVs: gollm.VariableKVs{
			{
				Key:   "Key1",
				Label: "given key 1",
			},
		},
		FixedKVs: prompt.StrBlocks{
			{
				Key: "Fixed Key 1",
				Val: "Fixed Value 1",
			},
		},
	}

	result, err := sut.String(map[string]string{"Key1": "Value1"})
	assert.Nil(err)

	want := "### FIXED KEY 1 START###\nFixed Value 1\n### FIXED KEY 1 END###\n\n\n\n### GIVEN KEY 1 START###\nValue1\n### GIVEN KEY 1 END###\n\n\n\nuser message"
	assert.Equal(want, result)
}

func TestMsg_String_MissingKeys(t *testing.T) {
	assert := assert.New(t)

	sut := gollm.Msg{
		Text: "user message",
		VariableKVs: gollm.VariableKVs{
			{
				Key:   "Key1",
				Label: "given key 1",
			},
		},
	}

	result, err := sut.String(map[string]string{})
	assert.Error(err)
	assert.Equal("", result)
	assert.ErrorContains(err, "missing values for input keys: block Key1 not found in blockOutputs")
}

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
		UserMsg: gollm.Msg{
			Text: "user message",
			VariableKVs: gollm.VariableKVs{
				{
					Key:   "Key1",
					Label: "given key 1",
				},
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
