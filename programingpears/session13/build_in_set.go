package pearls

import "sort"

type IntSetSTL struct {
	set map[int]struct{}
}

func NewIntSetSTL(maxElements, maxVal int) *IntSetSTL {
	return &IntSetSTL{set: make(map[int]struct{})}
}

func (i *IntSetSTL) Insert(t int) {
	i.set[t] = struct{}{}
}

func (i *IntSetSTL) Size() int {
	return len(i.set)
}

func (i *IntSetSTL) Report() []int {
	var list []int
	for k := range i.set {
		list = append(list, k)
	}
	sort.Slice(list, func(i int, j int) bool { return list[i] < list[j] })
	return list
}
