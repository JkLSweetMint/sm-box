package synchronization

import (
	"sync"
)

// Dev - инструменты синхронизации в режиме разработки.
type Dev struct {
	// WaitGroup - низкоуровневый инструмент для синхронизации горутин
	// без завершения которых система не должна завершить свою работу.
	WaitGroup *sync.WaitGroup
}

// Prod - инструменты синхронизации в боевом режиме.
type Prod struct {
	// WaitGroup - низкоуровневый инструмент для синхронизации горутин
	// без завершения которых система не должна завершить свою работу.
	WaitGroup *sync.WaitGroup
}

// Build - построение инструментов.
func (storage *Dev) Build() *Dev {
	return &Dev{
		WaitGroup: new(sync.WaitGroup),
	}
}

// Build - построение инструментов.
func (storage *Prod) Build() *Prod {
	return &Prod{
		WaitGroup: new(sync.WaitGroup),
	}
}
