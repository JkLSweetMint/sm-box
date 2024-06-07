package types

const (
	StatusUnknown Status = iota
	StatusFailed
	StatusError
	StatusFatal
)

var statusList = [...]string{
	StatusUnknown: "unknown",
	StatusFailed:  "failed",
	StatusError:   "error",
	StatusFatal:   "fatal",
}

// Status - статус ошибки.
type Status int

// String - получение строкового представления статуса ошибки.
func (s Status) String() (str string) {
	if s >= StatusUnknown && int(s) < len(statusList) {
		return statusList[s]
	}

	return statusList[StatusUnknown]
}

// ParseStatus - парсинг статуса ошибки из строки.
func ParseStatus(str string) (s Status) {
	s = StatusUnknown

	for i, s_ := range statusList {
		if s_ == str {
			s = Status(i)
			break
		}
	}

	return
}
