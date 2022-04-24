package list

type List interface {
	All() []interface{}
	Back() *Node
	Front() *Node
	InsertAfter(value interface{}, to *Node) *Node
	InsertBefore(value interface{}, to *Node) *Node
	Len() int
	MoveAfter(node, to *Node)
	MoveBefore(node, to *Node)
	MoveToBack(node *Node)
	MoveToFront(node *Node)
	PushBack(value interface{}) *Node
	PushFront(value interface{}) *Node
	PushBackList(list List)
	PushFrontList(list List)
	Remove(node *Node)
}

type Node struct {
	list       List
	value      interface{}
	prev, next *Node
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

func NewDoubList() List {
	return &doublist{
		nil, nil, 0,
	}
}
