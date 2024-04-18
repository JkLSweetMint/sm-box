package core

const (
	StateNil State = iota
	StateNew
	StateBooted
	StateServed
	StateOff
)

// State - состояние ядра системы.
// Возможные варианты:
//  0. StateNil    - "Nil";
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
type State int

// String - получение строкового представления состояния ядра системы.
func (index State) String() (v string) {
	switch index {
	case StateNil:
		v = "Nil"
	case StateNew:
		v = "New"
	case StateBooted:
		v = "Booted"
	case StateServed:
		v = "Served"
	case StateOff:
		v = "Off"
	}

	return
}
