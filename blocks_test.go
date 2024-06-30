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

	wantSysMsg := "system message"
	wantUsrMsg := "userMsg"
	variableKVs := map[string]string{"Key1": "Value1"}

	chatResp := &llm.ChatResp{
		Text: "response message",
	}
	mLLM := mocks.NewLLM(t)
	var ch chan string
	mLLM.EXPECT().Chat(wantSysMsg, wantUsrMsg, ch).Return(chatResp, nil)

	mUserMsg := mocks.NewUserMsg(t)
	mUserMsg.EXPECT().String(variableKVs).Return(wantUsrMsg, nil)

	sut := gollm.LLMBlock{
		LLM:       mLLM,
		Name:      "block1",
		SystemMsg: wantSysMsg,
		UserMsg:   mUserMsg,
	}

	result, err := sut.Run(variableKVs)
	assert.Nil(err)

	want := &gollm.BlockResult{
		Name: "block1",
		Text: "response message",
	}

	assert.Equal(want, result)
}
