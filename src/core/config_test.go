package core

import (
	"fmt"
	"reflect"
	"sm-box/src/core/components/closer"
	"sm-box/src/core/components/configurator/encoders"
	"testing"
)

func TestConfig_Default(t *testing.T) {
	type fields struct {
		Closer *closer.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Closer: tt.fields.Closer,
			}

			v, _ := encoders.XmlEncoder{}.Encode(conf.Default())

			fmt.Println(string(v))

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
