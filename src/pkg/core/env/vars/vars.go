package vars

import (
	"crypto/rsa"
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

	// Ключи шифрования системы.
	EncryptionKeys *struct {
		Private *rsa.PrivateKey // Приватный ключ.
		Public  *rsa.PublicKey  // Публичный ключ.
	}
}

// Prod - переменные окружения системы в боевом режиме.
type Prod struct {
	SystemName string    // Название системы.
	ID         string    // Идентификатор запуска
	PID        int       // Идентификатор процесса.
	LaunchTime time.Time // Время запуска системы.

	// Ключи шифрования системы.
	EncryptionKeys *struct {
		Private *rsa.PrivateKey // Приватный ключ.
		Public  *rsa.PublicKey  // Публичный ключ.
	}
}

// Build - построение хранилища.
func (storage *Dev) Build() *Dev {
	return &Dev{
		SystemName: "unknown",
		ID:         fmt.Sprintf("%d-%d", int(time.Now().UnixNano()), rand.Int()),
		PID:        syscall.Getpid(),
		LaunchTime: time.Now(),

		EncryptionKeys: new(struct {
			Private *rsa.PrivateKey
			Public  *rsa.PublicKey
		}),
	}
}

// Build - построение хранилища.
func (storage *Prod) Build() *Prod {
	return &Prod{
		SystemName: "unknown",
		ID:         fmt.Sprintf("%d-%d", int(time.Now().UnixNano()), rand.Int()),
		PID:        syscall.Getpid(),
		LaunchTime: time.Now(),

		EncryptionKeys: new(struct {
			Private *rsa.PrivateKey
			Public  *rsa.PublicKey
		}),
	}
}
