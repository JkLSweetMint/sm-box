package postman

// ListenType - определяет тип сценария, привязанного к событию.
type ListenType string

const (
	// PreRequest - сценарий обычно выполняется перед отправкой HTTP-запроса.
	PreRequest ListenType = "prerequest"
	// Test - сценарий обычно выполняется после отправки фактического HTTP-запроса и получения ответа.
	Test ListenType = "test"
)

// Script - это фрагмент кода Javascript, который может быть использован для выполнения операций настройки
// или удаления определенного ответа.
type Script struct {
	ID   string   `json:"id,omitempty"`
	Type string   `json:"type,omitempty"`
	Exec []string `json:"exec,omitempty"`
	Src  *URL     `json:"src,omitempty"`
	Name string   `json:"name,omitempty"`
}

// Event - определяет сценарий, связанный с соответствующим именем события.
type Event struct {
	ID       string     `json:"id,omitempty"`
	Listen   ListenType `json:"listen,omitempty"`
	Script   *Script    `json:"script,omitempty"`
	Disabled bool       `json:"disabled,omitempty"`
}

// NewEvent - создает новое событие типа text/javascript.
func NewEvent(listenType ListenType, script []string) *Event {
	return &Event{
		Listen: listenType,
		Script: &Script{
			Type: "text/javascript",
			Exec: script,
		},
	}
}
