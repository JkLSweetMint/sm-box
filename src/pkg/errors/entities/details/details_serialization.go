package details

import (
	"encoding/json"
	"encoding/xml"
)

// MarshalJSON - упаковать в формат JSON.
func (ds *Details) MarshalJSON() ([]byte, error) {
	ds.init()

	ds.rwMux.RLock()
	defer ds.rwMux.RUnlock()

	var w = make(map[string]any)

	for k, v := range ds.storage {
		w[k] = v
	}

	if len(ds.fields) > 0 {
		w["fields"] = ds.fields
	}

	return json.Marshal(w)
}

// MarshalXML - упаковать в формат XML.
func (ds *Details) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	ds.init()

	ds.rwMux.RLock()
	defer ds.rwMux.RUnlock()

	start = xml.StartElement{
		Name: xml.Name{
			Local: "Details",
		},
		Attr: nil,
	}

	var w = make(map[string]any)

	for k, v := range ds.storage {
		w[k] = v
	}

	if len(ds.fields) > 0 {
		w["Fields"] = ds.fields
	}

	if err = e.EncodeToken(start); err != nil {
		return
	}

	if err = ds.marshalXML(w, e, start); err != nil {
		return
	}

	if err = e.EncodeToken(start.End()); err != nil {
		return
	}

	return nil
}

// marshalXML - упаковать в формат XML.
func (ds *Details) marshalXML(w map[string]any, e *xml.Encoder, start xml.StartElement) (err error) {
	for k, v := range w {
		start = xml.StartElement{
			Name: xml.Name{
				Local: k,
			},
			Attr: nil,
		}

		if c, ok := v.(map[string]any); ok {
			if err = e.EncodeToken(start); err != nil {
				return
			}

			if err = ds.marshalXML(c, e, start); err != nil {
				return
			}

			if err = e.EncodeToken(start.End()); err != nil {
				return
			}

			continue
		}

		var subElement = xml.StartElement{
			Name: xml.Name{
				Local: "Item",
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

		if err = e.EncodeElement(v, subElement); err != nil {
			return
		}
	}

	return
}
