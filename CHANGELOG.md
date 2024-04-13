# ChangeLog:

### v24.0.6:
- Доработки компонента [ведения журнала трессировки вызовов](src/core/components/tracer/tracer.go);
- Разработан [инструмент для управления задачами ядра системы](src/core/tools/task_scheduler/task_scheduler.go);
- Компонент для [компонент для корректного завершения работы ядра системы](src/core/components/closer/closer.go) стал [инструментом](src/core/tools/closer/closer.go);
- Добавлено взаимодействие с [инструментами ядра](src/core/tools) в [ядре системы](src/core/core.go);
- Добавлен уровень ведения журнала трессировки для инструментов ядра в компоненте [ведения журнала трессировки вызовов](src/core/components/tracer/tracer.go);

---

### v24.0.5:
- Добавлена директория [/sbin](/sbin) в [хранилище путей системы](src/core/env/paths/paths.go);
- Разработан [компонент для корректного завершения работы ядра системы](src/core/components/closer/closer.go);
- Сделаны доработки [ядра системы](src/core/core.go);

---

### v24.0.4:
- Доработки компонента [управления конфигурациями проекта](src/core/components/configurator/configurator.go);
- Доработки компонента [ведения журнала трессировки вызовов](src/core/components/tracer/tracer.go);
- Разработка [ядра системы](src/core/core.go)

---

### v24.0.3:
- Разработан компонент [управления конфигурациями проекта](src/core/components/configurator/configurator.go);

---

### v24.0.2:
- Разработан компонент [ведения журнала трессировки вызовов](src/core/components/tracer/tracer.go);
- Правки по компоненту [ведения журнала](src/core/components/logger/logger.go);

---

### v24.0.1:
- Разработана основа [ядра системы](src/core/core.go);
- Разработан компонент [ведения журнала](src/core/components/logger/logger.go);

---