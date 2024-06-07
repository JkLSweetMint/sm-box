package types

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name    string
		s       Status
		wantStr string
	}{
		{
			name:    "Case 1",
			s:       StatusUnknown,
			wantStr: "unknown",
		},
		{
			name:    "Case 2",
			s:       StatusFailed,
			wantStr: "failed",
		},
		{
			name:    "Case 3",
			s:       StatusError,
			wantStr: "error",
		},
		{
			name:    "Case 4",
			s:       StatusFatal,
			wantStr: "fatal",
		},
		{
			name:    "Case 5",
			s:       -1,
			wantStr: "unknown",
		},
		{
			name:    "Case 6",
			s:       4,
			wantStr: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := tt.s.String(); gotStr != tt.wantStr {
				t.Errorf("String() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestParseStatus(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name  string
		args  args
		wantS Status
	}{
		{
			name: "Case 1",
			args: args{
				str: "unknown",
			},
			wantS: StatusUnknown,
		},
		{
			name: "Case 2",
			args: args{
				str: "failed",
			},
			wantS: StatusFailed,
		},
		{
			name: "Case 3",
			args: args{
				str: "error",
			},
			wantS: StatusError,
		},
		{
			name: "Case 4",
			args: args{
				str: "fatal",
			},
			wantS: StatusFatal,
		},
		{
			name: "Case 5",
			args: args{
				str: "",
			},
			wantS: StatusUnknown,
		},
		{
			name: "Case 6",
			args: args{
				str: "123",
			},
			wantS: StatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := ParseStatus(tt.args.str); gotS != tt.wantS {
				t.Errorf("ParseStatus() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
