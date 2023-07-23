package midjourney

type InteractionPayload[T any] struct {
	Type          int    `json:"type"`
	ApplicationID string `json:"application_id"`
	GuildID       string `json:"guild_id"`
	MessageFlags  int    `json:"message_flags,omitempty"`
	ChannelID     string `json:"channel_id"`
	SessionID     string `json:"session_id"`
	MessageID     string `json:"message_id,omitempty"`
	Data          T      `json:"data"`
}
