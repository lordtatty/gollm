package prompt_test

import (
	"testing"

	"github.com/lordtatty/gollm/prompt"
	"github.com/stretchr/testify/assert"
)

func TestMsg_String(t *testing.T) {
	assert := assert.New(t)

	sut := prompt.Msg{
		Text: "user message",
		VariableBlocks: prompt.VariableBlocks{
			{
				Key:   "Key1",
				Label: "given key 1",
			},
		},
		FixedBlocks: prompt.FixedBlocks{
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

	sut := prompt.Msg{
		Text: "user message",
		VariableBlocks: prompt.VariableBlocks{
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
