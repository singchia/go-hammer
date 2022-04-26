package list

import "testing"

func TestCircuListPush(t *testing.T) {
	clist := NewCircuList()
	for index, word := range wordsAdd {
		clist.PushBack(word)
		all := clist.All()
		if !stringsEqual(all, wordsAdd[0:index+1]) {
			t.Error("wrong elems in list")
			return
		}
		t.Log(index, all)
	}

	clist = NewCircuList()
	for index, word := range wordsAdd {
		clist.PushFront(word)
		all := clist.All()
		if !stringsEqual(all, wordsAdd[0:index+1]) {
			t.Error("wrong elems in list")
			return
		}
		t.Log(index, all)
	}
}

func TestCircuListInsert(t *testing.T) {
	clist := NewCircuList()
	a := clist.PushBack("a")
	clist.InsertAfter("b", a)
	all := clist.All()
	t.Log(all,
		a.Value(),
		a.Next().Value(),
		a.Next().Next().Value(),
		a.Next().Next().Next().Value())
	t.Log(all,
		a.Value(),
		a.Prev().Value(),
		a.Prev().Prev().Value(),
		a.Prev().Prev().Prev().Prev())
}
