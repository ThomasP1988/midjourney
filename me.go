package midjourney

import (
	"context"
	"net/http"
)

type MeResponse struct {
	ID string `json:"id"`
}

func (c *Client) Me(ctx context.Context) (*MeResponse, error) {
	output := &MeResponse{}

	err := c.Request(ctx, http.MethodGet, Endpoint_UserMe, nil, output)

	if err != nil {
		return nil, err
	}

	return output, nil
}
