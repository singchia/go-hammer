package list

type node struct {
	value      interface{}
	prev, next *node
}

type doublist struct {
	head, tail *node
	length     int
}

func (list *doublist) PushBack(value interface{}) *Node {
	node := &node{
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
	node := &node{
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
