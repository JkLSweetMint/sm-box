package configurator

import (
	"os"
	"path"
	"reflect"
	"sm-box/src/core/components/configurator/encoders"
	"sm-box/src/core/env"
	"testing"
)

func Test_publicConfigurator_Encoder(t *testing.T) {
	type args struct {
		encoder Encoder
	}

	type testCase[T any] struct {
		name string
		c    publicConfigurator[T]
		args args
		want Public[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: nil,
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.YamlEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.XmlEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				encoder: new(encoders.JsonEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				encoder: nil,
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				encoder: new(encoders.YamlEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				encoder: new(encoders.XmlEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				encoder: new(encoders.JsonEncoder),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.yaml",
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

func Test_publicConfigurator_File(t *testing.T) {
	type args struct {
		dir      string
		filename string
	}

	type testCase[T any] struct {
		name string
		c    publicConfigurator[T]
		args args
		want Public[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.yaml",
			},
		},

		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.xml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing1",
				filename: "test1.xml",
			},
			args: args{
				dir:      "",
				filename: "test.xml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
		},

		{
			name: "Case 9",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 10",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 11",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "/testing",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 12",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing1",
				filename: "test1.yaml",
			},
			args: args{
				dir:      "",
				filename: "test.yaml",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
		},

		{
			name: "Case 13",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "/testing",
				filename: "test.json",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.json",
			},
		},
		{
			name: "Case 14",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "",
			},
			args: args{
				dir:      "",
				filename: "test.json",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.json",
			},
		},
		{
			name: "Case 15",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing1",
				filename: "test1.json",
			},
			args: args{
				dir:      "/testing",
				filename: "test.json",
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing",
				filename: "test.json",
			},
		},
		{
			name: "Case 16",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing1",
				filename: "test1.json",
			},
			args: args{
				dir:      "",
				filename: "test.json",
			},
			want: &publicConfigurator[*TestConfig]{
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

func Test_publicConfigurator_Profile(t *testing.T) {
	type args struct {
		profile PublicProfile
	}

	type testCase[T any] struct {
		name string
		c    publicConfigurator[T]
		args args
		want Public[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				profile: PublicProfile{
					encoder:  nil,
					dir:      "/testing",
					filename: "test.yaml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
			args: args{
				profile: PublicProfile{
					encoder:  new(encoders.XmlEncoder),
					dir:      "/testing",
					filename: "test.xml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing_json",
				filename: "test.json",
			},
			args: args{
				profile: PublicProfile{
					encoder:  new(encoders.XmlEncoder),
					dir:      "/testing",
					filename: "test.xml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing",
				filename: "test.xml",
			},
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PublicProfile{
					filename: "test.yaml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PublicProfile{
					encoder: pbDefaultEncoder,
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing",
				filename: "test.yaml",
			},
			args: args{
				profile: PublicProfile{
					encoder:  pbDefaultEncoder,
					filename: "test.yaml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "",
			},
			args: args{
				profile: PublicProfile{
					filename: "test.yaml",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
			},
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.yaml",
			},
			args: args{
				profile: PublicProfile{
					dir: "/testing",
				},
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing",
				filename: "test.yaml",
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

func Test_publicConfigurator_Init(t *testing.T) {
	type testCase[T any] struct {
		name    string
		c       publicConfigurator[T]
		wantErr bool
	}

	// Подготовка
	{
		var dir = path.Join(env.Paths.SystemLocation, env.Paths.Etc, "/testing/init")

		// Директория
		{
			if err := os.MkdirAll(dir, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}

			defer func() {
				if err := os.RemoveAll(dir); err != nil {
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
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "test.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "test.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "test.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "test.json",
			},
			wantErr: false,
		},

		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 9",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 10",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 11",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 12",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 13",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 14",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 15",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 16",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},

		{
			name: "Case 17",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "test.1.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 18",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "test.2.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 19",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "test.3.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 20",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "test.4.json",
			},
			wantErr: false,
		},

		{
			name: "Case 21",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "test.5.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 22",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "test.6.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 23",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "test.7.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 24",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "test.8.json",
			},
			wantErr: true,
		},

		{
			name: "Case 25",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.9.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 26",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.10.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 27",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.11.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 28",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/init",
				filename: "test.12.json",
			},
			wantErr: true,
		},

		{
			name: "Case 29",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.13.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 30",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.14.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 31",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.15.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 32",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.16.json",
			},
			wantErr: false,
		},

		{
			name: "Case 33",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 34",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 35",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 36",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/init",
				filename: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_publicConfigurator_Read(t *testing.T) {
	type testCase[T any] struct {
		name    string
		c       publicConfigurator[T]
		wantErr bool
	}

	// Подготовка
	{
		var dir = path.Join(env.Paths.SystemLocation, env.Paths.Etc, "/testing/read")

		// Директория
		{
			if err := os.MkdirAll(dir, 0655); err != nil {
				t.Errorf("Failed to prepare data for testing: '%s'. ", err)
			}

			defer func() {
				if err := os.RemoveAll(dir); err != nil {
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
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: false,
		},

		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  pbDefaultEncoder,
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 9",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 10",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 11",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 12",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  nil,
				dir:      "/testing/read",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 13",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 14",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 15",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 16",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.json",
			},
			wantErr: true,
		},

		{
			name: "Case 17",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 18",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 19",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/read",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 20",
			c: publicConfigurator[*TestConfig]{
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

func Test_publicConfigurator_Write(t *testing.T) {
	type testCase[T any] struct {
		name    string
		c       publicConfigurator[T]
		wantErr bool
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/write",
				filename: "test.1.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/write",
				filename: "test.2.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/write",
				filename: "test.3.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 4",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/write",
				filename: "test.4.json",
			},
			wantErr: false,
		},

		{
			name: "Case 5",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  pbDefaultEncoder,
				dir:      "/testing/write",
				filename: "test.5.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 6",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/write",
				filename: "test.6.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 7",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/write",
				filename: "test.7.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 8",
			c: publicConfigurator[*TestConfig]{
				conf:     nil,
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/write",
				filename: "test.8.json",
			},
			wantErr: true,
		},

		{
			name: "Case 9",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/write",
				filename: "test.9.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 10",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/write",
				filename: "test.10.xml",
			},
			wantErr: true,
		},
		{
			name: "Case 11",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/write",
				filename: "test.11.yaml",
			},
			wantErr: true,
		},
		{
			name: "Case 12",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  nil,
				dir:      "/testing/write",
				filename: "test.12.json",
			},
			wantErr: true,
		},

		{
			name: "Case 13",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "test.13.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 14",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "",
				filename: "test.14.xml",
			},
			wantErr: false,
		},
		{
			name: "Case 15",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "",
				filename: "test.15.yaml",
			},
			wantErr: false,
		},
		{
			name: "Case 16",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "",
				filename: "test.16.json",
			},
			wantErr: false,
		},

		{
			name: "Case 17",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "/testing/write",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 18",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.XmlEncoder),
				dir:      "/testing/write",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 19",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.YamlEncoder),
				dir:      "/testing/write",
				filename: "",
			},
			wantErr: true,
		},
		{
			name: "Case 20",
			c: publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  new(encoders.JsonEncoder),
				dir:      "/testing/write",
				filename: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
