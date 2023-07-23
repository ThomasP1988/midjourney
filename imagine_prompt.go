package midjourney

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ImaginePrompt(ctx context.Context, prompt string) error {

	fmt.Printf("imagining: %v\n", prompt)

	payload := InteractionPayload[ImaginePayloadData]{
		Type:          2,
		ApplicationID: c.MidJourneyApplicationID,
		GuildID:       c.GuildID,
		ChannelID:     c.ChannelID,
		SessionID:     c.MidJourneySessionID,
		Data: ImaginePayloadData{
			Version: c.MidJourneyDataVersion,
			ID:      c.MidJourneyDataID,
			Name:    "imagine",
			Type:    1,
			Options: []ImaginePayloadDataOption{
				{
					Type:  3,
					Name:  "prompt",
					Value: prompt,
				},
			},
			ApplicationCommand: ImaginePayloadDataApplicationCommand{
				ID:                       c.MidJourneyDataID,
				ApplicationID:            c.MidJourneyApplicationID,
				Version:                  c.MidJourneyDataVersion,
				DefaultMemberPermissions: nil,
				Type:                     1,
				NSFW:                     false,
				Name:                     "imagine",
				Description:              "Create images with Midjourney",
				DMPermission:             true,
				Options: []ImaginePayloadDataApplicationCommandOption{
					{
						Type:        3,
						Name:        "prompt",
						Description: "the prompt to imagine",
						Required:    true,
					},
				},
			},
		},
	}
	fmt.Printf("payload: %+v\n", payload)
	payloadRequest, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	err = c.Request(ctx, http.MethodPost, Endpoint_Interactions, &payloadRequest, nil)

	if err != nil {
		return err
	}

	return nil
}
