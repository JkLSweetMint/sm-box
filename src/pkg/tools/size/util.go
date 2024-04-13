package size

import (
	"strconv"
	"strings"
)

// Size - размер.
type Size string

// Int64 - преобразовать размер в число.
func (size Size) Int64() (s int64) {
	var (
		value     = strings.ToUpper(strings.TrimSpace(string(size)))
		type_     string
		numberStr = string(size)
		number    float64
		err       error
	)

	if _, err = strconv.ParseFloat(string(size), 64); err != nil {
		type_ = value[len(value)-2:]
		numberStr = value[:len(value)-2]
		err = nil
	}

	if number, err = strconv.ParseFloat(numberStr, 64); err != nil {
		return
	}

	switch type_ {
	case "KB":
		s = int64(number * 1024)
	case "MB":
		s = int64(number * 1024 * 1024)
	case "GB":
		s = int64(number * 1024 * 1024 * 1024)
	case "TB":
		s = int64(number * 1024 * 1024 * 1024 * 1024)
	default:
		s = int64(number)
	}

	return
}
