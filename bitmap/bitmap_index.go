package bitindex

import (
	"sync/atomic"
	"unsafe"
)

type bitIndex struct {
	pIndexes unsafe.Pointer
}

func newBitIndex() Bitmap {
	return &bitIndex{}
}

func (bi *bitIndex) Add(x uint32) bool {
	indexes := bi.indexes(x)
	return bi.add(indexes)
}

func (bi *bitIndex) Del(x uint32) bool {
	return bi.del(bi.indexes(x))
}

func (bi *bitIndex) Contains(x uint32) bool {
	return bi.contains(bi.indexes(x))
}

func (bi *bitIndex) indexes(x uint32) []uint32 {
	last := (x << 26) >> 26
	third := (x << 18) >> 24
	second := (x << 10) >> 24
	first := x >> 22
	return []uint32{first, second, third, last}
}

func (bi *bitIndex) add(indexes []uint32) bool {
	if len(indexes) == 1 {
		offset := indexes[0]
		bm := uint64(0)
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
			pBM := unsafe.Pointer(&bm)
			atomic.CompareAndSwapPointer(&bi.pIndexes, unsafe.Pointer(nil), pBM)
		}
		pBM := atomic.LoadPointer(&bi.pIndexes)
		pOld := (*uint64)(pBM)
		old := atomic.LoadUint64(pOld)
		if old&(1<<offset) == (1 << offset) {
			return false
		}
		swapped := false
		for !swapped {
			new := old | (1 << offset)
			swapped = atomic.CompareAndSwapUint64(pOld, old, new)
			if swapped {
				return true
			} else {
				old = atomic.LoadUint64(pOld)
				if old&(1<<offset) == (1 << offset) {
					return false
				}
			}
		}
	} else {
		index := indexes[0]
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
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
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
			return false
		}
		pBM := atomic.LoadPointer(&bi.pIndexes)
		pOld := (*uint64)(pBM)
		old := atomic.LoadUint64(pOld)
		if !(old&(1<<offset) == (1 << offset)) {
			return false
		}
		swapped := false
		for !swapped {
			new := old & ^(1 << offset)
			swapped = atomic.CompareAndSwapUint64(pOld, old, new)
			if swapped {
				return true
			} else {
				old = atomic.LoadUint64(pOld)
				if !(old&(1<<offset) == (1 << offset)) {
					return false
				}
			}
		}
	} else {
		index := indexes[0]
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
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
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
			return false
		}
		pBM := atomic.LoadPointer(&bi.pIndexes)
		pOld := (*uint64)(pBM)
		old := atomic.LoadUint64(pOld)
		if !(old&(1<<offset) == (1 << offset)) {
			return false
		}
		return true
	} else {
		index := indexes[0]
		if atomic.LoadPointer(&bi.pIndexes) == unsafe.Pointer(nil) {
			return false
		}
		pIndexes := atomic.LoadPointer(&bi.pIndexes)
		return (*(*[]bitIndex)(pIndexes))[index].contains(indexes[1:])
	}
}
