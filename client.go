package midjourney

import (
	"context"
	"net/http"
)

type Client struct {
	APIURL                  string
	ChannelID               string
	AuthenticationToken     string
	GuildID                 string
	UserID                  string
	MidJourneyApplicationID string
	MidJourneyDataID        string
	MidJourneyDataVersion   string
	MidJourneySessionID     string

	HTTPClient *http.Client
}

type ClientOpts = func(client *Client) error

func NewClient(ctx context.Context, authenticationToken string, channelID string, opts ...ClientOpts) (*Client, error) {

	client := &Client{
		APIURL:                  "https://discord.com/api/v9",
		ChannelID:               channelID,
		AuthenticationToken:     authenticationToken,
		HTTPClient:              http.DefaultClient,
		MidJourneyApplicationID: "936929561302675456",
		MidJourneyDataID:        "938956540159881230",
		MidJourneyDataVersion:   "1118961510123847772",
		MidJourneySessionID:     "2fb980f65e5c9a77c96ca01f2c242cf6",
	}

	channel, err := client.Channel(ctx, channelID)

	if err != nil {
		return nil, err
	}

	client.GuildID = channel.GuildID

	me, err := client.Me(ctx)

	if err != nil {
		return nil, err
	}

	client.UserID = me.ID

	for _, option := range opts {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}
