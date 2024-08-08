package http_rest_api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/components/notification_notifier"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/internal/services/notifications/transport/servers/http/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"sm-box/pkg/tools/file"
	"sync"
	"time"
)

// server - сервер http rest api.
type server struct {
	app    *fiber.App
	router fiber.Router

	conf *config.Config
	ctx  context.Context

	controllers *controllers
	components  *components

	postman *postman.Collection
}

// controllers - контроллеры сервера.
type controllers struct {
	UserNotifications interface {
		GetList(ctx context.Context,
			recipientID common_types.ID,
			search *objects.UserNotificationSearch,
			pagination *objects.UserNotificationPagination,
			filters *objects.UserNotificationFilters,
		) (count, countNotRead int64, list []*models.UserNotificationInfo, cErr c_errors.RestAPI)

		CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *models.UserNotificationInfo, cErr c_errors.RestAPI)
		Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*models.UserNotificationInfo, cErr c_errors.RestAPI)

		RemoveOne(ctx context.Context, recipientID common_types.ID, id common_types.ID) (cErr c_errors.RestAPI)
		Remove(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (cErr c_errors.RestAPI)

		ReadOne(ctx context.Context, recipientID common_types.ID, id common_types.ID) (cErr c_errors.RestAPI)
		Read(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (cErr c_errors.RestAPI)
	}
}

// components - компоненты сервера.
type components struct {
	Logger               logger.Logger
	NotificationNotifier notification_notifier.Notifier
}

// Listen - запуск сервера.
func (srv *server) Listen() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("The http rest api server is listening... ").
		Field("config", srv.conf).Write()

	// Postman
	{
		if env.Mode == env_mode.Dev {
			var (
				p    = path.Join(env.Paths.SystemLocation, env.Paths.Src.Path, "/transport/postman", env.Vars.SystemName)
				name = fmt.Sprintf("service.%s@%s.json", srv.conf.Server.Name, srv.conf.Server.Version)
			)

			// Проверка наличия файла
			{
				var exist bool

				if exist, err = file.Exists(path.Join(p, name)); err != nil {
					srv.components.Logger.Error().
						Format("Could not verify the existence of the postman collection file: '%s'. ", err).Write()
					return
				}

				if exist {
					if err = os.Remove(path.Join(p, name)); err != nil {
						srv.components.Logger.Error().
							Format("The postman collection file could not be deleted: '%s'. ", err).Write()
						return
					}
				} else {
					if err = os.MkdirAll(p, 0666); err != nil {
						srv.components.Logger.Error().
							Format("Failed to create a directory for the postman collection file: '%s'. ", err).Write()
						return
					}
				}
			}

			var fl *os.File

			if fl, err = os.Create(path.Join(p, name)); err != nil {
				srv.components.Logger.Error().
					Format("Failed to create a file for the postman collection: '%s'. ", err).Write()
				return
			}

			defer fl.Close()

			if err = srv.postman.Write(fl, postman.V210); err != nil {
				srv.components.Logger.Error().
					Format("Failed to write postman collection data: '%s'. ", err).Write()
				return
			}
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = srv.app.Listen(srv.conf.Server.Addr); err != nil {
			srv.components.Logger.Error().
				Format("An error occurred when starting the http rest api server maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	srv.components.Logger.Info().
		Text("The http rest api server is listened. ").Write()

	wg.Wait()

	return
}

// Shutdown - завершение работы сервера.
func (srv *server) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Shutting down the http rest api server... ").Write()

	if err = srv.app.Shutdown(); err != nil {
		srv.components.Logger.Error().
			Format("An error occurred when completing http rest api server maintenance: '%s'. ", err).Write()
		return
	}

	srv.components.Logger.Info().
		Text("The http rest api server is turned off. ").Write()

	return
}
