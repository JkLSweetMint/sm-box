package mode

const (
	minMode = iota

	// Dev - режим разработки.
	Dev
	// Prod = боевой режим.
	Prod
)

var modeList = []string{
	"DEV",
	"PROD",
}

// Mode - режим работы системы.
//
// Возможные режимы:
//  1. Dev  - режим разработки;
//  2. Prod - боевой режим;
//
// Стандартное значение режима системы "UNKNOWN".
type Mode int

// String - получение строкового представления режима работы системы.
func (m Mode) String() (val string) {
	if m > minMode && int(m) <= len(modeList) {
		return modeList[m-1]
	}

	return "UNKNOWN"
}
