package list

type List interface {
	Back() *Node
	Front() *Node
	Len() int
	MoveAfter(node, to *Node)
	MoveBefore(node, to *Node)
	MoveToBack(node *Node)
	MoveToFront(node *Node)
	CompareInsert(value interface{}, compare func(value, next interface{}) int) (*Node, bool)
	CompareRemove(value interface{}, compare func(value, next interface{}) int) bool
	CompareGet(value interface{}, compare func(value, next interface{}) int) *Node
	InsertAfter(value interface{}, to *Node) *Node
	InsertBefore(value interface{}, to *Node) *Node
	PushBack(value interface{}) *Node
	PushFront(value interface{}) *Node
	PushBackList(list List)
	PushFrontList(list List)
	Remove(node *Node) bool
	All() []interface{}
	Iterate(cb func(node *Node) bool)
	ReceiveAfter(node *Node, to *Node)
	ReceiveBefore(node *Node, to *Node)
	ReceiveToBack(node *Node)
	ReceiveToFront(node *Node)
}

type Node struct {
	list       List
	value      interface{}
	prev, next *Node
}

func (node *Node) SetValue(value interface{}) {
	node.value = value
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Detach() {
	node.list.Remove(node)
}

func (node *Node) DetachTo(other List) {
	node.list.Remove(node)
	other.ReceiveToBack(node)
}

func NewDoubList() List {
	return &doublist{
		nil, nil, 0,
	}
}

func NewCircuList() List {
	return &circulist{
		nil, nil, 0,
	}
}
