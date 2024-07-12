package authentication_usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	app_entities "sm-box/internal/app/objects/entities"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	authentication_repository "sm-box/internal/services/authentication/infrastructure/repositories/authentication"
	"sm-box/internal/services/authentication/objects/entities"
	authentication_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/authentication_service"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
)

const (
	loggerInitiator = "infrastructure-[usecases]=authentication"
)

// UseCase - логика аутентификации пользователей.
type UseCase struct {
	components   *components
	gateways     *gateways
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// gateways - шлюзы логики.
type gateways struct {
	Authentication interface {
		BasicAuth(ctx context.Context, username, password string) (user *app_models.UserInfo, cErr c_errors.Error)
	}
	Projects interface {
		GetListByUser(ctx context.Context, userID types.ID) (list app_models.ProjectList, cErr c_errors.Error)
		Get(ctx context.Context, id types.ID) (project *app_models.ProjectInfo, cErr c_errors.Error)
	}
}

// repositories - репозитории логики.
type repositories struct {
	Authentication interface {
		GetToken(ctx context.Context, raw string) (tok *entities.JwtToken, err error)
		Register(ctx context.Context, tok *entities.JwtToken) (err error)
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

	// Шлюзы
	{
		usecase.gateways = new(gateways)

		// Authentication
		{
			if usecase.gateways.Authentication, err = authentication_service_gateway.New(ctx); err != nil {
				return
			}
		}

		// Projects
		{
			if usecase.gateways.Projects, err = projects_service_gateway.New(ctx); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// Authentication
		{
			if usecase.repositories.Authentication, err = authentication_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "authentication").
		Field("config", usecase.conf).Write()

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (usecase *UseCase) BasicAuth(ctx context.Context, rawToken, username, password string) (token *entities.JwtToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, rawToken, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(token) }()
	}

	usecase.components.Logger.Info().
		Text("User authorization is in progress... ").
		Field("username", username).
		Field("raw_token", rawToken).
		Field("password", password).Write()

	var (
		us  *app_models.UserInfo
		tok *entities.JwtToken
	)

	// Валидация
	{
		// Данные для авторизации
		{
			if len(username) == 0 {
				if cErr == nil {
					cErr = error_list.InvalidAuthorizationDataWasTransferred()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("username"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}

			if len(password) == 0 {
				if cErr == nil {
					cErr = error_list.InvalidAuthorizationDataWasTransferred()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("password"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}

			if len(username) > 256 {
				if cErr == nil {
					cErr = error_list.InvalidAuthorizationDataWasTransferred()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("username"),
					new(err_messages.TextMessage).Text("Is long. "),
				)
			}

			if len(password) > 256 {
				if cErr == nil {
					cErr = error_list.InvalidAuthorizationDataWasTransferred()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("password"),
					new(err_messages.TextMessage).Text("Is long. "),
				)
			}

			if cErr != nil {
				usecase.components.Logger.Error().
					Text("Invalid authorization data was transferred. ").
					Field("username", username).
					Field("password", password).Write()

				return
			}
		}

		// Токен
		{
			if len(rawToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty token string was passed. ").
					Field("raw", rawToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена
	{
		var err error

		if tok, err = usecase.repositories.Authentication.GetToken(ctx, rawToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("raw", rawToken).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.AnUnregisteredTokenWasTransferred()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's current token was received. ").
			Field("token", tok).Write()
	}

	// Проверка что уже авторизован
	{
		if tok.UserID != 0 {
			usecase.components.Logger.Warn().
				Text("The user is already logged in. ").
				Field("token", tok).Write()

			cErr = error_list.AlreadyAuthorized()

			return
		}
	}

	// Шифрование пароля
	{
		//var (
		//	err        error
		//	passwordDB []byte
		//)
		//
		//if passwordDB, err = rsa.EncryptOAEP(
		//	sha256.New(),
		//	rand.Reader,
		//	env.Vars.EncryptionKeys.Public,
		//	[]byte(password),
		//	[]byte("password")); err != nil {
		//	usecase.components.Logger.Error().
		//		Format("The decryption of the user's password failed: '%s'. ", err).
		//		Field("username", username).
		//		Field("password", password).Write()
		//
		//	cErr = error_list.InternalServerError()
		//	cErr.SetError(err)
		//	return
		//}
		//
		//password = string(passwordDB)
	}

	// Получение данных пользователя
	{
		var err error

		if us, err = usecase.gateways.Authentication.BasicAuth(ctx, username, password); err != nil {
			us = nil

			usecase.components.Logger.Error().
				Format("User authorization error: '%s'. ", err).
				Field("username", username).
				Field("password", password).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.UserNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's data has been successfully received. ").
			Field("user", us).Write()
	}

	// Создание нового токена
	{
		var err error

		token = &entities.JwtToken{
			ParentID: tok.ID,
			UserID:   us.ID,
			Params:   tok.Params,
		}

		if err = token.Generate(); err != nil {
			usecase.components.Logger.Error().
				Format("User token generation failed: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}

		// Сохранение в базе
		{
			if err = usecase.repositories.Authentication.Register(ctx, token); err != nil {
				usecase.components.Logger.Error().
					Format("The client's token could not be registered in the database: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}
		}

		usecase.components.Logger.Info().
			Text("Updating the user's token data after successful authorization is completed. ").
			Field("token_id", token.ID).
			Field("user_id", us.ID).Write()
	}

	usecase.components.Logger.Info().
		Text("The user's authorization has been successfully completed. ").
		Field("username", username).Write()

	return
}

// GetUserProjectList - получение списка проектов пользователя.
func (usecase *UseCase) GetUserProjectList(ctx context.Context, rawToken string) (list app_entities.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var tok = new(entities.JwtToken)

	// Валидация
	{
		// Токен
		{
			if len(rawToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty token string was passed. ").
					Field("raw", rawToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена
	{
		if err := tok.Parse(rawToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token data: '%s'. ", err).
				Field("raw", rawToken).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	// Проверки
	{
		// Авторизация
		{
			if tok.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token is not authorized, it is impossible to receive the user's projects. ").
					Field("token_id", tok.ID).Write()

				cErr = error_list.Unauthorized()
				return
			}
		}

		// Проект уже выбран
		{
			if tok.ProjectID != 0 {
				usecase.components.Logger.Error().
					Text("The project has already been selected, it is not possible to re-select it. ").
					Field("token_id", tok.ID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}
	}

	// Получение
	{
		var (
			err      error
			projects app_models.ProjectList
		)

		if projects, err = usecase.gateways.Projects.GetListByUser(ctx, tok.UserID); err != nil {
			usecase.components.Logger.Error().
				Format("The list of user's projects could not be retrieved: '%s'. ", err).
				Field("user_id", tok.UserID).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		list = make(app_entities.ProjectList, 0)

		for _, project := range projects {
			list = append(list, &app_entities.Project{
				ID:      project.ID,
				OwnerID: project.OwnerID,

				Name:        project.Name,
				Description: project.Description,
				Version:     project.Version,
			})
		}
	}

	return
}

// SetTokenProject - установить проект для токена.
func (usecase *UseCase) SetTokenProject(ctx context.Context, rawToken string, projectID types.ID) (token *entities.JwtToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, rawToken, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished(token) }()
	}

	usecase.components.Logger.Info().
		Text("The process of establishing a project for the token has been started... ").
		Field("project_id", projectID).
		Field("raw_token", rawToken).Write()

	var tok = new(entities.JwtToken)

	// Валидация
	{
		// ID проекта
		{
			if projectID == 0 {
				if cErr == nil {
					cErr = error_list.InvalidDataWasTransmitted()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}

			if cErr != nil {
				usecase.components.Logger.Error().
					Text("Invalid data was transmitted for the project selection. ").
					Field("project_id", projectID).Write()

				return
			}
		}

		// Токен
		{
			if len(rawToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty token string was passed. ").
					Field("raw", rawToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена
	{
		if err := tok.Parse(rawToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token data: '%s'. ", err).
				Field("raw", rawToken).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	// Проверки
	{
		// Авторизация
		{
			if tok.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token was not authorized, the project selection is not allowed. ").
					Field("token_id", tok.ID).
					Field("project_id", projectID).Write()

				cErr = error_list.Unauthorized()
				return
			}
		}

		// Проект уже выбран
		{
			if tok.ProjectID != 0 {
				usecase.components.Logger.Error().
					Text("The project has already been selected, it is not possible to re-select it. ").
					Field("token_id", tok.ID).
					Field("project_id", projectID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}

		// Существования проекта и доступа пользователя к нему
		{
			var (
				project *app_models.ProjectInfo
				err     error
			)

			if project, err = usecase.gateways.Projects.Get(ctx, projectID); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to get the project: '%s'. ", err).
					Field("id", projectID).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = error_list.ProjectNotFound()
					cErr.SetError(err)
					return
				}

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}

			fmt.Printf("%+v\n", project)
		}
	}

	// Создание нового токена
	{
		var err error

		token = &entities.JwtToken{
			ParentID:  tok.ID,
			UserID:    tok.UserID,
			ProjectID: projectID,

			Params: tok.Params,
		}

		if err = token.Generate(); err != nil {
			usecase.components.Logger.Error().
				Format("User token generation failed: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}

		// Сохранение в базе
		{
			if err = usecase.repositories.Authentication.Register(ctx, token); err != nil {
				usecase.components.Logger.Error().
					Format("The client's token could not be registered in the database: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}
		}

		usecase.components.Logger.Info().
			Text("The process of establishing a project for the token has been completed. ").
			Field("project_id", projectID).
			Field("raw_token", rawToken).Write()
	}

	return
}
