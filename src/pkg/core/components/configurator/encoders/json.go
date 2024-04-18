package encoders

import (
	"encoding/json"
)

// JsonEncoder - кодировщик для управления данными в формате json.
type JsonEncoder struct{}

// Encode - кодирование данных в формат json.
func (encoder JsonEncoder) Encode(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "\t")
}

// Decode - декодирование данных json формата.
func (encoder JsonEncoder) Decode(data []byte, v any) error {

	return json.Unmarshal(data, v)
}
