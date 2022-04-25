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
	node.prev = to
	node.next = to.next
	to.next.prev = node
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
	node.next = to
	node.prev = to.prev
	to.prev.next = node
	if list.head == to {
		list.head = node
	}
}

func (list *circulist) MoveToBack(node *Node) {
	if node.list != list || list.tail == node {
		return
	}
	if list.head == node {
		list.head = node.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	list.tail.next = node
	node.prev = list.tail
	node.next = list.tail.next
	list.tail = node
}

func (list *doublist) PushBack(value interface{}) *Node {
	node := &Node{
		list:  list,
		value: value,
		next:  nil,
		prev:  nil,
	}
	if list.length == 0 {
		list.head, list.tail = node, node
	} else {
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.length++
	return node
}
