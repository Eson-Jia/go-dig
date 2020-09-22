package dance

import (
	"sort"
	"testing"
)

//1589. 所有排列中的最大和
// 使用 C++ 的解决方案是先统计所有请求中所有下标的引用次数
// 然后将下标按照其引用次数排序(需要用到将map<int,int>中的pair 按照 value 排序的技术)
// 将引用次数最多的下标值设置为数组中的最大值，并以此类推
// 该解决方法是比较笨的方法，而且运行时间过长无法跑过所有的测试单例
func maxSumRangeQueryFirst(nums []int, requests [][]int) int {
	record := make(map[int]int)
	for _, request := range requests {
		for i := request[0]; i < request[1]; i++ {
			record[i]++
		}
	}
	// 按照引用次数对数组下标排序
	// *****将 golang 中的****
	//-----------------------
	// 将数组排序
	sort.Ints(nums)
	//按照排序结果将数组中最大值放入引用次数最多的下标中
	return 0
}

//1589. 所有排列中的最大和
// 官网提示使用贪心算法，现使用贪心算法来实现
// 在该题目与贪心算法之间是问题归约，如何将该问题归到贪心算法是使用贪心算法的第一步
// 在使用
func maxSumRangeQuerySecond(nums []int, requests [][]int) int {
	return 0
}

func TestMaxSUmRangeQuery(t *testing.T) {
	suits := []struct {
		Nums    []int
		Request [][]int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[][]int{
				[]int{
					1, 3,
				},
				[]int{0, 1},
			},
		},
	}
	//maxSumRangeQueryFirst()
}
