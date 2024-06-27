package tracer

// Tracer - компонент ведения журнала трессировки вызовов функций/методов.
//
// Пример:
//
//	defer func() {
//		var trc = New(LevelMain)
//
//		trc.FunctionCall(123)
//		trc.FunctionCallFinished(321)
//	}()
//
// С ошибкой:
//
//	defer func() {
//		var trc = New(LevelMain)
//
//		trc.FunctionCall(123)
//		trc.Error(errors.New("Test")).FunctionCallFinished(321)
//	}()
type Tracer interface {
	FunctionCall(args ...any)
	Error(err error) Tracer
	FunctionCallFinished(args ...any)
}

// New - создание компонента ведения журнала трессировки.
func New(levels ...Level) (t Tracer) {
	t = &tracer{
		params: make([]any, maxParams),
		levels: levels,
	}

	return
}
