package pearls

type IntSetList struct {
	sentinel, head *listNode
	size           int
}

func NewIntSetList(maxElements, maxVal int) *IntSetList {
	sentinel := newlistNode(maxVal, nil)
	head := sentinel
	return &IntSetList{
		sentinel: sentinel,
		head:     head,
	}
}

type listNode struct {
	value int
	next  *listNode
}

func newlistNode(value int, next *listNode) *listNode {
	return &listNode{value: value, next: next}
}

func (i *IntSetList) rinsert(node *listNode, value int) *listNode {
	if node.value < value {
		node.next = i.rinsert(node.next, value)
	} else if node.value > value {
		node = newlistNode(value, node)
		i.size++
	}
	return node
}

func (i *IntSetList) Insert(t int) {
	i.head = i.rinsert(i.head, t)
}

func (i *IntSetList) Size() int {
	return i.size
}

func (i *IntSetList) Report() []int {
	dist := make([]int, 0)
	begin := i.head
	for index := 0; index < i.size; index++ {
		dist = append(dist, begin.value)
		begin = begin.next
	}
	return dist
}
