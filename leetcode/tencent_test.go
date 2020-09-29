package dance

import (
	"fmt"
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
//func sortList(head *ListNode) *ListNode {
//	return nil
//}

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
