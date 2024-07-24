package users_adapter

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/infrastructure/controllers/users"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_Grpc = "infrastructure-[adapters]=users-(Grpc)"
)

// Adapter_Grpc - адаптер контроллера для grpc.
type Adapter_Grpc struct {
	components *components

	controller interface {
		Get(ctx context.Context, ids ...common_types.ID) (list []*models.UserInfo, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (user *models.UserInfo, cErr c_errors.Error)
	}

	ctx context.Context
}

// components - компоненты адаптера.
type components struct {
	Logger logger.Logger
}

// New_Grpc - создание контроллера для grpc.
func New_Grpc(ctx context.Context) (adapter *Adapter_Grpc, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(adapter) }()
	}

	adapter = new(Adapter_Grpc)
	adapter.ctx = ctx

	// Компоненты
	{
		adapter.components = new(components)

		// Logger
		{
			if adapter.components.Logger, err = logger.New(loggerInitiator_Grpc); err != nil {
				return
			}
		}
	}

	// Контроллер
	{
		if adapter.controller, err = users_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for Grpc has been created. ", "users").Write()

	return
}

// Get - получение пользователей по ID.
func (adapter *Adapter_Grpc) Get(ctx context.Context, ids ...common_types.ID) (list []*models.UserInfo, cErr c_errors.Grpc) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var proxyErr c_errors.Error

	if list, proxyErr = adapter.controller.Get(ctx, ids...); proxyErr != nil {
		cErr = c_errors.ToGrpc(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetOne - получение пользователя по ID.
func (adapter *Adapter_Grpc) GetOne(ctx context.Context, id common_types.ID) (user *models.UserInfo, cErr c_errors.Grpc) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var proxyErr c_errors.Error

	if user, proxyErr = adapter.controller.GetOne(ctx, id); proxyErr != nil {
		cErr = c_errors.ToGrpc(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
