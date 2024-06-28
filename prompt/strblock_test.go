package prompt_test

import (
	"testing"

	"github.com/lordtatty/gollm/prompt"
	"github.com/stretchr/testify/assert"
)

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

	result := sut.Build(sut.Key, sut.Val)

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

	result := sut.Build()

	want := "### PEOPLE START###\nRudolf\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nMrs Claus\n### PEOPLE END###\n\n"
	want += "### PEOPLE START###\nSanta Claus\n### PEOPLE END###\n\n"
	want += "### REINDEER START###\nPrancer\n### REINDEER END###\n\n"
	want += "### REINDEER START###\nVixen\n### REINDEER END###\n\n"
	want += "### REINDEER START###\nRudolf\n### REINDEER END###\n\n"

	assert.Equal(want, result)
}
