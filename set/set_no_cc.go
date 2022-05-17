package set

type MapSet struct {
	set map[interface{}]bool
}

func NewMapSet() Set {
	return &MapSet{
		set: make(map[interface{}]bool),
	}
}

func (ms *MapSet) Add(v interface{}) {
	ms.set[v] = true
}

func (ms *MapSet) Remove(v interface{}) {
	delete(ms.set, v)
}

func (ms *MapSet) Clear() {
	ms.set = make(map[interface{}]bool)
}

func (ms *MapSet) Contains(v interface{}) bool {
	_, ok := ms.set[v]
	return ok
}

func (ms *MapSet) IsEmpty() bool {
	if len(ms.set) == 0 {
		return true
	}
	return false
}

func (ms *MapSet) ToSlice() []interface{} {
	elems := make([]interface{}, len(ms.set))
	index := 0
	for elem, _ := range ms.set {
		elems[index] = elem
		index++
	}
	return elems
}

func (ms *MapSet) Size() int {
	return len(ms.set)
}
