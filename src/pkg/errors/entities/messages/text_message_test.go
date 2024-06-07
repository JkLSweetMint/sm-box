package messages

import (
	"reflect"
	"testing"
)

func TestTextMessage_Clone(t *testing.T) {
	type fields struct {
		content string
	}

	tests := []struct {
		name   string
		fields fields
		want   *TextMessage
	}{
		{
			name: "Case 1",
			fields: fields{
				content: "test",
			},
			want: &TextMessage{
				content: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TextMessage{
				content: tt.fields.content,
			}

			if got := m.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextMessage_Format(t *testing.T) {
	type fields struct {
		content string
	}

	type args struct {
		format string
		a      []any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TextMessage
	}{
		{
			name: "Case 1",
			fields: fields{
				content: "",
			},
			args: args{
				format: "test case %d",
				a:      []any{1},
			},
			want: &TextMessage{
				content: "test case 1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TextMessage{
				content: tt.fields.content,
			}

			if got := m.Format(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextMessage_String(t *testing.T) {
	type fields struct {
		content string
	}

	tests := []struct {
		name    string
		fields  fields
		wantStr string
	}{
		{
			name: "Case 1",
			fields: fields{
				content: "test",
			},
			wantStr: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TextMessage{
				content: tt.fields.content,
			}

			if gotStr := m.String(); gotStr != tt.wantStr {
				t.Errorf("String() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestTextMessage_Text(t *testing.T) {
	type fields struct {
		content string
	}

	type args struct {
		content string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TextMessage
	}{
		{
			name: "Case 1",
			fields: fields{
				content: "",
			},
			args: args{
				content: "test",
			},
			want: &TextMessage{
				content: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TextMessage{
				content: tt.fields.content,
			}

			if got := m.Text(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Text() = %v, want %v", got, tt.want)
			}
		})
	}
}
