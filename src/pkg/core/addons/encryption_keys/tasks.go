package encryption_keys

import (
	"context"
	"sm-box/pkg/core/tools/task_scheduler"
)

var (
	TaskInitEncryptionKeys = &task_scheduler.ImmediateTask{
		Name:     "Init encryption keys",
		Event:    task_scheduler.EventBeforeBoot,
		Priority: uint8(254),
		Func: func(ctx context.Context) (err error) {

			if err = Init(); err != nil {
				return
			}

			return
		},
	}
)
