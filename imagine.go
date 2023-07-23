package midjourney

import (
	"context"
	"fmt"
	"strings"
)

type ImagineResponse struct {
	OriginalMessage Message
	ImageURL        string
}

func (c *Client) Imagine(ctx context.Context, prompt string) (*Image, error) {
	err := c.ImaginePrompt(ctx, prompt)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	fmt.Printf("imagining finished: %v\n", prompt)

	return c.GetImage(ctx, func(msg *Message) bool {
		splittedMsg := strings.Split(msg.Content, c.UserID)
		fmt.Printf("splittedMsg: %v\n", splittedMsg)
		return !strings.Contains(splittedMsg[1], "%") && !strings.Contains(splittedMsg[1], "Waiting")
	})
}
