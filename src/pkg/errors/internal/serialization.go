package internal

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

type (
	// wrapper - структура обертка для упаковки ошибки.
	wrapper struct {
		ID     types.ID `json:"id"     xml:"id,attr"`
		Type   string   `json:"type"   xml:"type,attr"`
		Status string   `json:"status" xml:"status,attr"`

		Message any           `json:"message"           xml:"Message"`
		Details types.Details `json:"details,omitempty" xml:"Details,omitempty"`
	}
)

// MarshalJSON - упаковать в формат JSON.
func (i *Internal) MarshalJSON() ([]byte, error) {
	var w = &wrapper{
		ID:     i.Store.ID,
		Type:   i.Store.Type.String(),
		Status: i.Store.Status.String(),

		Message: i.Store.Message,
		Details: i.Store.Details,
	}

	// Сообщение
	{
		if _, ok := w.Message.(json.Marshaler); !ok {
			if v, ok := w.Message.(fmt.Stringer); ok {
				w.Message = v.String()
			} else {
				w.Message = nil
			}
		}
	}

	return json.Marshal(w)
}

// MarshalXML - упаковать в формат XML.
func (i *Internal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var w = &wrapper{
		ID:     i.Store.ID,
		Type:   i.Store.Type.String(),
		Status: i.Store.Status.String(),

		Message: i.Store.Message,
		Details: i.Store.Details,
	}

	// Сообщение
	{
		if _, ok := w.Message.(json.Marshaler); !ok {
			if v, ok := w.Message.(fmt.Stringer); ok {
				w.Message = v.String()
			} else {
				w.Message = nil
			}
		}
	}

	start = xml.StartElement{
		Name: xml.Name{
			Local: "Error",
		},
	}

	return e.EncodeElement(w, start)
}

// UnmarshalJSON - распаковать из формата JSON.
func (i *Internal) UnmarshalJSON(bytes []byte) (err error) {
	var w = make(map[string]any)

	if err = json.Unmarshal(bytes, &w); err != nil {
		return
	}

	i.ctx = context.Background()

	// Основные данные
	{
		// ID
		{
			switch v := w["id"].(type) {
			case string:
				{
					i.Store.ID = types.ID(v)
				}
			}
		}

		// Type
		{
			switch v := w["type"].(type) {
			case string:
				{
					i.Store.Type = types.ParseErrorType(v)
				}
			}
		}

		// Status
		{
			switch v := w["status"].(type) {
			case string:
				{
					i.Store.Status = types.ParseStatus(v)
				}
			}
		}
	}

	// Сообщение
	{
		switch v := w["message"].(type) {
		case string:
			{
				i.Store.Message = new(messages.TextMessage).Text(v)
			}
		}
	}

	// Детали
	{
		i.Store.Details = new(details.Details)

		if data, ok := w["details"].(map[string]any); ok {
			for k, v := range data {
				if k == "fields" {
					if data, ok := v.(map[string]any); ok {
						for k, v := range data {
							var m types.Message

							// Сообщение
							{
								switch q := v.(type) {
								case string:
									{
										m = new(messages.TextMessage).Text(q)
									}
								}
							}

							i.Store.Details.SetField(new(details.FieldKey).Add(k), m)
						}
					}

					continue
				}

				i.Store.Details.Set(k, v)
			}
		}
	}

	return
}
