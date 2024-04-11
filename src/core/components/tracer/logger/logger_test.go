package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"reflect"
	"sm-box/src/core/env"
	"sm-box/src/pkg/utils/size"
	"strings"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name    string
		wantLog Logger
		wantErr bool
	}{
		{
			name:    "Case 1. ",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(); (err != nil) != tt.wantErr {
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

		if log, err = newLogger(conf); err != nil {
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

func TestLogger_Error(t *testing.T) {
	type fields struct {
		internal *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(); err != nil {
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
				internal: log.(*logger).internal,
			},
			wantMsg: Message(&message{
				text:   "",
				lvl:    zapcore.ErrorLevel,
				fields: nil,
				write:  log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				internal: tt.fields.internal,
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

func TestLogger_Info(t *testing.T) {
	type fields struct {
		internal *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(); err != nil {
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
				internal: log.(*logger).internal,
			},
			wantMsg: Message(&message{
				text:   "",
				lvl:    zapcore.InfoLevel,
				fields: nil,
				write:  log.(*logger).instance.Log,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				internal: tt.fields.internal,
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

func TestLogger_Copy(t *testing.T) {
	type fields struct {
		internal *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(); err != nil {
			t.Errorf("An error occurred while creating the logging component: '%s'. ", err)
		}
	}

	tests := []struct {
		name   string
		fields fields
		want   Logger
	}{
		{
			name: "Case 1.",
			fields: fields{
				internal: &internal{
					conf: log.(*logger).internal.conf,
				},
			},
			want: &logger{
				internal: &internal{
					conf: log.(*logger).internal.conf,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				internal: tt.fields.internal,
			}

			if got := l.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Close(t *testing.T) {
	type fields struct {
		internal *internal
	}

	var log Logger

	// Build logger
	{
		var err error

		if log, err = New(); err != nil {
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
				internal: log.(*logger).internal,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l = &logger{
				internal: tt.fields.internal,
			}

			if err := l.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	once = new(sync.Once)
	instance = nil
}
