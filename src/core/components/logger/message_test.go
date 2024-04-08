package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
)

func TestMessage_Field(t *testing.T) {
	type fields struct {
		initiator string
		text      string
		lvl       zapcore.Level
		fields    []zap.Field
		write     func(lvl zapcore.Level, msg string, fields ...zap.Field)
	}

	type args struct {
		key string
		val any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.DebugLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 2. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.InfoLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 3. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.WarnLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 4. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.ErrorLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 5. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.PanicLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 6. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.DPanicLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
		{
			name: "Case 7. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"key",
				"value",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl: zap.FatalLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: nil,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message{
				initiator: tt.fields.initiator,
				text:      tt.fields.text,
				lvl:       tt.fields.lvl,
				fields:    tt.fields.fields,
				write:     tt.fields.write,
			}

			if got := msg.Field(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Format(t *testing.T) {
	type fields struct {
		initiator string
		text      string
		lvl       zapcore.Level
		fields    []zap.Field
		write     func(lvl zapcore.Level, msg string, fields ...zap.Field)
	}

	type args struct {
		format string
		a      []any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{1},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 1. ",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 2. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{2},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 2. ",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 3. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{3},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 3. ",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 4. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{4},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 4. ",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 5. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{5},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 5. ",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 6. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{6},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 6. ",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 7. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				"Test %d. ",
				[]any{7},
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 7. ",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message{
				initiator: tt.fields.initiator,
				text:      tt.fields.text,
				lvl:       tt.fields.lvl,
				fields:    tt.fields.fields,
				write:     tt.fields.write,
			}

			if got := msg.Format(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Initiator(t *testing.T) {
	type fields struct {
		initiator string
		text      string
		lvl       zapcore.Level
		fields    []zap.Field
		write     func(lvl zapcore.Level, msg string, fields ...zap.Field)
	}

	type args struct {
		name string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 2. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 3. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 4. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 5. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 6. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 7. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			},
			args: args{},
			want: Message(&message{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 8. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 9. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 10. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 11. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 12. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 13. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 14. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				name: "test",
			},
			want: Message(&message{
				initiator: "test",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message{
				initiator: tt.fields.initiator,
				text:      tt.fields.text,
				lvl:       tt.fields.lvl,
				fields:    tt.fields.fields,
				write:     tt.fields.write,
			}

			if got := msg.Initiator(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initiator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Text(t *testing.T) {
	type fields struct {
		initiator string
		text      string
		lvl       zapcore.Level
		fields    []zap.Field
		write     func(lvl zapcore.Level, msg string, fields ...zap.Field)
	}

	type args struct {
		text string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 1. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 1. ",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 2. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 2. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 2. ",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 3. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 3. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 3. ",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 4. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 4. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 4. ",

				lvl:    zap.ErrorLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 5. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 5. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 5. ",

				lvl:    zap.PanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 6. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 6. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 6. ",

				lvl:    zap.DPanicLevel,
				fields: nil,

				write: nil,
			}),
		},
		{
			name: "Case 7. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			},
			args: args{
				text: "Test 7. ",
			},
			want: Message(&message{
				initiator: "unknown",
				text:      "Test 7. ",

				lvl:    zap.FatalLevel,
				fields: nil,

				write: nil,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message{
				initiator: tt.fields.initiator,
				text:      tt.fields.text,
				lvl:       tt.fields.lvl,
				fields:    tt.fields.fields,
				write:     tt.fields.write,
			}

			if got := msg.Text(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Write(t *testing.T) {
	type fields struct {
		initiator string
		text      string
		lvl       zapcore.Level
		fields    []zap.Field
		write     func(lvl zapcore.Level, msg string, fields ...zap.Field)
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Fatalf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl: zap.DebugLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 2. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl: zap.InfoLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 3. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl: zap.WarnLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 4. ",
			fields: fields{
				initiator: "unknown",
				text:      "",

				lvl: zap.ErrorLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 5. ",
			fields: fields{
				initiator: "unknown",
				text:      "Test 8. ",

				lvl:    zap.DebugLevel,
				fields: nil,

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 6. ",
			fields: fields{
				initiator: "unknown",
				text:      "Test 9. ",

				lvl:    zap.InfoLevel,
				fields: nil,

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 7. ",
			fields: fields{
				initiator: "unknown",
				text:      "Test 10. ",

				lvl:    zap.WarnLevel,
				fields: nil,

				write: log.(*logger).instance.Log,
			},
		},
		{
			name: "Case 8. ",
			fields: fields{
				initiator: "unknown",
				text:      "Test 11. ",

				lvl: zap.ErrorLevel,
				fields: []zap.Field{
					zap.Any("key", "value"),
				},

				write: log.(*logger).instance.Log,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := &message{
				initiator: tt.fields.initiator,
				text:      tt.fields.text,
				lvl:       tt.fields.lvl,
				fields:    tt.fields.fields,
				write:     tt.fields.write,
			}

			msg.Write()
		})
	}
}
