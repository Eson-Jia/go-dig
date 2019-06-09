# 类型具有某个方法的具体含义

```go
package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

//实现了接受者是值类型的方法，相当于自动实现接收者是指针类型的方法。
func (g Gopher) code() {
	fmt.Printf("I am coding %s language\n", g.language)
}

//但是实现了接受者是指针类型的方法，不会自动生成对应接受者是值类型的方法。
func (g *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", g.language)
}

func main() {
	//如果将取地址符去掉，即将Gopher值类型字面量赋值给 coder接口会报错 Gopher 没有实现debug 方法
	var c coder = &Gopher{"GO"}
	c.code()
	c.debug()
}
```

接收者是指针类型的方法，很可能在方法中对接收者的属性进行更该操作，从而影响接收者。
而接收者是值类型的方法，会对调用者做一次拷贝，不会影响到接受者。

所以，当方法的接收者是值类型时，就可以自动生成一个接收者是对应指针类型的方法，因为两者都不会影响到调用者。
但是，当实现了一个接收者是指针类型的方法是，如果此时自动生成一个接收者是值类型的方法，原本期望对接受者的改变不会实现，因为值类型
会产生一个拷贝，不会真正影响到调用者。

类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
