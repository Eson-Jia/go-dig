package pearls

import "sort"

type IntSetBST struct {
	size           int
	travel         int
	root, sentinel *bTreeNode
}

func NewIntSetBST(maxElements, maxVal int) *IntSetBST {
	sentinel := NewBTreeNode(maxVal, nil, nil)
	return &IntSetBST{
		sentinel: sentinel,
		root:     sentinel,
	}
}

type bTreeNode struct {
	left, right *bTreeNode
	value       int
}

func NewBTreeNode(t int, left, right *bTreeNode) *bTreeNode {
	return &bTreeNode{
		value: t,
		left:  left,
		right: right,
	}
}

func (i *IntSetBST) rinsert(node *bTreeNode, t int) *bTreeNode {
	if node == i.sentinel {
		node = NewBTreeNode(t, i.sentinel, i.sentinel)
		i.size++
		return node
	}
	switch {
	case node.value == t:
		// nothing
	case node.value < t:
		node.right = i.rinsert(node.right, t)
	case node.value > t:
		node.left = i.rinsert(node.left, t)
	}
	return node
}

func (i *IntSetBST) Insert(t int) {
	i.root = i.rinsert(i.root, t)
}

func (i *IntSetBST) Size() int {
	return i.size
}

func (i *IntSetBST) traverse(node *bTreeNode, dst []int) {
	// 因为有哨兵，所以应该永远不会为 nil,否则逻辑有错误
	if node == nil {
		panic("should never be nil")
	}
	if node == i.sentinel {
		return
	}
	dst[i.travel] = node.value
	i.travel++
	i.traverse(node.left, dst)
	i.traverse(node.right, dst)
}

func (i *IntSetBST) Report() []int {
	i.travel = 0
	dst := make([]int, i.size, i.size)
	i.traverse(i.root, dst)
	dst = dst[:i.size]
	sort.Slice(dst, func(a, b int) bool { return dst[a] < dst[b] })
	return dst
}
