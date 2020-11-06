package dance

//面试题 10.09. 排序矩阵查找
//每一行每一列都按照升序
//左上有最小值，右下有最大值
//我们可以选择将右上角上的值 T 跟目标值 target 进行对比
//如果小于目标值那么T所在行可以剔除（因为T是该行最大的值），
//如果大于目标值那么T所在列可以被剔除（因为T是该列最小的值）
//这样我们每次比较都可一删除一行或者一列直到我们找到目标值，或者发现目标值不再矩阵中
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	row, column := 0, columns-1
	for row < rows && column >= 0 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}

//704. 二分查找
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := (low + high) / 2
		switch {
		case nums[middle] == target:
			return middle
		case nums[middle] < target:
			low = middle + 1
		case nums[middle] > target:
			high = middle - 1
		}
	}
	return -1
}

//剑指 Offer 11. 旋转数组的最小数字
func minArray(numbers []int) int {
	length := len(numbers)
	low, high := 0, length-1
	for low < high-1 {
		if numbers[low]<=numbers[high]{
			return numbers[low]
		}
		middle := (low + high) / 2
		if numbers[low] < numbers[middle] {
			low = middle
		} else {
			high = middle
		}
	}
	return numbers[high]
}
