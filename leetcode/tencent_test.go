package dance

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func reverseWords(s string) string {
	sByte := []byte(s)
	out := make([]byte, 0)
	beginIndex := 0
	previousSpace := false
	for i := 0; i < len(sByte); i++ {
		if sByte[i] == ' ' {
			for j := i - 1; i >= 0 && j >= beginIndex; j-- {
				out = append(out, sByte[j])
			}
			out = append(out, ' ')
			previousSpace = true
			beginIndex = 0
		} else {
			if previousSpace {
				beginIndex = i
				previousSpace = false
			}
		}
	}
	for j := len(sByte) - 1; j >= beginIndex; j-- {
		out = append(out, sByte[j])
	}
	return string(out)
}

//16. 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if sum := nums[i] + nums[j] + nums[k]; abs(sum-target) < abs(closestSum-target) {
					closestSum = sum
				}
			}
		}
	}
	return closestSum
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func TestTreeSumClosest(t *testing.T) {
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
}

//237. 删除链表中的节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//148. 排序链表
//对链表进行排序，而且要求常量级空间和 n log n 级的空间复杂度。
//链表排序与数组排序不同的是数组的游标可以左右移动，但是链表只有从头部向尾部移动
//有了这个限制，归并排序和快排这些有双向游标的排序算法就无法发挥作用了
//常量级空间的要求也决定了无法将数据读到数组中再进行排序
//现在使用排除法，先列出时间复杂度为 n log n 的排序算法
//归并排序，希尔排序，快速排序，堆排序

//看到相似题目中有合并两个有序链表，突然来了灵感
//我们可以将链表像拉链一样分开成两个子链表，然后再对子链表进行同样的操作，直到每个链表只剩一个元素，然后再将这些链表归并排序

func sortList(head *ListNode) *ListNode {
	next := head
	var first, firstLast, second, secondLast *ListNode = nil, nil, nil, nil
	position := 0
	for next != nil {
		if position%2 == 0 {
			if position == 0 {
				first = next
			} else {
				firstLast.Next = next
			}
			firstLast = next
		} else {
			if position == 1 {
				second = next
			} else {
				secondLast.Next = next
			}
			secondLast = next
		}
		position += 1
		next = next.Next
	}
	if firstLast != nil {
		firstLast.Next = nil
	}
	if secondLast != nil {
		secondLast.Next = nil
	}
	if position > 2 {
		first = sortList(first)
		second = sortList(second)
	}
	var sorted, sortedLast *ListNode = nil, nil
	for first != nil || second != nil {
		var currentMin *ListNode = nil
		if first == nil {
			currentMin = second
			second = second.Next
		} else if second == nil {
			currentMin = first
			first = first.Next
		} else {
			if first.Val > second.Val {
				currentMin = second
				second = second.Next
			} else {
				currentMin = first
				first = first.Next
			}
		}
		if sorted == nil {
			sorted = currentMin
		} else {
			sortedLast.Next = currentMin
		}
		sortedLast = currentMin
	}
	return sorted
}

func TestSortList(t *testing.T) {
	suits := [][]int{
		[]int{4, 2, 1, 3},
		[]int{-1, 5, 3, 4, 0},
	}
	head := constructorList(suits[1])
	sortedHead := sortList(head)
	fmt.Println(sortedHead)
}

//78. 子集
//这是一个 2的n次幂的问题，每一个字符在子集中有出现或者不出现两个选项，所以有 2 的 n 次幂
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	lengthPow := int(math.Pow(2, float64(length)))
	for i := 0; i < lengthPow; i++ {
		array := make([]int, 0)
		for j := 0; j < length; j++ {
			if ((1 << uint(j)) & i) > 0 {
				array = append(array, nums[j])
			}
		}
		ret = append(ret, array)
	}
	return ret
}

//78. 子集

func subsetsSecond(nums []int) [][]int {
	return nil
}

//238. 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	return nil
}

//136. 只出现一次的数字
//将所有数字与非到一起
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

//61. 旋转链表
// 先遍历一边获取长度 length  k = k % length，
// 将链表首尾相连，将 head 指针 向右移动 length - k 个位置
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	node, latest, length := head, head, 0
	for node != nil {
		latest = node
		length += 1
		node = node.Next
	}
	//注意这里，如果长度为 0 就直接返回 nil
	if length == 0 {
		return nil
	}
	k %= length
	if k == 0 {
		return head
	}
	latest.Next = head
	for i := 0; i < length-k; i++ {
		latest = head
		head = head.Next
	}
	latest.Next = nil
	return head
}

func TestRotateRight(t *testing.T) {
	head := constructorList([]int{1, 2, 3, 4, 5})
	ret := rotateRight(head, 2)
	fmt.Println(ret)
}

//15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	if length < 3 {
		return nil
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			//third := 0 - (nums[i] + nums[j])
		}
	}
	return nil
}

//11. 盛最多水的容器
func maxArea(height []int) int {
	return 0
}

//46. 全排列
// 这里要使用回溯算法

func permute(nums []int) [][]int {
	return nil
}

func dfs1(nums []int, left int) {
	for i := 0; i < left; i++ {

	}
}

func specialArray(nums []int) int {
	//sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if i < nums[j] {
				count += 1
			}
		}
		if len(nums)-count == i {
			return i
		}
	}
	return -1
}

func isEvenOddTree(root *TreeNode) bool {
	floor := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	next := make([]*TreeNode, 0)
	isOdd := func(a int) bool { return a%2 == 1 }
	isEven := func(a int) bool { return a%2 == 0 }
	isLess := func(a, b int) bool { return a < b }
	isGreat := func(a, b int) bool { return a > b }
	for {
		var LessGreat func(int, int) bool
		var OddEven func(int) bool
		if floor%2 == 1 {
			OddEven = isEven
			LessGreat = isGreat

		} else {
			OddEven = isOdd
			LessGreat = isLess
		}
		previous, isFirst := 0, true
		for i := 0; i < len(queue); i++ {
			theVal := queue[i].Val
			if !OddEven(theVal) {
				return false
			}
			if isFirst {
				isFirst = false
			} else if !LessGreat(previous, theVal) {
				return false
			}
			if queue[i].Left != nil {
				next = append(next, queue[i].Left)
			}
			if queue[i].Right != nil {
				next = append(next, queue[i].Right)
			}
			previous = theVal
		}
		if len(next) == 0 {
			break
		}
		queue = next
		next = make([]*TreeNode, 0)
		floor += 1
	}
	return true
}

func TestOdd(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Left.Left.Left = &TreeNode{Val: 12}
	root.Left.Left.Right = &TreeNode{Val: 8}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 2}
	t.Log(isEvenOddTree(root))
}
