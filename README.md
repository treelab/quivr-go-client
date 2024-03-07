# Introduction
Go client for [Quivr](https://github.com/QuivrHQ/quivr)

# Install
```shell
go get -u github.com/treelab/quivr-go-client
```

# Usage
```go
package main

import (
	"context"
	"fmt"
	quivr "github.com/treelab/quivr-go-client"
)

func main() {
	client := quivr.NewClient("http://localhost:5050", "226ce9e7e63cc389a5a035b64e4a305b")
	chat, err := client.CreateChat(context.Background(), &quivr.CreateChatInput{
		Name: "test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(chat.ChatId)
}

```