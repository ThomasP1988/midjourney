package midjourney

import (
	"context"
	"net/http"
	"strings"
)

type ChannelResponse struct {
	GuildID string `json:"guild_id"`
}

func (c *Client) Channel(ctx context.Context, channelID string) (*ChannelResponse, error) {
	output := &ChannelResponse{}

	var endpoint Endpoint = strings.Join([]string{string(Endpoint_Channels), channelID}, "/")

	err := c.Request(ctx, http.MethodGet, endpoint, nil, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}
