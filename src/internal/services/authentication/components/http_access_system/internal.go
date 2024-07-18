package http_access_system

import (
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/tracer"
)

// registerHttpRoutesInRedisDb - регистрация http маршрутов в базу данных redis.
func (acc *accessSystem) registerHttpRoutesInRedisDb() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponentInternal)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(acc) }()
	}

	acc.components.Logger.Info().
		Text("Http routes have been uploaded to the redis database... ").Write()

	var routes []*entities.HttpRoute

	// Получение маршрутов
	{
		if routes, err = acc.repositories.HttpRoutes.GetAll(acc.ctx); err != nil {
			acc.components.Logger.Error().
				Format("Failed to get a list of http routes: '%s'. ", err).Write()

			return
		}

		acc.components.Logger.Info().
			Text("The following http routes were received for uploading to the redis database. ").
			Field("routes", routes).Write()
	}

	// Очистка старых
	{
		if err = acc.repositories.HttpRoutesRedis.Clear(acc.ctx); err != nil {
			acc.components.Logger.Error().
				Format("Failed to clear http routes in the redis database: '%s'. ", err).Write()

			return
		}

		acc.components.Logger.Info().
			Text("The old http routes in the redis database have been successfully deleted. ").Write()
	}

	// Загрузка новых
	{
		if err = acc.repositories.HttpRoutesRedis.Register(acc.ctx, routes...); err != nil {
			acc.components.Logger.Error().
				Format("Failed to register http routes in the redis database: '%s'. ", err).Write()

			return
		}
	}

	acc.components.Logger.Info().
		Text("Http routes have been uploaded to the redis database. ").Write()

	return
}
