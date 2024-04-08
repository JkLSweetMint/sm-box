package mode

const (
	// Dev - режим разработки.
	Dev = iota + 1
	// Prod = боевой режим.
	Prod
)

// Mode - режим работы системы.
//
// Возможные режимы:
//  1. Dev  - режим разработки;
//  2. Prod - боевой режим;
//
// Стандартное значение режима системы "UNKNOWN".
type Mode int

// String - получение строкового представления режима работы системы.
func (mode Mode) String() (val string) {
	val = "UNKNOWN"

	switch mode {
	case Dev:
		val = "DEV"
	case Prod:
		val = "PROD"
	}

	return
}
