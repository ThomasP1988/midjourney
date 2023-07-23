package midjourney

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestJSONPayload[T any] struct {
	JSON T `json:"json"`
}

func (c *Client) Request(ctx context.Context, method string, endpoint Endpoint, payload *[]byte, output any) error {
	url := c.APIURL + endpoint

	var payloadReq io.Reader

	if payload != nil {
		payloadReq = bytes.NewReader(*payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, payloadReq)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", c.AuthenticationToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		fmt.Printf("err Request: %v\n", err)
		return err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Request body err: %v\n", err)
			return err
		}
		fmt.Printf("body: %+v\n", string(body))
		return err
	}

	if output == nil {
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Request body err: %v\n", err)
		return err
	}
	// fmt.Printf("body: %+v\n", string(body))

	return json.Unmarshal(body, output)

}
