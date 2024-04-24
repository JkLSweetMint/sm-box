package core

const (
	StateNil State = iota
	StateNew
	StateBooted
	StateServed
	StateOff
)

var statesList = [...]string{
	"Nil",
	"New",
	"Booted",
	"Served",
	"Off",
}

// State - состояние ядра системы.
// Возможные варианты:
//  0. StateNil    - "Nil";
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
type State int

// String - получение строкового представления состояния ядра системы.
func (i State) String() (v string) {
	if i >= StateNil && int(i) < len(statesList) {
		return statesList[i]
	}

	return
}
