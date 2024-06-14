package postman

import (
	"encoding/json"
	"fmt"
	"io"
)

type version string

const (
	// V210 : v2.1.0
	V210 version = "v2.1.0"
	// V200 : v2.0.0
	V200 version = "v2.0.0"
)

// Info - хранит данные о коллекции.
type Info struct {
	Name        string      `json:"name"`
	Description Description `json:"description"`
	Version     string      `json:"version"`
	Schema      string      `json:"schema"`
}

// Collection - представляет коллекцию Postman.
type Collection struct {
	Auth      *Auth       `json:"auth,omitempty"`
	Info      Info        `json:"info"`
	Items     []*Items    `json:"item"`
	Events    []*Event    `json:"event,omitempty"`
	Variables []*Variable `json:"variable,omitempty"`
}

// NewCollection - возвращает новую коллекцию.
func NewCollection(name string, desc string) *Collection {
	return &Collection{
		Info: Info{
			Name: name,
			Description: Description{
				Content: desc,
			},
		},
	}
}

// AddItem - добавляет элемент (Item или ItemGroup) к существующему фрагменту элементов.
func (c *Collection) AddItem(item *Items) {
	if c.Items == nil {
		c.Items = make([]*Items, 0)
	}

	c.Items = append(c.Items, item)
}

// AddItemGroup - создает новую ItemGroup и добавляет ее к существующему фрагменту items.
func (c *Collection) AddItemGroup(name string) (f *Items) {
	if c.Items == nil {
		c.Items = make([]*Items, 0)
	}

	f = &Items{
		Name:  name,
		Items: make([]*Items, 0),
	}

	c.Items = append(c.Items, f)

	return
}

// Write - кодирует структуру коллекции в формате JSON и записывает ее в предоставленный io.Writer.
func (c *Collection) Write(w io.Writer, v version) (err error) {

	c.Info.Schema = fmt.Sprintf("https://schema.getpostman.com/json/collection/%s/collection.json", string(v))
	setVersionForItems(c.Items, v)

	file, _ := json.MarshalIndent(c, "", "    ")

	_, err = w.Write(file)

	return
}

// setVersionForItems - установите версию для всех структур, поведение которых отличается в зависимости от
// версии коллекции Postman.
func setVersionForItems(items []*Items, v version) {
	for _, i := range items {
		if i.Auth != nil {
			i.Auth.setVersion(v)
		}
		if i.IsGroup() {
			setVersionForItems(i.Items, v)
		} else {
			if i.Request != nil {
				if i.Request.Auth != nil {
					i.Request.Auth.setVersion(v)
				}
				if i.Request.URL != nil {
					i.Request.URL.setVersion(v)
				}
			}
		}
	}
}

// ParseCollection - преобразует содержимое предоставленного потока данных в объект коллекции.
func ParseCollection(r io.Reader) (c *Collection, err error) {
	err = json.NewDecoder(r).Decode(&c)

	return
}
