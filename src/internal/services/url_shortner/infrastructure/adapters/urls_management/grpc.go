package urls_management_adapter

import (
	"context"
	urls_management_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls_management"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"time"
)

const (
	loggerInitiator_Grpc = "infrastructure-[adapters]=urls_management-(Grpc)"
)

// Adapter_Grpc - адаптер контроллера для grpc.
type Adapter_Grpc struct {
	components *components

	controller interface {
		Create(ctx context.Context,
			source string,
			type_ types.ShortUrlType,
			numberOfUses int64,
			startActive, endActive time.Time,
		) (url *models.ShortUrlInfo, cErr c_errors.Error)
	}

	ctx context.Context
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
		if adapter.controller, err = urls_management_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for Grpc has been created. ", "urls_management").Write()

	return
}

// Create - создание сокращенного url.
func (adapter *Adapter_Grpc) Create(ctx context.Context,
	source string,
	type_ types.ShortUrlType,
	numberOfUses int64,
	startActive, endActive time.Time) (url *models.ShortUrlInfo, cErr c_errors.Grpc) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, source, type_, numberOfUses, startActive, endActive)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	var proxyErr c_errors.Error

	if url, proxyErr = adapter.controller.Create(ctx, source, type_, numberOfUses, startActive, endActive); proxyErr != nil {
		cErr = c_errors.ToGrpc(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
