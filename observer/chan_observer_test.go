package observer

import (
	"testing"
)

func TestChanObserver(t *testing.T) {
	co := NewChanObserver(ChanObserverOptionMode(ModeEdge))

	type Subject struct {
		observer Observer
	}
	subject := &Subject{co}
	subject.observer.Update(1)
	subject.observer.Update(2)

	data := <-co.C
	t.Log(data.(int))
}
