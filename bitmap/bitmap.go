package bitindex

type Bitmap interface {
	Add(x uint32) bool
	Del(x uint32) bool
	Contains(x uint32) bool
}

func NewBitmap() Bitmap {
	return newBitIndex()
}
