package tracer

import (
	"testing"
	"time"
)

func Test_existLogLevel(t *testing.T) {
	type args struct {
		src  []Level
		list []Level
	}

	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Case 1",
			args: args{
				src:  []Level{},
				list: []Level{},
			},
			wantOk: false,
		},
		{
			name: "Case 2",
			args: args{
				src:  nil,
				list: nil,
			},
			wantOk: false,
		},
		{
			name: "Case 3",
			args: args{
				src: []Level{
					LevelMain,
					LevelDebug,
					LevelInternal,
					LevelEvent,
				},
				list: []Level{
					LevelInternal,
					LevelEvent,
				},
			},
			wantOk: true,
		},
		{
			name: "Case 4",
			args: args{
				src: []Level{
					LevelMain,
					LevelDebug,
					LevelInternal,
					LevelEvent,
				},
				list: []Level{
					LevelEvent,
				},
			},
			wantOk: true,
		},
		{
			name: "Case 5",
			args: args{
				src: []Level{
					LevelMain,
					LevelDebug,
					LevelInternal,
					LevelEvent,
				},
				list: []Level{
					LevelComponent,
				},
			},
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := existLogLevel(tt.args.src, tt.args.list); gotOk != tt.wantOk {
				t.Errorf("existLogLevel() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_formatWorkTime(t *testing.T) {
	type args struct {
		tm time.Duration
	}

	tests := []struct {
		name     string
		args     args
		wantText string
	}{
		{
			name: "Case 1",
			args: args{
				tm: time.Second,
			},
			wantText: "|        1s        |",
		},
		{
			name: "Case 2",
			args: args{
				tm: time.Millisecond * 7,
			},
			wantText: "|       7ms        |",
		},
		{
			name: "Case 3",
			args: args{
				tm: time.Millisecond,
			},
			wantText: "|       1ms        |",
		},
		{
			name: "Case 4",
			args: args{
				tm: time.Nanosecond,
			},
			wantText: "|       1ns        |",
		},
		{
			name: "Case 5",
			args: args{
				tm: time.Nanosecond * 100,
			},
			wantText: "|      100ns       |",
		},
		{
			name: "Case 6",
			args: args{
				tm: time.Microsecond,
			},
			wantText: "|        1µs       |",
		},
		{
			name: "Case 7",
			args: args{
				tm: time.Microsecond * 123,
			},
			wantText: "|       123µs      |",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := formatWorkTime(tt.args.tm); gotText != tt.wantText {
				t.Errorf("formatWorkTime() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func Test_functionName(t *testing.T) {
	tests := []struct {
		name     string
		wantName string
	}{
		{
			"Case 1",
			"testing.tRunner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotName := functionName(); gotName != tt.wantName {
				t.Errorf("functionName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
