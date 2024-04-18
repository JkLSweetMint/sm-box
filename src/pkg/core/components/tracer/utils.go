package tracer

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// functionName - получение наименования функции/метода из которого был вызван компонент.
func functionName() (name string) {
	var fpcs = make([]uintptr, 1)

	if n := runtime.Callers(3, fpcs); n == 0 {
		name = "unknown"
		return
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		name = "unknown"
		return
	}

	name = caller.Name()

	return
}

// existLogLevel - проверка пересечения уровней журнала с исходным.
func existLogLevel(src, list []Level) (ok bool) {
	if len(src) > 0 && len(list) > 0 {
		for _, el := range list {
			for _, e := range src {
				if ok = el == e; ok {
					return
				}
			}
		}
	}

	return
}

// formatWorkTime - форматирование времени работы функции/метода для вывода в журнал.
func formatWorkTime(tm time.Duration) (text string) {
	var count = 15

	text = fmt.Sprintf("%s", tm.String())

	if strings.Contains(text, "µs") {
		count++
	}

	for i := len(text); i <= count; i++ {
		switch i % 2 {
		case 0:
			text = " " + text
		case 1:
			text = text + " "
		}
	}

	text = fmt.Sprintf("| %s |", text)

	return
}
