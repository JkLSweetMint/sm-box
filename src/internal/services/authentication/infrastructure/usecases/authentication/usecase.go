package authentication_usecase

import (
	"context"
	"database/sql"
	"errors"
	error_list "sm-box/internal/common/errors"
	common_entities "sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	authentication_repository "sm-box/internal/services/authentication/infrastructure/repositories/authentication"
	projects_repository "sm-box/internal/services/authentication/infrastructure/repositories/projects"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=authentication"
)

// UseCase - логика аутентификации пользователей.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Authentication interface {
		GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error)
		GetTokenByID(ctx context.Context, id types.ID) (tok *common_entities.JwtToken, err error)
		SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error)
		SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (err error)
		BasicAuth(ctx context.Context, username, password string) (us *common_entities.User, err error)
	}
	Projects interface {
		GetListByUser(ctx context.Context, userID types.ID) (list common_entities.ProjectList, err error)
		GetByID(ctx context.Context, id types.ID) (project *common_entities.Project, err error)
		CheckAccess(ctx context.Context, userID, projectID types.ID) (exist bool, err error)
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

		// Authentication
		{
			if usecase.repositories.Authentication, err = authentication_repository.New(ctx); err != nil {
				return
			}
		}

		// Projects
		{
			if usecase.repositories.Projects, err = projects_repository.New(ctx); err != nil {
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
func (usecase *UseCase) BasicAuth(ctx context.Context, tokenData, username, password string) (tok *common_entities.JwtToken, us *common_entities.User, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, tokenData, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(tok, us) }()
	}

	usecase.components.Logger.Info().
		Text("User authorization is in progress... ").
		Field("token", tokenData).
		Field("username", username).
		Field("password", password).Write()

	// Получение токена
	{
		var err error

		if tok, err = usecase.repositories.Authentication.GetToken(ctx, tokenData); err != nil {
			tok = nil

			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("data", tokenData).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.TokenNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's token was received. ").
			Field("token", tok).Write()
	}

	// Проверка что уже авторизован
	{
		if tok.UserID != 0 {
			usecase.components.Logger.Warn().
				Text("The user is already logged in. ").
				Field("user_id", tok.ID).
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

		if us, err = usecase.repositories.Authentication.BasicAuth(ctx, username, password); err != nil {
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

	// Обновление данных токена
	{
		var err error

		if err = usecase.repositories.Authentication.SetTokenOwner(ctx, tok.ID, us.ID); err != nil {
			usecase.components.Logger.Error().
				Format("The token owner could not be identified: '%s'. ", err).
				Field("owner_id", us.ID).
				Field("token", tok).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}

		usecase.components.Logger.Info().
			Text("Updating the user's token data after successful authorization is completed. ").
			Field("token_id", tok.ID).
			Field("user_id", us.ID).Write()
	}

	tok.UserID = us.ID

	usecase.components.Logger.Info().
		Text("The user's authorization has been successfully completed. ").
		Field("token", tokenData).
		Field("username", username).Write()

	return
}

// SetTokenProject - установить проект для токена.
func (usecase *UseCase) SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, tokenID, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The installation of the project for the user token has begun... ").
		Field("token", tokenID).
		Field("project_id", projectID).Write()

	var tok *common_entities.JwtToken

	// Получение токена
	{
		var err error

		if tok, err = usecase.repositories.Authentication.GetTokenByID(ctx, tokenID); err != nil {
			tok = nil

			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("id", tokenID).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.TokenNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's token was received. ").
			Field("token", tok).Write()
	}

	// Проверки
	{
		// Авторизация
		{
			if tok.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token was not authorized, the project selection is not allowed. ").
					Field("token_id", tokenID).
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
					Field("token_id", tokenID).
					Field("project_id", projectID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}

		// Существования проекта и доступа пользователя к нему
		{
			if _, err := usecase.repositories.Projects.GetByID(ctx, projectID); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to get the project by ID: '%s'. ", err).
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

			if exist, err := usecase.repositories.Projects.CheckAccess(ctx, tok.UserID, projectID); err != nil {
				usecase.components.Logger.Error().
					Format("The user's access to the project could not be verified: '%s'. ", err).
					Field("project_id", projectID).
					Field("user_id", tok.UserID).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = error_list.ProjectNotFound()
					cErr.SetError(err)
					return
				}

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			} else if !exist {
				usecase.components.Logger.Error().
					Format("The user does not have access to the project: '%s'. ", err).
					Field("project_id", projectID).
					Field("user_id", tok.UserID).Write()

				cErr = error_list.NotAccessToProject()
				return
			}
		}
	}

	// Установить проект для токена
	{
		var err error

		if err = usecase.repositories.Authentication.SetTokenProject(ctx, tokenID, projectID); err != nil {
			usecase.components.Logger.Error().
				Format("The project value for the user token could not be set: '%s'. ", err).
				Field("token_id", tokenID).
				Field("project_id", projectID).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	usecase.components.Logger.Info().
		Text("The installation of the project for the user's token has been completed successfully. ").
		Field("token", tokenID).
		Field("project_id", projectID).Write()

	return
}

// GetUserProjectsList - получение списка проектов пользователя.
func (usecase *UseCase) GetUserProjectsList(ctx context.Context, tokenID, userID types.ID) (list common_entities.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, tokenID, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The list of user's projects has been started... ").
		Field("token", tokenID).
		Field("user_id", userID).Write()

	var tok *common_entities.JwtToken

	// Получение токена
	{
		var err error

		if tok, err = usecase.repositories.Authentication.GetTokenByID(ctx, tokenID); err != nil {
			tok = nil

			usecase.components.Logger.Error().
				Format("Failed to get token: '%s'. ", err).
				Field("id", tokenID).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.TokenNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's token was received. ").
			Field("token", tok).Write()
	}

	// Проверки
	{
		// Авторизация
		{
			if tok.UserID == 0 {
				usecase.components.Logger.Error().
					Text("The token was not authorized, the project selection is not allowed. ").
					Field("user_id", userID).
					Field("token_id", tokenID).
					Field("project_id", tok.ProjectID).Write()

				cErr = error_list.Unauthorized()
				return
			}
		}

		// Проект уже выбран
		{
			if tok.ProjectID != 0 {
				usecase.components.Logger.Error().
					Text("The project has already been selected, it is not possible to re-select it. ").
					Field("user_id", userID).
					Field("token_id", tokenID).
					Field("project_id", tok.ProjectID).Write()

				cErr = error_list.ProjectHasAlreadyBeenSelected()
				return
			}
		}
	}

	// Получение проектов
	{
		var err error

		if list, err = usecase.repositories.Projects.GetListByUser(ctx, userID); err != nil {
			usecase.components.Logger.Error().
				Format("The list of user's projects could not be retrieved: '%s'. ", err).
				Field("user_id", userID).Write()

			cErr = error_list.ListUserProjectsCouldNotBeRetrieved()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Format("The user has access to '%d' projects ", len(list)).
			Field("user_id", userID).
			Field("projects", list).Write()

	}

	usecase.components.Logger.Info().
		Text("Getting the list of the user's projects has been completed successfully. ").
		Field("token", tokenID).
		Field("user_id", userID).Write()

	return
}

// GetToken - получение jwt токена.
func (usecase *UseCase) GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The receipt of the user's token has begun... ").
		Field("data", data).Write()

	var err error

	if tok, err = usecase.repositories.Authentication.GetToken(ctx, data); err != nil {
		tok = nil

		usecase.components.Logger.Error().
			Format("Failed to get token: '%s'. ", err).
			Field("data", data).Write()

		if errors.Is(err, sql.ErrNoRows) {
			cErr = error_list.TokenNotFound()
			cErr.SetError(err)
			return
		}

		cErr = error_list.InternalServerError()
		cErr.SetError(err)
		return
	}

	usecase.components.Logger.Info().
		Text("The receipt of the user's token has been successfully completed. ").
		Field("token", tok).Write()

	return
}
