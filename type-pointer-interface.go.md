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
