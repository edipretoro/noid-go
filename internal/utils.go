package noid

type Counter struct {
	top   int
	value int
}

func NewCounter(top, value int) *Counter {
	c := Counter{top: top, value: value}
	return &c
}
