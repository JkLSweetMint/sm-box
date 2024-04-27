package mysql

import (
	"encoding/xml"
	"io"
)

// Tags - теги для подключения к базе данных.
type Tags map[string]string

func (m Tags) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if len(m) == 0 {
		return
	}

	for k, v := range m {
		start.Attr = append(start.Attr, xml.Attr{
			Name: xml.Name{
				Space: "",
				Local: k,
			},
			Value: v,
		})
	}

	if err = e.EncodeToken(start); err != nil {
		return
	}

	return e.EncodeToken(start.End())
}

func (m *Tags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tags = make(Tags)

	for _, attr := range start.Attr {
		tags[attr.Name.Local] = attr.Value
	}

	for {
		if err = d.Skip(); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
	}

	*m = tags

	return
}
