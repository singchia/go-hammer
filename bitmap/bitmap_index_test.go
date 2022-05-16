package bitindex

import (
	"testing"

	"github.com/RoaringBitmap/roaring"
)

func TestBitIndex(t *testing.T) {
	bi := newBitIndex()
	ok := bi.Add(1024)
	t.Log(ok)
	ok = bi.Add(1024)
	t.Log(ok)
	ok = bi.Contains(1024)
	t.Log(ok)
	ok = bi.Contains(1025)
	t.Log(ok)
	ok = bi.Del(1024)
	t.Log(ok)
	ok = bi.Contains(1024)
	t.Log(ok)
}

func BenchmarkBitIndex(b *testing.B) {
	bi := newBitIndex()
	for i := 0; i < b.N; i++ {
		bi.Add(uint32(i))
	}
}

func BenchmarkRoaring(b *testing.B) {
	rb := roaring.New()
	for i := 0; i < b.N; i++ {
		rb.Add(uint32(i))
	}
}
