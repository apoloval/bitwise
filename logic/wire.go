package logic

type Wire[T any] interface {
	Sample() T
}

type WireFunc[T any] func() T

func (fn WireFunc[T]) Sample() T {
	return fn()
}

func Fixed[T any](value T) Wire[T] {
	return fixed[T]{value}
}

type fixed[T any] struct {
	value T
}

func (fix fixed[T]) Sample() T {
	return fix.value
}
