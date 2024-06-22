package tracer

import (
	"sm-box/pkg/core/components/tracer/logger"
)

// Init - инициализация компонента ведения журнала трессировки вызовов функций/методов.
func Init() (err error) {
	var conf = new(Config)

	if err = conf.Read(); err != nil {
		return
	}

	if _, err = logger.New(conf.Logger); err != nil {
		return
	}

	config = conf

	return
}
