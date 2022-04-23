package list

type doublist struct {
	head, tail *Node
	length     int
}

func (list *doublist) Back() *Node {
	return list.tail
}

func (list *doublist) Front() *Node {
	return list.head
}

func (list *doublist) PushBack(value interface{}) *Node {
	node := &Node{
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

func (list *doublist) PushHead(value interface{}) *Node {
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
	if list.length == 0 {
		list.head, list.tail = node, node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.length++
	return node
}

func (list *doublist) InsertAfter(value interface{}, to *Node) *Node {
	if to.list != list {
		return nil
	}
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
	node.prev = to
	node.next = to.next
	if to.next != nil {
		to.next.prev = node
	}
	to.next = node
	return node
}

func (list *doublist) InsertBefore(value interface{}, to *Node) *Node {
	if to.list != list {
		return nil
	}
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
	node.next = to
	node.prev = to.prev
	if to.prev != nil {
		to.prev.next = node
	}
	to.prev = node
	return node
}

func (list *doublist) Len() int {
	return list.length
}

func (list *doublist) MoveAfter(node, to *Node) {
	if node.list != list || to.list != list ||
		node == to || to.next == node {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = to
	node.next = to.next
	if to.next != nil {
		too.next.prev = node
	}
	to.next = node

	if list.tail == to {
		list.tail = node
	}
}

func (list *doublist) MoveBefore(node, to *Node) {
	if node.list != list || to.list != list ||
		node == to || to.prev == node {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.next = to
	node.prev = to.prev
	if to.prev != nil {
		to.prev.next = node
	}
	to.prev = node
	if list.head == to {
		list.head = node
	}
}

func (list *doublist) MoveToBack(node *Node) {
	if node.list != list || list.tail == node {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	list.tail.next = node
	node.prev = list.tail
	node.next = nil
	list.tail = node
}

func (list *doublist) MoveToFront(node *Node) {
	if node.list != list || list.tail == node {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	list.head.prev = node
	node.next = list.head
	node.prev = nil
	list.head = node
}

func (list *doublist) PushBack(node *Node) {
}
