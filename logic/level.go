package logic

type Level bool

const (
	Low  Level = false
	High Level = true
)

type TriState[T comparable] struct {
	Value  T
	Driven bool
}

func Value[T comparable](value T) TriState[T] {
	return TriState[T]{Value: value, Driven: true}
}

func Z[T comparable]() TriState[T] {
	return TriState[T]{Driven: false}
}

func (z TriState[T]) Is(value T) bool {
	return z.Driven && z.Value == value
}

func (z TriState[T]) IsZ() bool {
	return !z.Driven
}

type TriStateLevel = TriState[Level]
