package postman

import (
	"encoding/json"
	"errors"
)

type AuthMechanism string

const (
	// InheritAuthFromParent - наследование механизма авторизации от родителя, является значением по умолчанию.
	InheritAuthFromParent = ""

	// NoAuth - не требуется авторизация.
	NoAuth AuthMechanism = "noauth"

	// Basic - базовый механизм авторизации, для авторизации используется логин и пароль.
	Basic AuthMechanism = "basic"

	// Bearer - механизм авторизации, который использует подписанный сервером токен (bearer token).
	Bearer AuthMechanism = "bearer"

	// Digest - один из общепринятых методов, используемых веб-сервером для обработки учётных
	// данных пользователя веб-браузера.
	Digest AuthMechanism = "digest"

	// OAuth1 - открытый стандарт авторизации, который позволяет пользователям выдавать доступ к своим данным в
	// соцсетях, почте или других приложениях, чтобы зайти на определённый сайт или ресурс.
	OAuth1 AuthMechanism = "oauth1"

	// OAuth2 - протокол авторизации, который используется для управления доступом к данным пользователя на одном
	// сервисе, который, в свою очередь, предоставляется клиентским приложениям на другом сервисе. Клиентскими
	// приложениями могут быть веб-сервисы, мобильные или десктопные приложения, а сервисами могут быть различные
	// платформы.
	OAuth2 AuthMechanism = "oauth2"

	// Hawk - схема аутентификации HTTP, которая использует алгоритм MAC для частичного криптографического
	// подтверждения запроса.
	Hawk AuthMechanism = "hawk"

	// AWSV4 - дополнительная мера безопасности для учётной записи Amazon. Она использует комбинацию механизмов,
	// таких как пароли, токены и одноразовые коды безопасности, для дополнительной защиты ресурсов облачных вычислений.
	AWSV4 AuthMechanism = "awsv4"

	// NTLM - используется в Windows-среде для аутентификации пользователей, устройств и сервисов в
	// доменной и недоменной инфраструктуре.
	NTLM AuthMechanism = "ntlm"

	// APIKey - механизм безопасности, который используется поставщиками API
	// для мониторинга и управления своими сервисами.
	//
	// API Key Authentication помогает гарантировать, что только авторизованные пользователи могут получить
	// доступ к API. Без правильной аутентификации злоумышленник может легко получить доступ к конфиденциальным данным.
	APIKey AuthMechanism = "apikey"

	// EdgeGrid - вспомогательный инструмент авторизации, разработанный и используемый компанией Akamai.
	EdgeGrid AuthMechanism = "edgegrid"

	// ASAP - токен на предъявителя JSON Web Token (JWT), который сервер API может использовать для аутентификации
	// запросов от клиента.
	ASAP AuthMechanism = "asap"
)

// AuthParam - представляет собой атрибут для любого метода аутентификации, предоставляемого Postman.
// Например, "username" и "password" задаются в качестве атрибутов auth для базового метода аутентификации.
type AuthParam struct {
	Key   string `json:"key,omitempty"`
	Value any    `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

// Auth - содержит используемый метод аутентификации и связанные с ним параметры.
type Auth authV210

type authV210 struct {
	version  version
	Type     AuthMechanism `json:"type,omitempty"`
	NoAuth   []*AuthParam  `json:"noauth,omitempty"`
	Basic    []*AuthParam  `json:"basic,omitempty"`
	Bearer   []*AuthParam  `json:"bearer,omitempty"`
	Digest   []*AuthParam  `json:"digest,omitempty"`
	OAuth1   []*AuthParam  `json:"oauth1,omitempty"`
	OAuth2   []*AuthParam  `json:"oauth2,omitempty"`
	Hawk     []*AuthParam  `json:"hawk,omitempty"`
	AWSV4    []*AuthParam  `json:"awsv4,omitempty"`
	NTLM     []*AuthParam  `json:"ntlm,omitempty"`
	APIKey   []*AuthParam  `json:"apikey,omitempty"`
	EdgeGrid []*AuthParam  `json:"edgegrid,omitempty"`
	ASAP     []*AuthParam  `json:"asap,omitempty"`
}

type authV200 struct {
	Type     AuthMechanism  `json:"type,omitempty"`
	NoAuth   map[string]any `json:"noauth,omitempty"`
	Basic    map[string]any `json:"basic,omitempty"`
	Bearer   map[string]any `json:"bearer,omitempty"`
	Digest   map[string]any `json:"digest,omitempty"`
	OAuth1   map[string]any `json:"oauth1,omitempty"`
	OAuth2   map[string]any `json:"oauth2,omitempty"`
	Hawk     map[string]any `json:"hawk,omitempty"`
	AWSV4    map[string]any `json:"awsv4,omitempty"`
	NTLM     map[string]any `json:"ntlm,omitempty"`
	APIKey   map[string]any `json:"apikey,omitempty"`
	EdgeGrid map[string]any `json:"edgegrid,omitempty"`
	ASAP     map[string]any `json:"asap,omitempty"`
}

// mAuth is used for marshalling/unmarshalling.
type mAuth struct {
	Type     AuthMechanism   `json:"type,omitempty"`
	NoAuth   json.RawMessage `json:"noauth,omitempty"`
	Basic    json.RawMessage `json:"basic,omitempty"`
	Bearer   json.RawMessage `json:"bearer,omitempty"`
	Digest   json.RawMessage `json:"digest,omitempty"`
	OAuth1   json.RawMessage `json:"oauth1,omitempty"`
	OAuth2   json.RawMessage `json:"oauth2,omitempty"`
	Hawk     json.RawMessage `json:"hawk,omitempty"`
	AWSV4    json.RawMessage `json:"awsv4,omitempty"`
	NTLM     json.RawMessage `json:"ntlm,omitempty"`
	APIKey   json.RawMessage `json:"apikey,omitempty"`
	EdgeGrid json.RawMessage `json:"edgegrid,omitempty"`
	ASAP     json.RawMessage `json:"asap,omitempty"`
}

func (a *Auth) setVersion(v version) {
	a.version = v
}

// GetParams - возвращает параметры, относящиеся к используемому методу аутентификации.
func (a Auth) GetParams() []*AuthParam {
	switch a.Type {
	case NoAuth:
		return a.NoAuth
	case Basic:
		return a.Basic
	case Bearer:
		return a.Bearer
	case Digest:
		return a.Digest
	case OAuth1:
		return a.OAuth1
	case OAuth2:
		return a.OAuth2
	case Hawk:
		return a.Hawk
	case AWSV4:
		return a.AWSV4
	case NTLM:
		return a.NTLM
	case APIKey:
		return a.APIKey
	case EdgeGrid:
		return a.EdgeGrid
	case ASAP:
		return a.ASAP
	}

	return nil
}

func (a *Auth) setParams(params []*AuthParam) {
	switch a.Type {
	case NoAuth:
		a.NoAuth = params
	case Basic:
		a.Basic = params
	case Bearer:
		a.Bearer = params
	case Digest:
		a.Digest = params
	case OAuth1:
		a.OAuth1 = params
	case OAuth2:
		a.OAuth2 = params
	case Hawk:
		a.Hawk = params
	case AWSV4:
		a.AWSV4 = params
	case NTLM:
		a.NTLM = params
	case APIKey:
		a.APIKey = params
	case EdgeGrid:
		a.EdgeGrid = params
	case ASAP:
		a.ASAP = params
	}
}

// UnmarshalJSON - анализирует данные, закодированные в формате JSON, и создает на их основе Auth.
// В зависимости от версии коллекции Postman свойство auth может быть либо массивом, либо объектом.
//   - v2.1.0 : Array
//   - v2.0.0 : Object
func (a *Auth) UnmarshalJSON(b []byte) (err error) {
	var tmp mAuth
	err = json.Unmarshal(b, &tmp)

	a.Type = tmp.Type

	if a.NoAuth, err = unmarshalAuthParam(tmp.NoAuth); err != nil {
		return
	}
	if a.Basic, err = unmarshalAuthParam(tmp.Basic); err != nil {
		return
	}
	if a.Bearer, err = unmarshalAuthParam(tmp.Bearer); err != nil {
		return
	}
	if a.Digest, err = unmarshalAuthParam(tmp.Digest); err != nil {
		return
	}
	if a.OAuth1, err = unmarshalAuthParam(tmp.OAuth1); err != nil {
		return
	}
	if a.OAuth2, err = unmarshalAuthParam(tmp.OAuth2); err != nil {
		return
	}
	if a.Hawk, err = unmarshalAuthParam(tmp.Hawk); err != nil {
		return
	}
	if a.AWSV4, err = unmarshalAuthParam(tmp.AWSV4); err != nil {
		return
	}
	if a.NTLM, err = unmarshalAuthParam(tmp.NTLM); err != nil {
		return
	}
	if a.APIKey, err = unmarshalAuthParam(tmp.APIKey); err != nil {
		return
	}
	if a.EdgeGrid, err = unmarshalAuthParam(tmp.EdgeGrid); err != nil {
		return
	}
	if a.ASAP, err = unmarshalAuthParam(tmp.ASAP); err != nil {
		return
	}

	return
}

func unmarshalAuthParam(b []byte) (a []*AuthParam, err error) {
	if len(b) > 0 {
		if b[0] == '{' { // v2.0.0
			var tmp map[string]string
			json.Unmarshal(b, &tmp)
			for k, v := range tmp {
				a = append(a, &AuthParam{
					Key:   k,
					Value: v,
				})
			}
		} else if b[0] == '[' { // v2.1.0
			json.Unmarshal(b, &a)
		} else {
			err = errors.New("Unsupported type")
		}
	}

	return
}

// MarshalJSON - возвращает кодировку автора в формате JSON.
// Если версия - v2.0.0, то она возвращается как объект, в противном случае - как массив (v2.1.0).
func (a Auth) MarshalJSON() ([]byte, error) {

	if a.version == V200 {
		return json.Marshal(authV200{
			Type:     a.Type,
			NoAuth:   authParamsToMap(a.NoAuth),
			Basic:    authParamsToMap(a.Basic),
			Bearer:   authParamsToMap(a.Bearer),
			Digest:   authParamsToMap(a.Digest),
			OAuth1:   authParamsToMap(a.OAuth1),
			OAuth2:   authParamsToMap(a.OAuth2),
			Hawk:     authParamsToMap(a.Hawk),
			AWSV4:    authParamsToMap(a.AWSV4),
			NTLM:     authParamsToMap(a.NTLM),
			APIKey:   authParamsToMap(a.APIKey),
			EdgeGrid: authParamsToMap(a.EdgeGrid),
			ASAP:     authParamsToMap(a.ASAP),
		})
	}

	return json.Marshal(authV210{
		Type:     a.Type,
		NoAuth:   a.NoAuth,
		Basic:    a.Basic,
		Bearer:   a.Bearer,
		Digest:   a.Digest,
		OAuth1:   a.OAuth1,
		OAuth2:   a.OAuth2,
		Hawk:     a.Hawk,
		AWSV4:    a.AWSV4,
		NTLM:     a.NTLM,
		APIKey:   a.APIKey,
		EdgeGrid: a.EdgeGrid,
		ASAP:     a.ASAP,
	})
}

func authParamsToMap(authParams []*AuthParam) map[string]any {
	authParamsMap := make(map[string]any)

	for _, authParam := range authParams {
		authParamsMap[authParam.Key] = authParam.Value
	}

	return authParamsMap
}

// NewAuth создает новую структуру Auth с заданными параметрами.
func NewAuth(a AuthMechanism, params ...*AuthParam) *Auth {
	auth := &Auth{
		Type: a,
	}
	auth.setParams(params)
	return auth
}

// NewAuthParam создает новый параметр Auth типа string.
func NewAuthParam(key string, value string) *AuthParam {
	return &AuthParam{
		Key:   key,
		Value: value,
		Type:  "string",
	}
}
