package bitmap

import (
	"fmt"
	"runtime"
	"testing"
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
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	bi := newBitIndex()
	for i := 0; i < b.N; i++ {
		bi.Add(uint32(i))
	}
	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}
