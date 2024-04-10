package encoders

type TestConfig struct {
	Elements []*TestConfigElement `json:"elements" yaml:"Elements" xml:"Elements>Element"`
	Date     string               `json:"date"     yaml:"Date"     xml:"date,attr"`
}

type TestConfigElement struct {
	Value string `json:"value" yaml:"Value" xml:"value,attr"`
}

func (conf *TestConfig) FillEmptyFields() *TestConfig {
	if conf.Elements == nil {
		conf.Elements = make([]*TestConfigElement, 0)
	}

	return conf
}

func (conf *TestConfig) Default() *TestConfig {
	conf.Elements = []*TestConfigElement{
		{Value: "1"},
		{Value: "2"},
		{Value: "3"},
		{Value: "4"},
		{Value: "5"},
	}

	conf.Date = "0001-01-01"

	return conf
}

func (conf *TestConfig) Validate() (err error) {
	return
}
