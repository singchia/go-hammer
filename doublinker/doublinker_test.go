package doublinker

import "testing"

var dl *Doublinker
var gt *testing.T

func Init(t *testing.T) {
	if dl == nil {
		dl = NewDoublinker()
	}
	if gt == nil {
		gt = t
	}
}

func ShowDetails(id DoubID) error {
	node := (*doubnode)(id)
	gt.Logf("%p, %p, %p, %v", node, node.prev, node.next, node.data)
	return nil
}

func Test_Add1Elems(t *testing.T) {
	Init(t)
	dl.Add(1)
	dl.Foreachnode(ShowDetails)
}

func Test_Add2Elems(t *testing.T) {
	Init(t)
	dl.Add(1)
	dl.Add(2)
	dl.Foreachnode(ShowDetails)
}

func Test_Add3Elems(t *testing.T) {
	Init(t)
	dl.Add(1)
	dl.Add(2)
	dl.Add(3)
	dl.Foreachnode(ShowDetails)
}

func Test_Del2elems(t *testing.T) {
	Init(t)
	id1 := dl.Add(1)
	id2 := dl.Add(2)
	dl.Add(3)
	dl.Delete(id1)
	dl.Delete(id2)
	dl.Foreachnode(ShowDetails)
}

func Test_Del3elems(t *testing.T) {
	Init(t)
	id1 := dl.Add(1)
	id2 := dl.Add(2)
	id3 := dl.Add(3)
	dl.Delete(id1)
	dl.Delete(id2)
	dl.Delete(id3)
	dl.Foreachnode(ShowDetails)
}

func Test_Update(t *testing.T) {
	Init(t)
	id1 := dl.Add(1)
	id2 := dl.Add(2)
	dl.Foreachnode(ShowDetails)
	dl.Update(id1, 3)
	dl.Update(id2, 4)
	dl.Foreachnode(ShowDetails)
}

func Test_Retrieve(t *testing.T) {
	Init(t)
	id1 := dl.Add(1)
	data := dl.Retrieve(id1)
	t.Log(data)
	dl.Foreachnode(ShowDetails)
}
