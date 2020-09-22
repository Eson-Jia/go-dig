package dance

import "fmt"

func ButtonUpCutRod(price []int, len int) int {
	result := make([]int, len+1)
	result[0] = 0
	for i := 1; i <= len; i++ {
		q := -1
		for j := 0; j < i; j++ {
			q = getMax(q, result[j]+price[i-j])
		}
		result[i] = q
	}
	return result[len]
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

func MemoizedCutRod(price []int, len int) int {
	result := make([]int, len+1)
	return memoizedCutRod(price, result, len)
}

func memoizedCutRod(price []int, result []int, len int) int {
	if len == 0 {
		return 0
	}
	if result[len] > 0 {
		return result[len]
	}
	max := -1
	for i := 0; i < len+1; i++ {
		max = getMax(max, memoizedCutRod(price, result, 0)+price[len-i])
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
