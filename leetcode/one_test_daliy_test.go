package dance

//每日一题

/**
日期: 2021/5/24
链接: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
状态: 通过
*/
// deleteDuplicates
func deleteDuplicates(head *ListNode) *ListNode {
	previous := head
	current := head
	for current != nil {
		if previous.Val == current.Val {
			previous.Next = current.Next
		} else {
			previous = current
		}
		current = current.Next
	}
	return head
}

/**
时间: 2021/5/25
链接: https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
状态: 未通过
子序列：是由原字符串删除部分字符构成的，注意字符的相对位置没有变化
子串：是由原字符串删除头部和尾部部分字符构成的
异同点：子序列可以不连续，子串必须连续
最长特殊序列：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）
*/
func findLULength(a string, b string) int {
	return 0
}

/**
155. 最小栈
时间: 2021/5/26
状态: 未通过
链接: https://leetcode-cn.com/problems/min-stack/
*/
type MinStack struct {
	Array []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{Array: make([]int, 0)}
}

func (this *MinStack) Push(val int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {
	return 0
}

func (this *MinStack) GetMin() int {
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/**
202. 快乐数
date: 2021/6/1 19:28
链接:https://leetcode-cn.com/problems/happy-number/
*/
func isHappy(n int) bool {
	current := n
	cache := map[int]struct{}{n: {}}
	for {
		sum := 0
		cache[current] = struct{}{}
		for current > 0 {
			sum += (current % 10) * (current % 10)
			current /= 10
		}
		current = sum
		if current == 1 {
			return true
		}
		if _, ok := cache[current]; ok {
			return false
		}
	}
}

/**
剑指 Offer 04. 二维数组中的查找
Date: 2021/6/2
https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
从右上角开始查找
*/
func findNumberIn2DArrayFromRightTop(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for rowIndex, columnIndex := 0, len(matrix[0])-1; rowIndex <= len(matrix)-1 && columnIndex >= 0; {
		if matrix[rowIndex][columnIndex] > target {
			columnIndex--
		} else if matrix[rowIndex][columnIndex] < target {
			rowIndex++
		} else {
			return true
		}
	}
	return false
}

/**
从左下向右上开始查找
*/
func findNumberIn2DArrayFromLeftButton(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for rowIndex, columnIndex := len(matrix)-1, 0; rowIndex >= 0 && columnIndex <= len(matrix[0])-1; {
		if matrix[rowIndex][columnIndex] > target {
			rowIndex--
		} else if matrix[rowIndex][columnIndex] < target {
			columnIndex++
		} else {
			return true
		}
	}
	return false
}
