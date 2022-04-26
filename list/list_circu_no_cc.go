package list

type circulist struct {
	head, tail *Node
	length     int
}

func (list *circulist) Back() *Node {
	return list.tail
}

func (list *circulist) Front() *Node {
	return list.head
}

func (list *circulist) InsertAfter(value interface{}, to *Node) *Node {
	if to.list != list {
		return nil
	}
	node := &Node{
		list:  list,
		value: value,
		next:  nil,
		prev:  nil,
	}
	node.prev = to
	node.next = to.next
	to.next.prev = node
	to.next = node
	if list.tail == to {
		list.tail = node
	}
	return node
}

func (list *circulist) InsertBefore(value interface{}, to *Node) *Node {
	if to.list != list {
		return nil
	}
	node := &Node{
		list:  list,
		value: value,
		next:  nil,
		prev:  nil,
	}
	node.next = to
	node.prev = to.prev
	to.prev.next = node

	if list.head == to {
		list.head = node
	}
	return node
}

func (list *circulist) Len() int {
	return list.length
}

func (list *circulist) MoveAfter(node, to *Node) {
	if node.list != list || to.list != list ||
		node == to || to.next == node {
		return
	}
	if list.head == node {
		list.head = node.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	// attach
	node.prev = to
	node.next = to.next
	to.next.prev = node
	to.next = node
	if list.tail == to {
		list.tail = node
	}
}

func (list *circulist) MoveBefore(node, to *Node) {
	if node.list != list || to.list != list ||
		node == to || to.prev == node {
		return
	}
	if list.tail == node {
		list.tail = node.prev
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	// attach
	node.next = to
	node.prev = to.prev
	to.prev.next = node
	to.prev = node
	if list.head == to {
		list.head = node
	}
}

func (list *circulist) MoveToBack(node *Node) {
	if node.list != list || list.tail == node {
		return
	}
	// detach
	node.prev.next = node.next
	node.next.prev = node.prev
	// attach
	node.prev = list.tail
	node.next = list.tail.next
	list.tail.next = node
	list.tail = node
	list.head.prev = list.tail
}

func (list *circulist) MoveToFront(node *Node) {
	if node.list != list || list.head == node {
		return
	}
	// detach
	node.prev.next = node.next
	node.next.prev = node.prev
	// attach
	node.next = list.head
	node.prev = list.head.prev
	list.head.prev = node
	list.head = node
	list.tail.next = list.head
}

func (list *circulist) PushBack(value interface{}) *Node {
	node := &Node{
		list:  list,
		value: value,
		next:  nil,
		prev:  nil,
	}
	if list.length == 0 {
		list.head, list.tail = node, node
		node.next, node.prev = node, node
	} else {
		// attach
		node.prev = list.tail
		node.next = list.tail.next
		list.tail.next = node
		list.tail = node
		list.head.prev = list.tail
	}
	list.length++
	return node
}

func (list *circulist) PushFront(value interface{}) *Node {
	node := &Node{
		list:  list,
		value: value,
		next:  nil,
		prev:  nil,
	}
	if list.length == 0 {
		list.head, list.tail = node, node
		node.next, node.prev = node, node
	} else {
		node.next = list.head
		node.prev = list.head.prev
		list.head.prev = node
		list.head = node
		list.tail.next = list.head
	}
	list.length++
	return node
}

func (list *circulist) PushBackList(other List) {
	for node := other.Front(); node != nil; node = node.Next() {
		list.PushBack(node.value)
	}
}

func (list *circulist) PushFrontList(other List) {
	for node := other.Back(); node != nil; node = node.Prev() {
		list.PushFront(node.value)
	}
}

func (list *circulist) Remove(node *Node) {
	if node.list != list {
		return
	}
	if list.length == 1 {
		list.head, list.tail = nil, nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.list, node.next, node.prev = nil, nil, nil
	list.length--
}

func (list *circulist) All() []interface{} {
	all := []interface{}{}
	for i, node := 0, list.Front(); i < list.length; i, node = i+1, node.Next() {
		all = append(all, node.Value())
	}
	return all
}

func (list *circulist) Iterate(iterator Iterator) {
	for i, node := 0, list.Front(); i < list.length; i, node = i+1, node.Next() {
		if iterator(node) == false {
			return
		}
	}
}

func (list *circulist) ReceiveAfter(node, to *Node) {
}

func (list *circulist) ReceiveBefore(node, to *Node) {
}

func (list *circulist) ReceiveToBack(node *Node) {
}

func (list *circulist) ReceiveToFront(node *Node) {
}
