package queue

import (
	"container/list"
	"errors"
	"math"
	"sync"
	"sync/atomic"
)

type OptionPrioQueue func(*PrioQueue) error

func OptionQueueLen(length int) OptionPrioQueue {
	return func(pq *PrioQueue) error {
		if length > math.MaxInt32 {
			return errors.New("length too large")
		}
		pq.length = length
		pq.ch = make(chan struct{}, length)
		return nil
	}
}

type prioQueue struct {
	*list.List
	prio  int
	mutex sync.Mutex
}

type PrioQueue struct {
	queues *list.List
	mutex  sync.RWMutex

	elems  int32
	length int
	ch     chan struct{}
}

// default queue priority 1, the higher the value, the higher the priority.
func NewPrioQueue(opts ...OptionPrioQueue) (*PrioQueue, error) {
	pq := &PrioQueue{
		queues: list.New(),
		length: 1024,
		elems:  0,
		ch:     make(chan struct{}, 1024),
	}
	queue := &prioQueue{
		prio: 1,
		List: list.New(),
	}
	pq.queues.PushFront(queue)

	err := error(nil)
	for _, opt := range opts {
		err = opt(pq)
		if err != nil {
			return nil, err
		}
	}
	return pq, nil
}

func (pq *PrioQueue) Len() int {
	return pq.length
}

func (pq *PrioQueue) Available() int {
	elems := atomic.LoadInt32(&pq.elems)
	return pq.length - int(elems)
}

func (pq *PrioQueue) PrioPush(prio int, data interface{}) error {
	select {
	case pq.ch <- struct{}{}:
	default:
		return errors.New("queue full")
	}

	queue := (*prioQueue)(nil)
	pq.mutex.RLock()
	for elem := pq.queues.Front(); elem != nil; elem = elem.Next() {
		value, _ := elem.Value.(*prioQueue)
		if value.prio == prio {
			queue = value
			break
		}
	}
	pq.mutex.RUnlock()

	if queue == nil {
		pq.mutex.Lock()
		for elem := pq.queues.Front(); elem != nil; elem = elem.Next() {
			value, _ := elem.Value.(*prioQueue)
			if value.prio == prio {
				queue = value
				break
			}
		}

		if queue == nil {
			queue = &prioQueue{
				prio: prio,
				List: list.New(),
			}
			inserted := false
			for elem := pq.queues.Front(); elem != nil; elem = elem.Next() {
				value, _ := elem.Value.(*prioQueue)
				if value.prio > prio {
					pq.queues.InsertBefore(queue, elem)
					inserted = true
					break
				}
			}
			if !inserted {
				pq.queues.PushBack(queue)
			}
		}
		pq.mutex.Unlock()
	}
	atomic.AddInt32(&pq.elems, 1)
	queue.mutex.Lock()
	queue.PushFront(data)
	queue.mutex.Unlock()
	return nil
}

func (pq *PrioQueue) Push(data interface{}) error {
	select {
	case pq.ch <- struct{}{}:
	default:
		return errors.New("queue full")
	}

	queue := (*prioQueue)(nil)
	pq.mutex.RLock()
	for elem := pq.queues.Front(); elem != nil; elem = elem.Next() {
		value, _ := elem.Value.(*prioQueue)
		if value.prio == 1 {
			queue = value
			break
		}
	}
	pq.mutex.RUnlock()

	if queue == nil {
		// critical
		return errors.New("queue not available")
	}

	atomic.AddInt32(&pq.elems, 1)
	queue.mutex.Lock()
	queue.PushFront(data)
	queue.mutex.Unlock()
	return nil
}

func (pq *PrioQueue) Pop() interface{} {
	queue := (*prioQueue)(nil)
	pq.mutex.RLock()
	for elem := pq.queues.Back(); elem != nil; elem = elem.Prev() {
		queue, _ = elem.Value.(*prioQueue)
		queue.mutex.Lock()
		if queue.Len() != 0 {
			data := queue.Back()
			value := queue.Remove(data)
			queue.mutex.Unlock()
			pq.mutex.RUnlock()
			atomic.AddInt32(&pq.elems, -1)
			return value
		}
		queue.mutex.Unlock()
	}
	pq.mutex.RUnlock()
	return nil
}

func (pq *PrioQueue) PopSync() interface{} {
	<-pq.ch
	for {
		queue := (*prioQueue)(nil)
		pq.mutex.RLock()
		for elem := pq.queues.Back(); elem != nil; elem = elem.Prev() {
			queue, _ = elem.Value.(*prioQueue)
			queue.mutex.Lock()
			if queue.Len() != 0 {
				data := queue.Back()
				value := queue.Remove(data)
				queue.mutex.Unlock()
				pq.mutex.RUnlock()
				atomic.AddInt32(&pq.elems, -1)
				return value
			}
			queue.mutex.Unlock()
		}
		pq.mutex.RUnlock()
	}
	return nil
}
