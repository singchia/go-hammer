package set

import (
	"math/rand"
	"testing"
)

func TestMapSet(t *testing.T) {
	set := NewMapSet()
	elemsAdd := []interface{}{
		1,
		"foo",
		struct{}{},
	}
	for _, elem := range elemsAdd {
		set.Add(elem)
	}
	index := rand.Intn(len(elemsAdd))
	ok := set.Contains(elemsAdd[index])
	if !ok {
		t.Log("elem not contained")
	}
}
