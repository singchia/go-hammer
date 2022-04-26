package prioqueue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPrioQueue(t *testing.T) {
	pq, err := NewPrioQueue()
	if err != nil {
		t.Error(err)
		return
	}
	pq.PrioPush(1, "foo")
	pq.PrioPush(-1, "bar")
	pq.PrioPush(99, "bza")
	value1 := pq.Pop()
	value2 := pq.Pop()
	value3 := pq.Pop()
	t.Log(value1, value2, value3)
}

func TestPrioQueuePopSync(t *testing.T) {
	// random
	pq, err := NewPrioQueue(OptionQueueLen(10))
	if err != nil {
		t.Error(err)
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := pq.PopSync()
		fmt.Printf("#1 pop 1: %v\n", data)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		data := pq.Pop()
		fmt.Printf("#1 pop 2: %v\n", data)
	}()

	err = pq.PrioPush(1, "foo")
	fmt.Printf("#1 push 1: %v\n", err)
	err = pq.PrioPush(99, "bar")
	fmt.Printf("#1 push 2: %v\n", err)
	err = pq.PrioPush(-1, "bza")
	fmt.Printf("#1 push 3: %v\n", err)
	wg.Wait()

	// first push then pop
	pq, err = NewPrioQueue(OptionQueueLen(1024))
	if err != nil {
		t.Error(err)
		return
	}
	wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		data := pq.PopSync()
		fmt.Printf("#2 pop 1: %v\n", data)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		data := pq.Pop()
		fmt.Printf("#2 pop 2: %v\n", data)
	}()

	err = pq.PrioPush(1, "foo")
	fmt.Printf("#2 push 1: %v\n", err)
	err = pq.PrioPush(99, "bar")
	fmt.Printf("#2 push 2: %v\n", err)
	err = pq.PrioPush(-1, "bza")
	fmt.Printf("#2 push 3: %v\n", err)
	wg.Wait()

}
