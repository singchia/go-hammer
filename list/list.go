package list

type List interface {
	Back() Node
	Front() Node
	InsertAfter(value interface{}, to *Node) *Node
	InsertBefore(value interface{}, to *Node) *Node
	Len() int
	MoveAfter(node, to *Node)
	MoveBefore(node, to *Node)
	MoveToBack(node *Node)
	MoveToFront(node *Node)
	PushBack(node *Node)
	PushBackList(list List)
	PushFront(node *Node)
	PushFrontList(list List)
	Remove(node Node) bool
}

type IsNode interface {
	Value() interface{}
	Prev() Node
	Next() Node
	Detach()
}

type Node struct {
	list       List
	value      interface{}
	prev, next *Node
}

func (node *Node) Value() interface{} {}
