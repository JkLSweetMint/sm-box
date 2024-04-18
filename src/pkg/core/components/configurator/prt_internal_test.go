package configurator

import (
	"os"
	"path"
	"reflect"
	"sm-box/src/pkg/core/components/configurator/encoders"
	"sm-box/src/pkg/core/env"
	"testing"
)

func Test_privateConfigurator_Encoder(t *testing.T) {
	type args struct {
		encoder Encoder
	}

	type testCase[T any] struct {
		name string
		c    privateConfigurator[T]
		args args
		want Private[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: nil,
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 2",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.YamlEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 3",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.XmlEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 4",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.JsonEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 5",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
			args: args{
				encoder: nil,
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 6",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				encoder: new(encoders.YamlEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 7",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
			args: args{
				encoder: new(encoders.XmlEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 8",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.json",
			},
			args: args{
				encoder: new(encoders.JsonEncoder),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Encoder(tt.args.encoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_privateConfigurator_File(t *testing.T) {
	type args struct {
		dir      string
		filename string
	}

	type testCase[T any] struct {
		name string
		c    privateConfigurator[T]
		args args
		want Private[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 2",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "test.xml",
			},
		},
		{
			name: "Case 3",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 4",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "test.xml",
			},
		},

		{
			name: "Case 5",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 6",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
		},
		{
			name: "Case 7",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 8",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
		},

		{
			name: "Case 9",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 10",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 11",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 12",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
		},

		{
			name: "Case 13",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.json",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.json",
			},
		},
		{
			name: "Case 14",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.json",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.json",
			},
		},
		{
			name: "Case 15",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing1",
				filename: "test1.json",
			},
			args: args{
				dir:      "/testing",
				filename: "test.json",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.json",
			},
		},
		{
			name: "Case 16",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing1",
				filename: "test1.json",
			},
			args: args{
				dir:      "",
				filename: "test.json",
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.File(tt.args.dir, tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_privateConfigurator_Profile(t *testing.T) {
	type args struct {
		profile PrivateProfile
	}

	type testCase[T any] struct {
		name string
		c    privateConfigurator[T]
		args args
		want Private[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				profile: PrivateProfile{
					Encoder:  nil,
					Dir:      "/testing",
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 2",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				profile: PrivateProfile{
					Encoder:  new(encoders.XmlEncoder),
					Dir:      "/testing",
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 3",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing_json",
				filename: "test.json",
			},
			args: args{
				profile: PrivateProfile{
					Encoder:  new(encoders.XmlEncoder),
					Dir:      "/testing",
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 4",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PrivateProfile{
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 5",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PrivateProfile{
					Encoder: prtDefaultEncoder,
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
		},
		{
			name: "Case 6",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				profile: PrivateProfile{
					Encoder:  prtDefaultEncoder,
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 7",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PrivateProfile{
					Filename: "test.xml",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 8",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "test.xml",
			},
			args: args{
				profile: PrivateProfile{
					Dir: "/testing",
				},
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing",
				filename: "test.xml",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Profile(tt.args.profile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_privateConfigurator_Read(t *testing.T) {
	type testCase[T any] struct {
		name    string
		c       privateConfigurator[T]
		wantErr bool
	}

	// Подготовка
	{
		var dir = path.Join(env.Paths.SystemLocation, env.Paths.System.Path, "/testing/read")

		// Директория
		{
			if err := os.MkdirAll(dir, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}

			defer func() {
				if err := os.RemoveAll(path.Join(env.Paths.SystemLocation, env.Paths.System.Path, "/testing")); err != nil {
					t.Errorf("Failed to prepare data for testing: '%s'. ", err)
				}
			}()
		}

		// XML
		{
			if data, err := new(encoders.XmlEncoder).Encode(new(TestConfig).Default()); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			} else if err = os.WriteFile(path.Join(dir, "test.xml"), data, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}
		}

		// YAML
		{
			if data, err := new(encoders.YamlEncoder).Encode(new(TestConfig).Default()); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			} else if err = os.WriteFile(path.Join(dir, "test.yaml"), data, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}
		}

		// JSON
		{
			if data, err := new(encoders.JsonEncoder).Encode(new(TestConfig).Default()); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			} else if err = os.WriteFile(path.Join(dir, "test.json"), data, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}
		}
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 4",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: false,
		},

		{
			name: "Case 5",
			c: privateConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  prtDefaultEncoder,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 6",
			c: privateConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 7",
			c: privateConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 8",
			c: privateConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 9",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 10",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 11",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 12",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 13",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 14",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 15",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 16",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 17",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 18",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 19",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 20",
			c: privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Read(); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
