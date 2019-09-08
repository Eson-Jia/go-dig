package main

import (
	"fmt"
	"testing"
)

func tempDir() []string {
	return []string{
		"/a",
		"/b",
		"/c"}
}

func TestCapture(t *testing.T) {
	var funcs []func()
	for _, dir := range tempDir() {
		dir := dir //如果不加这句话，数组中所有函数都捕获同一个dir变量，而dir在循环执行到最后的值为 "/c",则所有的删除匿名函数都尝试去删除"/c"目录
		funcs = append(funcs, func() {
			fmt.Println("dir:", dir)
		})
	}
	for _, pro := range funcs {
		pro()

	}
}
