package hashtable

import (
	"testing"
)

var ht Hashtable
var gt *testing.T

func initial(t *testing.T) {
	if ht == nil {
		ht = NewOAHashtable()
	}
	if gt == nil {
		gt = t
	}
}

func TestSet(t *testing.T) {
	initial(t)
	elemsAdd := []interface{}{
		1,
		"foo",
		struct{}{},
	}
	for _, elem := range elemsAdd {
		ht.Set(elem, "foo")
		ht.Remove(elem)
		_, ok := ht.Get(elem)
		if ok {
			t.Error("elem not deleted")
		}
	}

	for _, elem := range elemsAdd {
		ht.Set(elem, "foo")
	}
	t.Log(ht.Len())
	for _, elem := range elemsAdd {
		ht.Remove(elem)
	}
	t.Log(ht.Len())
}
