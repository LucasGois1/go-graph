package value_object

type Identifier[T any] interface {
	Value() T
}
