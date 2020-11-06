package coding_interview

import (
	"fmt"
	"log"
	"testing"
)

//剑指 Offer 12. 矩阵中的路径
//使用回溯算法

func exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows < 1 {
		return false
	}
	columns := len(board[0])
	if columns < 1 {
		return false
	}
	markRecord := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if hasPathTo(board, rows, columns, i, j, word, markRecord, 0) {
				return true
			}
		}
	}
	return false
}

func hasPathTo(board [][]byte, rows, columns, row, column int, word string, markRecord map[int]bool, pathLength int) bool {
	if pathLength == len(word) {
		return true
	}
	if !(row < rows && row >= 0 && column < columns && column >= 0) {
		return false
	}
	if marked, ok := markRecord[row*columns+column]; ok && marked {
		return false
	}
	if word[pathLength] != board[row][column] {
		return false
	}
	markRecord[row*columns+column] = true
	hasPath := false
	hasPath = hasPathTo(board, rows, columns, row+1, column, word, markRecord, pathLength+1) ||
		hasPathTo(board, rows, columns, row-1, column, word, markRecord, pathLength+1) ||
		hasPathTo(board, rows, columns, row, column+1, word, markRecord, pathLength+1) ||
		hasPathTo(board, rows, columns, row, column-1, word, markRecord, pathLength+1)
	if !hasPath {
		markRecord[row*columns+column] = false
	}
	return hasPath
}

func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	t.Log(exist(board, word))
}

//剑指 Offer 13. 机器人的运动范围
// 使用回溯

func movingCount(m int, n int, k int) int {
	marked := make([]bool, m*n)
	var dfs func(rows, columns, row, column int) int
	dfs = func(rows, columns, row, column int) int {
		count := 0
		if !(row < rows && row >= 0 && column < columns && column >= 0) {
			return count
		}
		if marked[row*columns+column] {
			return count
		}
		marked[row*columns+column] = true
		if getBitSum(row)+getBitSum(column) > k {
			return count
		}
		count += 1
		count = count +
			dfs(rows, columns, row+1, column) +
			dfs(rows, columns, row-1, column) +
			dfs(rows, columns, row, column+1) +
			dfs(rows, columns, row, column-1)
		return count

	}
	return dfs(m, n, 0, 0)
}

func getBitSum(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}
func TestMovingCount(t *testing.T) {
	t.Log(movingCount(1, 2, 1))
}

//剑指 Offer 14- I. 剪绳子

//f(n) 为 长度为 n的绳子的最大乘积
//f(n) = max(f(k)*f(n-k))(0<=k<=n)

func cuttingRope(n int) int {
	if n <= 4 {
		return n
	}
	max := n
	for i := 1; i < n; i++ {
		max = getMax(max, cuttingRope(i)*cuttingRope(n-1))
	}
	return max
}

func hammingWeight(num int32) int {
	count := 0
	for num != 0 {
		if (num & 1) == 1 {
			count += 1
		}
		num >>= 1
		log.Printf("%b", num)
	}
	return count
}

func TestHammingWeight(t *testing.T) {
	t.Log(hammingWeight(-100))
}

func printNumbers(n int) []string {
	buff := make([]byte, n)
	arr := make([]string, 0)
	for i := 0; i < 10; i++ {
		bfs(buff, 0, n, byte('0'+i), func(num []byte) {
			i := 0
			for ; i < len(num); i++ {
				if num[i] != '0' {
					break
				}
			}
			if i == len(num) {
				arr = append(arr, "0")
			} else {
				arr = append(arr, string(num[i:]))
			}
		})
	}
	return arr
}

func bfs(buff []byte, index, length int, c byte, print func(num []byte)) {
	if index <= length-1 {
		buff[index] = c
		if index == length-1 {
			print(buff)
		} else {
			for i := 0; i < 10; i++ {
				bfs(buff, index+1, length, byte('0'+i), print)
			}
		}
	}
}

func PrintNumber(num []byte) {
	i := 0
	for ; i < len(num); i++ {
		if num[i] != '0' {
			break
		}
	}
	if i == len(num) {
		fmt.Println("0")
	} else {
		fmt.Println(string(num[i:]))
	}
}

func TestPrintNumbers(t *testing.T) {
	results := printNumbers(10)
	theSet := make(map[string]bool)
	for _, result := range results {
		if _, ok := theSet[result]; ok {
			t.Errorf("duplicated:%s", result)
		}
	}
	t.Logf("the length is %d\n", len(results))
}
