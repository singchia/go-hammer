package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	n := 65536
	flag.IntVar(&n, "n", 65535, "number")
	flag.Parse()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	rb := roaring.New()
	for i := 0; i < n; i++ {
		rb.Add(uint32(i))
	}

	runtime.ReadMemStats(&m2)
	fmt.Printf("m2-m1: [n:%10d] [alloc:%10d] [totalalloc:%10d] "+
		"[mallocs:%10d] [frees:%10d] [heapalloc:%10d] [headobjects:%10d]\n",
		n, m2.Alloc-m1.Alloc, m2.TotalAlloc-m1.TotalAlloc,
		m2.Mallocs-m1.Mallocs, m2.Frees-m1.Frees,
		m2.HeapAlloc-m1.HeapAlloc, m2.HeapObjects-m1.HeapObjects)
}
