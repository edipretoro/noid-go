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
	total            int
	prefix           string
	mask             string
	template         string
}

func (m *Minter) initCounters() {
	return
}

func (m *Minter) parseTemplate(template string) {
	return
}

func (m *Minter) Mint() string {
	return ""
}

func NewMinter() *Minter {
	return &Minter{
		oacounter:        0,
		activeCounters:   make([]Counter, 283),
		inactiveCounters: make([]Counter),
	}
}
