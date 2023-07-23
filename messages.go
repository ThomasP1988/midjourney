package midjourney

import (
	"context"
	"net/http"
	"strings"
)

func (c *Client) GetMessages(ctx context.Context) (*[]Message, error) {

	output := &[]Message{}

	var endpoint Endpoint = strings.Join([]string{string(Endpoint_Channels), c.ChannelID, "messages"}, "/")

	err := c.Request(ctx, http.MethodGet, endpoint, nil, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}
