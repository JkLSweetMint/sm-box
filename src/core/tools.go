package core

import (
	"sm-box/src/core/tools/closer"
	"sm-box/src/core/tools/task_scheduler"
)

// tools - внутренние инструменты ядра системы.
type tools struct {
	closer        closer.Closer
	taskScheduler task_scheduler.Scheduler
}

// TaskScheduler - получение инструмента планировщика задач.
func (t *tools) TaskScheduler() task_scheduler.Scheduler {
	return t.taskScheduler
}
