package postman

// Эти константы представляют доступные необработанные языки.
const (
	HTML       string = "html"
	Javascript string = "javascript"
	JSON       string = "json"
	Text       string = "text"
	XML        string = "xml"
)

// Body - представляет данные, обычно содержащиеся в теле запроса.
type Body struct {
	Mode       string       `json:"mode"`
	Raw        string       `json:"raw,omitempty"`
	URLEncoded any          `json:"urlencoded,omitempty"`
	FormData   any          `json:"formdata,omitempty"`
	File       any          `json:"file,omitempty"`
	GraphQL    any          `json:"graphql,omitempty"`
	Disabled   bool         `json:"disabled,omitempty"`
	Options    *BodyOptions `json:"options,omitempty"`
}

// BodyOptions - содержит параметры тела.
type BodyOptions struct {
	Raw BodyOptionsRaw `json:"raw,omitempty"`
}

// BodyOptionsRaw - представляет собой фактический язык, который будет использоваться в postman.
type BodyOptionsRaw struct {
	Language string `json:"language,omitempty"`
}
