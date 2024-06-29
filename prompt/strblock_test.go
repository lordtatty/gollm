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
		VariableKVs: prompt.VariableKVs{
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

	sut := prompt.Msg{
		Text: "user message",
		VariableKVs: prompt.VariableKVs{
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

func TestStrBlock_Build(t *testing.T) {
	assert := assert.New(t)

	sut := prompt.StrBlock{
		Key: "People",
		Val: "Santa Claus",
		Vals: []string{
			"Rudolf",
			"Mrs Claus",
		},
	}

	result := sut.String(sut.Key, sut.Val)

	want := "### PEOPLE START###\nRudolf\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nMrs Claus\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nSanta Claus\n### PEOPLE END###\n\n"

	assert.Equal(want, result)
}

func TestStrBlocks_Build(t *testing.T) {
	assert := assert.New(t)

	sut := prompt.StrBlocks{
		{
			Key: "People",
			Val: "Santa Claus",
			Vals: []string{
				"Rudolf",
				"Mrs Claus",
			},
		},
		{
			Key: "Reindeer",
			Val: "Rudolf",
			Vals: []string{
				"Prancer",
				"Vixen",
			},
		},
	}

	result := sut.String()

	want := "### PEOPLE START###\nRudolf\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nMrs Claus\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nSanta Claus\n### PEOPLE END###\n\n"
	want += "### REINDEER START###\nPrancer\n### REINDEER END###\n\n"
	want += "### REINDEER START###\nVixen\n### REINDEER END###\n\n"
	want += "### REINDEER START###\nRudolf\n### REINDEER END###\n\n"

	assert.Equal(want, result)
}
