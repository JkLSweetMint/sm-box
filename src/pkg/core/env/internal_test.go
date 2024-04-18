package env

import "testing"

func Test_getSystemLocation(t *testing.T) {
	type args struct {
		testSystemLocation string
	}

	tests := []struct {
		name         string
		args         args
		wantLocation string
		wantErr      bool
	}{
		{
			name: "Case 1",
			args: args{
				testSystemLocation: testSystemLocation,
			},
			wantLocation: testSystemLocation,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocation, err := getSystemLocation(tt.args.testSystemLocation)

			if (err != nil) != tt.wantErr {
				t.Errorf("getSystemLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotLocation != tt.wantLocation {
				t.Errorf("getSystemLocation() gotLocation = %v, want %v", gotLocation, tt.wantLocation)
			}
		})
	}
}

func Test_initSystemDir(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Case 1",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initSystemDir(); (err != nil) != tt.wantErr {
				t.Errorf("initSystemDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
