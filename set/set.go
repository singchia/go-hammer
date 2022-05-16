package set

type Set interface {
	Add(interface{})
	Clear()
	Contains()
	IsEmpty() bool
	ToSlice() []interface{}
	Size() int
}
