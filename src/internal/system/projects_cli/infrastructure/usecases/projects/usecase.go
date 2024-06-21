package projects

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	g_uuid "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"path"
	"sm-box/internal/common/entities"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/internal/system/projects_cli/embed"
	repository_projects "sm-box/internal/system/projects_cli/infrastructure/repositories/projects"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
	c_errors "sm-box/pkg/errors"
	"strconv"
)

// UseCase - логика для управления проектами.
type UseCase struct {
	repositories *repositories
	components   *components

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Projects interface {
		Create(ctx context.Context, uuid, name, description, version string) (id types.ID, err error)
		GetAll(ctx context.Context) (projects []*entities.Project, err error)
		GetByID(ctx context.Context, id types.ID) (project *entities.Project, err error)
		GetByUUID(ctx context.Context, id string) (project *entities.Project, err error)
		RemoveByUUID(ctx context.Context, uuid string) (err error)

		GetEnvByID(ctx context.Context, id types.ID) (env entities.ProjectEnv, err error)
		GetEnvByUUID(ctx context.Context, uuid g_uuid.UUID) (env entities.ProjectEnv, err error)
		SetEnvByID(ctx context.Context, id types.ID, key, value string) (err error)
		SetEnvByUUID(ctx context.Context, uuid g_uuid.UUID, key, value string) (err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		if usecase.repositories.Projects, err = repository_projects.New(ctx); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "projects").
		Field("config", usecase.conf).Write()

	return
}

// Create - создание проекта.
func (usecase *UseCase) Create(ctx context.Context, name, description, version string) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, name, description, version)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("Starting the creation of a project... ").
		Field("name", name).
		Field("version", version).
		Field("description", description).Write()

	var (
		projectUUID = g_uuid.New()
		projectID   types.ID
	)

	// Добавление информации в системную базу данных
	{
		var err error

		if projectID, err = usecase.repositories.Projects.Create(ctx, projectUUID.String(), name, description, version); err != nil {
			cErr = error_list.FailedCreateProject()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The project data could not be written to the database: '%s'. ", err).Write()
			return
		}
	}

	// Процесс создания базы данных
	{
		var (
			connector sqlite3.Connector
			err       error
		)

		usecase.components.Logger.Info().
			Text("Creating a project database... ").Write()

		// Подключение/создание файла
		{
			var (
				filename = fmt.Sprintf("%s.db", projectUUID)
				fileDir  = env.Paths.Var.Lib.Projects
			)

			usecase.components.Logger.Info().
				Text("Creating a database file... ").
				Field("filename", filename).
				Field("dir", fileDir).Write()

			var conf = new(sqlite3.Config).Default()

			conf.Database = path.Join(fileDir, filename)

			if connector, err = sqlite3.New(ctx, conf); err != nil {
				cErr = error_list.FailedCreateProject()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("The system database file could not be created: '%s'. ", err).Write()
				return
			}

			usecase.components.Logger.Info().
				Text("CThe creation of the database file is completed. ").
				Field("filename", filename).
				Field("dir", fileDir).Write()
		}

		// Выполнение миграций
		{
			usecase.components.Logger.Info().
				Text("Starting migrations for the project database... ").Write()

			var query string

			// Чтение файла миграций
			{
				var data []byte

				if data, err = embed.Dir.ReadFile("migrations/project.sql"); err != nil {
					cErr = error_list.FailedCreateProject()
					cErr.SetError(err)

					usecase.components.Logger.Error().
						Format("The migration file for the system database could not be read: '%s'. ", err).Write()
					return
				}

				query = string(data)
			}

			if _, err = connector.Exec(query); err != nil {
				cErr = error_list.FailedCreateProject()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("Migrations for the system database failed: '%s'. ", err).Write()
				return
			}

			usecase.components.Logger.Info().
				Text("Migrations for the project database have been completed. ").Write()
		}

		// Запись окружения проекта
		{
			usecase.components.Logger.Info().
				Text("Starting data recording in the project environment... ").Write()

			var tx *sqlx.Tx

			tx, err = connector.BeginTxx(ctx, nil)

			var (
				query = `
					update 
						env 
					set
					    value = $1
					where 
					    key = $2
				`
			)

			var data = map[string]string{
				"id":          strconv.Itoa(int(projectID)),
				"uuid":        projectUUID.String(),
				"name":        name,
				"version":     version,
				"description": description,
			}

			for k, v := range data {
				if _, err = tx.Exec(query, v, k); err != nil {
					cErr = error_list.FailedCreateProject()
					cErr.SetError(err)

					usecase.components.Logger.Error().
						Format("The project environment variables could not be set: '%s'. ", err).Write()
					return
				}
			}

			if err = tx.Commit(); err != nil {
				cErr = error_list.FailedCreateProject()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("The project environment variables could not be set: '%s'. ", err).Write()
				return
			}

			usecase.components.Logger.Info().
				Text("Writing data to the project environment is completed. ").Write()
		}

		usecase.components.Logger.Info().
			Text("The creation of the project database has been completed. ").Write()
	}

	usecase.components.Logger.Info().
		Text("The creation of the project is complete. ").
		Field("name", name).
		Field("version", version).
		Field("description", description).Write()

	return
}

// GetAll - получение всех проектов системы.
func (usecase *UseCase) GetAll(ctx context.Context) (projects []*entities.Project, cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(cErr).FunctionCallFinished(projects) }()
	}

	var err error

	if projects, err = usecase.repositories.Projects.GetAll(ctx); err != nil {
		usecase.components.Logger.Error().
			Format("It was not possible to get data on system projects: '%s'. ", err).Write()

		cErr = error_list.ReceivingTheProjects()
		cErr.SetError(err)

		return
	}

	if projects == nil {
		projects = make([]*entities.Project, 0)
	}

	return
}

// RemoveByID - удаление проекта по ID.
func (usecase *UseCase) RemoveByID(ctx context.Context, id types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, id)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if project, err := usecase.repositories.Projects.GetByID(ctx, id); err != nil {
		usecase.components.Logger.Error().
			Format("Project data could not be retrieved from the database: '%s'. ", err).Write()

		if errors.Is(err, sql.ErrNoRows) {
			cErr = error_list.ProjectNotFound()
			cErr.SetError(err)

			return
		}

		cErr = error_list.ProjectDataCouldNotBeRetrieved()
		cErr.SetError(err)

		return
	} else if project == nil {
		err = errors.New("project == nil")

		usecase.components.Logger.Error().
			Format("Project data could not be retrieved from the database: '%s'. ", err).Write()

		cErr = error_list.ProjectNotFound()
		cErr.SetError(err)

		return
	} else {
		if cErr = usecase.remove(ctx, project); cErr != nil {
			usecase.components.Logger.Error().
				Format("The project could not be deleted: '%s'. ", cErr).Write()
			return
		}
	}

	return
}

// RemoveByUUID - удаление проекта по UUID.
func (usecase *UseCase) RemoveByUUID(ctx context.Context, uuid g_uuid.UUID) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, uuid)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if project, err := usecase.repositories.Projects.GetByUUID(ctx, uuid.String()); err != nil {
		usecase.components.Logger.Error().
			Format("Project data could not be retrieved from the database: '%s'. ", err).Write()

		if errors.Is(err, sql.ErrNoRows) {
			cErr = error_list.ProjectNotFound()
			cErr.SetError(err)

			return
		}

		cErr = error_list.ProjectDataCouldNotBeRetrieved()
		cErr.SetError(err)

		return
	} else if project == nil {
		err = errors.New("project == nil")

		usecase.components.Logger.Error().
			Format("Project data could not be retrieved from the database: '%s'. ", err).Write()

		cErr = error_list.ProjectNotFound()
		cErr.SetError(err)

		return
	} else {
		if cErr = usecase.remove(ctx, project); cErr != nil {
			usecase.components.Logger.Error().
				Format("The project could not be deleted: '%s'. ", cErr).Write()
			return
		}
	}

	return
}

// SetEnvByID - установить значение переменной окружения проекта по ID.
func (usecase *UseCase) SetEnvByID(ctx context.Context, id types.ID, key, value string) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, id, key, value)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if err := usecase.repositories.Projects.SetEnvByID(ctx, id, key, value); err != nil {
		usecase.components.Logger.Error().
			Format("The value of the project environment variable could not be set: '%s'. ", err).Write()

		cErr = error_list.FailedSetProjectEnv()
		cErr.SetError(err)

		return
	}

	return
}

// SetEnvByUUID - установить значение переменной окружения проекта по UUID.
func (usecase *UseCase) SetEnvByUUID(ctx context.Context, uuid g_uuid.UUID, key, value string) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, uuid, key, value)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if err := usecase.repositories.Projects.SetEnvByUUID(ctx, uuid, key, value); err != nil {
		usecase.components.Logger.Error().
			Format("The value of the project environment variable could not be set: '%s'. ", err).Write()

		cErr = error_list.FailedSetProjectEnv()
		cErr.SetError(err)

		return
	}

	return
}

// GetEnvByID - получить переменные окружения проекта по ID.
func (usecase *UseCase) GetEnvByID(ctx context.Context, id types.ID) (env entities.ProjectEnv, cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, id)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	var err error

	if env, err = usecase.repositories.Projects.GetEnvByID(ctx, id); err != nil {
		usecase.components.Logger.Error().
			Format("Failed to get the project environment variables: '%s'. ", err).Write()

		cErr = error_list.FailedGetProjectEnv()
		cErr.SetError(err)

		return
	}

	if env == nil {
		env = make(entities.ProjectEnv, 0)
	}

	return
}

// GetEnvByUUID - получить переменные окружения проекта по UUID.
func (usecase *UseCase) GetEnvByUUID(ctx context.Context, uuid g_uuid.UUID) (env entities.ProjectEnv, cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelUseCase)

		trace.FunctionCall(ctx, uuid)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	var err error

	if env, err = usecase.repositories.Projects.GetEnvByUUID(ctx, uuid); err != nil {
		usecase.components.Logger.Error().
			Format("Failed to get the project environment variables: '%s'. ", err).Write()

		cErr = error_list.FailedGetProjectEnv()
		cErr.SetError(err)

		return
	}

	if env == nil {
		env = make(entities.ProjectEnv, 0)
	}

	return
}
