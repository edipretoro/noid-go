package noid

type Counter struct {
	top   int
	value int
}

func NewCounter(top, value int) *Counter {
	c := Counter{top: top, value: value}
	return &c
}

type Minter struct {
	oacounter        int
	activeCounters   []Counter
	inactiveCounters []Counter
}

func NewMinter() *Minter {
	return &Minter{
		oacounter: 0,
		activeCounters: make([]Counter, 283),
		inactiveCounters: make([]Counter),
	}
}
