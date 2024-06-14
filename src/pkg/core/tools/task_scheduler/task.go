package task_scheduler

import (
	"context"
	"sm-box/pkg/core/components/tracer"
)

const (
	minEvent Event = iota

	// EventBeforeNew - вызов после создания ядра системы.
	EventBeforeNew

	// EventBeforeBoot - вызов перед запуском загрузки ядра системы.
	EventBeforeBoot
	// EventBoot - вызов одновременно с запуском загрузки ядра системы.
	EventBoot
	// EventAfterBoot - вызов после завершения загрузки ядра системы.
	EventAfterBoot

	// EventBeforeServe - вызов перед запуском обслуживания системы ядром.
	EventBeforeServe
	// EventServe - вызов одновременно c запуском обслуживания системы ядром.
	EventServe
	// EventAfterServe - вызов после запуском обслуживания системы ядром.
	EventAfterServe

	// EventBeforeShutdown - вызов перед завершением обслуживания системы ядром.
	EventBeforeShutdown
	// EventShutdown - вызов одновременно c завершением обслуживания системы ядром.
	EventShutdown
	// EventAfterShutdown - вызов после завершения обслуживания системы ядром.
	EventAfterShutdown

	maxEvent
)

var taskEventList = [...]string{
	"BeforeNew",

	"BeforeBoot",
	"Boot",
	"AfterBoot",

	"BeforeServe",
	"Serve",
	"AfterServe",

	"BeforeShutdown",
	"Shutdown",
	"AfterShutdown",
}

// Event - событие.
type Event int

// TaskFunc - функция задачи.
type TaskFunc func(ctx context.Context) (err error)

// Task - задача планировщика.
type Task interface {
	Exec(ctx context.Context) (err error)
}

// BackgroundTask - фоновая задача планировщика.
type BackgroundTask struct {
	Name     string
	Event    Event
	Priority uint8
	Func     TaskFunc
}

// ImmediateTask - задача планировщика.
type ImmediateTask struct {
	Name     string
	Event    Event
	Priority uint8
	Func     TaskFunc
}

// Exec - запуск выполнения задачи.
func (t *BackgroundTask) Exec(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return t.Func(ctx)
}

// Exec - запуск выполнения задачи.
func (t *ImmediateTask) Exec(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return t.Func(ctx)
}

// String - получение строкового представления типа задачи системы.
func (e Event) String() (val string) {
	if e > minEvent && int(e) <= len(taskEventList) {
		return taskEventList[e-1]
	}

	return
}
