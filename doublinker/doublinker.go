package doublinker

import (
	"errors"
	"sync"
)

type HasEqual interface {
	Equal(src interface{}) bool
}

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

func (d *Doublinker) Length() int64 {
	return d.length
}

//when node is deleted
type doubnode struct {
	data interface{}
	next *doubnode
	prev *doubnode
}

//append a node at tail
func (d *Doublinker) Add(data interface{}) DoubID {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	node := &doubnode{data: data, next: nil, prev: nil}

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

func (d *Doublinker) UniqueAdd(data interface{}) (error, DoubID) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	for itor := d.head; itor != nil; itor = itor.next {
		dst, ok := itor.data.(HasEqual)
		if ok && dst.Equal(data) {
			return errors.New("alrealy exists"), nil
		}
	}
	node := &doubnode{data: data, next: nil, prev: nil}

	if d.length == 0 {
		d.head, d.tail = node, node
		d.length++
		return nil, node
	}
	d.tail.next = node
	node.prev = d.tail
	d.tail = node
	d.length++
	return nil, node

}

func (d *Doublinker) UniqueDelete(data interface{}) error {
	if data == nil {
		return errors.New("data empty")
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.length == 0 {
		return errors.New("empty doublinker")
	}

	dst, ok := d.head.data.(HasEqual)
	if d.length == 1 && ok && dst.Equal(data) {
		d.head, d.tail = nil, nil
		d.length--
		return nil
	}
	if ok && dst.Equal(data) {
		d.head = d.head.next
		d.head.prev = nil
		d.length--
		return nil
	}

	dst, ok = d.tail.data.(HasEqual)
	if ok && dst.Equal(data) {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.length--
		return nil
	}
	//not first and last
	for itor := d.head; itor != nil; itor = itor.next {
		dst, ok := itor.data.(HasEqual)
		if ok && dst.Equal(data) {
			itor.prev.next = itor.next
			itor.next.prev = itor.prev
			d.length--
			return nil
		}
	}
	return errors.New("not found")
}

func (d *Doublinker) UniqueRetrieve(data interface{}) (error, interface{}) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	for itor := d.head; itor != nil; itor = itor.next {
		dst, ok := itor.data.(HasEqual)
		if ok && dst.Equal(data) {
			return nil, itor.data
		}
	}
	return errors.New("not found"), nil
}

func (d *Doublinker) Delete(id DoubID) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if id == nil {
		return errors.New("id empty")
	}

	if d.length == 0 {
		return errors.New("linker empty")
	}

	node := (*doubnode)(id)

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

//move to another doublinker
func (d *Doublinker) UniqueMove(data interface{}, dst *Doublinker) error {
	if data == nil || dst == nil {
		return errors.New("data or dst doublinker empty")
	}
	var node *doubnode
	d.mutex.Lock()
	defer d.mutex.Unlock()

	for itor := d.head; itor != nil; itor = itor.next {
		dst, ok := itor.data.(HasEqual)
		if ok && dst.Equal(data) {
			node = itor
			break
		}
	}

	if d.length == 1 && d.head == node {
		d.head, d.tail = nil, nil
		d.length--

	} else if d.head == node {
		d.head = node.next
		d.head.prev = nil
		d.length--

	} else if d.tail == node {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.length--

	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		d.length--
	}

	node.prev = nil
	node.next = nil

	return dst.Take(node)
}

//move to another doublinker
func (d *Doublinker) Move(id DoubID, dst *Doublinker) error {
	if id == nil || dst == nil {
		return errors.New("id or dst doublinker empty")
	}
	node := (*doubnode)(id)
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.length == 1 && d.head == node {
		d.head, d.tail = nil, nil
		d.length--

	} else if d.head == node {
		d.head = node.next
		d.head.prev = nil
		d.length--

	} else if d.tail == node {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.length--

	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		d.length--
	}

	return dst.Take(node)
}

func (d *Doublinker) Take(node *doubnode) error {
	if node == nil {
		return errors.New("node empty")
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.length == 0 {
		d.head, d.tail = node, node
		d.length++
		return nil
	}
	d.tail.next = node
	node.prev = d.tail
	d.tail = node
	d.length++
	return nil

}

type ForeachFunc func(data interface{}) error
type ForeachnodeFunc func(id DoubID) error

func (d *Doublinker) Foreachnode(f ForeachnodeFunc) error {
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

func (d *Doublinker) Foreach(f ForeachFunc) error {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	for itor := d.head; itor != nil; itor = itor.next {
		err := f(itor.data)
		if err != nil {
			return err
		}
	}
	return nil
}
