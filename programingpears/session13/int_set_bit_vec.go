package pearls

const SHIFT = 5
const MASK = 0x1f
const BITSPERWORD = 32

type IntSetBitVec struct {
	buff   []int
	size   int
	maxVal int
}

func NewIntSetBitVec(maxElements, maxVal int) *IntSetBitVec {
	return &IntSetBitVec{
		buff:   make([]int, (maxVal/BITSPERWORD + 1)),
		maxVal: maxVal,
	}
}

func (i *IntSetBitVec) Insert(t int) {
	i.insert(t)
}

func (i *IntSetBitVec) insert(t int) {
	if !i.test(t) {
		i.size++
		i.buff[t>>SHIFT] |= (1 << (t & MASK))
	}
}

func (i *IntSetBitVec) erase(t int) {
	i.buff[t>>SHIFT] &= ^(1 << (t & MASK))
}

func (i *IntSetBitVec) test(t int) bool {
	return i.buff[t>>SHIFT]&(1<<(t&MASK)) > 0
}

func (i *IntSetBitVec) Size() int {
	return i.size
}

func (i *IntSetBitVec) Report() []int {
	dst := make([]int, 0)
	for index, word := range i.buff {
		for j := 0; j < BITSPERWORD; j++ {
			if word&(1<<j) > 0 {
				dst = append(dst, index<<SHIFT+j)
			}
		}
	}
	return dst
}
