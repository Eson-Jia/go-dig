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

// 测试 [][]int 双重切片，第二层如何构造，使用 for _,row := range rows 拿到的第二层数组变量之后调用 make 是否能够正常构造一个二维数组，
// 我认为在使用 range 的情况下是不行的，因为 row 只是 原始数组的拷贝，对它进行修改不会进影响原始的数组，就跟在一个函数中修改 slice 一样，
// 函数外的 slice 不受影响（对 slice 重新赋值，外层数组不受影响。如果只是修改 slice元素内容，那么外层 slice 也会相应改变）。
// 经过我测试，发现情况的确如上所述，而且这段代码都不能通过编译，因为 row 变量只赋值的话会报错 Unused variable 'row'
func TestDoubleSlice(t *testing.T) {
	doubleSlice := make([][]int, 10)
	for _, row := range doubleSlice {
		row = make([]int, 10)
		row[0] = 100 // 如果不加这就会报错：Unused variable 'row'
	}
	t.Log(doubleSlice) // [[] [] [] [] [] [] [] [] [] []]
}

// 所以正确的用法就是使用 for index
func TestDoubleSliceInit(t *testing.T) {
	doubleSlice := make([][]int, 10)
	for i := 0; i < len(doubleSlice); i++ {
		doubleSlice[i] = make([]int, 10)
	}
	t.Log(doubleSlice)
}
