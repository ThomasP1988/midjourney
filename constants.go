package midjourney

type Endpoint = string

const (
	Endpoint_Channels     = "/channels"
	Endpoint_UserMe       = "/users/@me"
	Endpoint_Interactions = "/interactions"
)

type Upscale = string

const (
	Upscale_U1 = "U1"
	Upscale_U2 = "U2"
	Upscale_U3 = "U3"
	Upscale_U4 = "U4"
)
