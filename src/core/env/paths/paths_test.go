package paths

import (
	"reflect"
	"testing"
)

func TestDev_Build(t *testing.T) {
	type fields struct {
		SystemLocation string
		Bin            string
		Etc            string
		Src            *struct {
			Path string
		}
		System *struct {
			Path string
		}
		Var *struct {
			Path string

			Temp  string
			Logs  string
			Data  string
			Cache string
			Lib   string
			Run   string

			Test *struct {
				Path string

				Data  string
				Cache string
			}
		}
	}

	type args struct {
		options BuildOptions
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dev
	}{
		{
			name:   "Case 1",
			fields: fields{},
			args: args{
				options: BuildOptions{
					ID: "test-id",
				},
			},
			want: &Dev{
				Bin: "/bin",
				Etc: "/var/test/cache/test-id/etc",
				Src: &struct {
					Path string
				}{
					Path: "/src",
				},
				System: &struct {
					Path string
				}{
					Path: "/system",
				},
				Var: &struct {
					Path string

					Temp  string
					Logs  string
					Data  string
					Cache string
					Lib   string
					Run   string
					Test  *struct {
						Path  string
						Data  string
						Cache string
					}
				}{
					Path: "/var/test/cache/test-id/var",

					Temp:  "/var/test/cache/test-id/var/temp",
					Logs:  "/var/test/cache/test-id/var/logs",
					Data:  "/var/data",
					Cache: "/var/test/cache/test-id/var/cache",
					Lib:   "/var/lib",
					Run:   "/var/test/cache/test-id/var/run",
					Test: &struct {
						Path string

						Data  string
						Cache string
					}{
						Path: "/var/test",

						Data:  "/var/test/data",
						Cache: "/var/test/cache",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Dev{
				SystemLocation: tt.fields.SystemLocation,
				Bin:            tt.fields.Bin,
				Etc:            tt.fields.Etc,
				Src:            tt.fields.Src,
				System:         tt.fields.System,
				Var:            tt.fields.Var,
			}

			got := storage.Build(tt.args.options)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProd_Build(t *testing.T) {
	type fields struct {
		SystemLocation string
		Bin            string
		Etc            string
		Src            *struct {
			Path string
		}
		System *struct {
			Path string
		}
		Var *struct {
			Path string

			Temp  string
			Logs  string
			Data  string
			Cache string
			Lib   string
			Run   string
		}
	}

	type args struct {
		options BuildOptions
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Prod
	}{
		{
			name:   "Case 1",
			fields: fields{},
			args:   args{},
			want: &Prod{
				Bin: "/bin",
				Etc: "/etc",
				Src: &struct {
					Path string
				}{
					Path: "/src",
				},
				System: &struct {
					Path string
				}{
					Path: "/system",
				},
				Var: &struct {
					Path string

					Temp  string
					Logs  string
					Data  string
					Cache string
					Lib   string
					Run   string
				}{
					Path: "/var",

					Temp:  "/var/temp",
					Logs:  "/var/logs",
					Data:  "/var/data",
					Cache: "/var/cache",
					Lib:   "/var/lib",
					Run:   "/var/run",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Prod{
				SystemLocation: tt.fields.SystemLocation,
				Bin:            tt.fields.Bin,
				Etc:            tt.fields.Etc,
				Src:            tt.fields.Src,
				System:         tt.fields.System,
				Var:            tt.fields.Var,
			}

			if got := storage.Build(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
