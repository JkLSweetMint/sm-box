package details

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// MarshalJSON - упаковать в формат JSON.
func (list Fields) MarshalJSON() ([]byte, error) {
	var w = make(map[string]any)

	for _, f := range list {
		if _, ok := f.Message.(json.Marshaler); !ok {
			if v, ok := f.Message.(fmt.Stringer); ok {
				w[f.Key.String()] = v.String()
			} else {
				w[f.Key.String()] = nil
			}
		} else {
			w[f.Key.String()] = f
		}
	}

	return json.Marshal(w)
}

// MarshalXML - упаковать в формат XML.
func (list Fields) MarshalXML(encoder *xml.Encoder, start xml.StartElement) (err error) {
	var w = make(map[string]any)

	for _, f := range list {
		if _, ok := f.Message.(xml.Marshaler); !ok {
			if v, ok := f.Message.(fmt.Stringer); ok {
				w[f.Key.String()] = v.String()
			} else {
				w[f.Key.String()] = nil
			}
		} else {
			w[f.Key.String()] = f
		}
	}

	start = xml.StartElement{
		Name: xml.Name{
			Local: "Item",
		},
		Attr: []xml.Attr{
			{
				Name: xml.Name{
					Local: "key",
				},
				Value: "fields",
			},
		},
	}

	if err = encoder.EncodeToken(start); err != nil {
		return
	}

	for k, v := range w {
		var subElement = xml.StartElement{
			Name: xml.Name{
				Local: "Field",
			},
			Attr: []xml.Attr{
				{
					Name: xml.Name{
						Local: "key",
					},
					Value: k,
				},
			},
		}

		if err = encoder.EncodeElement(v, subElement); err != nil {
			return
		}
	}

	if err = encoder.EncodeToken(start.End()); err != nil {
		return
	}

	return
}
