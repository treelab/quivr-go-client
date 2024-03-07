package quivr_go_client

import (
	"context"
	"fmt"
)

type UploadInput struct {
	BrainId  string  `json:"brain_id,omitempty"`
	ChatId   *string `json:"chat_id,omitempty"`
	Filename string  `json:"filename,omitempty"`
	Data     []byte  `json:"data,omitempty"`
}

type UploadOutput struct {
	Message string `json:"message"`
}

func (c *Client) Upload(ctx context.Context, input *UploadInput) (*UploadOutput, error) {
	api := fmt.Sprintf("/upload?brain_id=%s", input.BrainId)
	if input.ChatId != nil {
		api += fmt.Sprintf("&chat_id=%s", *input.ChatId)
	}
	return Upload[UploadOutput](ctx, c, api, input.Filename, input.Data)
}
