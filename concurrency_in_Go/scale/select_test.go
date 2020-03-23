package scale

import (
	"fmt"
	"testing"
)

//TestFunc 测试函数调用顺序
// func A(param1,param2,...)函数参数的求值顺序是从左到右依次求值
// 并且函数 A 等所有参数都完成后才会被调用
//output:
// entry paramProduct 1
// defer paramProduct 1
// entry paramProduct 2
// defer paramProduct 2
// entry DestFunc
// get value: [1 2]
// defer DestFunc
func TestFunc(t *testing.T) {
	paramProduct := func(value int) int {
		defer fmt.Println("defer paramProduct", value)
		fmt.Println("entry paramProduct", value)
		return value
	}
	DestFunc := func(param ...int) {
		defer fmt.Println("defer DestFunc")
		fmt.Println("entry DestFunc")
		fmt.Println("get value:", param)
	}
	DestFunc(paramProduct(1), paramProduct(2))
}

//TestSelect 测试 select 相关语法
// select 中的操作的 channel 如果是从函数求值得来的话，那么这些函数会在调用 select之前被调用，而且调用顺序
// 是自上往下依次调用
// 在 select 中两个 case 可以声明同名的变量，如下面的 result
func TestSelect(t *testing.T) {
	productChan := func(param string) <-chan interface{} {
		fmt.Println("entry productChan: ", param)
		out := make(chan interface{}, 1)
		out <- "productChan: " + param
		return out
	}

	select {
	case result := <-productChan("one"):
		t.Log(result)
	case result := <-productChan("two"):
		t.Log(result)
	}
}
