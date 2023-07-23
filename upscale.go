package midjourney

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/exp/slices"
)

func (c *Client) Upscale(ctx context.Context, msg *Message, index Upscale) (*Image, error) {
	fmt.Printf("Upscaling: %v\n", index)

	var customID *string

	for _, mc := range msg.Components {
		indexComponent := slices.IndexFunc[MessageComponentNestedComponent](mc.Components, func(mc MessageComponentNestedComponent) bool {
			return index == mc.Label
		})
		if indexComponent >= 0 {
			customID = &mc.Components[indexComponent].CustomID
			break
		}
	}

	if customID == nil {
		return nil, errors.New("couldn't find upscale action")
	}

	payload := InteractionPayload[UpscalePayloadData]{
		Type:          3,
		ApplicationID: c.MidJourneyApplicationID,
		GuildID:       c.GuildID,
		ChannelID:     c.ChannelID,
		SessionID:     c.MidJourneySessionID,
		MessageID:     msg.ID,
		MessageFlags:  0,
		Data: UpscalePayloadData{
			ComponentType: 2,
			CustomID:      *customID,
		},
	}

	payloadRequest, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	err = c.Request(ctx, http.MethodPost, Endpoint_Interactions, &payloadRequest, nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("upscaling finished: %v\n", index)

	const keywordFinishedDetection = "Image #"
	occurenceKeyword := strings.Count(msg.Content, keywordFinishedDetection)

	return c.GetImage(ctx, func(msgIter *Message) bool {
		occurenceKeywordIter := strings.Count(msgIter.Content, keywordFinishedDetection)
		return occurenceKeywordIter > occurenceKeyword
	})
}
