package core

import (
	"sm-box/pkg/core/tools/closer"
	"sm-box/pkg/core/tools/task_scheduler"
)

// Tools - описание внутренних инструментов ядра системы.
type Tools interface {
	TaskScheduler() interface {
		Register(t task_scheduler.Task) (err error)
	}
}

// tools - внутренние инструменты ядра системы.
type tools struct {
	closer        closer.Closer
	taskScheduler task_scheduler.Scheduler
}

// TaskScheduler - получение инструмента планировщика задач.
func (t *tools) TaskScheduler() interface {
	Register(t task_scheduler.Task) (err error)
} {
	return t.taskScheduler
}
