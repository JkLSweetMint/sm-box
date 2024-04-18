package encoders

import (
	"reflect"
	"testing"
)

func TestYamlEncoder_Decode(t *testing.T) {
	type args struct {
		data []byte
		v    any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				data: []byte{69, 108, 101, 109, 101, 110, 116, 115, 58, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 49, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 50, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 51, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 52, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 53, 34, 10, 68, 97, 116, 101, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10},
				v:    new(TestConfig),
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				data: []byte{10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 108, 117, 101, 58, 32, 34, 50, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 51, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 52, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 53, 34, 10, 68, 97, 116, 101, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10},
				v:    new(TestConfig),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := YamlEncoder{}

			if err := encoder.Decode(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestYamlEncoder_Encode(t *testing.T) {
	type args struct {
		v any
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				v: new(TestConfig).Default(),
			},
			want:    []byte{69, 108, 101, 109, 101, 110, 116, 115, 58, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 49, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 50, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 51, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 52, 34, 10, 32, 32, 32, 32, 45, 32, 86, 97, 108, 117, 101, 58, 32, 34, 53, 34, 10, 68, 97, 116, 101, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10},
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				v: nil,
			},
			want:    []byte{110, 117, 108, 108, 10},
			wantErr: false,
		},
		{
			name: "Case 3",
			args: args{
				v: "1sa",
			},
			want:    []byte{49, 115, 97, 10},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := YamlEncoder{}
			got, err := encoder.Encode(tt.args.v)

			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
