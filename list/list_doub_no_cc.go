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

func (list *doublist) CompareInsert(value interface{},
	compare func(value, next interface{}) int) (*Node, bool) {
	for i, node := 0, list.Front(); i < list.length; i, node = i+1, node.Next() {
		ret := compare(value, node.value)
		if ret == 0 {
			return node, false
		}
		if ret < 0 {
			return list.InsertBefore(value, node), true
		}
	}
	return list.PushBack(value), true
}

func (list *doublist) CompareRemove(value interface{},
	compare func(value, next interface{}) int) bool {
	for i, node := 0, list.Front(); i < list.length; i, node = i+1, node.Next() {
		ret := compare(value, node.value)
		if ret == 0 {
			return list.Remove(node)
		}
	}
	return false
}

func (list *doublist) CompareGet(value interface{},
	compare func(value, next interface{}) int) *Node {
	for i, node := 0, list.Front(); i < list.length; i, node = i+1, node.Next() {
		ret := compare(value, node.value)
		if ret == 0 {
			return node
		}
	}
	return nil
}

func (list *doublist) InsertAfter(value interface{}, to *Node) *Node {
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
	if to.next != nil {
		to.next.prev = node
	}
	to.next = node
	if list.tail == to {
		list.tail = node
	}
	list.length++
	return node
}

func (list *doublist) InsertBefore(value interface{}, to *Node) *Node {
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
	if to.prev != nil {
		to.prev.next = node
	}
	to.prev = node
	if list.head == to {
		list.head = node
	}
	list.length++
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
		to.next.prev = node
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
	if list.tail == node {
		list.tail = node.prev
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
	if list.head == node {
		list.head = node.next
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
	if node.list != list || list.head == node {
		return
	}
	if list.tail == node {
		list.tail = node.prev
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

func (list *doublist) PushFront(value interface{}) *Node {
	node := &Node{
		list:  list,
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

func (list *doublist) PushBackList(other List) {
	for node := other.Front(); node != nil; node = node.Next() {
		list.PushBack(node.value)
	}
}

func (list *doublist) PushFrontList(other List) {
	for node := other.Back(); node != nil; node = node.Prev() {
		list.PushFront(node.value)
	}
}

func (list *doublist) Remove(node *Node) bool {
	if node.list != list {
		return false
	}
	if list.length == 1 {
		list.head, list.tail = nil, nil
	} else if node == list.head {
		node.next.prev = nil
		list.head = node.next
	} else if node == list.tail {
		node.prev.next = nil
		list.tail = node.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.list, node.next, node.prev = nil, nil, nil
	list.length--
	return true
}

func (list *doublist) All() []interface{} {
	all := []interface{}{}
	for node := list.Front(); node != nil; node = node.Next() {
		all = append(all, node.Value())
	}
	return all
}

func (list *doublist) Iterate(cb func(node *Node) bool) {
	for node := list.Front(); node != nil; node = node.Next() {
		if cb(node) == false {
			return
		}
	}
}

func (list *doublist) ReceiveAfter(node *Node, to *Node) {
	if node.list != nil || to.list != list {
		return
	}
	node.list = list
	node.prev = to
	node.next = to.next
	if to.next != nil {
		to.next.prev = node
	}
	to.next = node
	if list.tail == to {
		list.tail = node
	}
	list.length++
}

func (list *doublist) ReceiveBefore(node *Node, to *Node) {
	if node.list != nil || to.list != list {
		return
	}
	node.list = list
	node.prev = to
	node.next = to.next
	if to.prev != nil {
		to.prev.next = node
	}
	to.prev = node
	if list.head == to {
		list.head = node
	}
	list.length++
}

func (list *doublist) ReceiveToBack(node *Node) {
	if list.length == 0 {
		list.head, list.tail = node, node
	} else {
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.length++
}

func (list *doublist) ReceiveToFront(node *Node) {
	if list.length == 0 {
		list.head, list.tail = node, node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.length++
}
