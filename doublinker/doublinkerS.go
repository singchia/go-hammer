package doublinker

import (
	"errors"
)

type DoublinkerS struct {
	head   *doubnode
	tail   *doubnode
	length int64
}

func NewDoublinkerS() *DoublinkerS {
	return &DoublinkerS{head: nil, tail: nil, length: 0}
}

func (d *DoublinkerS) Length() int64 {
	return d.length
}

//append a node at tail
func (d *DoublinkerS) Add(data interface{}) DoubID {
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

func (d *DoublinkerS) UniqueAdd(data interface{}) (error, DoubID) {
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

func (d *DoublinkerS) UniqueDelete(data interface{}) error {
	if data == nil {
		return errors.New("data empty")
	}

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

func (d *DoublinkerS) UniqueRetrieve(data interface{}) (error, interface{}) {
	for itor := d.head; itor != nil; itor = itor.next {
		dst, ok := itor.data.(HasEqual)
		if ok && dst.Equal(data) {
			return nil, itor.data
		}
	}
	return errors.New("not found"), nil
}

func (d *DoublinkerS) Delete(id DoubID) error {
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

func (d *DoublinkerS) Update(id DoubID, data interface{}) error {
	if id == nil {
		return errors.New("id empty")
	}
	node := (*doubnode)(id)
	node.data = data
	return nil
}

func (d *DoublinkerS) Retrieve(id DoubID) interface{} {
	node := (*doubnode)(id)
	return node.data
}

//move to another doublinker
func (d *DoublinkerS) UniqueMove(data interface{}, dst *DoublinkerS) error {
	if data == nil || dst == nil {
		return errors.New("data or dst doublinker empty")
	}
	var node *doubnode

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
func (d *DoublinkerS) Move(id DoubID, dst *DoublinkerS) error {
	if id == nil || dst == nil {
		return errors.New("id or dst doublinker empty")
	}
	node := (*doubnode)(id)

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

func (d *DoublinkerS) Take(node *doubnode) error {
	if node == nil {
		return errors.New("node empty")
	}

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

func (d *DoublinkerS) Foreachnode(f ForeachnodeFunc) error {
	for itor := d.head; itor != nil; itor = itor.next {
		err := f(DoubID(itor))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DoublinkerS) Foreach(f ForeachFunc) error {
	for itor := d.head; itor != nil; itor = itor.next {
		err := f(itor.data)
		if err != nil {
			return err
		}
	}
	return nil
}
