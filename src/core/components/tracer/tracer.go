package tracer

// Tracer - компонент для ведения журнала истории вызова.
type Tracer interface {
	FunctionCall(args ...any)
	Error(err error) Tracer
	FunctionCallFinished(args ...any)
}
