package tracer

const (
	LevelMain Level = iota + 1
	LevelDebug
	LevelInternal
	LevelEvent

	LevelCore
	LevelCoreDebug
	LevelCoreInternal
	LevelCoreEvent
	LevelCoreComponent

	LevelCoreAddon
	LevelCoreAddonDebug
	LevelCoreAddonInternal
	LevelCoreAddonEvent

	LevelCoreTransport
	LevelCoreTransportDebug
	LevelCoreTransportInternal
	LevelCoreTransportEvent

	LevelPackage
	LevelPackageDebug
	LevelPackageInternal
	LevelPackageEvent

	LevelEntity
	LevelEntityDebug
	LevelEntityInternal
	LevelEntityEvent

	LevelRepository
	LevelRepositoryDebug
	LevelRepositoryInternal
	LevelRepositoryEvent

	LevelUseCase
	LevelUseCaseDebug
	LevelUseCaseInternal
	LevelUseCaseEvent

	LevelController
	LevelControllerDebug
	LevelControllerInternal
	LevelControllerEvent

	LevelComponent
	LevelComponentDebug
	LevelComponentInternal
	LevelComponentEvent
)

var (
	allLevelsString = [...]string{
		"Main",
		"Debug",
		"Internal",
		"Event",

		"Core",
		"CoreDebug",
		"CoreInternal",
		"CoreEvent",

		"CoreAddon",
		"CoreAddonDebug",
		"CoreAddonInternal",
		"CoreAddonEvent",

		"CoreTransport",
		"CoreTransportDebug",
		"CoreTransportInternal",
		"CoreTransportEvent",

		"Package",
		"PackageDebug",
		"PackageInternal",
		"PackageEvent",

		"Entity",
		"EntityDebug",
		"EntityInternal",
		"EntityEvent",

		"Repository",
		"RepositoryDebug",
		"RepositoryInternal",
		"RepositoryEvent",

		"UseCase",
		"UseCaseDebug",
		"UseCaseInternal",
		"UseCaseEvent",

		"Controller",
		"ControllerDebug",
		"ControllerInternal",
		"ControllerEvent",

		"Component",
		"ComponentDebug",
		"ComponentInternal",
		"ComponentEvent",
	}
	allLevels = []Level{
		LevelMain,
		LevelDebug,
		LevelInternal,
		LevelEvent,

		LevelCore,
		LevelCoreDebug,
		LevelCoreInternal,
		LevelCoreEvent,

		LevelCoreAddon,
		LevelCoreAddonDebug,
		LevelCoreAddonInternal,
		LevelCoreAddonEvent,

		LevelCoreTransport,
		LevelCoreTransportDebug,
		LevelCoreTransportInternal,
		LevelCoreTransportEvent,

		LevelPackage,
		LevelPackageDebug,
		LevelPackageInternal,
		LevelPackageEvent,

		LevelEntity,
		LevelEntityDebug,
		LevelEntityInternal,
		LevelEntityEvent,

		LevelRepository,
		LevelRepositoryDebug,
		LevelRepositoryInternal,
		LevelRepositoryEvent,

		LevelUseCase,
		LevelUseCaseDebug,
		LevelUseCaseInternal,
		LevelUseCaseEvent,

		LevelController,
		LevelControllerDebug,
		LevelControllerInternal,
		LevelControllerEvent,

		LevelComponent,
		LevelComponentDebug,
		LevelComponentInternal,
		LevelComponentEvent,
	}
)

// Level - уровень ведения журнала трессировки.
type Level int

// String - получение строкового представления уровнял журнала трессировки.
func (level Level) String() string {
	if len(allLevelsString) >= int(level) {
		return allLevelsString[level-1]
	}

	return ""
}
