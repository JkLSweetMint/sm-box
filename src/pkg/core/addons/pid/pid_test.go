package pid

import "testing"

func TestNewFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Case 1",
			wantErr: false,
		},
		{
			name:    "Case 2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewFile(); (err != nil) != tt.wantErr {
				t.Errorf("NewFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Case 1",
			wantErr: false,
		},
		{
			name:    "Case 2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFile(); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
