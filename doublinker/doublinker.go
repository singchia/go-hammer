package doublinker

import (
	"errors"
	"sync"
)

/*
*	I was trying to put Channel and Handler into doubnode,
*	but those should belong to business data, should be
*	done at application layer
**/

//we don't want generate a int64 to index the concrete node
//pointer is the best index
type DoubID *doubnode

type Doublinker struct {
	head   *doubnode
	tail   *doubnode
	length int64
	mutex  sync.RWMutex
}

func NewDoublinker() *Doublinker {
	return &Doublinker{head: nil, tail: nil, length: 0}
}

//when node is deleted
type doubnode struct {
	data interface{}
	next *doubnode
	prev *doubnode
}

//append a node at tail
func (d *Doublinker) Add(data interface{}) DoubID {
	node := &doubnode{data: data, next: nil, prev: nil}
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.length == 0 {
		d.head, d.tail = node, node
		d.length++
		return node
	}
	d.tail.next = node
	node.prev = d.tail
	d.tail = node
	d.length++
	return node
}

func (d *Doublinker) Delete(id DoubID) error {
	if id == nil {
		return errors.New("id empty")
	}
	node := (*doubnode)(id)
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.length == 1 && d.head == node {
		d.head, d.tail = nil, nil
		d.length--
		return nil
	}
	if d.head == node {
		d.head = node.next
		d.head.prev = nil
		d.length--
		return nil
	}
	if d.tail == node {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.length--
		return nil
	}
	if node.prev == nil || node.next == nil {
		return errors.New("isolated node")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	d.length--
	return nil
}

func (d *Doublinker) Update(id DoubID, data interface{}) error {
	if id == nil {
		return errors.New("id empty")
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	node := (*doubnode)(id)
	node.data = data
	return nil
}

func (d *Doublinker) Retrieve(id DoubID) interface{} {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	node := (*doubnode)(id)
	return node.data
}

type ForeachFunc func(id DoubID) error

func (d *Doublinker) Foreach(f ForeachFunc) error {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	for itor := d.head; itor != nil; itor = itor.next {
		err := f(DoubID(itor))
		if err != nil {
			return err
		}
	}
	return nil
}
