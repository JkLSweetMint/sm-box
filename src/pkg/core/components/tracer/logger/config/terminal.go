package config

// TerminalLog - конфигурация терминала для компонента ведения журнала трессировки.
type TerminalLog struct {
	Levels  *TerminalLogLevels  `json:"levels"  yaml:"Levels"  xml:"Levels"`  // Конфигурация уровней журнала.
	Options *TerminalLogOptions `json:"options" yaml:"Options" xml:"Options"` // Конфигурации параметров.
}

// TerminalLogLevels - конфигурация уровней журнала для терминала.
type TerminalLogLevels struct {
	Info  *TerminalLogLevel `json:"info" yaml:"Info" xml:"Info"`        // Уровень "Info" журнала.
	Error *TerminalLogLevel `json:"error"   yaml:"Error"   xml:"Error"` // Уровень "Error" журнала.
}

// TerminalLogOptions - конфигурация параметров журнала для терминала.
type TerminalLogOptions struct {
	TimeFormat string `json:"time_format"  yaml:"TimeFormat"  xml:"time_format,attr"` // Формат времени.
}

// TerminalLogLevel - конфигурация уровня журнала для терминала.
type TerminalLogLevel struct {
	Enable  bool                     `json:"enable"  yaml:"Enable"  xml:"enable,attr"` // Включение.
	Options *TerminalLogLevelOptions `json:"options" yaml:"Options" xml:"Options"`     // Конфигурации параметров.
}

// TerminalLogLevelOptions - конфигурация параметров уровня журнала для терминала.
type TerminalLogLevelOptions struct {
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
