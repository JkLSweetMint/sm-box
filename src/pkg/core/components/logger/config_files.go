package logger

import (
	"sm-box/pkg/tools/size"
)

// ConfigFilesLog - конфигурация файлов для компонента ведения журнала системы.
type ConfigFilesLog []*ConfigFileLog

// ConfigFileLog - конфигурация файла для компонента ведения журнала системы.
type ConfigFileLog struct {
	FileName string `json:"file_name" yaml:"FileName" xml:"file_name,attr"` // Название файла.
	Path     string `json:"path"      yaml:"Path"     xml:"path,attr"`      // Путь к файлу.

	Levels  *ConfigFileLogLevels       `json:"levels"  yaml:"Levels"  xml:"Levels"`  // Конфигурация уровней журнала.
	Options *ConfigFilesLogFileOptions `json:"options" yaml:"Options" xml:"Options"` // Конфигурации параметров.
}

// ConfigFileLogLevels - конфигурация уровней журнала для файла.
type ConfigFileLogLevels struct {
	Debug  *ConfigFileLogLevel `json:"debug"  yaml:"Debug"  xml:"Debug"`  // Уровень "Debug" журнала.
	Info   *ConfigFileLogLevel `json:"info"   yaml:"Info"   xml:"Info"`   // Уровень "Info" журнала.
	Warn   *ConfigFileLogLevel `json:"warn"   yaml:"Warn"   xml:"Warn"`   // Уровень "Warn" журнала.
	Error  *ConfigFileLogLevel `json:"error"  yaml:"Error"  xml:"Error"`  // Уровень "Error" журнала.
	DPanic *ConfigFileLogLevel `json:"dpanic" yaml:"DPanic" xml:"DPanic"` // Уровень "DPanic" журнала.
	Panic  *ConfigFileLogLevel `json:"panic"  yaml:"Panic"  xml:"Panic"`  // Уровень "Panic" журнала.
	Fatal  *ConfigFileLogLevel `json:"fatal"  yaml:"Fatal"  xml:"Fatal"`  // Уровень "Fatal" журнала.
}

// ConfigFileLogLevel - конфигурация уровня журнала для файла.
type ConfigFileLogLevel struct {
	Enable  bool                        `json:"enable"  yaml:"Enable"  xml:"enable,attr"` // Включение.
	Options *ConfigFilesLogLevelOptions `json:"options" yaml:"Options" xml:"Options"`     // Конфигурации параметров.
}

// ConfigFilesLogLevelOptions - конфигурация параметров уровня журнала для файла.
type ConfigFilesLogLevelOptions struct {
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

	// EnableTrace - включение трессировки вызова.
	EnableTrace bool `json:"enable_trace" yaml:"EnableTrace" xml:"enable_trace,attr"`
}

// ConfigFilesLogFileOptions - конфигурация параметров журнала для файла.
type ConfigFilesLogFileOptions struct {
	Rotation   *ConfigFilesLogFileOptionRotation `json:"rotation" yaml:"Rotation" xml:"Rotation"`                // Конфигурация ротации файла.
	TimeFormat string                            `json:"time_format"  yaml:"TimeFormat"  xml:"time_format,attr"` // Формат времени.
}

// ConfigFilesLogFileOptionRotation - конфигурация параметра ротации журнала для файла.
type ConfigFilesLogFileOptionRotation struct {
	Enable   bool      `json:"enable"    yaml:"Enable"   xml:"enable,attr"`    // Включение.
	FileSize size.Size `json:"file_size" yaml:"FileSize" xml:"file_size,attr"` // Макс. размер файла.
}
