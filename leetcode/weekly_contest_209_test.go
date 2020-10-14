package dance

import (
	"testing"
)

//1614. 括号的最大嵌套深度
func maxDepth1(s string) int {
	sByte := []byte(s)
	stack := make([]byte, 0)
	max := 0
	for _, element := range sByte {
		switch element {
		case '(':
			stack = append(stack, element)
			if length := len(stack); length > max {
				max = length
			}
		case ')':
			stack = stack[:len(stack)-1]
		default:
			// nothing
		}
	}
	return max
}

func TestMaxDepth(t *testing.T) {
	if max := maxDepth1("(1+(2*3)+((8)/4))+1"); max != 3 {
		t.Errorf("expected %d got %d", 3, max)
	}
}

//1615. 最大网络秩
//使用邻接表记录所有点与其连接的的点
//将临接表按照长度排序
//f(x)为 x点临接表的长度
//最大的网络秩就是 m = max(f(x)+ f(y)+(x<->y?1:0))
func maximalNetworkRank(n int, roads [][]int) int {
	graph := make([][]int, n)
	for _, road := range roads {
		c1, c2 := road[0], road[1]
		graph[c1] = append(graph[c1], c2)
		graph[c2] = append(graph[c2], c1)
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			base := len(graph[i]) + len(graph[j])
			for _, p := range graph[i] {
				if p == j {
					base -= 1
				}
			}
			if max < base {
				max = base
			}
		}
	}
	return max
}

func TestMaximalNetworkRank(t *testing.T) {
	maximalNetworkRank(4, [][]int{[]int{0, 1}, []int{0, 3}, []int{1, 2}, []int{1, 3}})
}

//1616. 分割两个字符串得到回文串
func checkPalindromeFormation(a string, b string) bool {
	aByte, bByte := []byte(a), []byte(b)
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if isLoop(string(aByte[:i])+string(bByte[i:])) ||
			isLoop(string(bByte[:i])+string(aByte[i:])) {
			return true
		}
	}
	return false
}
func isLoop(a string) bool {
	aLen := len(a)
	aByte := []byte(a)
	for i := 0; i < aLen; i++ {
		if aByte[i] != aByte[aLen-1-i] {
			return false
		}
	}
	return true
}
