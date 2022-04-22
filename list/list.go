package list

type List interface {
	Back() Node
	Front() Node
	InsertAfter(value interface{}, node *Node) bool
	InsertBefore(value interface{}, node *Node) bool
	Len() int
	MoveAfter(value interface{}, node *Node)
	MoveBefore(value interface{}, node *Node)
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
