package postman

import "encoding/json"

// Items - являются базовой единицей для сбора почтальоном.
// Это может быть либо запрос (Item), либо папка (ItemGroup).
type Items struct {
	// Общие поля.
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Events                  []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior any         `json:"protocolProfileBehavior,omitempty"`

	// Поля, относящиеся к Item
	ID        string      `json:"id,omitempty"`
	Request   *Request    `json:"request,omitempty"`
	Responses []*Response `json:"response,omitempty"`

	// Поля, относящиеся к ItemGroup
	Items []*Items `json:"item"`
	Auth  *Auth    `json:"auth,omitempty"`
}

// Item - это объект, содержащий фактический HTTP-запрос и прикрепленные к нему примеры ответов.
type Item struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Events                  []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior any         `json:"protocolProfileBehavior,omitempty"`
	ID                      string      `json:"id,omitempty"`
	Request                 *Request    `json:"request,omitempty"`
	Responses               []*Response `json:"response,omitempty"`
}

// ItemGroup - представляет собой упорядоченный набор запросов.
type ItemGroup struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Events                  []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior any         `json:"protocolProfileBehavior,omitempty"`
	Items                   []*Items    `json:"item"`
	Auth                    *Auth       `json:"auth,omitempty"`
}

// NewItem - является помощником для создания нового Item.
func NewItem(i Item) *Items {
	return &Items{
		Name:                    i.Name,
		Description:             i.Description,
		Variables:               i.Variables,
		Events:                  i.Events,
		ProtocolProfileBehavior: i.ProtocolProfileBehavior,
		ID:                      i.ID,
		Request:                 i.Request,
		Responses:               i.Responses,
	}
}

// NewItemGroup - является помощником для создания нового ItemGroup.
func NewItemGroup(ig ItemGroup) *Items {
	return &Items{
		Name:                    ig.Name,
		Description:             ig.Description,
		Variables:               ig.Variables,
		Events:                  ig.Events,
		ProtocolProfileBehavior: ig.ProtocolProfileBehavior,
		Items:                   ig.Items,
		Auth:                    ig.Auth,
	}
}

// IsGroup - возвращает значение false, поскольку Item не является группой.
func (i Items) IsGroup() bool {
	if i.Items != nil {
		return true
	}

	return false
}

// AddItem - добавляет Item к срезу существующих Items.
func (i *Items) AddItem(item *Items) {
	i.Items = append(i.Items, item)
}

// AddItemGroup - создает новую папку Items и добавляет ее к фрагменту существующих Items.
func (i *Items) AddItemGroup(name string) (f *Items) {
	f = &Items{
		Name:  name,
		Items: make([]*Items, 0),
	}

	i.Items = append(i.Items, f)

	return
}

// MarshalJSON - возвращает кодировку Item/ItemGroup в формате JSON.
func (i Items) MarshalJSON() ([]byte, error) {
	if i.IsGroup() {
		return json.Marshal(ItemGroup{
			Name:                    i.Name,
			Description:             i.Description,
			Variables:               i.Variables,
			Events:                  i.Events,
			ProtocolProfileBehavior: i.ProtocolProfileBehavior,
			Items:                   i.Items,
			Auth:                    i.Auth,
		})
	}

	return json.Marshal(Item{
		Name:                    i.Name,
		Description:             i.Description,
		Variables:               i.Variables,
		Events:                  i.Events,
		ProtocolProfileBehavior: i.ProtocolProfileBehavior,
		ID:                      i.ID,
		Request:                 i.Request,
		Responses:               i.Responses,
	})
}
