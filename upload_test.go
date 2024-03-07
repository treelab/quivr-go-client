package quivr_go_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpload(t *testing.T) {
	c := _testClient()

	resp, err := c.Upload(context.Background(), &UploadInput{
		BrainId:  "c12669e9-9780-4ca2-a218-ddd36884dfc0",
		Filename: "test2.txt",
		Data:     []byte("文件里说了啥"),
	})
	assert.Nil(t, err)
	_testPrint(resp)
}
