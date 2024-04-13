package task_scheduler

import (
	"context"
	"sm-box/src/core/components/tracer"
)

const (
	minTaskType TaskType = iota

	// TaskBeforeNew - вызов после создания ядра системы.
	TaskBeforeNew

	// TaskAfterBoot - вызов перед запуском загрузки ядра системы.
	TaskAfterBoot
	// TaskBoot - вызов одновременно с запуском загрузки ядра системы.
	TaskBoot
	// TaskBeforeBoot - вызов после завершения загрузки ядра системы.
	TaskBeforeBoot

	// TaskAfterServe - вызов перед запуском обслуживания системы ядром.
	TaskAfterServe
	// TaskServe - вызов одновременно c запуском обслуживания системы ядром.
	TaskServe
	// TaskBeforeServe - вызов после запуском обслуживания системы ядром.
	TaskBeforeServe

	// TaskAfterShutdown - вызов перед завершением обслуживания системы ядром.
	TaskAfterShutdown
	// TaskShutdown - вызов одновременно c завершением обслуживания системы ядром.
	TaskShutdown
	// TaskBeforeShutdown - вызов после завершения обслуживания системы ядром.
	TaskBeforeShutdown

	maxTaskType
)

var allTaskTypesString = [...]string{
	"BeforeNew",

	"AfterBoot",
	"Boot",
	"BeforeBoot",

	"AfterServe",
	"Serve",
	"BeforeServe",

	"AfterShutdown",
	"Boot",
	"BeforeBoot",
}

// TaskType - тип задачи.
type TaskType int

// TaskFunc - функция задачи.
type TaskFunc func(ctx context.Context) (err error)

// Task - задача планировщика.
type Task struct {
	Name string
	Type TaskType
	Func TaskFunc
}

// Exec - запуск выполнения задачи планировщика.
func (t *Task) Exec(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished()
	}

	return t.Func(ctx)
}

// String - получение строкового представления типа задачи системы.
func (e TaskType) String() (val string) {
	if len(allTaskTypesString) >= int(e) {
		return allTaskTypesString[e-1]
	}

	return
}
