package dance

import "testing"

// 2021/8/25
// 排序算法

/**
归并排序
2021/8/20
*/
func mergeSort(nums []int) []int {
	l := len(nums)
	//递归终止条件
	if l == 1 {
		return nums
	}
	internal := l / 2
	return merge(mergeSort(nums[:internal]), mergeSort(nums[internal:]))
}

func merge(a1 []int, a2 []int) []int {
	l1, l2 := len(a1), len(a2)
	temp := make([]int, l1+l2)
	for i1, i2, t := 0, 0, 0; i1 < l1 || i2 < l2; t++ {
		switch {
		case i1 < l1 && i2 < l2:
			// 这里的比较符号决定升序还是降序
			if a1[i1] < a2[i2] {
				temp[t] = a1[i1]
				i1++
			} else {
				temp[t] = a2[i2]
				i2++
			}
		case i1 < l1:
			temp[t] = a1[i1]
			i1++
		case i2 < l2:
			temp[t] = a2[i2]
			i2++
		}
	}
	return temp
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	result := mergeSort(nums)
	t.Log(result)
}

/**
快速排序
*/
func quickSort(nums []int) []int {
	helper(nums, 0, len(nums)-1)
	return nums
}

func helper(nums []int, begin, end int) {
	if end-begin == 0 {
		return
	}
	if end-begin == 1 {
		/*
			todo
		*/
	}
	index := partial(nums, begin, end)
	if index-1 > begin {
		helper(nums, begin, index-1)
	}
	if index+1 < end {
		helper(nums, index+1, end)
	}
}

/*
分区函数是快速排序算法的灵魂所有
分区函数首先选出一个元素 x,使用 x 将数组分为 A1, x, A2 三部分
其中 A1 中所有的元素 <= x,A2 中所有元素 >= x.
*/
func partial(nums []int, begin, end int) int {
	pritvol := (begin + end) / 2
	i, j := begin, end
	for i != j {
		if nums[i] < nums[pritvol] {
			i++
			continue
		}
		if nums[j] > nums[pritvol] {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return i
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 2, 3, 1}
	result := quickSort(nums)
	t.Log(result)
}
