package encoders

import (
	"reflect"
	"testing"
)

func TestJsonEncoder_Decode(t *testing.T) {
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
				data: []byte{123, 10, 9, 34, 101, 108, 101, 109, 101, 110, 116, 115, 34, 58, 32, 91, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 49, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 50, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 51, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 52, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 53, 34, 10, 9, 9, 125, 10, 9, 93, 44, 10, 9, 34, 100, 97, 116, 101, 34, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10, 125},
				v:    new(TestConfig),
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				data: []byte{10, 9, 34, 101, 108, 101, 109, 101, 110, 116, 115, 34, 58, 32, 91, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 49, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 50, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 51, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 52, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 53, 34, 10, 9, 9, 125, 10, 9, 93, 44, 10, 9, 34, 100, 97, 116, 101, 34, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10, 125},
				v:    new(TestConfig),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := JsonEncoder{}

			if err := encoder.Decode(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonEncoder_Encode(t *testing.T) {
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
			want:    []byte{123, 10, 9, 34, 101, 108, 101, 109, 101, 110, 116, 115, 34, 58, 32, 91, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 49, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 50, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 51, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 52, 34, 10, 9, 9, 125, 44, 10, 9, 9, 123, 10, 9, 9, 9, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 53, 34, 10, 9, 9, 125, 10, 9, 93, 44, 10, 9, 34, 100, 97, 116, 101, 34, 58, 32, 34, 48, 48, 48, 49, 45, 48, 49, 45, 48, 49, 34, 10, 125},
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				v: nil,
			},
			want:    []byte{110, 117, 108, 108},
			wantErr: false,
		},
		{
			name: "Case 3",
			args: args{
				v: "1sa",
			},
			want:    []byte{34, 49, 115, 97, 34},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := JsonEncoder{}

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
