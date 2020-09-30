package dance

import (
	"fmt"
	"testing"
)

func xor() int {
	result := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("%08b^%08b=", result, i)
		result ^= i
		fmt.Printf("%08b\n", result)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%08b^%08b=", result, i)
		result ^= i
		fmt.Printf("%08b\n", result)
	}
	return result
}

func TestXor(t *testing.T) {
	t.Log(xor())
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	stack := make([]*ListNode, 0)
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	for i := 1; i < len(stack); i++ {
		stack[i].Next = stack[i-1]
	}
	stack[0].Next = nil
	return stack[len(stack)-1]
}

func sumOfLeftLeaves(root *TreeNode) int {
	return get(root, false)
}

func get(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}
	return get(node.Left, true) + get(node.Right, false)
}

//968. 监控二叉树
func minCameraCover(root *TreeNode) int {
	return 0
}

//501. 二叉搜索树中的众数

/*给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：
	结点左子树中所含结点的值小于等于当前结点的值
	结点右子树中所含结点的值大于等于当前结点的值
	左子树和右子树都是二叉搜索树*/

func findMode(root *TreeNode) []int {
	return nil
}

//145. 二叉树的后序遍历
// 这里采用递归的方式遍历
func postorderTraversalRecursive(root *TreeNode) []int {
	ret := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		ret = append(ret, node.Val)
	}
	traversal(root)
	return ret
}

//145. 二叉树的后序遍历
// 这里采用迭代的方式遍历
type stackStruct struct {
	Node  *TreeNode
	Left  bool
	Right bool
}

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := []stackStruct{{
		Node:  root,
		Left:  false,
		Right: false,
	},
	}
	length := len(stack)
	for length > 0 {
		if stack[length-1].Node == nil {
			stack = stack[:length-1]
		} else if !stack[length-1].Left {
			stack = append(stack, stackStruct{
				Node: stack[length-1].Node.Left,
			})
			stack[length-1].Left = true
		} else if !stack[length-1].Right {
			stack = append(stack, stackStruct{
				Node: stack[length-1].Node.Right,
			})
			stack[length-1].Right = true
		} else {
			ret = append(ret, stack[length-1].Node.Val)
			stack = stack[:length-1]
		}
		length = len(stack)
	}
	return ret
}
