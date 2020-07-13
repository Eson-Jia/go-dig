package dance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Max struct {
	Current  int
	Modified bool
}

func isValidBST(root *TreeNode) bool {
	max := Max{
		Current:  0,
		Modified: false,
	}
	_, result := walk(root, max)
	return result
}

func walk(node *TreeNode, max Max) (Max, bool) {
	if node == nil {
		return max, true
	}
	returnMax, result := walk(node.Left, max)
	if !result {
		return returnMax, result
	}
	if returnMax.Modified && returnMax.Current >= node.Val {
		return returnMax, false
	}
	returnMax.Modified = true
	returnMax.Current = node.Val
	returnMax, result = walk(node.Right, returnMax)
	if !result {
		return returnMax, result
	}
	return returnMax, true
}
