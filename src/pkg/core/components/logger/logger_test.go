package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"reflect"
	"sm-box/src/pkg/core/env"
	"sm-box/src/pkg/tools/size"
	"strings"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		initiator string
	}

	tests := []struct {
		name    string
		args    args
		wantLog Logger
		wantErr bool
	}{
		{
			name: "Case 1. ",
			args: args{
				initiator: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.args.initiator); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestLogger_Rotation(t *testing.T) {
	var conf = new(Config).Default()

	// Edit configuration
	{
		conf.Terminal = nil

		conf.Files = conf.Files[0:1]
		conf.Files[0].Options.Rotation.FileSize = "4KB"

		conf.FillEmptyFields()
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = newLogger("", conf); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	// Запись тестовых данных
	{
		for i := 0; i < 100; i++ {
			log.Info().Text("Test message ").Write()
		}
	}

	// Проверка ротации
	{
		var checker = func(s size.Size, p string) bool {
			var stat, err = os.Stat(p)

			if err != nil {
				return false
			}

			return stat.Size() <= conf.Files[0].Options.Rotation.FileSize.Int64()+((conf.Files[0].Options.Rotation.FileSize.Int64()/100)*10)
		}

		var (
			dir []os.DirEntry
			err error
		)

		if dir, err = os.ReadDir(path.Join(env.Paths.SystemLocation, conf.Files[0].Path)); err != nil {
			t.Errorf("Could not read data from the directory: '%s'. ", err)
		}

		for _, fl := range dir {
			if !strings.HasPrefix(fl.Name(), conf.Files[0].FileName) {
				continue
			}

			if !checker(conf.Files[0].Options.Rotation.FileSize, path.Join(env.Paths.SystemLocation, conf.Files[0].Path, fl.Name())) {
				t.Error("Invalid log file size value, rotation failed. ")
			}
		}
	}
}

func TestLogger_DPanic(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.DPanicLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.DPanic()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("DPanic() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.DebugLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Debug()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Debug() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.ErrorLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Error()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Error() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Fatal(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.FatalLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Fatal()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Error() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.InfoLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Info()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Info() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Panic(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.PanicLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Panic()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Panic() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantMsg Message
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantMsg: Message(&message{
				initiator: log.(*logger).initiator,
				text:      "",
				lvl:       zapcore.WarnLevel,
				fields:    nil,
				write:     log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			var gotMsg = l.Warn()

			if gotMsg.(*message).write == nil {
				t.Errorf("Message write function is nil. ")
			}

			gotMsg.(*message).write = nil
			tt.wantMsg.(*message).write = nil

			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("Warn() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestLogger_Copy(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	type args struct {
		initiator string
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Logger
	}{
		{
			name: "Case 1.",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal: &internal{
					conf: log.(*logger).internal.conf,
				},
			},
			args: args{
				initiator: "test",
			},
			want: &logger{
				initiator: "test",
				internal: &internal{
					conf: log.(*logger).internal.conf,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			if got := l.Copy(tt.args.initiator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Close(t *testing.T) {
	type fields struct {
		initiator string
		internal  *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(""); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Case 1. ",
			fields: fields{
				initiator: log.(*logger).initiator,
				internal:  log.(*logger).internal,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				initiator: tt.fields.initiator,
				internal:  tt.fields.internal,
			}

			if err := l.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	once = new(sync.Once)
	instance = nil
}
