# ToDo:

### Global:
- [ ] Добавить валидацию конфигураций;
- [ ] Разработать компонент для управления событиями системы;
- [ ] Доработки тестов системы;

---

### v24.0.17:
- [x] Доработки [ядра системы](src/pkg/core/core.go);
- [x] Доработки [компонента для корректного завершения работы ядра системы](src/pkg/core/tools/closer/closer.go);
- [x] Доработки [инструмента для управления задачами ядра системы](src/pkg/core/tools/task_scheduler/task_scheduler.go);
- [x] Доработки [системного скрипта для инициализации коробки;](src/internal/system/init_script);

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
  - [x] [JWT токен](src/internal/common/db_models/jwt_token.go);
- [x] Разработать модели:
  - [x] [JWT токен](src/internal/common/models/jwt_token.go);
- [x] Разработать сущности:
  - [x] [JWT токен](src/internal/common/entities/jwt_token.go);
- [x] Доработки архитектуры проекта;
- [x] Доработки архитектуры базы данных;
- [x] Доработки по компоненту [системы доступа http rest api](src/internal/app/transports/rest_api/components/access_system);

---

### v24.0.13:
- [x] Разработать модели базы данных:
    - [x] [Проект](src/internal/common/db_models/project.go);
    - [x] [Роль](src/internal/common/db_models/role.go);
    - [x] [Http Маршрут](src/internal/common/db_models/http_route.go);
    - [x] [Пользователь](src/internal/common/db_models/user.go);
- [x] Разработать модели:
    - [x] [Проект](src/internal/common/models/project.go);
    - [x] [Роль](src/internal/common/models/role.go);
    - [x] [Http Маршрут](src/internal/common/models/http_route.go);
    - [x] [Пользователь](src/internal/common/models/user.go);
- [x] Разработать сущности:
    - [x] [Проект](src/internal/common/entities/project.go);
    - [x] [Роль](src/internal/common/entities/role.go);
    - [x] [Http Маршрут](src/internal/common/entities/http_route.go);
    - [x] [Пользователь](src/internal/common/entities/user.go);
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