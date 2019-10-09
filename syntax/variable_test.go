package syntax

import (
	"fmt"
	"testing"
)

// TestVariable 形参变量都是函数的局部变量，初始值由调用者提供的实参传递。
// 函数形参以及命名返回值同属于函数最外层作用域的局部变量。
// f 函数的形参 a 是其最外层的局部变量,就相当于在 函数体第一行的如下声明: var a int
// 这两个声明处于同一作用域
func TestVariable(t *testing.T) {
	f := func(a int) *int {
		// var a int ！a redeclared in this block
		return &a
	}
	for i := 0; i < 10; i++ {
		result := f(10)
		fmt.Printf("%d,%v\n", result, result)
	}
}

// Closure 由上面测试案例可知，amount 是该函数中最外层作用域的局部变量。
// 该函数返回一个匿名函数，这个匿名函数是个闭包(closure),
// 因为它引用了外层作用域中的 amount 变量
// 闭包就是定义在一个函数内部的函数
func Closure(amount int) func() int {
	return func() int {
		amount -= 1
		return amount
	}
}

func TestClosure(t *testing.T) {
	AMOUNT := 5
	f := Closure(AMOUNT)
	for i := 0; i < AMOUNT; i++ {
		fmt.Println("f():", f())
	}
}

// TestForCompare
// for initialization; condition; post {
// 零个或多个语句
// }
// initialization 在循环开始之前执行，condition 在每次迭代开始之前执行，post 在每次迭代完成后执行
func TestForCompare(t *testing.T) {
	f := func(amount int) int {
		fmt.Println("get amount:", amount)
		return amount
	}
	for i := 0; i < f(4); i++ {
		fmt.Println("the value of i:", i)
	}
}
