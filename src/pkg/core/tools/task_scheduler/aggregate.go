package task_scheduler

import (
	"sort"
	"sync"
)

// Aggregate - описание коллекции для хранения задач планировщика.
type Aggregate interface {
	Iterator(te Event) Iterator
	Add(t Task)
}

// baseShelf - коллекция для хранения задач планировщика.
type baseShelf struct {
	Tasks []Task
	rwMx  *sync.RWMutex
}

// Iterator - создает и возвращает итератор по коллекции.
func (s *baseShelf) Iterator(e Event) (iter Iterator) {
	var copyShelf = &baseShelf{
		Tasks: make([]Task, 0),
		rwMx:  new(sync.RWMutex),
	}

	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	for _, t_ := range s.Tasks {
		switch t := t_.(type) {
		case *BackgroundTask:
			{
				if t.Event == e {
					copyShelf.Tasks = append(copyShelf.Tasks, t)
				}
			}
		case *ImmediateTask:
			{
				if t.Event == e {
					copyShelf.Tasks = append(copyShelf.Tasks, t)
				}
			}
		}
	}

	sort.SliceStable(copyShelf.Tasks, func(i, j int) bool {
		var p1, p2 uint8

		for index, t := range []Task{
			copyShelf.Tasks[i],
			copyShelf.Tasks[j],
		} {
			var p uint8

			switch v := t.(type) {
			case *BackgroundTask:
				{
					if v.Event == e {
						p = v.Priority
					}
				}
			case *ImmediateTask:
				{
					if v.Event == e {
						p = v.Priority
					}
				}
			}

			if index == 0 {
				p1 = p
			} else if index == 1 {
				p2 = p
			}
		}

		return p1 < p2
	})

	iter = &baseIterator{shelf: copyShelf}

	return
}

// Add - добавляет элемент в коллекцию.
func (s *baseShelf) Add(t Task) {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	s.Tasks = append(s.Tasks, t)
}
