package bitmap

import (
	"sync/atomic"
	"unsafe"
)

type bitIndex struct {
	pIndexes unsafe.Pointer // *[]bitIndex
	end      uint64
}

func newBitIndex() *bitIndex {
	return &bitIndex{}
}

// 2^(10+8+8+6)
func (bi *bitIndex) Add(x uint32) bool {
	return bi.add(bi.indexes(x))
}

func (bi *bitIndex) Del(x uint32) bool {
	return bi.del(bi.indexes(x))
}

func (bi *bitIndex) Contains(x uint32) bool {
	return bi.contains(bi.indexes(x))
}

func (bi *bitIndex) indexes(x uint32) []uint32 {
	first := x >> 22
	second := (x << 10) >> 24
	third := (x << 18) >> 24
	end := (x << 28) >> 26
	return []uint32{first, second, third, end}
}

func (bi *bitIndex) add(indexes []uint32) bool {
	if len(indexes) == 1 {
		offset := indexes[0]
		old := atomic.LoadUint64(&bi.end)
		if old&(1<<offset) == (1 << offset) {
			return false
		}
		swapped := false
		for !swapped {
			new := old | (1 << offset)
			swapped = atomic.CompareAndSwapUint64(&bi.end, old, new)
			if swapped {
				return true
			} else {
				old = atomic.LoadUint64(&bi.end)
				if old&(1<<offset) == (1 << offset) {
					return false
				}
			}
		}
	} else {
		index := indexes[0]
		if bi.pIndexes == nil {
			indexes := make([]bitIndex, 1024)
			pIndexes := unsafe.Pointer(&indexes)
			atomic.CompareAndSwapPointer(&bi.pIndexes, unsafe.Pointer(nil), pIndexes)
		}
		pIndexes := atomic.LoadPointer(&bi.pIndexes)
		return (*(*[]bitIndex)(pIndexes))[index].add(indexes[1:])
	}
	return false
}

func (bi *bitIndex) del(indexes []uint32) bool {
	if len(indexes) == 1 {
		offset := indexes[0]
		old := atomic.LoadUint64(&bi.end)
		if !(old&(1<<offset) == (1 << offset)) {
			// already set
			return false
		}
		swapped := false
		for !swapped {
			new := old & ^(1 << offset)
			swapped = atomic.CompareAndSwapUint64(&bi.end, old, new)
			if swapped {
				return true
			} else {
				old = atomic.LoadUint64(&bi.end)
				if !(old&(1<<offset) == (1 << offset)) {
					// already set
					return false
				}
			}
		}
	} else {
		index := indexes[0]
		if bi.pIndexes == nil {
			return false
		}
		pIndexes := atomic.LoadPointer(&bi.pIndexes)
		return (*(*[]bitIndex)(pIndexes))[index].del(indexes[1:])
	}
	return false
}

func (bi *bitIndex) contains(indexes []uint32) bool {
	if len(indexes) == 1 {
		offset := indexes[0]
		old := atomic.LoadUint64(&bi.end)
		if !(old&(1<<offset) == (1 << offset)) {
			return false
		}
		return true
	} else {
		index := indexes[0]
		if bi.pIndexes == nil {
			return false
		}
		pIndexes := atomic.LoadPointer(&bi.pIndexes)
		return (*(*[]bitIndex)(pIndexes))[index].contains(indexes[1:])
	}
}
