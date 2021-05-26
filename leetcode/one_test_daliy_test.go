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
