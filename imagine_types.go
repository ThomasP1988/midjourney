package midjourney

type ImaginePayloadDataOption struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ImaginePayloadDataApplicationCommandOption struct {
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ImaginePayloadDataApplicationCommand struct {
	ID                       string                                       `json:"id"`
	ApplicationID            string                                       `json:"application_id"`
	Version                  string                                       `json:"version"`
	DefaultMemberPermissions *string                                      `json:"default_member_permissions"`
	Type                     int                                          `json:"type"`
	NSFW                     bool                                         `json:"nsfw"`
	Name                     string                                       `json:"name"`
	Description              string                                       `json:"description"`
	DMPermission             bool                                         `json:"dm_permission"`
	Options                  []ImaginePayloadDataApplicationCommandOption `json:"options"`
	Attachements             struct{}                                     `json:"attachments"`
}

type ImaginePayloadData struct {
	Version            string                               `json:"version"`
	ID                 string                               `json:"id"`
	Name               string                               `json:"name"`
	Type               int                                  `json:"type"`
	Options            []ImaginePayloadDataOption           `json:"options"`
	ApplicationCommand ImaginePayloadDataApplicationCommand `json:"application_command"`
}
