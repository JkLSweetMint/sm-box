package vars

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

// Dev - переменные окружения системы в режиме разработки.
type Dev struct {
	SystemName string    // Название системы.
	ID         string    // Идентификатор запуска
	PID        int       // Идентификатор процесса.
	LaunchTime time.Time // Время запуска системы.
}

// Prod - переменные окружения системы в боевом режиме.
type Prod struct {
	SystemName string    // Название системы.
	ID         string    // Идентификатор запуска
	PID        int       // Идентификатор процесса.
	LaunchTime time.Time // Время запуска системы.
}

// Build - построение хранилища.
func (storage *Dev) Build() *Dev {
	return &Dev{
		SystemName: "unknown",
		ID:         fmt.Sprintf("%d-%d", int(time.Now().UnixNano()), rand.Int()),
		PID:        syscall.Getpid(),
		LaunchTime: time.Now(),
	}
}

// Build - построение хранилища.
func (storage *Prod) Build() *Prod {
	return &Prod{
		SystemName: "unknown",
		ID:         fmt.Sprintf("%d-%d", int(time.Now().UnixNano()), rand.Int()),
		PID:        syscall.Getpid(),
		LaunchTime: time.Now(),
	}
}
