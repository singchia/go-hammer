package circulinker

import (
	"errors"
	"sync"
)

type CircuID *circunode

type Circulinker struct {
	head   *circunode
	tail   *circunode
	length int
	cur    *circunode
	mutex  sync.RWMutex
}

func NewCirculinker() *Circulinker {
	return &Circulinker{}
}

type circunode struct {
	data interface{}
	next *circunode
}

func (c *Circulinker) Add(data interface{}) CircuID {
	node := &circunode{data: data, next: nil}
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.length == 0 {
		//first time to append node
		c.head, c.tail = node, node
		c.tail.next = c.head
		c.cur = node
		c.length++
		return node
	}

	c.tail.next = node
	c.tail = c.tail.next
	c.tail.next = c.head
	c.length++
	return node

}

func (c *Circulinker) Delete(id CircuID) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var nextItor *circunode
	itor := c.tail
	node := (*circunode)(id)
	if c.length == 1 && node == c.head {
		c.head, c.tail = nil, nil
		c.cur = nil
		c.length--
		return nil
	}

	if node == c.head {
		c.tail.next = c.head.next
		c.head = c.head.next
		if c.cur == node {
			c.cur = c.cur.next
		}
		c.length--
		return nil
	}
	itor = c.head

	for i := 1; i < c.length; i++ {
		nextItor = itor.next
		if nextItor == node {
			//if the node is current pointed node, then move forward
			//if c.length == 0, don't need to specify corresponding action,
			//it will be repointed at AppendNode
			if c.cur == node {
				c.cur = c.cur.next
			}
			//if the node is tail, tail should move backward
			if c.tail == node {
				c.tail = itor
			}
			itor.next = nextItor.next
			c.length--
			return nil
		}
		itor = itor.next
	}
	return errors.New("no such node")

}

func (c *Circulinker) Update(id CircuID, data interface{}) bool {
	if id == nil {
		return false
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	node := (*circunode)(id)
	node.data = data
	return true
}

func (c *Circulinker) Retrieve(id CircuID) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	node := (*circunode)(id)
	return node.data
}

func (c *Circulinker) RetrieveCur() interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.length == 0 {
		return nil
	}
	node := *c.cur
	return node.data

}
func (c *Circulinker) Rightshift() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.length == 0 {
		return errors.New("empty linker")
	}
	c.cur = c.cur.next
	return nil
}

type ForeachFunc func(data interface{}) error
type ForeachnodeFunc func(id CircuID) error

func (c *Circulinker) Foreachnode(f ForeachnodeFunc) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	itor := c.head
	for i := 0; i < c.length; i++ {
		err := f(CircuID(itor))
		if err != nil {
			return err
		}
		itor = itor.next
	}
	return nil
}
func (c *Circulinker) Foreach(f ForeachFunc) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	itor := c.head
	for i := 0; i < c.length; i++ {
		err := f(itor.data)
		if err != nil {
			return err
		}
		itor = itor.next
	}
	return nil
}
