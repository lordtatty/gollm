package gollm_test

import (
	"testing"

	"github.com/lordtatty/gollm"
	"github.com/lordtatty/gollm/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRunnableBlocks_Names(t *testing.T) {
	assert := assert.New(t)

	mBlock1 := mocks.NewRunnableBlock(t)
	mBlock1.EXPECT().UniqName().Return("block1")
	mBlock2 := mocks.NewRunnableBlock(t)
	mBlock2.EXPECT().UniqName().Return("block2")

	sut := gollm.RunnableBlocks{
		mBlock1,
		mBlock2,
	}

	want := []string{"block1", "block2"}
	result := sut.Names()

	assert.Equal(want, result)
}

func TestRunnableBlocks_Valid_IsInvalid(t *testing.T) {
	assert := assert.New(t)

	mBlock1 := mocks.NewRunnableBlock(t)
	mBlock1.EXPECT().UniqName().Return("block1")
	mBlock2 := mocks.NewRunnableBlock(t)
	mBlock2.EXPECT().UniqName().Return("block2")
	mBlock3 := mocks.NewRunnableBlock(t)
	mBlock3.EXPECT().UniqName().Return("block2") // Duplicate name

	sut := gollm.RunnableBlocks{
		mBlock1,
		mBlock2,
		mBlock3,
	}

	result := sut.Valid()

	assert.Error(result)
	assert.Contains(result.Error(), "block name block2 is not unique")
}

func TestFlow(t *testing.T) {
	assert := assert.New(t)

	mBlock1 := mocks.NewRunnableBlock(t)
	mBlock1.EXPECT().UniqName().Return("block1")
	mBlock1.EXPECT().Run(map[string]string{}).Return(&gollm.BlockResult{Text: "response1"}, nil)
	mBlock2 := mocks.NewRunnableBlock(t)
	mBlock2.EXPECT().UniqName().Return("block2")
	mBlock2.EXPECT().Run(map[string]string{"block1": "response1"}).Return(&gollm.BlockResult{Text: "response2"}, nil)

	blocks := []gollm.RunnableBlock{mBlock1, mBlock2}

	want := &gollm.FlowResp{
		Run:    []string{"block1", "block2"},
		NotRun: []string{},
		Outputs: map[string]gollm.BlockResult{
			"block1": {
				Text: "response1",
			},
			"block2": {
				Text: "response2",
			},
		},
	}
	result, err := gollm.Flow(blocks)
	assert.Nil(err)
	assert.Equal(want, result)
}
