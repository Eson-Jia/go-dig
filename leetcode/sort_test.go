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
	return nil
}
