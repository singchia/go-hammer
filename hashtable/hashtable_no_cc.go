package hashtable

import (
	"github.com/singchia/go-hammer/list"
)

type kv struct {
	keyhash uint64
	key     interface{}
	value   interface{}
}

func kvcompare(value, next interface{}) int {
	valuekv := value.(*kv)
	otherkv := next.(*kv)

	if valuekv.keyhash == otherkv.keyhash {
		if valuekv.key == otherkv.key {
			return 0
		}
		return 1
	} else if valuekv.keyhash < otherkv.keyhash {
		return -1
	} else {
		return 1
	}
}

type OAHashtableOption func(*OAHashtable)

func OptionOAHashtableLength(length uint32) OAHashtableOption {
	return func(h *OAHashtable) {
		h.length = length
	}
}

func OptionOAHashtableHash(hash func(key interface{}) uint64) OAHashtableOption {
	return func(h *OAHashtable) {
		h.hash = hash
	}
}

// Open addressing hash table
type OAHashtable struct {
	table  []list.List
	length uint32
	hash   Hash
}

func NewOAHashtable(options ...OAHashtableOption) Hashtable {
	h := &OAHashtable{
		length: 1024,
		hash:   SamplingHash,
	}
	for _, option := range options {
		option(h)
	}
	h.table = make([]list.List, h.length)
	for i := uint32(0); i < h.length; i++ {
		h.table[i] = list.NewDoubList()
	}
	return h
}

func (h *OAHashtable) Set(key interface{}, value interface{}) {
	hash := h.hash(key)
	tophash := hash >> 32
	bottomhash := hash << 32 >> 32
	index := bottomhash % uint64(h.length)
	valuekv := &kv{keyhash: tophash, key: key, value: value}
	node, ok := h.table[index].CompareInsert(valuekv, kvcompare)
	if !ok {
		node.SetValue(valuekv)
	}
}

func (h *OAHashtable) Remove(key interface{}) bool {
	hash := h.hash(key)
	tophash := hash >> 32
	bottomhash := hash << 32 >> 32
	index := bottomhash % uint64(h.length)
	valuekv := &kv{keyhash: tophash, key: key}
	return h.table[index].CompareRemove(valuekv, kvcompare)
}

func (h *OAHashtable) Get(key interface{}) (interface{}, bool) {
	hash := h.hash(key)
	tophash := hash >> 32
	bottomhash := hash << 32 >> 32
	index := bottomhash % uint64(h.length)
	valuekv := &kv{keyhash: tophash, key: key}
	node := h.table[index].CompareGet(valuekv, kvcompare)
	if node == nil {
		return nil, false
	}
	return node.Value().(*kv).value, true
}

func (h *OAHashtable) Iterate(cb func(key, value interface{}) bool) {
	for _, item := range h.table {
		item.Iterate(func(node *list.Node) bool {
			kv := node.Value().(*kv)
			return cb(kv.key, kv.value)
		})
	}
}

func (h *OAHashtable) Len() int {
	count := 0
	for _, list := range h.table {
		count += list.Len()
	}
	return count
}
