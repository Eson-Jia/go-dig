package dance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	theList := make([]int, 0)
	_, result := walk(root, theList)
	return result
}

func walk(node *TreeNode, list []int) ([]int, bool) {
	if node == nil {
		return list, true
	}
	returnList, result := walk(node.Left, list)
	if !result {
		return returnList, result
	}
	if len(returnList) > 0 && returnList[len(returnList)-1] >= node.Val {
		return returnList, false
	}
	returnList = append(returnList, node.Val)
	returnList, result = walk(node.Right, returnList)
	if !result {
		return returnList, result
	}
	return returnList, true
}
