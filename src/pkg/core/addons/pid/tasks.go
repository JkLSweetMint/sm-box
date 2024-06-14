package pid

import (
	"context"
	"sm-box/pkg/core/tools/task_scheduler"
)

var (
	TaskCreatePIDFile = &task_scheduler.ImmediateTask{
		Name:     "Create PID file",
		Event:    task_scheduler.EventBeforeBoot,
		Priority: uint8(255),
		Func: func(ctx context.Context) (err error) {

			if err = NewFile(); err != nil {
				return
			}

			return
		},
	}
	TaskRemovePIDFile = &task_scheduler.ImmediateTask{
		Name:     "Remove PID file",
		Event:    task_scheduler.EventAfterShutdown,
		Priority: uint8(255),
		Func: func(ctx context.Context) (err error) {

			if err = RemoveFile(); err != nil {
				return
			}

			return
		},
	}
)
