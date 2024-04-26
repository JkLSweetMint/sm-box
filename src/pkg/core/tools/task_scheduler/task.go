package task_scheduler

import (
	"context"
	"sm-box/pkg/core/components/tracer"
)

const (
	minTaskType TaskType = iota

	// TaskBeforeNew - вызов после создания ядра системы.
	TaskBeforeNew

	// TaskBeforeBoot - вызов перед запуском загрузки ядра системы.
	TaskBeforeBoot
	// TaskBoot - вызов одновременно с запуском загрузки ядра системы.
	TaskBoot
	// TaskAfterBoot - вызов после завершения загрузки ядра системы.
	TaskAfterBoot

	// TaskBeforeServe - вызов перед запуском обслуживания системы ядром.
	TaskBeforeServe
	// TaskServe - вызов одновременно c запуском обслуживания системы ядром.
	TaskServe
	// TaskAfterServe - вызов после запуском обслуживания системы ядром.
	TaskAfterServe

	// TaskBeforeShutdown - вызов перед завершением обслуживания системы ядром.
	TaskBeforeShutdown
	// TaskShutdown - вызов одновременно c завершением обслуживания системы ядром.
	TaskShutdown
	// TaskAfterShutdown - вызов после завершения обслуживания системы ядром.
	TaskAfterShutdown

	maxTaskType
)

var taskTypesList = [...]string{
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
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return t.Func(ctx)
}

// String - получение строкового представления типа задачи системы.
func (e TaskType) String() (val string) {
	if e > minTaskType && int(e) <= len(taskTypesList) {
		return taskTypesList[e-1]
	}

	return
}
