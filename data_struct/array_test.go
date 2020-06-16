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

// 如果元素类型是不可比较的那么这个数组也是不可比较的
func TestArrayWithIncomparableElement(t *testing.T) {
	arrayOfSlice1 := [3][]int{
		{
			1, 2, 3,
		},
		{
			1, 2, 3, 4, 5, 6,
		},
		{
			4, 5, 6,
		},
	}
	arrayOfSlice2 := arrayOfSlice1
	fmt.Println(arrayOfSlice1, arrayOfSlice2)
	/**
	slice 是不可以比较的类型,用来作 array 的元素的话,这个数组也是不可比较的
	下面的 == 比较会报错 Invalid operation: arrayOfSlice1 == arrayOfSlice2 (operator == is not defined on [3][]int)
	if arrayOfSlice1 == arrayOfSlice2 {
		fmt.Println("equal")
	}
	*/

}
