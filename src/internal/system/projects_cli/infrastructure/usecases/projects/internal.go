package projects

import (
	"context"
	"fmt"
	"os"
	"path"
	"sm-box/internal/common/entities"
	error_list "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
)

// remove - удаление проекта.
func (usecase *UseCase) remove(ctx context.Context, project *entities.Project) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCaseInternal)

		trace.FunctionCall(ctx, project)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	// Удаление из системной базы данных
	{
		if err := usecase.repositories.Projects.RemoveByUUID(ctx, project.UUID.String()); err != nil {
			cErr = error_list.FailedRemoveProject()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The project could not be deleted from the database: '%s'. ", cErr).Write()
			return
		}
	}

	// Удаление базы данных проекта
	{
		if err := os.Remove(path.Join(env.Paths.SystemLocation, env.Paths.Var.Lib.Projects, fmt.Sprintf("%s.db", project.UUID.String()))); err != nil {
			cErr = error_list.FailedRemoveProject()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The project database could not be deleted: '%s'. ", err).Write()
			return
		}
	}

	return
}
