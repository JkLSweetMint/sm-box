package encoders

import (
	"gopkg.in/yaml.v3"
)

// YamlEncoder - кодировщик для управления данными в формате yaml.
type YamlEncoder struct{}

// Encode - кодирование данных в формат yaml.
func (encoder YamlEncoder) Encode(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode - декодирование данных yaml формата.
func (encoder YamlEncoder) Decode(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
