package access_system

import (
	"sm-box/internal/services/authentication/transport/servers/http/rest_api/components/access_system/repositories/http_routes"
	"sm-box/internal/services/authentication/transport/servers/http/rest_api/components/access_system/repositories/jwt_tokens"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
)

// Config - конфигурация компонента системы доступа.
type Config struct {
	CookieKeyForSessionToken string `json:"cookie_key_for_session_token" yaml:"CookieKeyForSessionToken" xml:"cookie_key_for_session_token,attr"`
	CookieKeyForAccessToken  string `json:"cookie_key_for_access_token"   yaml:"CookieKeyForAccessToken"  xml:"cookie_key_for_access_token,attr"`

	Repositories *RepositoriesConfig `json:"repositories" yaml:"Repositories" xml:"Repositories"`
}

// RepositoriesConfig - конфигурация репозиториев компонента системы доступа.
type RepositoriesConfig struct {
	HttpRoutes *http_routes_repository.Config `json:"http_routes" yaml:"HttpRoutes" xml:"HttpRoutes"`
	JwtTokens  *jwt_tokens_repository.Config  `json:"jwt_tokens"  yaml:"jwt_tokens" xml:"JwtTokens"`
}

// Read - чтение конфигурации.
func (conf *Config) Read() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		c       configurator.Configurator[*Config]
		profile = configurator.PrivateProfile{
			Dir:      "/components/",
			Filename: "system_access.xml",
		}
	)

	if c, err = configurator.New[*Config](conf); err != nil {
		return
	} else if err = c.Private().Profile(profile).Init(); err != nil {
		return
	}

	if err = conf.FillEmptyFields().Validate(); err != nil {
		return
	}

	return
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Repositories == nil {
		conf.Repositories = new(RepositoriesConfig)
	}

	conf.Repositories.FillEmptyFields()

	return conf
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *RepositoriesConfig) FillEmptyFields() *RepositoriesConfig {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.HttpRoutes == nil {
		conf.HttpRoutes = new(http_routes_repository.Config)
	}

	if conf.JwtTokens == nil {
		conf.JwtTokens = new(jwt_tokens_repository.Config)
	}

	conf.HttpRoutes.FillEmptyFields()
	conf.JwtTokens.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.CookieKeyForSessionToken = "box_session"
	conf.CookieKeyForAccessToken = "box_access"

	conf.Repositories = new(RepositoriesConfig).Default()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *RepositoriesConfig) Default() *RepositoriesConfig {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.HttpRoutes = new(http_routes_repository.Config).Default()
	conf.JwtTokens = new(jwt_tokens_repository.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = conf.Repositories.Validate(); err != nil {
		return
	}

	return
}

// Validate - валидация конфигурации.
func (conf *RepositoriesConfig) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = conf.HttpRoutes.Validate(); err != nil {
		return
	}

	if err = conf.JwtTokens.Validate(); err != nil {
		return
	}

	return
}
