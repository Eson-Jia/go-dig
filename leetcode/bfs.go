package dance

// 二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	nextLists := []*TreeNode{}
	if root != nil {
		nextLists = append(nextLists, root)
	}
	reverse := false
	for len(nextLists) > 0 {
		currentLists := nextLists
		nextLists = make([]*TreeNode, 0)
		layer := make([]int, 0)
		for _, currentNode := range currentLists {
			layer = append(layer, currentNode.Val)
			if currentNode.Left != nil {
				nextLists = append(nextLists, currentNode.Left)
			}
			if currentNode.Right != nil {
				nextLists = append(nextLists, currentNode.Right)
			}
		}
		if reverse {
			layer = reverseSlice(layer)
		}
		reverse = !reverse
		result = append(result, layer)
	}
	return result
}
func reverseSlice(origin []int) []int {
	result := make([]int, 0)
	for i := len(origin) - 1; i >= 0; i-- {
		result = append(result, origin[i])
	}
	return result
}
