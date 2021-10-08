package observer

type Mode int

const (
	ModeLevel = iota
	ModeEdge
)

type ChanObserverOption func(*ChanObserver)

func ChanObserverOptionMode(mode Mode) ChanObserverOption {
	return func(co *ChanObserver) {
		co.mode = mode
	}
}

type ChanObserver struct {
	mode Mode
	C    chan interface{}
}

func NewChanObserver(options ...ChanObserverOption) *ChanObserver {
	co := &ChanObserver{
		mode: ModeLevel,
		C:    make(chan interface{}, 1),
	}
	for _, option := range options {
		option(co)
	}
	return co
}

func (co *ChanObserver) Update(data interface{}) {
	if co.mode == ModeLevel {
		co.C <- data
		return
	}

	select {
	case co.C <- data:
	default:
	}
}
