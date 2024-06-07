package details

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"sm-box/pkg/errors/entities/messages"
	"testing"
)

func TestFields_MarshalJSON(t *testing.T) {
	type fields struct {
		list Fields
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				list: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
			},
			want:    []byte{123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				list: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test").AddArray("arr", 1),
						Message: new(messages.TextMessage).Text("hzz"),
					},
				},
			},
			want:    []byte{123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 44, 34, 116, 101, 115, 116, 46, 97, 114, 114, 91, 49, 93, 34, 58, 34, 104, 122, 122, 34, 125},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				list: Fields{},
			},
			want:    []byte{123, 125},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := tt.fields.list

			got, err := json.Marshal(fs)

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFields_MarshalXML(t *testing.T) {
	type fields struct {
		list Fields
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				list: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
			},
			want:    []byte{60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 70, 105, 101, 108, 100, 62, 60, 47, 73, 116, 101, 109, 62},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				list: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test").AddArray("arr", 1),
						Message: new(messages.TextMessage).Text("hzz"),
					},
				},
			},
			want:    []byte{60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 70, 105, 101, 108, 100, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 46, 97, 114, 114, 91, 49, 93, 34, 62, 104, 122, 122, 60, 47, 70, 105, 101, 108, 100, 62, 60, 47, 73, 116, 101, 109, 62},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				list: Fields{},
			},
			want:    []byte{60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 47, 73, 116, 101, 109, 62},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := tt.fields.list

			got, err := xml.Marshal(fs)

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalXML() got = %v, want %v", got, tt.want)
			}
		})
	}
}
