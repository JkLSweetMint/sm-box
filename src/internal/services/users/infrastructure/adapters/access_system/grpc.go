package access_system_adapter

import (
	"context"
	common_types "sm-box/internal/common/types"
	access_system_controller "sm-box/internal/services/users/infrastructure/controllers/access_system"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_Grpc = "infrastructure-[adapters]=access_system-(Grpc)"
)

// Adapter_Grpc - адаптер контроллера для grpc.
type Adapter_Grpc struct {
	components *components

	controller interface {
		CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Error)
	}

	ctx context.Context
}

// New_Grpc - создание контроллера для grpc.
func New_Grpc(ctx context.Context) (adapter *Adapter_Grpc, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelAdapter)

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
		if adapter.controller, err = access_system_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for Grpc has been created. ", "access_system").Write()

	return
}

// CheckUserAccess - проверка доступов пользователя.
func (adapter *Adapter_Grpc) CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Grpc) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, rolesID, permissionsID)
		defer func() { trc.Error(cErr).FunctionCallFinished(allowed) }()
	}

	var proxyErr c_errors.Error

	if allowed, proxyErr = adapter.controller.CheckUserAccess(ctx, userID, rolesID, permissionsID); proxyErr != nil {
		cErr = c_errors.ToGrpc(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
