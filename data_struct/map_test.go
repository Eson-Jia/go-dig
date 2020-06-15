package data_struct

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

// 在对 map 进行取值操作的时候,即使键不存在也不会报错
func TestMapKeyNotExist(t *testing.T) {
	m := make(map[string]int)
	//1. 使用一元取值符取值不存在时会返回类型零值
	v1 := m["dog"]
	fmt.Println(v1)
	//2. 键不存在会返回零值,本操作使得 dog 值为 1
	m["dog"]++
	//3. 当我们需要知道是键不存在还是值就是零值的时候我们可以使用第二个返回值,
	// 第二个 bool 返回值当键不存在的时候返回 false 否则返回 true
	v, ok := m["dog"]
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

// 尝试往值为 nil 的 map 中设置值会触发 panic
func TestSetNilMap(t *testing.T) {
	var m map[string]int
	//panic: assignment to entry in nil map [recovered]
	//	panic: assignment to entry in nil map
	m["foo"] = 0
	t.Log(m)
}

// 创建 map 的几种方式
func TestCreateMap(t *testing.T) {
	//1 使用字面量表达式创建空的 map
	m0 := map[string]int{
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
	fmt.Println(m0, m1, m2)
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

// 每次遍历 map 的顺序是不固定的
func TestMapRangeKey(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	buff := bytes.NewBuffer(nil)
	for k, _ := range m {
		fmt.Fprint(buff, k, " ")
	}
	fmt.Println(buff.String())
	buff.Reset()
	for k, _ := range m {
		fmt.Fprint(buff, k, " ")
	}
	fmt.Println(buff.String())
	//output:
	// d e a b c
	// c d e a b
}

// 因为 map 遍历顺序是不固定的,所以如果想顺序遍历的话需要在外面先对 key 进行排序
func TestMapSorted(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s ==> %d\n", k, m[k])
	}
}
