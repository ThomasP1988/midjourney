package midjourney

import (
	"context"
	"fmt"
	"testing"
)

func TestImagineAndUpscale(t *testing.T) {
	botToken := ""
	channelID := ""

	client, err := NewClient(context.TODO(), botToken, channelID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	image, err := client.Imagine(context.TODO(), "portraying a gentle touch, bathed in shades of beige, from a birds-eye-view perspective, beautiful romantic gesture, all hand-drawn with a touch of animated magic")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	upscaledImage, err := client.Upscale(context.TODO(), &image.OriginalMessage, Upscale_U2)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}

	fmt.Printf("upscaledImage: %+v\n", upscaledImage)

	t.Fail()
}
