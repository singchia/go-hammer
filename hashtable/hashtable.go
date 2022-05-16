package hashtable

type Hashtable interface {
	Set(key, value interface{})
	Remove(key interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Iterate(cb func(key, value interface{}) bool)
	Len() int
}

type Hash func(interface{}) uint64

type HasHash interface {
	Hash() uint64
}
