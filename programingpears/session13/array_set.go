package pearls

type IntSetArr struct {
	array []int
	size  int
}

func NewIntSetArr(maxElements, maxVal int) *IntSetArr {
	array := make([]int, maxElements)

	array[0] = maxVal
	return &IntSetArr{array: array}
}

func (i *IntSetArr) Insert(t int) {
	var index int
	for index = 0; i.array[index] < t; index++ {
	}
	if i.array[index] == t {
		return
	}
	// 因为有 sentinel(哨兵)，所以初始值是 i.size 而不是 i-size
	// index 处在第一个不小于 t 的数字的下标
	for j := i.size; j >= index; j-- {
		i.array[j+1] = i.array[j]
	}
	i.array[index] = t
	i.size++
}

func (i *IntSetArr) Size() int {
	return i.size
}

func (i *IntSetArr) Report() []int {
	var dst []int
	copy(dst, i.array)
	return dst
}
