package dance

import "fmt"

// 自顶向下递归实现
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

// 自底向上法
// 现在这个函数有问题，当 length 超过 price 的长度以后就会报错 index out of range
func ButtonUpCutRod(price []int, length int) int {
	result := make([]int, length+1)
	result[0] = 0
	for i := 1; i <= length; i++ {
		q := -1
		for j := 0; j < i; j++ {
			q = getMax(q, result[j]+price[i-j])
		}
		result[i] = q
	}
	return result[length]
}

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

// 现在这个函数有问题，当 length 超过 price 的长度以后就会报错 index out of range
// 这是因为切去部分的长度不能超过 price 表中最大的长度
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
	for i := 0; i < length+1; i++ {
		max = getMax(max, memoizedCutRod(price, result, i)+price[length-i])
	}
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

func waysToStep(n int) int {
	//stepRecord := map[int]int{1: 1,2:2,3:}
	for i := 1; i <= n; i++ {
		if n == 1 {
		}
	}
	return 0
}
