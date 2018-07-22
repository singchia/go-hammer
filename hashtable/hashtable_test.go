package hashtable

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/singchia/go-hammer/doublinker"
)

var ht *Hashtable
var gt *testing.T

func initial(t *testing.T) {
	if ht == nil {
		Length = 2
		ht = NewHashtable()
	}
	if gt == nil {
		gt = t
	}
}

func ShowItem(dl *doublinker.Doublinker) error {
	dl.Foreachnode(ShowNode)
	gt.Logf("^^^^^^^^^^^^^^^^^^")
	return nil
}

func ShowNode(di doublinker.DoubID) error {
	rid := reflect.Indirect(reflect.ValueOf(di)).Interface()
	str := fmt.Sprintf("%v", rid)
	str = strings.Trim(str, "{")
	str = strings.Trim(str, "}")
	strs := strings.Split(str, " ")

	gt.Logf("%v, %p, %v, %s\n", strs[2], di, strs[1], strs[0])
	return nil
}

func Test_Add(t *testing.T) {
	initial(t)
	err := ht.Add(1, 1)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(2, 2)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(3, 3)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(3, 4)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(4, 4)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(5, 5)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")
}

func Test_Delete(t *testing.T) {
	initial(t)
	err := ht.Add(1, 1)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(2, 2)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(3, 3)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Delete(1)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")
}

func Test_Retrieve(t *testing.T) {
	initial(t)
	err := ht.Add(1, 1)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(2, 2)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}
	ht.Foreachitem(ShowItem)
	gt.Log("-----------------\n")

	err = ht.Add(3, 3)
	if err != nil {
		t.Log("singchia watching: ", err.Error())
	}

	err, value := ht.Retrieve(1)
	gt.Log(value)

	err, value = ht.Retrieve(2)
	gt.Log(value)

	err, value = ht.Retrieve(3)
	gt.Log(value)
}
