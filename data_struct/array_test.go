package data_struct

import (
	"fmt"
	"testing"
)

// 如果数组中的元素是可以比较的,那么这个数组也是可以比较
// 比较的结果是两边元素的值是否完全相同
func TestArrayCompare(t *testing.T) {
	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{3, 4}
	fmt.Println(a1 == a2, a1 == a3, a2 == a3)
	/* 元素个数不相等的数组不是同一类型,当然不能进行比较
	*  Invalid operation: a1 ==a4 (mismatched types [2]int and [3]int)
	*  a4 := [3]int{1,2}
	*  fmt.Println(a1 ==a4)
	 */
}
