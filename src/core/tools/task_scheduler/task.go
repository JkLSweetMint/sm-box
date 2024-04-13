package task_scheduler

import (
	"context"
	"sm-box/src/core/components/tracer"
)

const (
	minTaskType TaskType = iota

	// TaskBeforeNew - вызов после создания ядра системы.
	TaskBeforeNew

	// TaskBeforeBoot - вызов после завершения загрузки ядра системы.
	TaskBeforeBoot
	// TaskBoot - вызов одновременно с запуском загрузки ядра системы.
	TaskBoot
	// TaskAfterBoot - вызов перед запуском загрузки ядра системы.
	TaskAfterBoot

	// TaskBeforeServe - вызов после запуском обслуживания системы ядром.
	TaskBeforeServe
	// TaskServe - вызов одновременно c запуском обслуживания системы ядром.
	TaskServe
	// TaskAfterServe - вызов перед запуском обслуживания системы ядром.
	TaskAfterServe

	// TaskBeforeShutdown - вызов после завершения обслуживания системы ядром.
	TaskBeforeShutdown
	// TaskShutdown - вызов одновременно c завершением обслуживания системы ядром.
	TaskShutdown
	// TaskAfterShutdown - вызов перед завершением обслуживания системы ядром.
	TaskAfterShutdown

	maxTaskType
)

var allTaskTypesString = [...]string{
	"BeforeNew",

	"BeforeBoot",
	"Boot",
	"AfterBoot",

	"BeforeServe",
	"Serve",
	"AfterServe",

	"BeforeBoot",
	"Boot",
	"AfterShutdown",
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
