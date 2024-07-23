package models

import (
	"encoding/xml"
)

type (
	// Dictionary - внешняя модель словаря локализации.
	Dictionary map[string]any
)

// MarshalXML - упаковка словаря в XML.
func (dictionary Dictionary) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if err = e.EncodeToken(start); err != nil {
		return
	}

	if err = dictionary.marshalXML(dictionary, e, start); err != nil {
		return err
	}

	if err = e.EncodeToken(start.End()); err != nil {
		return
	}

	return
}

// marshalXML - упаковка словаря в XML.
func (dictionary Dictionary) marshalXML(data map[string]any, e *xml.Encoder, start xml.StartElement) (err error) {

	for key, value := range data {
		start = xml.StartElement{Name: xml.Name{"", key}}

		switch v := value.(type) {
		case map[string]any:
			if err = e.EncodeToken(start); err != nil {
				return
			}

			if err = dictionary.marshalXML(v, e, start); err != nil {
				return
			}
		case string:
			{
				start.Attr = append(start.Attr, xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "value",
					},
					Value: v,
				})

				if err = e.EncodeToken(start); err != nil {
					return
				}
			}
		}

		if err = e.EncodeToken(start.End()); err != nil {
			return
		}
	}

	return e.Flush()
}
