package config

import (
	"sm-box/pkg/tools/size"
)

// FilesLog - конфигурация файлов для компонента ведения журнала системы.
type FilesLog []*FileLog

// FileLog - конфигурация файла для компонента ведения журнала системы.
type FileLog struct {
	FileName string `json:"file_name" yaml:"FileName" xml:"file_name,attr"` // Название файла.
	Path     string `json:"path"      yaml:"Path"     xml:"path,attr"`      // Путь к файлу.

	Levels  *FileLogLevels       `json:"levels"  yaml:"Levels"  xml:"Levels"`  // Конфигурация уровней журнала.
	Options *FilesLogFileOptions `json:"options" yaml:"Options" xml:"Options"` // Конфигурации параметров.
}

// FileLogLevels - конфигурация уровней журнала для файла.
type FileLogLevels struct {
	Debug  *FileLogLevel `json:"debug"  yaml:"Debug"  xml:"Debug"`  // Уровень "Debug" журнала.
	Info   *FileLogLevel `json:"info"   yaml:"Info"   xml:"Info"`   // Уровень "Info" журнала.
	Warn   *FileLogLevel `json:"warn"   yaml:"Warn"   xml:"Warn"`   // Уровень "Warn" журнала.
	Error  *FileLogLevel `json:"error"  yaml:"Error"  xml:"Error"`  // Уровень "Error" журнала.
	DPanic *FileLogLevel `json:"dpanic" yaml:"DPanic" xml:"DPanic"` // Уровень "DPanic" журнала.
	Panic  *FileLogLevel `json:"panic"  yaml:"Panic"  xml:"Panic"`  // Уровень "Panic" журнала.
	Fatal  *FileLogLevel `json:"fatal"  yaml:"Fatal"  xml:"Fatal"`  // Уровень "Fatal" журнала.
}

// FileLogLevel - конфигурация уровня журнала для файла.
type FileLogLevel struct {
	Enable  bool                  `json:"enable"  yaml:"Enable"  xml:"enable,attr"` // Включение.
	Options *FilesLogLevelOptions `json:"options" yaml:"Options" xml:"Options"`     // Конфигурации параметров.
}

// FilesLogLevelOptions - конфигурация параметров уровня журнала для файла.
type FilesLogLevelOptions struct {
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

// FilesLogFileOptions - конфигурация параметров журнала для файла.
type FilesLogFileOptions struct {
	Rotation   *FilesLogFileOptionRotation `json:"rotation" yaml:"Rotation" xml:"Rotation"`                // Конфигурация ротации файла.
	TimeFormat string                      `json:"time_format"  yaml:"TimeFormat"  xml:"time_format,attr"` // Формат времени.
}

// FilesLogFileOptionRotation - конфигурация параметра ротации журнала для файла.
type FilesLogFileOptionRotation struct {
	Enable   bool      `json:"enable"    yaml:"Enable"   xml:"enable,attr"`    // Включение.
	FileSize size.Size `json:"file_size" yaml:"FileSize" xml:"file_size,attr"` // Макс. размер файла.
}
