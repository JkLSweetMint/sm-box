package pid

import (
	"context"
	"sm-box/pkg/core/tools/task_scheduler"
)

var (
	TaskCreatePIDFile = task_scheduler.Task{
		Name: "Create PID file",
		Type: task_scheduler.TaskBeforeBoot,
		Func: func(ctx context.Context) (err error) {

			if err = NewFile(); err != nil {
				return
			}

			return
		},
	}
	TaskRemovePIDFile = task_scheduler.Task{
		Name: "Remove PID file",
		Type: task_scheduler.TaskAfterShutdown,
		Func: func(ctx context.Context) (err error) {

			if err = RemoveFile(); err != nil {
				return
			}

			return
		},
	}
)
