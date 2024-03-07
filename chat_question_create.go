package quivr_go_client

import (
	"context"
	"fmt"
)

type CreateChatQuestionInput struct {
	ChatId   string `json:"chat_id,omitempty"`
	BrainId  string `json:"brain_id,omitempty"`
	Question string `json:"question,omitempty"`
}

type CreateChatQuestionOutput struct {
	ChatId      string `json:"chat_id"`
	MessageId   string `json:"message_id"`
	UserMessage string `json:"user_message"`
	Assistant   string `json:"assistant"`
	MessageTime string `json:"message_time"`
	PromptTitle any    `json:"prompt_title"`
	BrainName   string `json:"brain_name"`
	BrainId     string `json:"brain_id"`
	Metadata    any    `json:"metadata"`
}

func (c *Client) CreateChatQuestion(ctx context.Context, input *CreateChatQuestionInput) (*CreateChatQuestionOutput, error) {
	return Do[CreateChatQuestionOutput](ctx, c, POST, fmt.Sprintf("/chat/%s/question?brain_id=%s", input.ChatId, input.BrainId), input)
}
