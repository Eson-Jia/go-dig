package data_struct

import (
	"fmt"
	"testing"
)

func TestMapKeyNotExist(t *testing.T) {
	m := map[string]int{}
	fmt.Println(m["dog"])
	v, ok := m["dog"]
	fmt.Println(v, ok)
	m["dog"]++
	v, ok = m["dog"]
	fmt.Println(v, ok)
}

// 值为 nil 的 map 可以安全的执行 查找 删除 range 遍历 和 len
func TestMapIsNil(t *testing.T) {
	var m map[string]int
	fmt.Println("is nil:", m == nil)
	//1. 查找
	v, ok := m["foo"]
	fmt.Println(v, ok)
	//2. 删除
	delete(m, "foo")
	//3. 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	//4. len()
	fmt.Println(len(m))
}

// 创建 map 的几种方式
func TestCreateMap(t *testing.T) {
	//1 使用字面量表达式创建空的 map
	m0:=map[string]int{
		// empty
	}
	//1.1 使用字面量表达式同时填入值
	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	//2. 使用 make,第二个参数作用是预先 n 个元素的空间
	//An empty map is allocated with enough space to hold the specified number of elements.
	m2 := make(map[string]int, 10)
	fmt.Println(m0,m1, m2)
}

// python 中的 dict 不支持遍历删除 会报错 dictionary changed size during iteration
// 但是 golang 就没有这个限制,可以一边遍历一边删除
func TestRangeAndDelete(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	for k, v := range m {
		t.Log("delete", k, v)
		delete(m, k)
	}
}
