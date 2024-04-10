package encoders

import (
	"encoding/xml"
)

// XmlEncoder - кодировщик для управления данными в формате xml.
type XmlEncoder struct{}

// Encode - кодирование данных в формат xml.
func (encoder XmlEncoder) Encode(v any) ([]byte, error) {
	return xml.MarshalIndent(v, "", "\t")
}

// Decode - декодирование данных xml формата.
func (encoder XmlEncoder) Decode(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
