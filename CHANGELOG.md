# ChangeLog:

### v24.0.28:
- Разработка системы i18n;
- Разработка [сервиса i18n](src/internal/services/i18n/service.go);
- Разработка [сервиса управления проектом](src/internal/services/project_manager/service.go);
- Разработка [сервиса управления пользователя](src/internal/services/user_manager/service.go);

---

### v24.0.27:
- Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);

---

### v24.0.26:
- Доработки архитектуры системы;
- Доработки системы ошибок;
- Доработки [системы доступа http rest api](src/internal/common/transports/rest_api/components/access_system);
---

### v24.0.25:
- Доработки архитектуры системы;
- Создание запроса авторизации;
- Проверка и правки по компоненту [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- Доработки [сервиса аутентификации](src/internal/services/authentication/service.go);
- Доработки системы ошибок;

---

### v24.0.24:
- Переработка архитектуры системы, переезд в Docker;

---

### v24.0.23:
- Переработка архитектуры системы;
- project-cli (новое название project manager) теперь микросервис на который направляет запросы [коробка](src/internal/app/box.go);
- Создан сервис для [аутентификации пользователей](src/internal/services/authentication/service.go);
- Доработки [коробки](src/internal/app/box.go) для проксирования на микросервисы;

---

### v24.0.22:
- Исправлен баг в компоненте [ведения журнала](src/pkg/core/components/logger/logger.go);
- Переработка архитектуры системы;
- [Приложение (Коробка)](src/internal/app/box.go) теперь занимается проксированием на микросервисы;
- [init-cli](src/internal/tools/init_cli) перенесен в инструменты;
- Доработки компонента [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go);

---

### v24.0.21:
- Доработать [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.20:
- Доработать [CLI для инициализации коробки](src/internal/system/init_cli);
- Доработать [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.19:
- Доработать [CLI для инициализации коробки](src/internal/system/init_cli);

---

### v24.0.18:
- Переработка [системного скрипта для инициализации коробки](src/internal/system/init_script) в [CLI для инициализации коробки](src/internal/system/init_cli);
- Добавление [CLI для управления проектами](src/internal/system/project_cli);

---

### v24.0.17:
- Доработки [ядра системы](src/pkg/core/core.go);
- Доработки [компонента для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);
- Доработки [инструмента для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);
- Доработки [системного скрипта для инициализации коробки;](src/internal/system/init_script);

---

### v24.0.16:
- Доработки [инструментов для работы с http](src/pkg/http);
- Доработка архитектуры проекта;
- Доработки архитектуры базы данных;
- Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);
- Добавлена [система ошибок](src/pkg/errors/errors.go);

---

### v24.0.15:
- Разработан [скрипт для инициализации системы](src/internal/system/init/script.go) (первый запуск если);
- Доработка архитектуры проекта;
- Доработки архитектуры базы данных;
- Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);
- Добавлено [хранилище файлов системы](src/pkg/core/env/files/files.go) в [окружение](src/pkg/core/env/env.go);

---

### v24.0.14:
- Разработаны модели базы данных:
  - [JWT токен](src/internal/common/objects/db_models/jwt_token.go);
- Разработаны модели:
  - [JWT токен](src/internal/common/objects/models/jwt_token.go);
- Разработаны сущности:
  - [JWT токен](src/internal/common/objects/entities/jwt_token.go);
- Доработка архитектуры проекта;
- Доработки архитектуры базы данных;
- Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);
- Доработки [коннектора для sqlite3 баз данных](src/pkg/databases/connectors/sqlite3/connector.go);
- Доработки [коннектора для postgresql баз данных](src/pkg/databases/connectors/postgresql/connector.go);
- Доработки [коннектора для mysql баз данных](src/pkg/databases/connectors/mysql/connector.go);

---

### v24.0.13:
- Разработаны модели базы данных:
  - [Проект](src/internal/common/objects/db_models/project.go);
  - [Роль](src/internal/common/objects/db_models/role.go);
  - [Http Маршрут](src/internal/common/objects/db_models/http_route.go);
  - [Пользователь](src/internal/common/objects/db_models/user.go);
- Разработаны модели:
  - [Проект](src/internal/common/objects/models/project.go);
  - [Роль](src/internal/common/objects/models/role.go);
  - [Http Маршрут](src/internal/common/objects/models/http_route.go);
  - [Пользователь](src/internal/common/objects/models/user.go);
- Разработаны сущности:
  - [Проект](src/internal/common/objects/entities/project.go);
  - [Роль](src/internal/common/objects/entities/role.go);
  - [Http Маршрут](src/internal/common/objects/entities/http_route.go);
  - [Пользователь](src/internal/common/objects/entities/user.go);
- Добавлен тип данных [ID](src/internal/common/types/id.go) для сущностей и моделей; 
- Доработка архитектуры проекта;
- Доработки архитектуры базы данных;

---

### v24.0.12:
- Доработки компонента [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go);
- Доработки основы [коробки](src/internal/app/box.go);
- Доработка архитектуры проекта;
- Доработка [http rest api коробки](src/internal/app/transports/rest_api/engine.go);
- Доработка системный базы данных коробки;
- Убрал все тесты для дальнейшей переработки;
- Добавлен уровень ведения журнала трессировки для коннекторов баз данных и конфигураций в компоненте [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- Доработки по вызову компонента [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go) во всё проекте;
- Разработан [универсальный коннектор для sql баз данных](src/pkg/databases/connectors/universal_sqlx/connector.go);
- Разработан [коннектор для sqlite3 баз данных](src/pkg/databases/connectors/sqlite3/connector.go);
- Разработан [коннектор для postgresql баз данных](src/pkg/databases/connectors/postgresql/connector.go);
- Разработан [коннектор для mysql баз данных](src/pkg/databases/connectors/mysql/connector.go);

---

### v24.0.11:
- Доработка архитектуры проекта;
- Доработки основы [коробки](src/internal/app/box.go);
- Создание [транспортной части коробки](src/internal/app/transports);
- Создание [http rest api коробки](src/internal/app/transports/rest_api/engine.go);

---

### v24.0.10:
- Доработка архитектуры проекта;
- Доработки основы [коробки](src/internal/app/box.go);
- Разработка [дополнения ядра для использования ключей шифрования в работе системы](src/pkg/core/addons/encryption_keys/encryption_keys.go);
- Разработка [дополнения ядра для записи PID файла системы](src/pkg/core/addons/pid/pid.go);
- Доработки [переменных окружения](src/pkg/core/env/vars/vars.go);

---

### v24.0.9:
- Переработка архитектуры проекта;
- Доработки основы [коробки](src/internal/app/box.go);
- Проектирование базы данных;

---

### v24.0.8:
- Разработана основа [коробки](src/internal/app/box.go);
- Доработки [ядра системы](src/pkg/core/core.go);
- Доработки [компонента для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);
- Доработки [инструмента для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);

---

### v24.0.7:
- Доработки [ядра системы](src/pkg/core/core.go);
- Доработки [компонента для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);

---

### v24.0.6:
- Доработки компонента [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- Разработан [инструмент для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);
- Компонент для [компонент для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go) стал [инструментом](src/pkg/core/tools/closer/closer.go);
- Добавлено взаимодействие с [инструментами ядра](src/pkg/core/tools) в [ядре системы](src/pkg/core/core.go);
- Добавлен уровень ведения журнала трессировки для инструментов ядра в компоненте [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);

---

### v24.0.5:
- Добавлена директория [/sbin](/sbin) в [хранилище путей системы](src/pkg/core/env/paths/paths.go);
- Разработан [компонент для корректного завершения работы ядра системы](src/pkg/core/components/closer/closer.go);
- Сделаны доработки [ядра системы](src/pkg/core/core.go);

---

### v24.0.4:
- Доработки компонента [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go);
- Доработки компонента [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- Разработка [ядра системы](src/pkg/core/core.go)

---

### v24.0.3:
- Разработан компонент [управления конфигурациями проекта](src/pkg/core/components/configurator/configurator.go);

---

### v24.0.2:
- Разработан компонент [ведения журнала трессировки вызовов](src/pkg/core/components/tracer/tracer.go);
- Правки по компоненту [ведения журнала](src/pkg/core/components/logger/logger.go);

---

### v24.0.1:
- Разработана основа [ядра системы](src/pkg/core/core.go);
- Разработан компонент [ведения журнала](src/pkg/core/components/logger/logger.go);

---