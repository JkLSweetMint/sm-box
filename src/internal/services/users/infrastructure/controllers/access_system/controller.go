package access_system_controller

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/infrastructure/usecases/access_system"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=access_system"
)

// Controller - контроллер для работы с системой доступа.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	AccessSystem interface {
		GetRolesListForSelect(ctx context.Context) (list []*entities.Role, cErr c_errors.Error)
		GetPermissionsListForSelect(ctx context.Context) (list []*entities.Permission, cErr c_errors.Error)

		CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Error)
	}
}

// components - компоненты контроллера.
type components struct {
	Logger logger.Logger
}

// New - создание контроллера.
func New(ctx context.Context) (controller *Controller, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(controller) }()
	}

	controller = new(Controller)
	controller.ctx = ctx

	// Конфигурация
	{
		controller.conf = new(Config).Default()

		if err = controller.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		controller.components = new(components)

		// Logger
		{
			if controller.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		// AccessSystem
		{
			if controller.usecases.AccessSystem, err = access_system_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "access_system").
		Field("config", controller.conf).Write()

	return
}

// GetRolesListForSelect - получение списка ролей для select'ов.
func (controller *Controller) GetRolesListForSelect(ctx context.Context) (list []*models.RoleInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var roles []*entities.Role

	// Выполнения инструкций
	{
		if roles, cErr = controller.usecases.AccessSystem.GetRolesListForSelect(ctx); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		list = make([]*models.RoleInfo, 0, len(roles))

		if roles != nil {
			for _, rl := range roles {
				list = append(list, rl.ToModel())
			}
		}
	}

	return
}

// GetPermissionsListForSelect - получение списка прав для select'ов.
func (controller *Controller) GetPermissionsListForSelect(ctx context.Context) (list []*models.PermissionInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var permissions []*entities.Permission

	// Выполнения инструкций
	{
		if permissions, cErr = controller.usecases.AccessSystem.GetPermissionsListForSelect(ctx); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		list = make([]*models.PermissionInfo, 0, len(permissions))

		if permissions != nil {
			for _, perm := range permissions {
				list = append(list, perm.ToModel())
			}
		}
	}

	return
}

// CheckUserAccess - проверка доступов пользователя.
func (controller *Controller) CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, rolesID, permissionsID)
		defer func() { trc.Error(cErr).FunctionCallFinished(allowed) }()
	}

	// Выполнения инструкций
	{
		if allowed, cErr = controller.usecases.AccessSystem.CheckUserAccess(ctx, userID, rolesID, permissionsID); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}
