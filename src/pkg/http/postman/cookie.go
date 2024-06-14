package postman

// Cookie - представляет собой файл cookie, который соответствует формату
// Google Chrome (https://developer.chrome.com/extensions/cookies)
type Cookie struct {
	Domain     string `json:"domain"`
	Expires    string `json:"expires,omitempty"`
	MaxAge     string `json:"maxAge,omitempty"`
	HostOnly   bool   `json:"hostOnly,omitempty"`
	HTTPOnly   bool   `json:"httpOnly,omitempty"`
	Name       string `json:"name,omitempty"`
	Path       string `json:"path"`
	Secure     string `json:"secure,omitempty"`
	Session    bool   `json:"session,omitempty"`
	Value      string `json:"value,omitempty"`
	Extensions any    `json:"extensions,omitempty"`
}
