package dance

import (
	"sort"
	"testing"
)

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

/**
剑指 Offer 05. 替换空格
https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
Date:6/3
*/

func replaceSpace(s string) string {
	space := 0
	for _, i2 := range []byte(s) {
		if i2 == ' ' {
			space++
		}
	}
	dest := make([]byte, len(s)+2*(space))
	for i, b := range []byte(s) {
		dest[i] = b
	}
	originIndex := len(s) - 1
	destIndex := len(dest) - 1
	for originIndex >= 0 {
		if dest[originIndex] == ' ' {
			dest[destIndex] = '0'
			destIndex--
			dest[destIndex] = '2'
			destIndex--
			dest[destIndex] = '%'
			destIndex--
		} else {
			dest[destIndex] = dest[originIndex]
			destIndex--
		}
		originIndex--
	}
	return string(dest)
}

/**
https://leetcode-cn.com/problems/target-sum/
Date: 6/7
思路: 动态规划用来寻找最优解,但是本例中只是用来遍历所有解法,所以应该不用动态规划.
因为不断尝试在原来结果基础上加上或者减去某个数,所以我觉得需要使用回溯算法.
在寻找答案的过程中,我们不断在原来的基础上做选择(- or +),并且在最后的结果不符合要求的时候,
回溯到上一个选项并从新选择另一个选项.
16:43 需要遍历 2 ** len(nums)种选项
17:01 突然有个想法,使用 一个数 0 <= i < 2 ** len(nums),每次 +1 递增,然后每一个数组下标与其与,结果是 1 就加上,0 就减去
18:20 能得到正确结果,但是运行时间太长了.总结一下是因为遍历的方式不像回溯法那样回退到之前步骤的结果.
例如,1+1+1+1 最后一步是 +1 或者 -1,回溯法就可以使用之前结果直接执行 4+1 或者 4-1 .但是遍历法需要从头开始计算,当数组越长,回溯法的优势越明显.
6/8
9:47 使用回溯算法,利用调用栈实现中间结果的存储
*/
func findTargetSumWays(nums []int, target int) int {
	theLen := len(nums)
	ope := []int{-1, 1}
	count := 0
	for i := 0; i < 1<<theLen; i++ {
		sum := 0
		for j := 0; j < theLen; j++ {
			sum += ope[(i&(1<<j))>>j] * nums[j]
		}
		if sum == target {
			count++
		}
	}
	return count
}

/**
6/8 10:41 使用我以为的回溯算法,虽然所有的测试用例都跑过了(138),但最终还是超过运行时间,暂时还没想起来更好的方法
*/
func findTargetSumWaysWithBack(nums []int, target int) int {
	operators := []int{-1, 1}
	count := 0
	for _, nextOperator := range operators {
		count += calculateCurrent(nums, 0, operators, nextOperator, 0, target)
	}
	return count
}

func calculateCurrent(nums []int, numsIndex int, operators []int, operator int, previousSum int, target int) int {
	sum := previousSum + operator*nums[numsIndex]
	if numsIndex == len(nums)-1 {
		if target == sum {
			return 1
		}
		return 0
	}
	count := 0
	for _, nextOperator := range operators {
		count += calculateCurrent(nums, numsIndex+1, operators, nextOperator, sum, target)
	}
	return count
}

func TestFindTargetSumWays(t *testing.T) {
	if result := findTargetSumWaysWithBack([]int{1, 1, 1, 1, 1}, 3); result == 5 {
		t.Log("good")
	} else {
		t.Errorf("error want:%d got:%d", 3, result)
	}

}

/**
1046. 最后一块石头的重量
https://leetcode-cn.com/problems/last-stone-weight/
6/8 10:46
可以通过优先队列来实现,但是 golang 没有优先队列,只能每次修改完元素之后,重新排序
6/9 17:32 通过
*/

func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		sort.Ints(stones)
		if y, x := stones[len(stones)-1], stones[len(stones)-2]; x == y {
			stones = stones[:len(stones)-2]
		} else {
			stones = stones[:len(stones)-2]
			stones = append(stones, y-x)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

/**
204. 计数质数
https://leetcode-cn.com/problems/count-primes/
Date: 6/10
埃拉托斯特尼筛法: 要得到自然数 n 以内的全部素数,必须把不大于根号 n 的所有素数的倍数剔除,剩下的就是素数.
*/

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	flags := make([]bool, n+1)
	sqrtN := 0
	for i := 0; i < n; i++ {
		if i*i > n {
			sqrtN = i
			break
		}
	}
	for i := 2; i <= sqrtN; i++ {
		if flags[i] == true {
			continue
		}
		for j := 2 * i; j < n; j += i {
			flags[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if flags[i] == false {
			count++
		}
	}
	return count
}
