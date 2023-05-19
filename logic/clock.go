package logic

type ClockEvent int

const (
	RisingEdge ClockEvent = iota
	FallingEdge
)

type ClockObserver interface {
	OnClockEvent(ev ClockEvent)
}

type ClockObserverFunc func(ClockEvent)

func (fn ClockObserverFunc) OnClockEvent(ev ClockEvent) {
	fn(ev)
}

type Clock struct {
	observers []ClockObserver
}

func (clk *Clock) Register(obs ClockObserver) {
	clk.observers = append(clk.observers, obs)
}

func (clk *Clock) Step(n int) {
	for i := 0; i < n; i++ {
		clk.Cycle()
	}
}

func (clk *Clock) Cycle() {
	for _, obs := range clk.observers {
		obs.OnClockEvent(FallingEdge)
	}
	for _, obs := range clk.observers {
		obs.OnClockEvent(RisingEdge)
	}
}
