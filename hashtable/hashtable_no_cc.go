package hashtable

import (
	"errors"
	"fmt"
	"hash/fnv"
	"reflect"
	"strconv"
	"sync"

	"github.com/singchia/go-hammer/linker"
)

type Hashable interface {
	Hash() string
}

type node struct {
	key   interface{}
	value interface{}
}

func (n *node) Equal(data interface{}) bool {
	return n.key == (data.(*node)).key
}

var Length uint32 = 64

type Hashtable struct {
	table []*linker.Doublinker
	mutex *sync.RWMutex
}

func NewHashtable() *Hashtable {
	h := &Hashtable{}
	h.table = make([]*linker.Doublinker, 0, Length)
	for i := uint32(0); i < Length; i++ {
		h.table = append(h.table, linker.NewDoublinker())
	}
	h.mutex = new(sync.RWMutex)
	return h
}

func (h *Hashtable) hash(key interface{}) (uint32, error) {
	var str string
	var slice []byte
	switch key.(type) {
	case bool:
		str = strconv.FormatBool(key.(bool))
		slice = []byte(str)
	case float32:
		str = strconv.FormatFloat(float64(key.(float32)), 'f', 10, 64)
		slice = []byte(str)
	case float64:
		str = strconv.FormatFloat(key.(float64), 'f', 10, 64)
		slice = []byte(str)
	case int:
		str = strconv.FormatInt(int64(key.(int)), 10)
		slice = []byte(str)
	case int8:
		str = strconv.FormatInt(int64(key.(int8)), 10)
		slice = []byte(str)
	case int16:
		str = strconv.FormatInt(int64(key.(int16)), 10)
		slice = []byte(str)
	case int32:
		str = strconv.FormatInt(int64(key.(int32)), 10)
		slice = []byte(str)
	case int64:
		str = strconv.FormatInt(key.(int64), 10)
		slice = []byte(str)
	case uint:
		str = strconv.FormatUint(uint64(key.(uint)), 10)
		slice = []byte(str)
	case uint8:
		str = strconv.FormatUint(uint64(key.(uint8)), 10)
		slice = []byte(str)
	case uint16:
		str = strconv.FormatUint(uint64(key.(uint16)), 10)
		slice = []byte(str)
	case uint32:
		str = strconv.FormatUint(uint64(key.(uint32)), 10)
		slice = []byte(str)
	case uint64:
		str = strconv.FormatUint(key.(uint64), 10)
		slice = []byte(str)
	case string:
		str = key.(string)
		slice = []byte(str)
	case []byte:
		slice = key.([]byte)
	case Hashable:
		str = key.(Hashable).Hash()
		slice = []byte(str)
	case uintptr:
		str = fmt.Sprintf("%p", key.(uintptr))
		slice = []byte(str)
	default:
		//reflect
		if reflect.TypeOf(key).Kind() != reflect.Uintptr {
			return 0, errors.New("unhashable key")
		}
		str = fmt.Sprintf("%p", key.(uintptr))
		slice = []byte(str)
	}
	hash := fnv.New32()
	hash.Write(slice)
	return hash.Sum32() % Length, nil
}

func (h *Hashtable) Add(key interface{}, value interface{}) error {
	index, err := h.hash(key)
	if err != nil {
		return err
	}
	n := &node{key: key, value: value}
	_, err = h.table[index].UniqueAdd(n)
	return err
}

func (h *Hashtable) Delete(key interface{}) error {
	index, err := h.hash(key)
	if err != nil {
		return err
	}
	n := &node{key: key}
	return h.table[index].UniqueDelete(n)
}

func (h *Hashtable) Retrieve(key interface{}) (interface{}, error) {
	index, err := h.hash(key)
	if err != nil {
		return nil, err
	}
	n := &node{key: key}
	value, err := h.table[index].UniqueRetrieve(n)
	return value.(*node).value, err
}

func (h *Hashtable) Foreachitem(f ForeachitemFunc) error {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, item := range h.table {
		err := f(item)
		if err != nil {
			return err
		}
	}
	return nil
}

type ForeachitemFunc func(dl *linker.Doublinker) error
