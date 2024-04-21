# ChangeLog:

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