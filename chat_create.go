package quivr_go_client

import "context"

type CreateChatInput struct {
	Name string `json:"name"`
}

type CreateChatOutput struct {
	ChatId       string `json:"chat_id"`
	UserId       string `json:"user_id"`
	CreationTime string `json:"creation_time"`
	ChatName     string `json:"chat_name"`
}

func (c *Client) CreateChat(ctx context.Context, input *CreateChatInput) (*CreateChatOutput, error) {
	return Do[CreateChatOutput](ctx, c, POST, "/chat", input)
}
