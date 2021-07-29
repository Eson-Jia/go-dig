package dance

import (
	"fmt"
	"math/bits"
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

6/8 10:41
使用我以为的回溯算法 findTargetSumWaysRecursive ,虽然所有的测试用例都跑过了(138),但最终还是超过运行时间,暂时还没想起来更好的方法

6/29 15:30 优化了回溯算法(也就是递归遍历法) findTargetSumWaysRecursive,精简了代码实现.现在可以跑过测试

7/2 9:57
看了官方解题,可以使用动态规划来解题
dp 有两个要素: 1.有限问题域 2. 最优子结构
很明显枚举是有限问题域,现在来看最优子结构,额,怎么算最优呢?正好等于 target,然后可以寻找子结构 target = (+/-)第一元素 + 其他元素的组合
如果我们选定第一个元素的符号为(+) 那么子结构就是 = (target - 第一元素) = (2...n)元素组合,但是现在有问题就是,一共有 1>>n 个选择,而不是 n 个
选择,其实最终也是遍历枚举了(1>>n)

7/2 10:38
看了官方 dp 解法的提示,使用了一个值 sums 即所有元素的和.然后就有灵感了,因为所有元素都为正整数,我们可以对这些元素求和,然后在一个个减去,直到最终结果为 target.

7/6 10:05
使用 numbs 的所有和 sums 依次减去元素的方式到这就没有思路了.
查看官方解题,思路如下:因为每个元素都是正整数, target = positive - negative = (positive + negative) - 2*negative
= sums - 2*negative; 2 * negative = sums -target; neg = (sums - target)/2;
那么问题就可以转化为:在 nums 中找到所有和为 neg 的所有组合.
这个问题就可以转化为背包问题了:有若干个物品重量在数组 nums 中,一个容量为 neg 的背包,求恰好能装满背包的所有组合方式.
利用动态规划来解背包问题
dp[i][j] 表示前 i 个数 和为 j 的个数有多少个

7/7 15:25  根据官方提示实现了 findTargetSumWaysDP
使用动态规划来解背包问题

dp[i][j] 表示在前 i 个物品中,组合后质量为 j 的组合方式有多少个
状态转移方程:

dp[i][j]的值由两部分构成:
1. 当不选择最后一个物品时的组合个数: dp[i-1][j]
2. 选择最后一个物品(nums[i])时的前 i-1 个物品质量为 j-nums[i] 的组合个数: dp[i-1][j - nums[i]]
dp[0][0] = 1
dp[0][>0] = 0

7/9 10:02
官方解题提示算法还可以进一步优化,因为 dp[i][j] 只会依赖 dp[i-1][<=j]的数据,所以我们可以滚动地使用一行数据而非 len(nums) 行
findTargetSumWaysDPOptimal 一直无法得到正确的结果

7/12 11:04
明白为什么无法得到正确的结果了,为了节省内存采用了一行数据滚动使用, dp[i][j] 依赖 dp[i-1][k](k<j)的结果,但是如果从 0 到 negative 的方向
遍历的话,dp[i-1][k] 的值会被 dp[i][k]覆盖掉,导致 dp[i][j] 的值不正确.所有我们应该采取从 negative 到 0遍历的方式
*/
func findTargetSumWays(nums []int, target int) int {
	theLen := len(nums)
	ope := []int{-1, 1}
	count := 0
	for i := 0; i < 1<<theLen; i++ {
		sum := 0
		for j := 0; j < theLen; j++ {
			sum += ope[(i>>j)&1] * nums[j]
			/**
			6/10 今天再次看代码的时候发现,可以将位操作优化到上面这个,更简单,更容易理解
			sum += ope[(i&(1<<j))>>j] * nums[j]
			*/
		}
		if sum == target {
			count++
		}
	}
	return count
}

func findTargetSumWaysRecursive(nums []int, target int) int {
	return calculateCurrent(nums, 0, 1, 0, target) + calculateCurrent(nums, 0, -1, 0, target)
}

func calculateCurrent(nums []int, numsIndex, ope int, previousSum int, target int) int {
	sum := previousSum + ope*nums[numsIndex]
	if numsIndex == len(nums)-1 {
		if target == sum {
			return 1
		}
		return 0
	}
	return calculateCurrent(nums, numsIndex+1, 1, sum, target) + calculateCurrent(nums, numsIndex+1, -1, sum, target)
}

func findTargetSumWaysDP(nums []int, target int) int {
	sums := 0
	for _, num := range nums {
		sums += num
	}
	result := 0
	if result = sums - target; result < 0 || result%2 == 1 {
		return 0
	}
	result /= 2
	length := len(nums)
	dp := make([][]int, length+1)
	for i, _ := range dp {
		dp[i] = make([]int, result+1)
	}
	dp[0][0] = 1
	for i := 1; i <= length; i++ {
		for j := 0; j <= result; j++ {
			dp[i][j] = dp[i-1][j]
			if temp := j + nums[i-1]; temp <= result {
				dp[i][temp] += dp[i-1][j]
			}
		}
	}
	return dp[length][(sums-target)/2]
}

func findTargetSumWaysDPOptimal(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	negative := 0
	if negative = sum - target; negative < 0 || negative%2 == 1 {
		return 0
	}
	negative /= 2
	dp := make([]int, negative+1)
	dp[0] = 1
	for i := 1; i <= len(nums); i++ {
		current := nums[i-1]
		for j := negative; j >= 0; j-- {
			if current <= j {
				dp[j] += dp[j-current]
			}
		}
	}
	return dp[negative]
}

func TestFindTargetSumWays(t *testing.T) {
	if result := findTargetSumWaysDPOptimal([]int{1, 2, 1}, 0); result == 5 {
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

/**
6/10
快速排序
*/

func quickSort(nums []int) {
}

/**
剑指 Offer 15. 二进制中1的个数
https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
Date: 6/15
*/

func hammingWeight(num uint32) int {
	count := 0
	for {
		if num <= 0 {
			break
		}
		count++
		num &= num - 1
	}
	return count
}

func hammingWeightBasic(num uint32) int {
	count := 0
	for i := 0; i < 2<<5; i++ {
		if num&(1<<i) > 0 {
			count++
		}
	}
	return count
}

/**
https://leetcode-cn.com/problems/binary-watch/
2021/6/21
回溯算法
对树形或者图形结构执行一次深度优先遍历,实际上类似枚举的搜索尝试过程,在遍历的过程中寻找问题的解.
深度优先遍历有个特点:当发现已不满足求解条件时,就返回,尝试别的路径.此时对象变量就需要重置成为和之前一样,称为"状态重置".
许多复杂的,规模较大的问题都可以使用回溯算法.实际上,回溯算法就是暴力搜索算法.
代表小时的和代表分钟的,小时的可能亮灯数为:0-4 一共五个可能每一种一种或有多种排列组合
1. 第一步使用朴素的方式尝试解题
6/23
查看提示说,遍历所有时间可能的组合(12*60)然后在组合里面查找符合灯数的条目
这不就是遍历吗,能用得到回溯算法吗?
可以在某一组合小时的灯数超过总灯数直接不再遍历分钟,也算是一种回溯算法吧.
潜意识里,使用遍历算法就会遍历所有值,对于输入参数的不同值使用遍历算法总觉得会浪费.
这是不正确的,因为不同的输入参数,可以使遍历算法剪除不符合条件的分支,不会全部遍历.
小时可能值为 0-11,分钟可能值为 0-59.

6/23 11:25
暴露的问题是解题时脑子不会拐弯,不会逆向思维.当时的解题思路:
将 turnedOn 个灯分配给小时和分钟.假如 小时有 h 个灯,那么 h 个灯总共有多少组合方式以及如何遍历(这个是计算整数 n 中有多少个 1 的逆运算,
一个数中有 n 个 1,求这个数的所有可能值),而且还有不符合条件的,因为有的亮灯组合会超过 11,这样的结果要被剔除.
分钟会有同样的情况,这些边界情况都会使代码复杂.
但是换个方向,逻辑就很简单了:遍历所有可能的时间组合,计算组合亮灯数(计算整数 n 中有多少个 1)选出符合条件的时间组合.

6/23  11:49
使用 bits.OneCount 代替自己实现的计算数中的 二进制1

6/25 11:02
另一个思路就是遍历 10 个灯开合的所有组合(共有 1<<10 种),i 从 0 到 1<<10,其中 i 的 前四位代表小时,后六位代表分钟,
把不符合亮灯数和时间不合法的的组合过滤掉即可得到最终结果.该解法参考 readBinaryWatchSecond
*/
func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)
	for i := uint8(0); i < 12; i++ {
		hourCount := bits.OnesCount8(i)
		if hourCount > turnedOn {
			continue
		}
		for j := uint8(0); j < 60; j++ {
			if hourCount+bits.OnesCount8(j) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}

func readBinaryWatchSecond(turnedOn int) []string {
	result := make([]string, 0)
	for i := uint32(0); i < (1 << 10); i++ {
		hour, minute := i>>6, i&0b111111
		if turnedOn != bits.OnesCount32(i) {
			continue
		}
		if hour >= 12 || minute >= 60 {
			continue
		}
		result = append(result, fmt.Sprintf("%d:%02d", hour, minute))
	}
	return result
}

/**
https://leetcode-cn.com/problems/sum-of-all-subset-xor-totals/
6/28
9:56
寻找数组 array 中元素的子集的遍历,就涉及到排列组合,这时,遍历从 0 到 1>>len(array)
之间的每一个数 i,i 的二级制格式中的 j 是 0/1,代表着 i 当前子集是否包含 array[j] 元素
6/29 10:29
看了官方解题,一共有三种解法,这个是第二个解法叫:迭代法枚举子集 subsetXORSum
现在来试试第一种解法:递归枚举子集法 subsetXORSumRecursive

6/30 9:56
第三种解法:按位考虑 + 二项式展开

*/
func subsetXORSum(nums []int) int {
	length := len(nums)
	sum := 0
	for i := 0; i < 1<<length; i++ {
		result := 0
		for j := 0; j < length; j++ {
			if (1<<j)&i > 0 {
				result ^= nums[j]
			}
		}
		sum += result
	}
	return sum
}

func subsetXORSumRecursive(nums []int) int {
	result := new(int)
	subsetXORSumR(nums, 0, true, 0, result)
	subsetXORSumR(nums, 0, false, 0, result)
	return *result
}

func subsetXORSumR(nums []int, index int, include bool, xorResult int, sums *int) {
	if len(nums)-1 >= index {
		if include {
			xorResult ^= nums[index]
		}
	}
	if len(nums)-1 == index {
		*sums += xorResult
		return
	}
	subsetXORSumR(nums, index+1, true, xorResult, sums)
	subsetXORSumR(nums, index+1, false, xorResult, sums)
}

/**
https://leetcode-cn.com/problems/bracket-lcci/
https://leetcode-cn.com/problems/generate-parentheses/
generateParenthesis
2021/7/13 10:07
自然想到 n 对括号的组合就是 n-1 对的括号组合再加上一对括号,那么如何将第 n 个括号加到原有的组合中呢
或者使用回溯,或者使用枚举,先生成所有可能的组合,然后判断是否合法
除了枚举法之外,需要注意的问题是如何在原有子序列的基础上生成合法的序列
现在尝试使用回溯方法构建
1. ()
2. ()() (())
3. ()()() (()()) ()(()) (())() ((()))
2021/7/14 10:30
在看了 allPossibleFBT 问题的时候,突然发现该问题就是一个树(非二叉树)形状结构,这也是一种解题思路,那么问题就可以转化为 n 个节点组成森林有多少种方式?

17:02
原来想着将上面的
*/
func generateParenthesis(n int) []string {
	return nil
}

/**
https://leetcode-cn.com/problems/all-possible-full-binary-trees/
2021/7/14 10:35
现在有个疑问:奇数个肯定是可以构造成满二叉树,是否节点为偶数个无法形成满二叉树?
采用归纳法可以证明所有的奇数都是可以构造成满二叉树.
先尝试使用朴素的方式解题:
额,没想起来
使用动态规划的话需要找到最优子结构,然后找到状态转移方程:
额,没想起来

2021/7/14 15:37
突然有个

7/14 16:45
allPossibleFBT generateParenthesis 都是树相关的
*/
func allPossibleFBT(n int) []*TreeNode {
	return nil
}

/**
https://leetcode-cn.com/study-plan/dynamic-programming/?progress=qm448d
开启动态规划学习计划
*/

/**
***Day 1 ***
 */

/**
https://leetcode-cn.com/problems/fibonacci-number/

使用一维动态规划即可解出 fibDP
可以看出只是用了数组最后那两个,所有为了节省内存可以不使用数组改为滚动使用两个变量 fibDPOptimal
*/
func fibDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibDPOptimal(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		second, first = second+first, second
	}
	return second
}

func TestFibDPOptimal(t *testing.T) {
	t.Log(fibDPOptimal(8))
}

/**
https://leetcode-cn.com/problems/n-th-tribonacci-number/
*/
func tribonacciDP(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func tribonacciDPOptimal(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	f, s, t := 0, 1, 1
	for i := 3; i <= n; i++ {
		t, s, f = t+s+f, t, s
	}
	return t
}

/**
*** Day 2 ***
 */

/**
https://leetcode-cn.com/problems/climbing-stairs/
70. 爬楼梯
状态转移方程 f(n) = f(n-1) + f(n-2)
f(0)=1
f(1)=1
f(2)=2
*/

func climbStairs(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	f, s := 1, 1
	for i := 2; i <= n; i++ {
		s, f = f+s, s
	}
	return s
}

/**
https://leetcode-cn.com/problems/min-cost-climbing-stairs/
最优子结构:到达顶层 n 的方式有两种,从 n-1 花费 cost[n-1],或者从 n-2 花费 cost[n-2].
f(n) 为到达 n 层的最小花费
状态转移方程:f(n) =min(f(n-1)+cost(n-1),f(n-2)+cost(n-2))
f(0)

一维数组解法 minCostClimbingStairs
节省内存采取滚动使用两个变量的方法 minCostClimbingStairsMemoryOptimal
因为都要到达楼层顶部,但是顶部 minCostClimbingStairsOptimal
*/
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	best := make([]int, n)
	best[0] = cost[0]
	best[1] = cost[1]
	for i := 2; i < n; i++ {
		best[i] = min(best[i-1], best[i-2]) + cost[i]
	}
	return min(best[n-1], best[n-2])
}

func minCostClimbingStairsMemoryOptimal(cost []int) int {
	n := len(cost)
	f, s := cost[0], cost[1]
	for i := 2; i < n; i++ {
		s, f = min(s, f)+cost[i], s
	}
	return min(s, f)
}

func minCostClimbingStairsOptimal(cost []int) int {
	newCost := append(cost, 0)
	n := len(newCost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + newCost[i]
	}
	return dp[n-1]
}

/**
*** Day 3 ***
 */

/**
https://leetcode-cn.com/problems/house-robber/
f(n) = max(f(n-1),f(n-2)+nums[n])
*/
func rob(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		second, first = max(second, first+nums[i]), second
	}
	return second
}

/**
https://leetcode-cn.com/problems/house-robber-ii/
f(0,n)  = max(f(0,n-1),f(1,n-2)+nums[n])

7/20 10:40
根据上面的公式,可以看出该问题可以分解成求两个不同范围的打家劫舍
*/
func rob2(nums []int) int {
	n := len(nums)
	switch n {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	return max(rob(nums[:n-1]), rob(nums[1:n-2])+nums[n-1])
}

/**
https://leetcode-cn.com/problems/delete-and-earn/
1. 每次操作中，选择任意一个 nums[i]
2. 删除之后必须删除所有等于 nums[i]-1/nums[i]+1 的元素。
3. 根据标签提示感觉需要使用哈希表 map

现在问题是：
1. 如何求问题的子结构？
2. 如何进行问题转化：求和吗？
3. 如何删除所有等于 nums[i]-/+1的元素？将其重置为 0 吗？

2021/7/19
现有思路：
1. 可以将数放入 map 中 key 是数 value 是出现次数，因为需要权衡该数与其左右邻居谁的和更大
2. 从一个方向开始删除 ,所以将 key 排序,
3. 该问题可以转化为一组数组里面有不同的值,数组是挨着的(0代表左右不挨着),如果删除一个值就需要把其左右的值清掉,求能达到的最大值.

17:57
呀,问题转化以后,不就成了打家劫舍的问题了吗?哈哈哈
*/
func deleteAndEarn(nums []int) int {
	zip := make(map[int]int, 0)
	for _, num := range nums {
		zip[num]++
	}
	keys := make([]int, 0)
	for k, _ := range zip {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	final := make([]int, 0)
	previous := 0
	for _, key := range keys {
		if previous != 0 && previous != key-1 {
			final = append(final, 0)
		}
		final = append(final, zip[key]*key)
		previous = key
	}
	// 然后这里就跟 rob的解题思路一样了
	return rob(final)
}

/**
*** Day 4 ***
 */

/**
https://leetcode-cn.com/problems/jump-game/
最后一个元素没有作用
只能往前跳
有向无环图
二维数组
深度优先遍历算法:遍历有向图,直到遇到最后一个节点
深度优先遍历消耗太多的内存 canJumpDFS
7/22 10:10 还是得需要使用动态规划 canJumpDP
假设 f(n) 为 n 数组的最远距离,那么 <n 的所有下标是否都可以到达?
是的,可以用反证法证明所有下标都可以到达:假设有一点 0<x<n 无法到达
f(n)表示在数组下标 <=n 的元素中,能到的最大的下标.
状态转移方程:
f(n) == n 说明在这里就断了,无法再前进 f(n+1) = f(n)
f(n) > n 的话说明没有断,可以跳到下一个位置 n+1
那么 f(n+1) = max(f(n),(n+1)+nums[n+1])
*/
func canJumpDFS(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	df := make([][]int, length)
	for i := 0; i < length; i++ {
		df[i] = make([]int, 0)
	}
	for i := 0; i < length; i++ {
		for j := 1; j <= nums[i]; j++ {
			df[i] = append(df[i], i+j)
		}
	}
	var directedDFS func(int) bool
	directedDFS = func(index int) bool {
		for _, p := range df[index] {
			if p == length-1 {
				return true
			}
			if directedDFS(p) {
				return true
			}
		}
		return false
	}
	return directedDFS(0)
}

func canJumpDP(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	for i := 0; i < length; i++ {
		if dp[i-1] == i-1 {
			return false
		}
		dp[i] = max(i+nums[i], dp[i-1])
		if dp[i] >= length-1 {
			return true
		}
	}
	return false
}

func TestCanJump(t *testing.T) {
	t.Log(canJumpDP([]int{1, 2}))
}

/**
https://leetcode-cn.com/problems/jump-game-ii/
f(n) 表示跳到下标 n 所需的最小步数
*/
func jump(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	dp := make([]int, length)
	for i := 1; i < length; i++ {
		minJump := -1
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				if minJump == -1 {
					minJump = dp[j]
				} else {
					minJump = min(minJump, dp[j])
				}
			}
		}
		dp[i] = minJump + 1
	}
	return dp[length-1]
}

/**
https://leetcode-cn.com/problems/maximum-subarray/
*/
func maxSubArray(nums []int) int {
	return 0
}
