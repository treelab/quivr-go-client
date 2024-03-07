package quivr_go_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_CreateChatQuestion(t *testing.T) {
	c := _testClient()

	question, err := c.CreateChatQuestion(context.Background(), &CreateChatQuestionInput{
		ChatId:   "c655f7d0-c8e4-4906-9c33-43b1b2970bfd",
		BrainId:  "c12669e9-9780-4ca2-a218-ddd36884dfc0",
		Question: "nico的中文名是什么？",
	})
	assert.Nil(t, err)
	_testPrint(question)
}
