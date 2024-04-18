package app

import (
	"sm-box/src/internal/app/transports/graphql"
)

// Transports - описание транспортной части коробки.
type Transports interface {
	Graphql() graphql.Engine
}

// components - транспортная часть коробки.
type transports struct {
	graphql graphql.Engine
}

// Graphql - получение graphql http api.
func (t *transports) Graphql() graphql.Engine {
	return t.graphql
}
