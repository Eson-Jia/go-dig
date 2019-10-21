package pearls

import "sort"

type IntSetBST struct {
	size           int
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

func (i *IntSetBST) report(node *bTreeNode, dst map[int]struct{}) {
	if node == nil || node == i.sentinel {
		return
	}
	dst[node.value] = struct{}{}
	i.report(node.left, dst)
	i.report(node.right, dst)
}

func (i *IntSetBST) Report() (ret []int) {
	dst := make(map[int]struct{})
	i.report(i.root, dst)
	for k := range dst {
		ret = append(ret, k)
	}
	sort.Slice(ret, func(a, b int) bool { return ret[a] < ret[b] })
	return
}
