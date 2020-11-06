package dance

import (
	"fmt"
	"testing"
)

//5543. 两个相同字符之间的最长子字符串
func maxLengthBetweenEqualCharacters(s string) int {
	sByte := []byte(s)
	marked := make(map[byte]int)
	theMax := -1
	for i, theByte := range sByte {
		if v, ok := marked[theByte]; ok {
			if distance := i - v - 1; distance >= 0 && distance > theMax {
				theMax = distance
			}
		} else {
			marked[theByte] = i
		}
	}
	return theMax
}

//5544. 执行操作后字典序最小的字符串
func findLexSmallestString(s string, a int, b int) string {
	return ""
}

//5545. 无矛盾的最佳球队
//不适合使用动态规划，因为假设 f(n) 为在前 n 个人中得分最高那支队伍
//f(n) = f(n-1)+n
//使用回溯
func bestTeamScore(scores []int, ages []int) int {
	//scoreSum := 0
	//length := len(scores)
	//best := make(map[int]int)
	//for i := 0; i < length; i++ {
	//
	//}
	return 0
}

func TestCharIndexForZero(t *testing.T) {
	fmt.Printf("%d--%d", int('0'), int('1'))
}
