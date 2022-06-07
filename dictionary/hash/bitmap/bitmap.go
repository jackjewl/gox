package bitmap

type Bitmap struct {
	elems []uint8
	size  int
}

func NewBitmap(size int) *Bitmap {
	bitmap := &Bitmap{
		elems: make([]uint8, (size+7)>>3),
		size:  size,
	}
	return bitmap
}

func (this *Bitmap) Exist(key int) bool {
	return this.elems[key>>3]&(0x80>>(key&0x07)) != 0
}

func (this *Bitmap) Set(key int) {
	this.elems[key>>3] |= (0x80 >> (key & 0x07))
}

func (this *Bitmap) Remove(key int) {
	this.elems[key>>3] &= ^(0x80 >> (key & 0x70))
}

func (this *Bitmap) Size() int {
	return this.size
}
