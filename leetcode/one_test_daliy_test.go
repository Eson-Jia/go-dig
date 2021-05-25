package dance

import "testing"

//每日一题

// 2021/5/24
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	previous := head
	current := head
	for current != nil {
		if previous.Val == current.Val {

		}
		current = current.Next
	}
	return head
}
func Test_DeleteDuplicates(t *testing.T) {

}

// 2021/5/25
//https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
/**
子序列：是由原字符串删除部分字符构成的，注意字符的相对位置没有变化
子串：是由原字符串删除头部和尾部部分字符构成的
异同点：子序列可以不连续，子串必须连续
最长特殊序列：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）
*/
// findLULength
func findLULength(a string, b string) int {
	return 0
}
