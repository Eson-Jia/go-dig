package dance

import "fmt"

// 1.自顶向下递归实现
// 现在需要明确一个概念，切割完成之后的钢条中的每段的长度都不会超过 price 表中最大的长度
func cutRod(price []int, length int) int {
	if length == 0 {
		return 0
	}
	max := -1
	// i 表示的是切去的部分，切去的部分肯定是可以定价的，不会超过 price 表中最大的长度
	// i 又必须不能大于总长度 length
	for i := 1; i < len(price) && i <= length; i++ {
		max = getMax(max, price[i]+cutRod(price, length-i))
	}
	return max
}

// 3.自底向上法
func ButtonUpCutRod(price []int, length int) int {
	result := make([]int, length+1)
	result[0] = 0
	for i := 1; i <= length; i++ {
		q := -1
		// i 表示的是切去的部分，切去的部分肯定是可以定价的，不会超过 price 表中最大的长度
		//又必须不能大于总长度 length
		for j := 1; j <= i && j <= len(price)-1; j++ {
			q = getMax(q, price[j]+result[i-j])
		}
		result[i] = q
	}
	return result[length]
}

// 扩展自底向上法
func ExtendedButtonUpCutRod(price []int, len int) ([]int, []int) {
	result, firstCut := make([]int, len+1), make([]int, len+1)
	for i := 1; i < len+1; i++ {
		q := -1
		for j := 0; j < i; j++ {
			t := result[j] + price[i-j]
			if t >= q {
				q = t
				firstCut[i] = i - j
			}
		}
		result[i] = q
	}
	return result, firstCut
}

func PrintButtonUpCutRod(price []int, len int) {
	result, firstCut := ExtendedButtonUpCutRod(price, len)
	fmt.Println("amount:", result[len])
	fmt.Println("list:")
	for len > 0 {
		current := firstCut[len]
		fmt.Println(current)
		len -= current
	}
}

// 2. 自带备忘的自顶向下发法
func MemoizedCutRod(price []int, length int) int {
	result := make([]int, length+1)
	return memoizedCutRod(price, result, length)
}

func memoizedCutRod(price []int, result []int, length int) int {
	if length == 0 {
		return 0
	}
	if result[length] > 0 {
		return result[length]
	}
	max := -1
	// i 在这里是被切去的部分，i不大于总长度且不大于 price 表中最大的长度
	for i := 1; i <= length && i < len(price); i++ {
		max = getMax(max, price[i]+memoizedCutRod(price, result, length-i))
	}
	result[length] = max
	return max
}

func ExtendedMemoizedCutRod(price []int, len int) ([]int, []int) {
	firstCut, result := make([]int, len+1), make([]int, len+1)
	extendedMemoizedCutRod(price, firstCut, result, len)
	return firstCut, result
}

func extendedMemoizedCutRod(price []int, firstCut []int, result []int, len int) int {
	max := -1
	for i := 0; i < len+1; i++ {
		max = getMax(max, extendedMemoizedCutRod(price, firstCut, result, i)+1)
	}
	return 0
}

//面试题 08.01. 三步问题
// 这是很基础的动态规划
// 我们考虑现在已经n阶上面，有几种方法一步走到n阶呢，我们可以在n-1阶上走一步1阶到n,也可以在 n-2 阶走一步2阶，也可以在n-3阶走一步3阶
// 那么 一共有 f(n) = f(n-1)+f(n-2)+f(n-3)种走法
// 动态规划有两要素：最优子结构，重叠子问题
// 这里有多少种走法就是该问题的唯一解，也可以理解为最优解
// n 的解是 n-1,n-2,n-3 子问题的解的和
// 重叠子问题，如果递归算法反复求解子问题，我们就称最优化问题具有重叠子问题性质
// f(n) = f(n-1)+f(n-2)+f(n-3)
// f(n-1) = f(n-2)+f(n-3)+f(n-4)
// 可以看出求解子问题并不会引出更多的子问题
func waysToStep(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= n; i++ {
		count := 0
		// 注意这里 j可以等于i,因为如果 n <=3 我们一步就可以走到
		for j := 1; j <= i && j <= 3; j++ {
			count += result[i-j]
			count %= 1000000007
		}
		result[i] = count
	}
	return result[n]
}

// 1143. 最长公共子序列
// 最长公共子序列与最长公共子串的区别就是公共子序列中的元素在原始字符串 str1,str2 不需要连续，
// 例如： A,C,A 就是 A,B,D,C,B,A 和 A,C,B,A 的一个公共子串，但不是最长公共子序列。A,C,B,A 是最长公共子序列
// 而最长公共子串是 B,A
// 因为公共子序列的性质，我们可以刻画最长公共子序列的性质，str1,str2 的最长子序列为 k，
// 如果 str1[-1] == str2[-1], 那么 k[:-1] 是 str1[:-1] 和 str2[:-1]的最长公共子串
// 如果 str1[-1] != str2[-1]
// 		如果 str1[-1] != k[-1] 那么 k 是 str1[:-1] 和 str2 的 LCS
//      如果 str2[-1] != k[-1] 那么 k 是 str1 和 str2[:-1] 的 LCS
// 假设 str1,str2 长度分别为 rows,column,我们创建一个 c[rows][column] 的两维数组,
// c[i][j]表示 str1 前 i 个元素与 str2 前 j 个元素的 LCS 的长度 ，例如，str1 前4个元素是 A,B,D,C ,str2 前 3个元素是 A,C,B
// 那么他们的 LCS 就是 A,B   c[4][3] = 2

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	c := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		c[i] = make([]int, len2+1)
	}
	for i := 0; i < len1+1; i++ {
		c[i][0] = 0
	}
	for j := 0; j < len2+1; j++ {
		c[0][j] = 0
	}
	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if text1[i-1] == text2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else if c[i-1][j] > c[i][j-1] {
				c[i][j] = c[i-1][j]
			} else {
				c[i][j] = c[i][j-1]
			}
		}
	}
	return c[len1][len2]
}

// 计算 LCS 的长度
// c[i][j]为 Xi 和 Yj LCS 的长度，Xi 和 Yj 有一个为 0 c[i][j]==0
//
//
func LCSLength(text1, text2 string) ([][]int, [][]byte) {
	len1, len2 := len(text1), len(text2)
	c, b := make([][]int, len1+1), make([][]byte, len1+1)
	for i := 0; i < len1+1; i++ {
		c[i] = make([]int, len2+1)
		b[i] = make([]byte, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = '\\'
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = '^'
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = '<'
			}
		}
	}
	return c, b
}

func PrintLCS(b [][]byte, text1 string, i, j int) []byte {
	buffer := make([]byte, 0)
	var printLCS func([][]byte, string, int, int)
	printLCS = func(b [][]byte, text1 string, i, j int) {
		switch b[i][j] {
		case '^':
			printLCS(b, text1, i-1, j)
		case '<':
			printLCS(b, text1, i, j-1)
		case '\\':
			buffer = append(buffer, text1[i-1])
			printLCS(b, text1, i-1, j-1)
		}
	}
	printLCS(b, text1, i, j)
	buffLen := len(buffer)
	reverse := make([]byte, buffLen)
	for i := buffLen - 1; i >= 0; i-- {
		reverse[buffLen-1-i] = buffer[i]
	}
	return reverse
}

// 疑问，最长公共子序列和最长公共子串的共同点在哪里，是否可以转化？
// 我认为最长公共子串是最长公共子序列的特殊情况。特殊点是公共子串每个字符之间没有其他不同的字符。所以公共子序列的算法在
// 公共子串上同样试用。

// 最长公共子串
// 最优子结构  有限子问题域
// 现在想如何实现最优子结构
func longestCommonSubstring(text1 string, text2 string) int {
	return 0
}
