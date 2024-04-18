package logger

// ConfigTerminalLog - конфигурация терминала для компонента ведения журнала трессировки.
type ConfigTerminalLog struct {
	Levels  *ConfigTerminalLogLevels  `json:"levels"  yaml:"Levels"  xml:"Levels"`  // Конфигурация уровней журнала.
	Options *ConfigTerminalLogOptions `json:"options" yaml:"Options" xml:"Options"` // Конфигурации параметров.
}

// ConfigTerminalLogLevels - конфигурация уровней журнала для терминала.
type ConfigTerminalLogLevels struct {
	Info  *ConfigTerminalLogLevel `json:"info" yaml:"Info" xml:"Info"`        // Уровень "Info" журнала.
	Error *ConfigTerminalLogLevel `json:"error"   yaml:"Error"   xml:"Error"` // Уровень "Error" журнала.
}

// ConfigTerminalLogOptions - конфигурация параметров журнала для терминала.
type ConfigTerminalLogOptions struct {
	TimeFormat string `json:"time_format"  yaml:"TimeFormat"  xml:"time_format,attr"` // Формат времени.
}

// ConfigTerminalLogLevel - конфигурация уровня журнала для терминала.
type ConfigTerminalLogLevel struct {
	Enable  bool                           `json:"enable"  yaml:"Enable"  xml:"enable,attr"` // Включение.
	Options *ConfigTerminalLogLevelOptions `json:"options" yaml:"Options" xml:"Options"`     // Конфигурации параметров.
}

// ConfigTerminalLogLevelOptions - конфигурация параметров уровня журнала для терминала.
type ConfigTerminalLogLevelOptions struct {
	// Encoder - это не зависящий от формата интерфейс для всех маршалеров записей журнала.
	// Возможные варианты:
	//	1. "raw"  - в одну строку;
	// 	2. "json" - в json строку;
	Encoder string `json:"encoder" yaml:"Encoder" xml:"encoder,attr"`

	// Format - формат отображения.
	// 	1. "capital"         - большими буквами;
	// 	2. "capital_color"   - большими буквами с цветом;
	//	3. "lowercase"       - маленькими буквами;
	//	4. "lowercase_color" - маленькими буквами с цветом;
	Format string `json:"format" yaml:"Format" xml:"format,attr"`
}