package task_scheduler

import (
	"context"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// scheduler - инструмент ядра системы для выполнения запланированных задач.
type scheduler struct {
	aggregate aggregate
	channel   chan TaskType

	components *components
}

// components - компоненты инструмента ядра системы.
type components struct {
	Logger logger.Logger
}

// Register - регистрация задачи в планировщике.
func (s *scheduler) Register(t Task) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(t)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if t.Type <= minTaskType || t.Type >= maxTaskType {
		err = ErrInvalidTaskType
		return
	}

	s.aggregate.Add(t)

	return
}

// tracking - запуск фонового процесса для отслеживания сигнала для вызова задач из планировщика.
func (s *scheduler) tracking(ctx context.Context) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		defer func() { trc.FunctionCallFinished() }()
	}

For:
	for {
		select {
		case tt := <-s.channel:
			{
				for iter := s.aggregate.Iterator(tt); iter.Has(); iter.Next() {
					var task = iter.Value()

					if task == nil {
						continue
					}

					env.Synchronization.WaitGroup.Add(1)

					go func() {
						defer env.Synchronization.WaitGroup.Done()

						s.components.Logger.Info().
							Format("Task '%s' with type '%s' started... ", task.Name, task.Type)

						if err := task.Exec(context.Background()); err != nil {
							s.components.Logger.Error().
								Format("An error occurred during the execution of the task '%s' of type '%s': '%s'. ",
									task.Name,
									task.Type,
									err).Write()
							return
						}

						s.components.Logger.Info().
							Format("Task '%s' with type '%s' completed. ", task.Name, task.Type)
					}()
				}

				if tt == TaskAfterShutdown {
					break For
				}
			}
		}
	}

	return
}
