package users_controller

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/infrastructure/usecases/users"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=users"
)

// Controller - контроллер работы с пользователями.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Users interface {
		Get(ctx context.Context, ids ...common_types.ID) (list []*entities.User, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (us *entities.User, cErr c_errors.Error)
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

		// Users
		{
			if controller.usecases.Users, err = users_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "users").
		Field("config", controller.conf).Write()

	return
}

// Get - получение пользователей по ID.
func (controller *Controller) Get(ctx context.Context, ids ...common_types.ID) (list []*models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var users []*entities.User

	// Выполнения инструкций
	{
		if users, cErr = controller.usecases.Users.Get(ctx, ids...); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		list = make([]*models.UserInfo, 0, len(users))

		if users != nil {
			for _, us := range users {
				list = append(list, us.ToModel())
			}
		}
	}

	return
}

// GetOne - получение пользователя по ID.
func (controller *Controller) GetOne(ctx context.Context, id common_types.ID) (user *models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var us *entities.User

	// Выполнения инструкций
	{
		if us, cErr = controller.usecases.Users.GetOne(ctx, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модель
	{
		if us != nil {
			user = us.ToModel()
		}
	}

	return
}
