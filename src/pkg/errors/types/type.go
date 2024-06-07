package types

const (
	TypeUnknown ErrorType = iota
	TypeSystem
)

var errorTypesList = [...]string{
	TypeUnknown: "unknown",
	TypeSystem:  "system",
}

type (
	// ErrorType - тип ошибки.
	ErrorType int
)

// String - получение строкового представления типа ошибки.
func (t ErrorType) String() (str string) {
	if t >= TypeUnknown && int(t) < len(statusList) {
		return errorTypesList[t]
	}

	return errorTypesList[TypeUnknown]
}

// ParseErrorType - парсинг типа ошибки из строки.
func ParseErrorType(str string) (t ErrorType) {
	t = TypeUnknown

	for i, t_ := range errorTypesList {
		if t_ == str {
			t = ErrorType(i)
			break
		}
	}

	return
}
