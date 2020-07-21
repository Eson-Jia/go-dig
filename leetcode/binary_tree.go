package dance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Previous struct {
	Value    int
	Modified bool
}

func isValidBST(root *TreeNode) bool {
	previous := Previous{
		Value:    0,
		Modified: false,
	}
	_, result := walk(root, previous)
	return result
}

func walk(node *TreeNode, previous Previous) (Previous, bool) {
	if node == nil {
		return previous, true
	}
	newPrevious, valid := walk(node.Left, previous)
	if !valid {
		return newPrevious, false
	}
	if newPrevious.Modified && newPrevious.Value >= node.Val {
		return newPrevious, false
	}
	newPrevious = Previous{
		Value:    node.Val,
		Modified: true,
	}
	newPrevious, valid = walk(node.Right, newPrevious)
	if !valid {
		return newPrevious, false
	}
	return newPrevious, true
}

func ConstructBT([]int) (root *TreeNode, err error) {
	return nil, nil
}
