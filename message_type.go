package midjourney

type Message struct {
	Type        int                  `json:"type"`
	ID          string               `json:"id"`
	Content     string               `json:"content"`
	ChannelID   string               `json:"channel_id"`
	Author      MessageAuthor        `json:"author"`
	Attachments []MessageAttachement `json:"attachments"`
	Mentions    []MessageMention     `json:"mentions"`
	Components  []MessageComponent   `json:"components"`
}

type MessageAuthor struct {
	ID string `json:"id"`
}

type MessageAttachement struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	ContentType string `json:"content_type"`
}

type MessageMention struct {
	ID string `json:"id"`
}

type MessageComponent struct {
	Type       int                               `json:"type"`
	Components []MessageComponentNestedComponent `json:"components"`
}

type MessageComponentNestedComponent struct {
	Type     int    `json:"type"`
	CustomID string `json:"custom_id"`
	Style    int    `json:"style"`
	Label    string `json:"label"`
}
