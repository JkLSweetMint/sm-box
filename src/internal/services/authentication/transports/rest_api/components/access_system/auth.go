package access_system

// BasicUserAuthData - данные базовой авторизации.
type BasicUserAuthData struct {
	Username string `json:"username" xml:"Username"`
	Password string `json:"password" xml:"Password"`
}
