package types

import "testing"

func TestErrorType_String(t *testing.T) {
	tests := []struct {
		name    string
		t       ErrorType
		wantStr string
	}{
		{
			name:    "Case 1",
			t:       TypeUnknown,
			wantStr: "unknown",
		},
		{
			name:    "Case 2",
			t:       TypeSystem,
			wantStr: "system",
		},
		{
			name:    "Case 3",
			t:       -1,
			wantStr: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := tt.t.String(); gotStr != tt.wantStr {
				t.Errorf("String() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestParseErrorType(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name  string
		args  args
		wantT ErrorType
	}{
		{
			name: "Case 1",
			args: args{
				str: "unknown",
			},
			wantT: TypeUnknown,
		},
		{
			name: "Case 2",
			args: args{
				str: "system",
			},
			wantT: TypeSystem,
		},
		{
			name: "Case 3",
			args: args{
				str: "",
			},
			wantT: TypeUnknown,
		},
		{
			name: "Case 4",
			args: args{
				str: "123",
			},
			wantT: TypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := ParseErrorType(tt.args.str); gotT != tt.wantT {
				t.Errorf("ParseErrorType() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
