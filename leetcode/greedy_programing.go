package dance

import (
	"sort"
)

// 递归贪心算法
func recursiveActivitySelector(s, f []int, k, n int) []int {
	var res []int
	m := k + 1
	for m <= n && s[m] < f[k] {
		m += 1
	}
	if m <= n {
		res = append(res, m)
		return append(res, recursiveActivitySelector(s, f, m, n)...)
	}
	return nil
}

// 迭代贪心算法
func greedyActivitySelector(s, f []int, n int) []int {
	var theSet []int
	finish := f[0]
	for m := 1; m <= n; m++ {
		if s[m] >= finish {
			theSet = append(theSet, m)
			finish = f[m]
		}
	}
	return theSet
}

//1589. 所有排列中的最大和
// 根据
func maxSumRangeQuery(nums []int, requests [][]int) int {
	return 0
}

//1589. 所有排列中的最大和
//先使用 map 统计所有下标的请求次数，然后将 map 中的数据存入到 slice 中并按照请求次数排序，
//最后将请求次数最多的下标设置为最大的数字，并依次类推计算总和。这个解决方案不行，最后通过了 72/84 个测试用例，没有通过的原因是超出时间限制
func maxSumRangeQueryFirst(nums []int, requests [][]int) int {
	record := make(map[int]int)
	for _, request := range requests {
		for i := request[0]; i <= request[1]; i++ {
			record[i]++
		}
	}
	pairSlice := make(PairSlice, len(record))
	{
		i := 0
		for k, v := range record {
			pairSlice[i] = Pair{
				K: k,
				V: v,
			}
			i++
		}
	}
	sort.Sort(sort.Reverse(pairSlice))
	sort.Ints(nums)
	i := len(nums) - 1
	sum := 0
	for _, pair := range pairSlice {
		sum += pair.V * nums[i]
		sum %= 1000000007
		i--
	}
	return sum
}

type Pair struct {
	K int
	V int
}

type PairSlice []Pair

func (p PairSlice) Len() int { return len(p) }
func (p PairSlice) Less(i, j int) bool {
	if p[i].V == p[j].V {
		return p[i].K < p[j].K
	}
	return p[i].V < p[j].V
}
func (p PairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//1589. 所有排列中的最大和
// 官网提示使用贪心算法，现使用贪心算法来实现
// 在该题目与贪心算法之间是问题归约，如何将该问题归到贪心算法是使用贪心算法的第一步
// 在使用
func maxSumRangeQuerySecond(nums []int, requests [][]int) int {
	return 0
}

//1221. 分割平衡字符串
// 这里使用贪心算法,满足两大关键要素：贪心选择性质和最优子结构性质。
// 贪心选择性质，在原字符串O中从头开始找到最短的平衡字符串s并分割,将原字符串分割成[s,L]。现在需要证明这个选择是最优选择，
// 考虑非空平衡字符串s，s0是其从第一个字符开始最短的平衡字符串，则s0是s的最多数量平衡字符串按照原始顺序排序的数组A中的第一个。
// 反证证明，假设sx是A中第一个元素，假设sx!=s0,那么 sx一定比s0长，那么可以将 sx拆分为 s0和另一个平衡字符串sx-0，将这两个字符串替换掉 sx,
// 则可以得到一个数组A1,那么A1中的平衡字符串数要比A的大1.这与假设A是数量最多的相悖。所以 sx=s0 即 s0是最大数组中的第一个元素。
// 最优子结构性质，如果一个子字符串 s 最优解数组第一个元素为 s0，那么 s的最优解 = s0+ (s-s0)的最优解
func balancedStringSplit(s string) int {
	balance := 0
	count := 0
	for _, alpha := range s {
		switch alpha {
		case 'L':
			balance -= 1
		case 'R':
			balance += 1
		}
		if balance == 0 {
			count += 1
		}
	}
	return count
}

// 392. 判断子序列
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}
	if j < len(s) {
		return false
	}
	return true
}
