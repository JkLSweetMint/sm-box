package configurator

// configurator - внутренняя реализация диспетчера конфигураций.
type configurator[T any] struct {
	conf Config[T]
}

// Public - получение диспетчера публичных конфигураций.
func (c *configurator[T]) Public() Public[T] {
	return &publicConfigurator[T]{
		conf:    c.conf,
		encoder: pbDefaultEncoder,
	}
}

// Private - получение диспетчера приватных конфигураций.
func (c *configurator[T]) Private() Private[T] {
	return &privateConfigurator[T]{
		conf:    c.conf,
		encoder: prtDefaultEncoder,
	}
}
