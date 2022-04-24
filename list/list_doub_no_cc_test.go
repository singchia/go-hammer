package list

import (
	"math/rand"
	"testing"
	"time"
)

var (
	wordsAdd = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}
)

func stringsEqual(a []interface{}, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, aElem := range a {
		found := false
		for _, bElem := range b {
			if aElem.(string) == bElem {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestDoubListAdd(t *testing.T) {
	dlist := NewDoubList()
	for index, word := range wordsAdd {
		dlist.PushBack(word)
		all := dlist.All()
		if !stringsEqual(all, wordsAdd[0:index+1]) {
			t.Error("wrong elems in list")
			return
		}
		t.Log(index, all)
	}
}

func TestDoubListRemove(t *testing.T) {
	for i := 1; i < len(wordsAdd); i++ {
		dlist := NewDoubList()
		rand.Seed(time.Now().Unix())
		index := rand.Intn(i)
		node := (*Node)(nil)
		for j, word := range wordsAdd[0:i] {
			tmp := dlist.PushBack(word)
			if index == j {
				node = tmp
			}
		}
		dlist.Remove(node)
		all := dlist.All()
		for _, elem := range all {
			if node.Value() == elem {
				t.Error("elem not deleted")
				return
			}
		}
		t.Log(all)
	}
}
