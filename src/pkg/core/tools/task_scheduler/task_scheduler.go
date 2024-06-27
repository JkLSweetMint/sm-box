package task_scheduler

import (
	"context"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sync"
)

// scheduler - инструмент ядра системы для выполнения запланированных задач.
type scheduler struct {
	aggregate Aggregate

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

	var e Event

	switch v := t.(type) {
	case *BackgroundTask:
		{
			e = v.Event
		}
	case *ImmediateTask:
		{
			e = v.Event
		}
	default:
		{
			err = ErrInvalidEvent
			return
		}
	}

	if e <= minEvent || e >= maxEvent {
		err = ErrInvalidEvent
		return
	}

	s.aggregate.Add(t)

	return
}

// Call - вызов события.
func (s *scheduler) Call(e Event) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall(e)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var wg = new(sync.WaitGroup)

	for iter := s.aggregate.Iterator(e); iter.Has(); iter.Next() {
		var (
			task      = iter.Value()
			priority  uint8
			processes bool
		)

		if err != nil {
			break
		}

		if task == nil {
			continue
		}

		switch t := task.(type) {
		case *BackgroundTask:
			{
				if !processes {
					priority = t.Priority
				}

				switch {
				case t.Priority == priority:
					{
						env.Synchronization.WaitGroup.Add(1)

						go func() {
							defer func() {
								env.Synchronization.WaitGroup.Done()
							}()

							s.components.Logger.Info().
								Format("Task '%s' with event '%s' and priority '%d' started... ",
									t.Name,
									t.Event,
									t.Priority).Write()

							if err_ := task.Exec(context.Background()); err_ != nil {
								s.components.Logger.Error().
									Format("An error occurred during the execution of the task '%s' of event '%s' and priority '%d': '%s'. ",
										t.Name,
										t.Event,
										t.Priority,
										err_).Write()
								return
							}

							s.components.Logger.Info().
								Format("Task '%s' with event '%s' and priority '%d' completed. ",
									t.Name,
									t.Event,
									t.Priority).Write()
						}()
					}
				}
			}
		case *ImmediateTask:
			{
				if !processes {
					priority = t.Priority
					processes = true
				}

				switch {
				case t.Priority == priority:
					{
						wg.Add(1)
						env.Synchronization.WaitGroup.Add(1)

						go func() {
							defer func() {
								env.Synchronization.WaitGroup.Done()
								wg.Done()
							}()

							s.components.Logger.Info().
								Format("Task '%s' with event '%s' and priority '%d' started... ",
									t.Name,
									t.Event,
									t.Priority).Write()

							if err_ := task.Exec(context.Background()); err_ != nil {
								err = err_

								s.components.Logger.Error().
									Format("An error occurred during the execution of the task '%s' of event '%s' and priority '%d': '%s'. ",
										t.Name,
										t.Event,
										t.Priority,
										err_).Write()
								return
							}

							s.components.Logger.Info().
								Format("Task '%s' with event '%s' and priority '%d' completed. ",
									t.Name,
									t.Event,
									t.Priority).Write()
						}()
					}
				case t.Priority != priority:
					{
						wg.Wait()

						if err != nil {
							return
						}

						processes = false
						priority = 0
					}
				}
			}
		}
	}

	wg.Wait()

	if err != nil {
		return
	}

	return
}
