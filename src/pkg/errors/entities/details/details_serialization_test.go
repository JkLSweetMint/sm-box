package details

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"sm-box/pkg/errors/entities/messages"
	"sync"
	"testing"
)

func TestDetails_MarshalJSON(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
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
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want:    []byte{123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125, 44, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				fields: nil,
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want:    []byte{123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want:    []byte{123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 4",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: nil,
			},
			want:    []byte{123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 123, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125, 44, 34, 116, 101, 115, 116, 34, 58, 34, 49, 50, 51, 34, 125},
			wantErr: false,
		},
		{
			name: "Case 5",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want:    []byte{123, 125},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}
			got, err := json.Marshal(ds)

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

func TestDetails_MarshalXML(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
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
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want:    []byte{60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 73, 116, 101, 109, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 70, 105, 101, 108, 100, 62, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				fields: nil,
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want:    []byte{60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want:    []byte{60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 70, 105, 101, 108, 100, 62, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62},
			wantErr: false,
		},
		{
			name: "Case 4",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: nil,
			},
			want:    []byte{60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 73, 116, 101, 109, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 102, 105, 101, 108, 100, 115, 34, 62, 60, 70, 105, 101, 108, 100, 32, 107, 101, 121, 61, 34, 116, 101, 115, 116, 34, 62, 49, 50, 51, 60, 47, 70, 105, 101, 108, 100, 62, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62},
			wantErr: false,
		},
		{
			name: "Case 5",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want:    []byte{60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}
			got, err := xml.Marshal(ds)

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
