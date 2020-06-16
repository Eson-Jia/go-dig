package data_struct

import (
	"fmt"
	"log"
	"testing"
)

// TestCreateSlice
// 学习 slice make 语法的使用
// slice 有三个属性: 指针,长度,容量
// 指针指向 slice 在 底层数组中的第一个元素
// 长度指的是 slice 中元素的个数,不能超过 slice 的容量
// 容量指的是从 slice 起始元素到底层数组最后一个元素之间元素的个数
func TestCreateSlice(t *testing.T) {
	// 创建新的 slice 有 3 种方式
	// 1. 使用内置函数 make
	// 第一个 size 指定 slice 的 length,第二个 size 指定 capacity,注意 capacity <= length,否则编译报错
	s1 := make([]int, 10, 15)
	s1 = append(s1, 12)
	// 2.使用字面量表达式
	s2 := []int{1, 2, 3, 4, 5, 6}
	// 注意 与下面使用字面量表达式创建数组的区别
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	// 3.使用切片表达式在原有 slice或者 array 的基础上创建新的
	s3 := s2[2:]
	t.Log(s1, s2, s3, a1)
}

// TestSliceIsReference
// slice 包含指针指向底层的数组,多个 slice 可以引用同一个底层数组的不同位置,其中的元素可以交叠
func TestSliceIsReference(t *testing.T) {
	a1 := [...]string{"cat", "dog", "mouse", "pig", "monkey"}
	s1, s2 := a1[1:], a1[1:]
	//1. 当一个 slice 修改了某个元素,另一个 slice 中同一元素的值当然会改变
	s1[2] = "PIG"
	log.Print(s1, s2)
	//2. 当修改了多个 slice 共同的底层数组的值,这些 slice 中对应的元素值也会变化
	a1[3] = "pig"
	log.Print(s1, s2)
	//3. 因为 slice 包含了指向底层数组的指针,所以将 slice 传给函数,函数可以修改底层数组元素中的值
	upper := func(s []string) {
		s[3] = "PIG"
	}
	upper(s1)
	log.Println(a1)
	//4. 使用 append 修改 slice 的时候,如果 slice 包含的位置位于底层数组中间,那么位于 slice 后面的元素就会被覆盖掉
	{
		a1 := [...]string{"cat", "dog", "mouse", "pig", "monkey"}
		s4 := a1[0:1]
		s4 = append(s4, "tidy")
		log.Println("a1:", a1, "s4:", s4)
	}
}

// TestSliceAppend
// 使用

type IntSlice struct {
	ptr *int
	len int
	cap int
}

func TestSliceAppend(t *testing.T) {

	// slice
}

// TestNewCapacity
// 当通过 append 添加元素时,当底层数组的长度不够时,会创建一个长度为原始长度*2的新数组
func TestNewCapacity(t *testing.T) {
	var s1 []int
	for i := 0; i < 10; i++ {
		t.Logf("%d len:%d\t cap:%d\t %v\n", i, len(s1), cap(s1), s1)
		s1 = append(s1, i)
	}
}

// TestSliceOperator
func TestSliceOperator(t *testing.T) {
	// s[i:j]( 0<=i<=j<=cap(s) )
	// slice 操作符创建一个新的 slice ,包含 s 中[i,j)区间的元素
	// 注意对 i,j 的限制是 <=cap(s) 而不受 s length 的限制
	a := [...]int{1, 2, 3, 4, 5, 6}
	// 所以我们可以使用切片操作得到更大长度的 slice
	// s 只有一个元素
	s := a[0:1]
	t.Logf("len:%d\t cap:%d\t %v\n", len(s), cap(s), s)
	// 对 s 进行切片操作,注意 i j 大于 len(s) 是合法的,如果 slice 的引用超过被引用对象的长度,即len(s),那么最终的 slice 会比原 slice 长
	s1 := s[4:6]
	t.Logf("len:%d\t cap:%d\t %v\n", len(s1), cap(s1), s1)
}

// slice 无法进行比较操作,除了跟 nil 进行比较
func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	fmt.Println(s1, s2)
	/**
	slice 无法跟 nil 之外的表达式比较,否则会报语法错误:
	Invalid operation: s1 == s2 (operator == is not defined on []int)
		if s1 == s2 {
	}
	**/

}
