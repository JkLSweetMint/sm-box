package task_scheduler

import (
	"sync"
)

// aggregate - описание коллекции для хранения задач планировщика.
type aggregate interface {
	Iterator(tt TaskType) iterator
	Add(t Task)
}

// baseShelf - коллекция для хранения задач планировщика.
type baseShelf struct {
	Tasks []*Task
	rwMx  *sync.RWMutex
}

// Iterator - создает и возвращает итератор по коллекции.
func (s *baseShelf) Iterator(tt TaskType) (iter iterator) {
	var copyShelf = &baseShelf{
		Tasks: make([]*Task, 0),
		rwMx:  new(sync.RWMutex),
	}

	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	for _, e := range s.Tasks {
		if e.Type == tt {
			copyShelf.Tasks = append(copyShelf.Tasks, e)
		}
	}

	iter = &baseIterator{shelf: copyShelf}

	return
}

// Add - добавляет элемент в коллекцию.
func (s *baseShelf) Add(t Task) {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	s.Tasks = append(s.Tasks, &t)
}
