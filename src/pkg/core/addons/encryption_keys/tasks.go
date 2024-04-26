package encryption_keys

import (
	"context"
	"sm-box/pkg/core/tools/task_scheduler"
)

var (
	TaskInitEncryptionKeys = task_scheduler.Task{
		Name: "Init encryption keys",
		Type: task_scheduler.TaskBeforeBoot,
		Func: func(ctx context.Context) (err error) {

			if err = Init(); err != nil {
				return
			}

			return
		},
	}
)
