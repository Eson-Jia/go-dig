package dance

//跟定一个二叉树，找出其再最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	var rightMaxDepth, leftMaxDepth int
	if node.Right != nil {
		rightMaxDepth = dfs(node.Right)
	}
	if node.Left != nil {
		leftMaxDepth = dfs(node.Left)
	}
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	}
	return rightMaxDepth + 1
}

func maxDepthEasy(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
