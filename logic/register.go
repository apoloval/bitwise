package logic

type Register[T any] struct {
	value T
	next  T
}

func (reg *Register[T]) Drive(val T) {
	reg.next = val
}

func (reg Register[T]) Sample() T {
	return reg.value
}
func (reg *Register[T]) OnClockEvent(ev ClockEvent) {
	if ev == RisingEdge {
		reg.value = reg.next

	}
}

type TriStateRegister[T comparable] struct {
	Register[TriState[T]]
}

func (reg *TriStateRegister[T]) Drive(val T) {
	reg.Register.Drive(TriState[T]{Driven: true, Value: val})
}

func (reg *TriStateRegister[T]) Undrive() {
	reg.Register.Drive(TriState[T]{Driven: false})
}
