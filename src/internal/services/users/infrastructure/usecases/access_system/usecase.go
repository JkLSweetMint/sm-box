package access_system_usecase

import (
	"context"
	common_errors "sm-box/internal/common/errors"
	access_system_repository "sm-box/internal/services/users/infrastructure/repositories/access_system"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=access_system"
)

// UseCase - логика для работы с системой доступа.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	AccessSystem interface {
		GetRolesListForSelect(ctx context.Context) (list []*entities.Role, err error)
		GetPermissionsListForSelect(ctx context.Context) (list []*entities.Permission, err error)
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
		var trace = tracer.New(tracer.LevelMain, tracer.LevelUseCase)

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

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// AccessSystem
		{
			if usecase.repositories.AccessSystem, err = access_system_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "access_system").
		Field("config", usecase.conf).Write()

	return
}

// GetRolesListForSelect - получение списка ролей для select'ов.
func (usecase *UseCase) GetRolesListForSelect(ctx context.Context) (list []*entities.Role, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a list of access system roles for select has been started... ").Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of obtaining the list of access system roles for the selects is completed. ").
			Field("list", list).Write()
	}()

	// Получение данных
	{
		var err error

		if list, err = usecase.repositories.AccessSystem.GetRolesListForSelect(ctx); err != nil {
			usecase.components.Logger.Error().
				Format("The list of access system roles for the selects could not be retrieved: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The list of access system roles for the selects has been successfully received. ").
			Field("users", list).Write()
	}

	return
}

// GetPermissionsListForSelect - получение списка прав для select'ов.
func (usecase *UseCase) GetPermissionsListForSelect(ctx context.Context) (list []*entities.Permission, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a list of access system permissions for select has been started... ").Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of obtaining the list of access system permissions for the selects is completed. ").
			Field("list", list).Write()
	}()

	// Получение данных
	{
		var err error

		if list, err = usecase.repositories.AccessSystem.GetPermissionsListForSelect(ctx); err != nil {
			usecase.components.Logger.Error().
				Format("The list of access system permissions for the selects could not be retrieved: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The list of access system permissions for the selects has been successfully received. ").
			Field("users", list).Write()
	}

	return
}