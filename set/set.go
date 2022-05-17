package set

type Set interface {
	Add(interface{})
	Remove(interface{})
	Clear()
	Contains(interface{}) bool
	IsEmpty() bool
	ToSlice() []interface{}
	Size() int
}
