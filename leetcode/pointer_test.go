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
