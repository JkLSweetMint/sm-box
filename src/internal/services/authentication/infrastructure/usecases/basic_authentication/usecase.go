package basic_authentication_usecase

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	app_entities "sm-box/internal/app/objects/entities"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	basic_authentication_repository "sm-box/internal/services/authentication/infrastructure/repositories/basic_authentication"
	"sm-box/internal/services/authentication/objects/entities"
	basic_authentication_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/basic_authentication_service"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
	users_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/users_service"
	users_models "sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
)

const (
	loggerInitiator = "infrastructure-[usecases]=basic_authentication"
)

// UseCase - логика базовой аутентификации пользователей.
type UseCase struct {
	components   *components
	gateways     *gateways
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// gateways - шлюзы логики.
type gateways struct {
	BasicAuthentication interface {
		Auth(ctx context.Context, username, password string) (user *users_models.UserInfo, cErr c_errors.Error)
	}
	Projects interface {
		Get(ctx context.Context, ids ...types.ID) (list app_models.ProjectList, cErr c_errors.Error)
		GetOne(ctx context.Context, id types.ID) (project *app_models.ProjectInfo, cErr c_errors.Error)
	}
	Users interface {
		Get(ctx context.Context, ids ...types.ID) (list []*users_models.UserInfo, cErr c_errors.Error)
		GetOne(ctx context.Context, id types.ID) (project *users_models.UserInfo, cErr c_errors.Error)
	}
}

// repositories - репозитории логики.
type repositories struct {
	BasicAuthentication interface {
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

		// BasicAuthentication
		{
			if usecase.gateways.BasicAuthentication, err = basic_authentication_service_gateway.New(ctx); err != nil {
				return
			}
		}

		// Projects
		{
			if usecase.gateways.Projects, err = projects_service_gateway.New(ctx); err != nil {
				return
			}
		}

		// Users
		{
			if usecase.gateways.Users, err = users_service_gateway.New(ctx); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// BasicAuthentication
		{
			if usecase.repositories.BasicAuthentication, err = basic_authentication_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "basic_authentication").
		Field("config", usecase.conf).Write()

	return
}

// Auth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (usecase *UseCase) Auth(ctx context.Context, rawSessionToken, username, password string) (sessionToken *entities.JwtSessionToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, rawSessionToken, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken) }()
	}

	usecase.components.Logger.Info().
		Text("User authorization is in progress... ").
		Field("username", username).
		Field("raw_session_token", rawSessionToken).
		Field("password", password).Write()

	var (
		user                *users_models.UserInfo
		currentSessionToken *entities.JwtSessionToken
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

		// Токен сессии
		{
			if len(rawSessionToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty session token string was passed. ").
					Field("raw", rawSessionToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена
	{
		var err error

		currentSessionToken = new(entities.JwtSessionToken)

		if err = currentSessionToken.Parse(rawSessionToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("raw", rawSessionToken).Write()

			cErr = error_list.InvalidToken()
			cErr.SetError(err)
			return
		}

		if currentSessionToken.JwtToken, err = usecase.repositories.BasicAuthentication.GetToken(ctx, rawSessionToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("raw", rawSessionToken).Write()

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
			Text("The user's current session token was received. ").
			Field("token", currentSessionToken).Write()
	}

	// Проверка что уже авторизован
	{
		if currentSessionToken.UserID != 0 {
			usecase.components.Logger.Warn().
				Text("The user is already logged in. ").
				Field("session_token", currentSessionToken).Write()

			cErr = error_list.AlreadyAuthorized()

			return
		}
	}

	// Шифрование пароля
	{
		var (
			err        error
			passwordDB []byte
		)

		if passwordDB, err = rsa.EncryptOAEP(
			sha256.New(),
			rand.Reader,
			env.Vars.EncryptionKeys.Public,
			[]byte(password),
			[]byte("password")); err != nil {
			usecase.components.Logger.Error().
				Format("The encryption of the user's password failed: '%s'. ", err).
				Field("username", username).
				Field("password", password).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		password = base64.StdEncoding.EncodeToString(passwordDB)
	}

	// Получение данных пользователя
	{
		if user, cErr = usecase.gateways.BasicAuthentication.Auth(ctx, username, password); cErr != nil {
			usecase.components.Logger.Error().
				Format("User authorization error: '%s'. ", cErr).
				Field("username", username).
				Field("password", password).Write()

			if errors.Is(cErr, sql.ErrNoRows) {
				cErr = error_list.UserNotFound()
				return
			}

			cErr = error_list.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The user's data has been successfully received. ").
			Field("user", user).Write()
	}

	// Обновление токена сессии
	{
		var err error

		sessionToken = &entities.JwtSessionToken{
			JwtToken: &entities.JwtToken{
				ParentID: currentSessionToken.ID,
				UserID:   user.ID,

				Params: currentSessionToken.Params,
			},
			Claims: currentSessionToken.Claims,
		}

		if err = sessionToken.Generate(); err != nil {
			usecase.components.Logger.Error().
				Format("User session token generation failed: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}

		// Сохранение в базе
		{
			if err = usecase.repositories.BasicAuthentication.Register(ctx, sessionToken.JwtToken); err != nil {
				usecase.components.Logger.Error().
					Format("The client's token could not be registered in the database: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}
		}

		usecase.components.Logger.Info().
			Text("The user's jwt session token has been updated. ").
			Field("old", currentSessionToken).
			Field("new", sessionToken).Write()
	}

	usecase.components.Logger.Info().
		Text("The user's authorization has been successfully completed. ").
		Field("username", username).Write()

	return
}

// GetUserProjectList - получение списка проектов пользователя.
func (usecase *UseCase) GetUserProjectList(ctx context.Context, rawSessionToken string) (list app_entities.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawSessionToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var sessionToken *entities.JwtSessionToken

	usecase.components.Logger.Info().
		Text("The process of receiving user projects has been started... ").
		Field("raw_session_token", rawSessionToken).Write()

	// Валидация
	{
		// Токен
		{
			if len(rawSessionToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty session token string was passed. ").
					Field("raw", rawSessionToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена сессии
	{
		sessionToken = new(entities.JwtSessionToken)

		if err := sessionToken.Parse(rawSessionToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get session token data: '%s'. ", err).
				Field("raw", rawSessionToken).Write()

			cErr = error_list.InvalidToken()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's session token has been successfully received. ").
			Field("token", rawSessionToken).Write()
	}

	// Проверки
	{
		// Авторизация
		{
			if sessionToken.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token is not authorized, it is impossible to receive the user's projects. ").
					Field("token_id", sessionToken.ID).Write()

				cErr = error_list.Unauthorized()
				return
			}
		}

		// Проект уже выбран
		{
			if sessionToken.ProjectID != 0 {
				usecase.components.Logger.Error().
					Text("The project has already been selected, it is not possible to re-select it. ").
					Field("token_id", sessionToken.ID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}
	}

	// Получение
	{
		var (
			user *users_models.UserInfo
			ids  = make([]types.ID, 0)
		)

		// Данные пользователя
		{
			if user, cErr = usecase.gateways.Users.GetOne(ctx, sessionToken.UserID); cErr != nil {
				usecase.components.Logger.Error().
					Format("User data could not be retrieved: '%s'. ", cErr).
					Field("user_id", sessionToken.UserID).Write()

				cErr = error_list.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("The user's data has been successfully received. ").
				Field("user", user).Write()
		}

		// Список id проектов
		{
			var ids_ = make(map[types.ID]struct{})

			var writeInheritance func(rl *users_models.RoleInfo)

			writeInheritance = func(rl *users_models.RoleInfo) {
				if id := rl.ProjectID; id != 0 {
					ids_[id] = struct{}{}
				}

				for _, child := range rl.Inheritances {
					writeInheritance(child.RoleInfo)
				}
			}

			for _, rl := range user.Accesses {
				writeInheritance(rl.RoleInfo)
			}

			for k, _ := range ids_ {
				ids = append(ids, k)
			}
		}

		usecase.components.Logger.Info().
			Text("The list of the user's project IDs has been collected. ").
			Field("ids", ids).Write()

		var projects app_models.ProjectList

		if projects, cErr = usecase.gateways.Projects.Get(ctx, ids...); cErr != nil {
			usecase.components.Logger.Error().
				Format("The list of user's projects could not be retrieved: '%s'. ", cErr).
				Field("user_id", sessionToken.UserID).Write()

			cErr = error_list.InternalServerError()
			return
		}

		list = make(app_entities.ProjectList, 0)

		for _, project := range projects {
			list = append(list, &app_entities.Project{
				ID: project.ID,

				Name:        project.Name,
				Description: project.Description,
				Version:     project.Version,
			})
		}
	}

	usecase.components.Logger.Info().
		Text("The process of receiving user projects is completed. ").
		Field("raw_session_token", rawSessionToken).
		Field("projects", list).Write()

	return
}

// SetTokenProject - установить проект для токена.
func (usecase *UseCase) SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
	sessionToken *entities.JwtSessionToken,
	accessToken *entities.JwtAccessToken,
	refreshToken *entities.JwtRefreshToken,
	cErr c_errors.Error) {

	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, rawSessionToken, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken, accessToken, refreshToken) }()
	}

	var (
		user                *users_models.UserInfo
		project             *app_models.ProjectInfo
		currentSessionToken *entities.JwtSessionToken
	)

	usecase.components.Logger.Info().
		Text("The process of establishing a project for the token has been started... ").
		Field("project_id", projectID).
		Field("raw_session_token", rawSessionToken).Write()

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

		// Токен сессии
		{
			if len(rawSessionToken) == 0 {
				usecase.components.Logger.Error().
					Text("An empty session token string was passed. ").
					Field("raw_session_token", rawSessionToken).Write()

				cErr = error_list.TokenWasNotTransferred()
				return
			}
		}
	}

	// Получение токена
	{
		var err error

		currentSessionToken = new(entities.JwtSessionToken)

		if err = currentSessionToken.Parse(rawSessionToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get session token: '%s'. ", err).
				Field("raw", rawSessionToken).Write()

			cErr = error_list.InvalidToken()
			cErr.SetError(err)
			return
		}

		if currentSessionToken.JwtToken, err = usecase.repositories.BasicAuthentication.GetToken(ctx, rawSessionToken); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get session token: '%s'. ", err).
				Field("raw", rawSessionToken).Write()

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
			Text("The user's current session token was received. ").
			Field("token", currentSessionToken).Write()
	}

	// Проверки
	{
		// Авторизация
		{
			if currentSessionToken.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token was not authorized, the project selection is not allowed. ").
					Field("session_token", currentSessionToken).
					Field("project_id", projectID).Write()

				cErr = error_list.Unauthorized()
				return
			}
		}

		// Проект уже выбран
		{
			if currentSessionToken.ProjectID != 0 {
				usecase.components.Logger.Error().
					Text("The project has already been selected, it is not possible to re-select it. ").
					Field("session_token", currentSessionToken).
					Field("project_id", projectID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}
	}

	// Получение данных пользователя
	{
		if user, cErr = usecase.gateways.Users.GetOne(ctx, currentSessionToken.UserID); cErr != nil {
			usecase.components.Logger.Error().
				Format("Failed to get the user data: '%s'. ", cErr).
				Field("id", projectID).Write()

			if errors.Is(cErr, sql.ErrNoRows) {
				cErr = error_list.UserNotFound()
				return
			}

			cErr = error_list.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The user's data has been successfully received. ").
			Field("user", user).Write()
	}

	// Получение данных проекта
	{
		if project, cErr = usecase.gateways.Projects.GetOne(ctx, projectID); cErr != nil {
			usecase.components.Logger.Error().
				Format("Failed to get the project: '%s'. ", cErr).
				Field("id", projectID).Write()

			if errors.Is(cErr, sql.ErrNoRows) {
				cErr = error_list.ProjectNotFound()
				return
			}

			cErr = error_list.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The project data has been successfully received. ").
			Field("project", project).Write()
	}

	// Проверка доступа
	{
		var ids = make(map[types.ID]struct{})

		// Список id проектов
		{
			var writeInheritance func(rl *users_models.RoleInfo)

			writeInheritance = func(rl *users_models.RoleInfo) {
				if id := rl.ProjectID; id != 0 {
					ids[id] = struct{}{}
				}

				for _, child := range rl.Inheritances {
					writeInheritance(child.RoleInfo)
				}
			}

			for _, rl := range user.Accesses {
				writeInheritance(rl.RoleInfo)
			}
		}

		if _, ok := ids[project.ID]; !ok {
			usecase.components.Logger.Error().
				Format("The user does not have access to the project: '%s'. ", cErr).
				Field("project_id", project.ID).
				Field("user_id", user.ID).Write()

			cErr = error_list.NotAccessToProject()
			return
		}
	}

	// Создание новых токенов
	{
		// Обновление токена сессии
		{
			var err error

			sessionToken = &entities.JwtSessionToken{
				JwtToken: &entities.JwtToken{
					ParentID:  currentSessionToken.ID,
					ProjectID: projectID,
					UserID:    user.ID,

					Params: currentSessionToken.Params,
				},
				Claims: currentSessionToken.Claims,
			}

			if err = sessionToken.Generate(); err != nil {
				usecase.components.Logger.Error().
					Format("User session token generation failed: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}

			// Сохранение в базе
			{
				if err = usecase.repositories.BasicAuthentication.Register(ctx, sessionToken.JwtToken); err != nil {
					usecase.components.Logger.Error().
						Format("The client's token could not be registered in the database: '%s'. ", err).Write()

					cErr = error_list.InternalServerError()
					cErr.SetError(err)

					return
				}
			}

			usecase.components.Logger.Info().
				Text("The user's jwt session token has been updated. ").
				Field("old", currentSessionToken).
				Field("new", sessionToken).Write()
		}

		// Создание токена обновления
		{
			var err error

			refreshToken = &entities.JwtRefreshToken{
				JwtToken: &entities.JwtToken{
					ProjectID: projectID,
					UserID:    user.ID,

					Params: currentSessionToken.Params,
				},
				Claims: nil,
			}

			if err = refreshToken.Generate(); err != nil {
				usecase.components.Logger.Error().
					Format("User session token generation failed: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}

			// Сохранение в базе
			{
				if err = usecase.repositories.BasicAuthentication.Register(ctx, refreshToken.JwtToken); err != nil {
					usecase.components.Logger.Error().
						Format("The client's token could not be registered in the database: '%s'. ", err).Write()

					cErr = error_list.InternalServerError()
					cErr.SetError(err)

					return
				}
			}

			usecase.components.Logger.Info().
				Text("The user's jwt refresh token has been updated. ").
				Field("token", refreshToken).Write()
		}

		// Создание токена доступа
		{
			var err error

			accessToken = &entities.JwtAccessToken{
				JwtToken: &entities.JwtToken{
					ProjectID: projectID,
					UserID:    user.ID,

					Params: currentSessionToken.Params,
				},
				Claims: nil,
			}

			// Запись доступов пользователя
			{
				accessToken.Claims = &entities.JwtAccessTokenClaims{
					Accesses: user.Accesses.ListIDs(),
				}
			}

			if err = accessToken.Generate(); err != nil {
				usecase.components.Logger.Error().
					Format("User session token generation failed: '%s'. ", err).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)

				return
			}

			// Сохранение в базе
			{
				if err = usecase.repositories.BasicAuthentication.Register(ctx, accessToken.JwtToken); err != nil {
					usecase.components.Logger.Error().
						Format("The client's token could not be registered in the database: '%s'. ", err).Write()

					cErr = error_list.InternalServerError()
					cErr.SetError(err)

					return
				}
			}

			usecase.components.Logger.Info().
				Text("The user's jwt access token has been updated. ").
				Field("token", accessToken).Write()
		}
	}

	usecase.components.Logger.Info().
		Text("The process of establishing a project for the token has been completed. ").
		Field("project_id", projectID).
		Field("raw_session_token", rawSessionToken).Write()

	return
}

// RefreshAccessToken - обновление токена доступа.
func (usecase *UseCase) RefreshAccessToken(ctx context.Context, rawRefreshToken string) (
	accessToken *entities.JwtAccessToken,
	refreshToken *entities.JwtRefreshToken,
	cErr c_errors.Error) {

	return
}
