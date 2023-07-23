package midjourney

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

type Image struct {
	OriginalMessage Message
	URL             string
}

func (c *Client) GetImage(ctx context.Context, checkIfFinished func(msg *Message) bool) (*Image, error) {
	fmt.Printf("%v\n", "sleeping 10 seconds...")
	time.Sleep(time.Second * 10)
	msgs, err := c.GetMessages(ctx)

	if err != nil {
		return nil, err
	}

	var msg *Message
	var isFinished bool

	for {
		isFinished, msg, err = c.DetectMidjourneyResponse(msgs, checkIfFinished)
		if err != nil {
			return nil, err
		}

		if isFinished {
			break
		}

		fmt.Printf("%v\n", "Image currently processing, sleeping 10 seconds before retry")
		time.Sleep(time.Second * 10)
		msgs, err = c.GetMessages(ctx)

		if err != nil {
			return nil, err
		}
	}

	res := Image{
		OriginalMessage: *msg,
		URL:             msg.Attachments[0].Url,
	}

	return &res, nil
}

func (c *Client) DetectMidjourneyResponse(msgs *[]Message, checkIfFinished func(msg *Message) bool) (bool, *Message, error) {
	for _, msg := range *msgs {
		isSendByMidjourney := msg.Author.ID == c.MidJourneyApplicationID

		if !isSendByMidjourney {
			continue
		}

		indexMention := slices.IndexFunc[MessageMention](msg.Mentions, func(mm MessageMention) bool {
			return mm.ID == c.UserID
		})

		isMessageMentioningMe := indexMention < 0

		if isMessageMentioningMe {
			continue
		}

		isFinished := checkIfFinished(&msg)

		return isFinished, &msg, nil
	}

	return false, nil, errors.New("couldn't find processing image")
}
