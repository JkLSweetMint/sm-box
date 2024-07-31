# ToDo:

### Global:
- [ ] Добавить валидацию конфигураций;
- [ ] Разработать компонент для управления событиями системы;
- [ ] Доработки тестов системы;
- [ ] Доработать систему ошибок;
- [ ] Сделать проверку на доступ к короткому маршруту;
 
---

### v24.0.49:
- [x] Разработка [сервиса коротких ссылок](src/internal/services/url_shortner/service.go);

---

### v24.0.48:
- [x] Разработка [сервиса коротких ссылок](src/internal/services/url_shortner/service.go);

---

### v24.0.47:
- [x] Переработка системы доступа;
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.46:
- [x] Разработка [сервиса коротких ссылок](src/internal/services/url_shortner/service.go);
- [x] Доработки [сервиса i18n](src/internal/services/i18n/service.go);
- [x] Проектирование базы данных для dashboard;
- [x] Переработка системы доступа;
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.45:
- [x] Разработка [сервиса коротких ссылок](src/internal/services/url_shortner/service.go);
- [x] Доработки [сервиса i18n](src/internal/services/i18n/service.go);
- [x] Проектирование базы данных для dashboard;

---

### v24.0.44:
- [x] Разработка [сервиса ресурсов](src/internal/services/resources/service.go);
- [x] Разработка [сервиса коротких ссылок](src/internal/services/url_shortner/service.go);

---

### v24.0.43:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Навести порядок в коде;

---

### v24.0.42:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.42:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Добавление [контейнера панели управления](src/docker/Dockerfile.dashboard);

---

### v24.0.40:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.39:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Добавления redis в архитектуру;
- [x] Реализация [коннектора для redis](src/pkg/databases/connectors/redis);

---

### v24.0.38:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.37:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.36:
- [x] Доработки [приложения](src/internal/app/box.go)
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Разработка [сервиса пользователей](src/internal/services/users/service.go);
- [x] Доработки взаимодействия сервисов через grpc;
- [x] Доработки системы ошибок;

---

### v24.0.35:
- [x] Доработки [приложения](src/internal/app/box.go)
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Доработки взаимодействия сервисов через grpc;

---

### v24.0.34:
- [x] Доработки системы доступа через nginx auth module;
- [x] Доработки [приложения](src/internal/app/box.go)
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Доработки архитектуры базы данных;

---

### v24.0.33:
- [x] Доработки системы доступа через nginx auth module;
- [x] Доработки [приложения](src/internal/app/box.go)
- [x] Доработки [сервиса управления пользователями](src/internal/services/users/service.go)
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Доработки архитектуры базы данных;

---

### v24.0.32:
- [x] Доработки системы доступа через nginx auth module;
- [x] Добавление кэширования в nginx;
- [x] Создание [сервиса управления пользователями](src/internal/services/users/service.go)
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.31:
- [x] Переработка всей системы;
- [x] Реализация системы доступа через nginx auth module;

---

### v24.0.30:
- [x] Разработка системы i18n;
- [x] Разработка [сервиса i18n](src/internal/services/i18n/service.go);

---

### v24.0.29:
- [x] Разработка системы i18n;
- [x] Разработка [сервиса i18n](src/internal/services/i18n/service.go);

---

### v24.0.28:
- [x] Разработка системы i18n;
- [x] Разработка [сервиса i18n](src/internal/services/i18n/service.go);
- [x] Разработка [сервиса управления проектом](src/internal/services/project_manager/service.go);
- [x] Разработка [сервиса управления пользователя](src/internal/services/user_manager/service.go);

---

### v24.0.27:
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---


### v24.0.26:
- [x] Доработки архитектуры системы;
- [x] Доработки системы ошибок;
- [x] Доработки [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);

---

### v24.0.25:
- [x] Доработки архитектуры системы;
- [x] Создание запроса авторизации;
- [x] Проверка и правки по компоненту [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- [x] Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- [x] Доработки системы ошибок;

---

### v24.0.24:
- [x] Переработка архитектуры системы, переезд в Docker;

---

### v24.0.23:
- [x] Переработка архитектуры системы;
- [x] Доработки [коробки](src/internal/app/box.go) для проксирования на микросервисы;
- [x] Создать сервис для [аутентификации пользователей](src/internal/services/authentication/service.go);

---

### v24.0.22:
- [x] Переработка архитектуры системы;

---

### v24.0.21:
- [x] Доработки [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.20:
- [x] Доработки [CLI для инициализации коробки](src/internal/system/init_cli);
- [x] Доработки [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.19:
- [x] Доработки [CLI для инициализации коробки](src/internal/system/init_cli);

---

### v24.0.18:
- [x] Переработка [системного скрипта для инициализации коробки](src/internal/system/init_script) в [CLI для инициализации коробки](src/internal/system/init_cli);
- [x] Добавление [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.17:
- [x] Доработки [ядра системы](src/pkg/core/core.go);
- [x] Доработки [компонента для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);
- [x] Доработки [инструмента для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);
- [x] Доработки [системного скрипта для инициализации коробки](src/internal/system/init_script);

---

### v24.0.16:
- [x] Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);
- [x] Доработки [инструментов для работы с http](src/pkg/http);
- [x] Добавить [систему ошибок](src/pkg/errors/errors.go);

---

### v24.0.15:
- [x] Разработать [скрипт для инициализации системы](src/internal/system/init/script.go) (первый запуск если);
- [x] Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);
- [x] Добавить [хранилище файлов системы](src/pkg/core/env/files/files.go) в [окружение](src/pkg/core/env/env.go);

---

### v24.0.14:
- [x] Разработать модели базы данных:
  - [x] [JWT токен](src/internal/common/objects/db_models/jwt_token.go);
- [x] Разработать модели:
  - [x] [JWT токен](src/internal/common/objects/models/jwt_token.go);
- [x] Разработать сущности:
  - [x] [JWT токен](src/internal/common/objects/entities/jwt_token.go);
- [x] Доработки архитектуры проекта;
- [x] Доработки архитектуры базы данных;
- [x] Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);

---

### v24.0.13:
- [x] Разработать модели базы данных:
    - [x] [Проект](src/internal/common/objects/db_models/project.go);
    - [x] [Роль](src/internal/common/objects/db_models/role.go);
    - [x] [Http Маршрут](src/internal/common/objects/db_models/http_route.go);
    - [x] [Пользователь](src/internal/common/objects/db_models/user.go);
- [x] Разработать модели:
    - [x] [Проект](src/internal/common/objects/models/project.go);
    - [x] [Роль](src/internal/common/objects/models/role.go);
    - [x] [Http Маршрут](src/internal/common/objects/models/http_route.go);
    - [x] [Пользователь](src/internal/common/objects/models/user.go);
- [x] Разработать сущности:
    - [x] [Проект](src/internal/common/objects/entities/project.go);
    - [x] [Роль](src/internal/common/objects/entities/role.go);
    - [x] [Http Маршрут](src/internal/common/objects/entities/http_route.go);
    - [x] [Пользователь](src/internal/common/objects/entities/user.go);
- [x] Доработки архитектуры проекта;
- [x] Доработки архитектуры базы данных;

---

### v24.0.12:
- [x] Доработать архитектуру проекта;
- [x] Доработать основу [коробки](src/internal/app/box.go);
- [x] Разработать [универсальный коннектор для sql баз данных](src/pkg/databases/connectors/universal_sqlx/connector.go);
- [x] Разработать [коннектор для sqlite3 баз данных](src/pkg/databases/connectors/sqlite3/connector.go);
- [x] Разработать [коннектор для postgresql баз данных](src/pkg/databases/connectors/postgresql/connector.go);
- [x] Разработать [коннектор для mysql баз данных](src/pkg/databases/connectors/mysql/connector.go);

---

### v24.0.11:
- [x] Доработать архитектуру проекта;
- [x] Доработать основу [коробки](src/internal/app/box.go);
- [x] Сделать [транспортную часть коробки](src/internal/app/transports);
- [x] Сделать [http rest api коробки](src/internal/app/transports/rest_api/engine.go);

---

### v24.0.10:
- [x] Доработать архитектуру проекта;
- [x] Доработать основу [коробки](src/internal/app/box.go);
- [x] Разработать [дополнение ядра для использования ключей шифрования в работе системы](src/pkg/core/addons/encryption_keys/encryption_keys.go);
- [x] Разработать [дополнение ядра для записи PID файла системы](src/pkg/core/addons/pid/pid.go);

---

### v24.0.9:
- [x] Доработки основы [коробки](src/internal/app/box.go);
- [x] Проектирование базы данных;

---

### v24.0.8:
- [x] Разработка основы [коробки](src/internal/app/box.go);

---

### v24.0.7:
- [x] Доработки [ядра системы](src/pkg/core/core.go);

---

### v24.0.6:
- [x] Разработать [инструмент для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);

---

### v24.0.5:
- [x] Разработать [компонент для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);
- [x] Доработки [ядра системы](src/pkg/core/core.go);

---

### v24.0.4:
- [x] Разработка [ядра системы](src/pkg/core/core.go)
- [x] Прикрутить использования компонента [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go) для компонентов [ведения журнала](src/pkg/core/components/logger/logger.go) и [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);

---

### v24.0.3:
- [x] Разработать компонент [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go);

---

### v24.0.2:
- [x] Разработать компонент [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);

---

### v24.0.1:
- [x] Разработать основу [ядра системы](src/pkg/core/core.go);
- [x] Разработать компонент [ведения журнала](src/pkg/core/components/logger/logger.go);

---