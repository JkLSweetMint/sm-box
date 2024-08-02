package tracer

const (
	minLevel Level = iota

	LevelMain
	LevelDebug
	LevelInternal
	LevelEvent
	LevelDatabaseConnector
	LevelConfig

	LevelCore
	LevelCoreDebug
	LevelCoreInternal
	LevelCoreEvent
	LevelCoreComponent
	LevelCoreTool

	LevelCoreAddon
	LevelCoreAddonDebug
	LevelCoreAddonInternal
	LevelCoreAddonEvent

	LevelTransport
	LevelTransportDebug
	LevelTransportInternal
	LevelTransportEvent

	LevelTransportGateway
	LevelTransportGatewayDebug
	LevelTransportGatewayInternal
	LevelTransportGatewayEvent

	LevelTransportGrpc
	LevelTransportGatewayGrpc

	LevelTransportHttp
	LevelTransportGatewayHttp

	LevelPackage
	LevelPackageDebug
	LevelPackageInternal
	LevelPackageEvent

	LevelEntity
	LevelEntityDebug
	LevelEntityInternal
	LevelEntityEvent

	LevelConstructor
	LevelConstructorDebug
	LevelConstructorInternal
	LevelConstructorEvent

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

	LevelAdapter
	LevelAdapterDebug
	LevelAdapterInternal
	LevelAdapterEvent

	LevelComponent
	LevelComponentDebug
	LevelComponentInternal
	LevelComponentEvent
)

var (
	levelsList = [...]string{
		"Main",
		"Debug",
		"Internal",
		"Event",
		"DatabaseConnector",
		"Config",

		"Core",
		"CoreDebug",
		"CoreInternal",
		"CoreEvent",
		"CoreComponent",
		"CoreTool",

		"CoreAddon",
		"CoreAddonDebug",
		"CoreAddonInternal",
		"CoreAddonEvent",

		"Transport",
		"TransportDebug",
		"TransportInternal",
		"TransportEvent",

		"TransportGateway",
		"TransportGatewayDebug",
		"TransportGatewayInternal",
		"TransportGatewayEvent",

		"TransportGrpc",
		"TransportGatewayGrpc",

		"TransportHttp",
		"TransportGatewayHttp",

		"Package",
		"PackageDebug",
		"PackageInternal",
		"PackageEvent",

		"Entity",
		"EntityDebug",
		"EntityInternal",
		"EntityEvent",

		"Constructor",
		"ConstructorDebug",
		"ConstructorInternal",
		"ConstructorEvent",

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

		"Adapter",
		"AdapterDebug",
		"AdapterInternal",
		"AdapterEvent",

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
		LevelDatabaseConnector,
		LevelConfig,

		LevelCore,
		LevelCoreDebug,
		LevelCoreInternal,
		LevelCoreEvent,
		LevelCoreComponent,
		LevelCoreTool,

		LevelCoreAddon,
		LevelCoreAddonDebug,
		LevelCoreAddonInternal,
		LevelCoreAddonEvent,

		LevelTransport,
		LevelTransportDebug,
		LevelTransportInternal,
		LevelTransportEvent,

		LevelTransportGateway,
		LevelTransportGatewayDebug,
		LevelTransportGatewayInternal,
		LevelTransportGatewayEvent,

		LevelTransportGrpc,
		LevelTransportGatewayGrpc,

		LevelTransportHttp,
		LevelTransportGatewayHttp,

		LevelPackage,
		LevelPackageDebug,
		LevelPackageInternal,
		LevelPackageEvent,

		LevelEntity,
		LevelEntityDebug,
		LevelEntityInternal,
		LevelEntityEvent,

		LevelConstructor,
		LevelConstructorDebug,
		LevelConstructorInternal,
		LevelConstructorEvent,

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

		LevelAdapter,
		LevelAdapterDebug,
		LevelAdapterInternal,
		LevelAdapterEvent,

		LevelComponent,
		LevelComponentDebug,
		LevelComponentInternal,
		LevelComponentEvent,
	}
)

// Level - уровень ведения журнала трессировки.
type Level int

// String - получение строкового представления уровнял журнала трессировки.
func (l Level) String() string {
	if l > minLevel && int(l) <= len(levelsList) {
		return levelsList[l-1]
	}

	return ""
}
