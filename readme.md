## Introduction

**YOU MIGHT GET YOUR MIDJOURNET ACCOUNT OR DISCORD ACCOUNT BANNED**

This library **DOES NOT** use discord public API or midjourney API.

I initially wanted to develop a bot but discord is against bot to bot communication, this API is using the API dedicated to their front end, which makes this library very unstable.

Discord does not like automated user account.
Midjourney does not like automation.

This library is inspired by the PHP library [midjourney-discord-api-php](https://github.com/ferranfg/midjourney-discord-api-php)

I expect this library to be broken soon or later.

## How to use?

You need to get an authorization token [How to get your Discord token?](https://www.androidauthority.com/get-discord-token-3149920/)

Install the library 
`go get github.com/ThomasP1988/midjourney` 

Basically the library has two functions Imagine to set the prompt and upscale to chose the one you want.


	botToken := ""
	channelID := ""

	client, err := midjourney.NewClient(context.TODO(), botToken, channelID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
        return
	}

	image, err := client.Imagine(context.TODO(), "portraying a gentle touch, bathed in shades of beige, from a birds-eye-view perspective, beautiful romantic gesture, all hand-drawn with a touch of animated magic")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	upscaledImage, err := client.Upscale(context.TODO(), &image.OriginalMessage, midjourney.Upscale_U2)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("upscaledImage: %+v\n", upscaledImage)


## Concurrency

This library is designed in a way preventing concurrent imagine/upscale from the same user **and** to the same channel.
If you need to achieve concurrency, you need to set multiple client with each a different channel or user.

## Advice

Set your own discord server, this library is polling the messages but if there is a too high number of messages, your imagine/upscale answer can get lost.

## Modifying client constants

you can pass functions as +3rd argument on NewClient to modify the values of Client.


	func(client *Client) error



