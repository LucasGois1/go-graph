package domain

import "graph/src/value_object"

type Aggregate interface {
	Id() value_object.Identifier[any]
}
