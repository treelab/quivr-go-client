package quivr_go_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func _testPrint(o any) {
	bs, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(bs))
}

func _testClient() *Client {
	return NewClient("http://localhost:5050", "226ce9e7e63cc389a5a035b64e4a305b")
}

func TestNewClient(t *testing.T) {
	c := _testClient()

	ctx := context.Background()
	chat, err := c.CreateChat(ctx, &CreateChatInput{
		Name: "test",
	})
	assert.Nil(t, err)
	_testPrint(chat)
	brain, err := c.CreateBrain(ctx, &CreateBrainInput{
		Name:      "test",
		BrainType: "doc",
		Status:    "private",
	})
	assert.Nil(t, err)
	_testPrint(brain)

	resp, err := c.Upload(ctx, &UploadInput{
		BrainId:  brain.Id,
		ChatId:   Ptr(chat.ChatId),
		Filename: "test.txt",
		Data:     []byte("nico的中文名是袋鼠"),
	})
	assert.Nil(t, err)
	_testPrint(resp)

	question, err := c.CreateChatQuestion(ctx, &CreateChatQuestionInput{
		BrainId:  brain.Id,
		Question: "nico的中文名是什么？",
	})
	assert.Nil(t, err)
	_testPrint(question)
}
