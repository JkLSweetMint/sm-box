package postman

// Variable - позволяет вам сохранять и повторно использовать значения в ваших запросах и скриптах.
type Variable struct {
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	System      bool   `json:"system,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
}

// NewVariable - создает новую переменную типа string.
func NewVariable(name string, value string) *Variable {
	return &Variable{
		Name:  name,
		Value: value,
		Type:  "string",
	}
}
